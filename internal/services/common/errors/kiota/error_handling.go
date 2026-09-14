package errors

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
)

// ErrorDescription contains standardized error messages and summaries
type ErrorDescription struct {
	Summary string
	Detail  string
}

// GraphErrorInfo contains extracted information from a Graph API error
type GraphErrorInfo struct {
	StatusCode      int
	ErrorCode       string
	ErrorMessage    string
	Target          string
	IsODataError    bool
	AdditionalData  map[string]any
	Headers         *abstractions.ResponseHeaders
	RequestDetails  string
	RetryAfter      string
	RequestID       string
	ClientRequestID string
	ErrorDate       string
	InnerErrors     []InnerErrorInfo
	ErrorDetails    []ErrorDetailInfo
	CorrelationID   string
	ThrottledReason string
	Category        ErrorCategory
	DiagnosticInfo  string
}

// InnerErrorInfo contains information from nested inner errors
type InnerErrorInfo struct {
	Code        string
	Message     string
	ODataType   string
	RequestID   string
	ClientReqID string
	Date        string
	Target      string
}

// ErrorDetailInfo contains information from error details array
type ErrorDetailInfo struct {
	Code    string
	Message string
	Target  string
}

// ErrorCategory represents different types of errors
type ErrorCategory string

const (
	CategoryAuthentication ErrorCategory = "authentication"
	CategoryAuthorization  ErrorCategory = "authorization"
	CategoryValidation     ErrorCategory = "validation"
	CategoryThrottling     ErrorCategory = "throttling"
	CategoryService        ErrorCategory = "service"
	CategoryNetwork        ErrorCategory = "network"
	CategoryUnknown        ErrorCategory = "unknown"
)

// GraphErrorOptions enables endpoint-specific handling without changing existing callers.
type GraphErrorOptions struct {
	// PreserveStateOnReadBadRequest keeps validation errors from removing existing state.
	PreserveStateOnReadBadRequest bool
	// PreserveStateOnReadNotFound keeps unavailable singleton endpoints in state.
	PreserveStateOnReadNotFound bool
}

// HandleKiotaGraphError processes Graph API errors with the existing default semantics.
func HandleKiotaGraphError(
	ctx context.Context,
	err error,
	resp any,
	operation string,
	requiredPermissions []string,
) {
	HandleKiotaGraphErrorWithOptions(
		ctx,
		err,
		resp,
		operation,
		requiredPermissions,
		GraphErrorOptions{},
	)
}

// HandleKiotaGraphErrorWithOptions applies endpoint-specific semantics to shared diagnostics.
func HandleKiotaGraphErrorWithOptions(
	ctx context.Context,
	err error,
	resp any,
	operation string,
	requiredPermissions []string,
	options GraphErrorOptions,
) {
	errorInfo := GraphError(ctx, err)
	errorDesc := getErrorDescription(errorInfo.StatusCode)

	tflog.Debug(ctx, "Handling Graph error:", map[string]any{
		"status_code":    errorInfo.StatusCode,
		"operation":      operation,
		"error_code":     errorInfo.ErrorCode,
		"error_message":  errorInfo.ErrorMessage,
		"target":         errorInfo.Target,
		"is_odata_error": errorInfo.IsODataError,
		"request_id":     errorInfo.RequestID,
		"correlation_id": errorInfo.CorrelationID,
		"inner_errors":   len(errorInfo.InnerErrors),
		"error_details":  len(errorInfo.ErrorDetails),
		"category":       errorInfo.Category,
	})

	// Record error metrics for observability
	recordErrorMetrics(ctx, &errorInfo, operation)

	// Handle special cases first
	switch errorInfo.StatusCode {
	case 400:
		// URL errors use a synthetic 400; they are not evidence that a resource was deleted.
		var urlError *url.Error
		if operation == constants.TfOperationRead && !options.PreserveStateOnReadBadRequest &&
			!errors.As(err, &urlError) {
			if _, ok := resp.(*resource.ReadResponse); ok {
				tflog.Warn(
					ctx,
					"Resource appears to no longer exist (400 Response), removing from state",
				)
				removeResourceFromState(ctx, resp)
				return
			}
		}
		addErrorToDiagnostics(ctx, resp, errorDesc.Summary,
			constructDetailedErrorMessage(errorDesc.Detail, &errorInfo))

	case 401, 403:
		tflog.Warn(
			ctx,
			fmt.Sprintf(
				"Permission error on %s operation, check required Graph permissions",
				operation,
			),
		)
		handlePermissionError(ctx, errorInfo, resp, operation, requiredPermissions)
		return

	case 404:
		if operation == constants.TfOperationRead && !options.PreserveStateOnReadNotFound {
			if _, ok := resp.(*resource.ReadResponse); ok {
				tflog.Warn(ctx, "Resource not found (404 Response), removing from state")
				removeResourceFromState(ctx, resp)
				return
			}
		}
		if operation == constants.TfOperationDelete {
			tflog.Info(
				ctx,
				"Resource already deleted or does not exist (404 Response), treating as successful deletion",
			)
			return
		}
		addErrorToDiagnostics(ctx, resp, errorDesc.Summary,
			constructDetailedErrorMessage(errorDesc.Detail, &errorInfo))

	case 429:
		if operation == constants.TfOperationRead {
			tflog.Warn(ctx, "Rate limit exceeded on read operation")
			handleRateLimitError(ctx, errorInfo, resp)
			return
		}
		addErrorToDiagnostics(ctx, resp, errorDesc.Summary,
			constructDetailedErrorMessage(errorDesc.Detail, &errorInfo))

	case 503:
		if operation == constants.TfOperationRead {
			tflog.Warn(
				ctx,
				"Service Unavailable (503 Response), service is temporarily unavailable",
			)
			handleServiceUnavailableError(ctx, errorInfo, resp)
			return
		}
		addErrorToDiagnostics(ctx, resp, errorDesc.Summary,
			constructDetailedErrorMessage(errorDesc.Detail, &errorInfo))

	default:
		addErrorToDiagnostics(ctx, resp, errorDesc.Summary,
			constructDetailedErrorMessage(errorDesc.Detail, &errorInfo))
	}
}

// HandleKiotaGraphErrorForAction processes Graph API errors and returns a formatted error string for actions
// This is used by action implementations which use SendProgress instead of Diagnostics
// Returns a concise single-line format suitable for progress messages
func HandleKiotaGraphErrorForAction(ctx context.Context, err error) string {
	if err == nil {
		return "unknown error occurred"
	}

	errorInfo := GraphError(ctx, err)

	tflog.Debug(ctx, "Formatting Graph error for action:", map[string]any{
		"status_code":    errorInfo.StatusCode,
		"error_code":     errorInfo.ErrorCode,
		"error_message":  errorInfo.ErrorMessage,
		"target":         errorInfo.Target,
		"is_odata_error": errorInfo.IsODataError,
		"request_id":     errorInfo.RequestID,
		"correlation_id": errorInfo.CorrelationID,
		"category":       errorInfo.Category,
	})

	var parts []string

	if errorInfo.ErrorMessage != "" {
		parts = append(parts, errorInfo.ErrorMessage)
	}

	if errorInfo.ErrorCode != "" && errorInfo.ErrorCode != "UnknownError" {
		if description, exists := comprehensiveODataErrorCodes[errorInfo.ErrorCode]; exists {
			parts = append(parts, fmt.Sprintf("Code: %s (%s)", errorInfo.ErrorCode, description))
		} else {
			parts = append(parts, fmt.Sprintf("Code: %s", errorInfo.ErrorCode))
		}
	}

	if errorInfo.Target != "" {
		parts = append(parts, fmt.Sprintf("Target: %s", errorInfo.Target))
	}

	if len(parts) > 0 {
		return strings.Join(parts, ". ")
	}

	if errorInfo.StatusCode > 0 {
		errorDesc := getErrorDescription(errorInfo.StatusCode)
		return errorDesc.Summary
	}

	return "API error occurred"
}

// GraphError extracts and analyzes error information from Graph API errors
func GraphError(ctx context.Context, err error) GraphErrorInfo {
	errorInfo := GraphErrorInfo{
		AdditionalData: make(map[string]any),
		InnerErrors:    []InnerErrorInfo{},
		ErrorDetails:   []ErrorDetailInfo{},
	}

	if err == nil {
		return errorInfo
	}

	errorInfo.ErrorMessage = err.Error()

	tflog.Debug(ctx, "Extracting error information", map[string]any{
		"error_type": fmt.Sprintf("%T", err),
		"error":      err.Error(),
	})

	{
		var urlError *url.Error
		var odataError *odataerrors.ODataError
		var graphResponseError interface {
			GetStatusCode() int
			GetErrorEscaped() odataerrors.MainErrorable
		}
		var apiError abstractions.ApiErrorable
		switch {
		case errors.As(err, &urlError):
			extractURLError(ctx, urlError, &errorInfo)
		case errors.As(err, &odataError):
			extractAPIError(ctx, odataError, &errorInfo)
		case errors.As(err, &graphResponseError):
			errorInfo.StatusCode = graphResponseError.GetStatusCode()
			mainError := graphResponseError.GetErrorEscaped()
			extractMainError(ctx, mainError, &errorInfo)
		case errors.As(err, &apiError):
			extractAPIError(ctx, apiError, &errorInfo)
		default:
			errorInfo.StatusCode = 500
			errorInfo.ErrorCode = "UnknownError"
			errorInfo.Category = CategoryUnknown
		}
	}

	errorInfo.Category = categorizeError(&errorInfo)

	logErrorDetails(ctx, &errorInfo)
	return errorInfo
}

// extractURLError handles URL specific errors
func extractURLError(ctx context.Context, urlErr *url.Error, errorInfo *GraphErrorInfo) {
	tflog.Debug(ctx, "URL error detected", map[string]any{
		"url": urlErr.URL,
		"op":  urlErr.Op,
		"err": urlErr.Err.Error(),
	})

	switch {
	case strings.Contains(urlErr.Error(), "context deadline exceeded"):
		errorInfo.StatusCode = 504 // Gateway Timeout
		errorInfo.ErrorCode = "RequestTimeout"
		errorInfo.Category = CategoryNetwork
	case strings.Contains(urlErr.Error(), "connection refused"):
		errorInfo.StatusCode = 503 // Service Unavailable
		errorInfo.ErrorCode = "ConnectionRefused"
		errorInfo.Category = CategoryNetwork
	case strings.Contains(urlErr.Error(), "no such host"):
		errorInfo.StatusCode = 503 // Service Unavailable
		errorInfo.ErrorCode = "HostNotFound"
		errorInfo.Category = CategoryNetwork
	case strings.Contains(urlErr.Error(), "network is unreachable"):
		errorInfo.StatusCode = 503 // Service Unavailable
		errorInfo.ErrorCode = "NetworkUnreachable"
		errorInfo.Category = CategoryNetwork
	case strings.Contains(urlErr.Error(), "certificate"):
		errorInfo.StatusCode = 503 // Service Unavailable
		errorInfo.ErrorCode = "CertificateError"
		errorInfo.Category = CategoryNetwork
	default:
		errorInfo.StatusCode = 400 // Bad Request
		errorInfo.ErrorCode = "URLError"
		errorInfo.Category = CategoryNetwork
	}

	errorInfo.AdditionalData["original_error"] = urlErr.Error()
	errorInfo.AdditionalData["url"] = urlErr.URL
	errorInfo.AdditionalData["operation"] = urlErr.Op

	errorInfo.ErrorMessage = urlErr.Error()
}

// extractAPIError handles Microsoft Graph API specific errors with enhanced extraction
func extractAPIError(
	ctx context.Context,
	apiErr abstractions.ApiErrorable,
	errorInfo *GraphErrorInfo,
) {
	errorInfo.StatusCode = apiErr.GetStatusCode()
	errorInfo.Headers = apiErr.GetResponseHeaders()

	extractHeaders(apiErr, errorInfo)

	switch typedApiErr := apiErr.(type) {
	case *odataerrors.ODataError:
		extractODataError(ctx, typedApiErr, errorInfo)
	case *abstractions.ApiError:
		if typedApiErr.Message != "" {
			errorInfo.ErrorMessage = typedApiErr.Message
		}
		errorInfo.ErrorCode = "ApiError"
	}
}

// extractHeaders extracts comprehensive header information from the API error
func extractHeaders(apiErr abstractions.ApiErrorable, errorInfo *GraphErrorInfo) {
	if headers := apiErr.GetResponseHeaders(); headers != nil {
		for _, key := range headers.ListKeys() {
			values := headers.Get(key)
			if len(values) > 0 {
				// Extract specific headers that are useful for debugging
				switch strings.ToLower(key) {
				case "request-id":
					errorInfo.RequestID = values[0]
				case "client-request-id":
					errorInfo.ClientRequestID = values[0]
				case "ms-correlation-id":
					errorInfo.CorrelationID = values[0]
				case "retry-after":
					errorInfo.RetryAfter = values[0]
				case "x-throttled-reason":
					errorInfo.ThrottledReason = values[0]
				case "date":
					errorInfo.ErrorDate = values[0]
				}
				errorInfo.RequestDetails += fmt.Sprintf("%s: %v\n", key, values)
			}
		}
	}
}

// extractODataError handles OData specific errors with complete extraction
func extractODataError(
	ctx context.Context,
	odataErr *odataerrors.ODataError,
	errorInfo *GraphErrorInfo,
) {
	errorInfo.IsODataError = true

	if mainError := odataErr.GetErrorEscaped(); mainError != nil {

		extractMainError(ctx, mainError, errorInfo)
		extractErrorDetails(ctx, mainError, errorInfo)

		if innerError := mainError.GetInnerError(); innerError != nil {
			extractInnerError(ctx, innerError, errorInfo)
		}
	}
}

// extractMainError extracts comprehensive information from the main error object
func extractMainError(
	ctx context.Context,
	mainError odataerrors.MainErrorable,
	errorInfo *GraphErrorInfo,
) {
	if code := mainError.GetCode(); code != nil && *code != "" {
		errorInfo.ErrorCode = *code
		tflog.Debug(ctx, "Found main error code", map[string]any{
			"code": *code,
		})

		if description, exists := comprehensiveODataErrorCodes[*code]; exists {
			errorInfo.AdditionalData["error_description"] = description
		}
	}

	if message := mainError.GetMessage(); message != nil && *message != "" {
		errorInfo.ErrorMessage = *message
		tflog.Debug(ctx, "Found main error message", map[string]any{
			"message": *message,
		})
	}

	if target := mainError.GetTarget(); target != nil && *target != "" {
		errorInfo.Target = *target
		tflog.Debug(ctx, "Found error target", map[string]any{
			"target": *target,
		})
	}
}

// extractErrorDetails extracts the details array from the main error
func extractErrorDetails(
	ctx context.Context,
	mainError odataerrors.MainErrorable,
	errorInfo *GraphErrorInfo,
) {
	details := mainError.GetDetails()
	if len(details) == 0 {
		return
	}

	tflog.Debug(ctx, "Extracting error details", map[string]any{
		"detail_count": len(details),
	})

	for i, detail := range details {
		detailInfo := ErrorDetailInfo{}

		if code := detail.GetCode(); code != nil && *code != "" {
			detailInfo.Code = *code
		}

		if msg := detail.GetMessage(); msg != nil && *msg != "" {
			detailInfo.Message = *msg
		}

		if target := detail.GetTarget(); target != nil && *target != "" {
			detailInfo.Target = *target
		}

		errorInfo.ErrorDetails = append(errorInfo.ErrorDetails, detailInfo)

		tflog.Debug(ctx, "Extracted error detail", map[string]any{
			"index":  i,
			"code":   detailInfo.Code,
			"target": detailInfo.Target,
		})
	}
}

// extractInnerError extracts the inner error (only one level in current SDK)
func extractInnerError(
	ctx context.Context,
	innerError odataerrors.InnerErrorable,
	errorInfo *GraphErrorInfo,
) {
	if innerError == nil {
		return
	}

	innerInfo := InnerErrorInfo{}

	if reqID := innerError.GetRequestId(); reqID != nil && *reqID != "" {
		innerInfo.RequestID = *reqID
		if errorInfo.RequestID == "" {
			errorInfo.RequestID = *reqID
		}
	}

	if clientReqID := innerError.GetClientRequestId(); clientReqID != nil && *clientReqID != "" {
		innerInfo.ClientReqID = *clientReqID
		if errorInfo.ClientRequestID == "" {
			errorInfo.ClientRequestID = *clientReqID
		}
	}

	if date := innerError.GetDate(); date != nil {
		innerInfo.Date = date.String()
		if errorInfo.ErrorDate == "" {
			errorInfo.ErrorDate = date.String()
		}
	}

	if odataType := innerError.GetOdataType(); odataType != nil && *odataType != "" {
		innerInfo.ODataType = *odataType
		// Inner error OData type often contains more specific error codes
		if errorInfo.ErrorCode == "" {
			errorInfo.ErrorCode = *odataType
		}
	}

	errorInfo.InnerErrors = append(errorInfo.InnerErrors, innerInfo)

	tflog.Debug(ctx, "Extracted inner error", map[string]any{
		"odata_type": innerInfo.ODataType,
		"request_id": innerInfo.RequestID,
	})
}

// constructDetailedErrorMessage creates a comprehensive error message
func constructDetailedErrorMessage(standardDetail string, errorInfo *GraphErrorInfo) string {
	var parts []string
	parts = append(parts, standardDetail)

	if errorInfo.ErrorMessage != "" {
		parts = append(parts, fmt.Sprintf("Error: %s", errorInfo.ErrorMessage))
	}

	if errorInfo.ErrorCode != "" {
		parts = append(parts, fmt.Sprintf("Code: %s", errorInfo.ErrorCode))
		if description, exists := comprehensiveODataErrorCodes[errorInfo.ErrorCode]; exists {
			parts = append(parts, fmt.Sprintf("Description: %s", description))
		}
	}

	if errorInfo.Target != "" {
		parts = append(parts, fmt.Sprintf("Target: %s", errorInfo.Target))
	}

	if len(errorInfo.ErrorDetails) > 0 {
		var detailParts []string
		for _, detail := range errorInfo.ErrorDetails {
			detailStr := ""
			if detail.Code != "" {
				detailStr += fmt.Sprintf("Code: %s", detail.Code)
			}
			if detail.Message != "" {
				if detailStr != "" {
					detailStr += " - "
				}
				detailStr += detail.Message
			}
			if detail.Target != "" {
				if detailStr != "" {
					detailStr += " "
				}
				detailStr += fmt.Sprintf("(Target: %s)", detail.Target)
			}
			if detailStr != "" {
				detailParts = append(detailParts, detailStr)
			}
		}
		if len(detailParts) > 0 {
			parts = append(parts, fmt.Sprintf("Details: %s", strings.Join(detailParts, "; ")))
		}
	}

	if len(errorInfo.InnerErrors) > 0 {
		innerParts := make([]string, 0, len(errorInfo.InnerErrors))
		for i, inner := range errorInfo.InnerErrors {
			innerStr := fmt.Sprintf("Level %d", i+1)
			if inner.ODataType != "" {
				innerStr += fmt.Sprintf(" - Type: %s", inner.ODataType)
			}
			if inner.Code != "" {
				innerStr += fmt.Sprintf(" - Code: %s", inner.Code)
			}
			if inner.Message != "" {
				innerStr += fmt.Sprintf(" - Message: %s", inner.Message)
			}
			innerParts = append(innerParts, innerStr)
		}
		if len(innerParts) > 0 {
			parts = append(parts, fmt.Sprintf("Inner Errors: %s", strings.Join(innerParts, "; ")))
		}
	}

	var trackingParts []string
	if errorInfo.RequestID != "" {
		trackingParts = append(trackingParts, fmt.Sprintf("Request ID: %s", errorInfo.RequestID))
	}
	if errorInfo.ClientRequestID != "" {
		trackingParts = append(
			trackingParts,
			fmt.Sprintf("Client Request ID: %s", errorInfo.ClientRequestID),
		)
	}
	if errorInfo.CorrelationID != "" {
		trackingParts = append(
			trackingParts,
			fmt.Sprintf("Correlation ID: %s", errorInfo.CorrelationID),
		)
	}
	if errorInfo.ErrorDate != "" {
		trackingParts = append(trackingParts, fmt.Sprintf("Date: %s", errorInfo.ErrorDate))
	}
	if len(trackingParts) > 0 {
		parts = append(parts, fmt.Sprintf("Tracking: %s", strings.Join(trackingParts, ", ")))
	}

	if errorInfo.Category != "" {
		parts = append(parts, fmt.Sprintf("Category: %s", errorInfo.Category))
	}

	return strings.Join(parts, "\n")
}

// categorizeError categorizes errors based on status code and error information
func categorizeError(errorInfo *GraphErrorInfo) ErrorCategory {
	switch errorInfo.StatusCode {
	case 401:
		return CategoryAuthentication
	case 403:
		return CategoryAuthorization
	case 400, 422:
		return CategoryValidation
	case 429:
		return CategoryThrottling
	case 503, 502, 504, 500:
		return CategoryService
	case 0: // Network errors typically have status code 0
		return CategoryNetwork
	default:
		// Check error codes for more specific categorization
		if errorInfo.ErrorCode != "" {
			switch {
			case strings.Contains(strings.ToLower(errorInfo.ErrorCode), "auth"):
				return CategoryAuthentication
			case strings.Contains(strings.ToLower(errorInfo.ErrorCode), "forbidden"):
				return CategoryAuthorization
			case strings.Contains(strings.ToLower(errorInfo.ErrorCode), "throttle"):
				return CategoryThrottling
			case strings.Contains(strings.ToLower(errorInfo.ErrorCode), "network"):
				return CategoryNetwork
			}
		}
		return CategoryService
	}
}

// recordErrorMetrics records error metrics for observability
func recordErrorMetrics(ctx context.Context, errorInfo *GraphErrorInfo, operation string) {
	// Log structured metrics that can be consumed by monitoring systems
	tflog.Info(ctx, "Graph API Error Metrics", map[string]any{
		"metric_type":       "graph_api_error",
		"status_code":       errorInfo.StatusCode,
		"error_code":        errorInfo.ErrorCode,
		"operation":         operation,
		"category":          errorInfo.Category,
		"is_odata_error":    errorInfo.IsODataError,
		"has_inner_errors":  len(errorInfo.InnerErrors) > 0,
		"has_error_details": len(errorInfo.ErrorDetails) > 0,
		"request_id":        errorInfo.RequestID,
		"timestamp":         time.Now().Unix(),
	})
}

// handlePermissionError processes permission-related errors with enhanced details
func handlePermissionError(
	ctx context.Context,
	errorInfo GraphErrorInfo,
	resp any,
	operation string,
	requiredPermissions []string,
) {
	var permissionMsg string

	switch {
	case len(requiredPermissions) == 1:
		permissionMsg = fmt.Sprintf(
			"%s operation requires permission: %s",
			operation,
			requiredPermissions[0],
		)
	case len(requiredPermissions) > 1:
		permissionMsg = fmt.Sprintf(
			"%s operation requires one or more of the following permissions: %s",
			operation,
			strings.Join(requiredPermissions, ", "),
		)
	default:
		permissionMsg = fmt.Sprintf(
			"%s operation: No specific permissions defined. Please check Microsoft documentation.",
			operation,
		)
	}

	errorDesc := getErrorDescription(errorInfo.StatusCode)
	detail := fmt.Sprintf("%s\n%s\n%s",
		errorDesc.Detail,
		permissionMsg,
		constructDetailedErrorMessage("Graph API Error Details:", &errorInfo))

	addErrorToDiagnostics(ctx, resp, errorDesc.Summary, detail)
}

// handleRateLimitError processes rate limit errors with enhanced information
func handleRateLimitError(ctx context.Context, errorInfo GraphErrorInfo, resp any) {
	tflog.Warn(ctx, "Rate limit exceeded", map[string]any{
		"retry_after":      errorInfo.RetryAfter,
		"throttled_reason": errorInfo.ThrottledReason,
		"request_id":       errorInfo.RequestID,
		"details":          errorInfo.ErrorMessage,
	})

	errorDesc := getErrorDescription(429)
	detail := constructDetailedErrorMessage(errorDesc.Detail, &errorInfo)

	if errorInfo.RetryAfter != "" {
		detail += fmt.Sprintf("\nRetry-After: %s seconds", errorInfo.RetryAfter)
	}

	if errorInfo.ThrottledReason != "" {
		detail += fmt.Sprintf("\nThrottled Reason: %s", errorInfo.ThrottledReason)
	}

	addErrorToDiagnostics(ctx, resp, errorDesc.Summary, detail)
}

// handleServiceUnavailableError processes 503 Service Unavailable errors
func handleServiceUnavailableError(ctx context.Context, errorInfo GraphErrorInfo, resp any) {
	retryAfter := errorInfo.RetryAfter
	if retryAfter == "" {
		retryAfter = "unspecified"
	}

	tflog.Warn(ctx, "Service temporarily unavailable", map[string]any{
		"retry_after": retryAfter,
		"request_id":  errorInfo.RequestID,
		"details":     errorInfo.ErrorMessage,
	})

	errorDesc := getErrorDescription(503)
	detail := fmt.Sprintf(
		"%s\nThe service may be experiencing high load or undergoing maintenance.\nRetry-After: %s\n%s",
		errorDesc.Detail,
		retryAfter,
		constructDetailedErrorMessage("Additional Details:", &errorInfo),
	)

	addErrorToDiagnostics(ctx, resp, errorDesc.Summary, detail)
}

// getErrorDescription returns a standard error description based on the HTTP status code
func getErrorDescription(statusCode int) ErrorDescription {
	if desc, ok := standardErrorDescriptions[statusCode]; ok {
		return desc
	}
	return ErrorDescription{
		Summary: fmt.Sprintf("HTTP Error %d", statusCode),
		Detail:  "An unexpected error occurred. Please check the request and try again.",
	}
}

// addErrorToDiagnostics adds an error to the response diagnostics
func addErrorToDiagnostics(ctx context.Context, resp any, summary, detail string) {
	switch r := resp.(type) {
	case *resource.CreateResponse:
		r.Diagnostics.AddError(summary, detail)
	case *resource.ReadResponse:
		r.Diagnostics.AddError(summary, detail)
	case *resource.UpdateResponse:
		r.Diagnostics.AddError(summary, detail)
	case *resource.DeleteResponse:
		r.Diagnostics.AddError(summary, detail)
	case *datasource.ReadResponse:
		r.Diagnostics.AddError(summary, detail)
	default:
		tflog.Error(ctx, "Unknown response type in addErrorToDiagnostics", map[string]any{
			"response_type": fmt.Sprintf("%T", resp),
			"summary":       summary,
			"detail":        detail,
		})
	}
}

// removeResourceFromState removes a resource from the state
func removeResourceFromState(ctx context.Context, resp any) {
	switch r := resp.(type) {
	case *resource.ReadResponse:
		r.State.RemoveResource(ctx)
		tflog.Debug(ctx, "Resource removed from state due to not found error")
	default:
		tflog.Error(ctx, "Cannot remove resource from state for this response type", map[string]any{
			"response_type": fmt.Sprintf("%T", resp),
		})
	}
}

// logErrorDetails logs comprehensive error details for debugging
func logErrorDetails(ctx context.Context, errorInfo *GraphErrorInfo) {
	details := map[string]any{
		"status_code":       errorInfo.StatusCode,
		"error_code":        errorInfo.ErrorCode,
		"is_odata_error":    errorInfo.IsODataError,
		"error_message":     errorInfo.ErrorMessage,
		"target":            errorInfo.Target,
		"category":          errorInfo.Category,
		"request_id":        errorInfo.RequestID,
		"client_request_id": errorInfo.ClientRequestID,
		"correlation_id":    errorInfo.CorrelationID,
		"error_date":        errorInfo.ErrorDate,
		"retry_after":       errorInfo.RetryAfter,
		"throttled_reason":  errorInfo.ThrottledReason,
	}

	if len(errorInfo.AdditionalData) > 0 {
		details["additional_data"] = errorInfo.AdditionalData
	}

	if len(errorInfo.InnerErrors) > 0 {
		details["inner_errors_count"] = len(errorInfo.InnerErrors)
		details["inner_errors"] = errorInfo.InnerErrors
	}

	if len(errorInfo.ErrorDetails) > 0 {
		details["error_details_count"] = len(errorInfo.ErrorDetails)
		details["error_details"] = errorInfo.ErrorDetails
	}

	if errorInfo.RequestDetails != "" {
		details["request_headers"] = strings.Split(errorInfo.RequestDetails, "\n")
	}

	tflog.Debug(ctx, "Comprehensive error details", details)
}
