package graphBetaConditionalAccessPolicy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// constructResource maps the Terraform schema to the SDK model
func constructResource(ctx context.Context, data *ConditionalAccessPolicyResourceModel) (*models.ConditionalAccessPolicy, error) {
	requestBody := models.NewConditionalAccessPolicy()

	displayName := data.DisplayName.ValueString()
	requestBody.SetDisplayName(&displayName)

	if !data.Description.IsNull() {
		description := data.Description.ValueString()
		requestBody.SetDescription(&description)
	}

	if !data.State.IsNull() {
		stateStr := data.State.ValueString()
		stateAny, err := models.ParseConditionalAccessPolicyState(stateStr)
		if err != nil {
			return nil, fmt.Errorf("invalid state: %s", err)
		}
		if stateAny != nil {
			state, ok := stateAny.(*models.ConditionalAccessPolicyState)
			if !ok {
				return nil, fmt.Errorf("unexpected type for state: %T", stateAny)
			}
			requestBody.SetState(state)
		}
	}

	if data.Conditions != nil {
		conditions, err := constructConditions(data.Conditions)
		if err != nil {
			return nil, fmt.Errorf("error constructing conditions: %s", err)
		}
		requestBody.SetConditions(conditions)
	}

	if data.GrantControls != nil {
		grantControls, err := constructGrantControls(data.GrantControls)
		if err != nil {
			return nil, fmt.Errorf("error constructing grant controls: %s", err)
		}
		requestBody.SetGrantControls(grantControls)
	}

	if data.SessionControls != nil {
		sessionControls, err := constructSessionControls(data.SessionControls)
		if err != nil {
			return nil, fmt.Errorf("error constructing session controls: %s", err)
		}
		requestBody.SetSessionControls(sessionControls)
	}

	requestBodyJSON, err := json.MarshalIndent(map[string]interface{}{
		"displayName":     requestBody.GetDisplayName(),
		"description":     requestBody.GetDescription(),
		"state":           requestBody.GetState(),
		"conditions":      requestBody.GetConditions(),
		"grantControls":   requestBody.GetGrantControls(),
		"sessionControls": requestBody.GetSessionControls(),
	}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body to JSON: %s", err)
	}

	tflog.Debug(ctx, "Constructed conditional access policy resource:\n"+string(requestBodyJSON))

	return requestBody, nil
}

// Helper functions to construct nested objects
func constructConditions(data *ConditionalAccessConditionsModel) (*models.ConditionalAccessConditionSet, error) {
	if data == nil {
		return nil, nil
	}

	conditions := models.NewConditionalAccessConditionSet()

	// Applications
	if data.Applications != nil {
		applications, err := constructApplications(data.Applications)
		if err != nil {
			return nil, fmt.Errorf("error constructing applications: %v", err)
		}
		conditions.SetApplications(applications)
	}

	// Authentication Flows
	if data.AuthenticationFlows != nil {
		authFlows, err := constructAuthenticationFlows(data.AuthenticationFlows)
		if err != nil {
			return nil, fmt.Errorf("error constructing authentication flows: %v", err)
		}
		conditions.SetAuthenticationFlows(authFlows)
	}

	// Client Applications
	if data.ClientApplications != nil {
		clientApps, err := constructClientApplications(data.ClientApplications)
		if err != nil {
			return nil, fmt.Errorf("error constructing client applications: %v", err)
		}
		conditions.SetClientApplications(clientApps)
	}

	// Client App Types
	if len(data.ClientAppTypes) > 0 {
		clientAppTypes := make([]models.ConditionalAccessClientApp, len(data.ClientAppTypes))
		for i, appType := range data.ClientAppTypes {
			clientAppType, err := models.ParseConditionalAccessClientApp(appType.ValueString())
			if err != nil {
				return nil, fmt.Errorf("error parsing client app type: %v", err)
			}
			clientAppTypes[i] = *clientAppType
		}
		conditions.SetClientAppTypes(clientAppTypes)
	}

	// Devices
	if data.Devices != nil {
		devices, err := constructDevices(data.Devices)
		if err != nil {
			return nil, fmt.Errorf("error constructing devices: %v", err)
		}
		conditions.SetDevices(devices)
	}

	// Device States (deprecated)
	if data.DeviceStates != nil {
		deviceStates, err := constructDeviceStates(data.DeviceStates)
		if err != nil {
			return nil, fmt.Errorf("error constructing device states: %v", err)
		}
		conditions.SetDeviceStates(deviceStates)
	}

	// Insider Risk Levels
	if !data.InsiderRiskLevels.IsNull() {
		insiderRiskLevel, err := models.ParseConditionalAccessInsiderRiskLevels(data.InsiderRiskLevels.ValueString())
		if err != nil {
			return nil, fmt.Errorf("error parsing insider risk level: %v", err)
		}
		conditions.SetInsiderRiskLevels(insiderRiskLevel)
	}

	// Locations
	if data.Locations != nil {
		locations, err := constructLocations(data.Locations)
		if err != nil {
			return nil, fmt.Errorf("error constructing locations: %v", err)
		}
		conditions.SetLocations(locations)
	}

	// Platforms
	if data.Platforms != nil {
		platforms, err := constructPlatforms(data.Platforms)
		if err != nil {
			return nil, fmt.Errorf("error constructing platforms: %v", err)
		}
		conditions.SetPlatforms(platforms)
	}

	// Service Principal Risk Levels
	if len(data.ServicePrincipalRiskLevels) > 0 {
		riskLevels := make([]models.RiskLevel, len(data.ServicePrincipalRiskLevels))
		for i, level := range data.ServicePrincipalRiskLevels {
			riskLevel, err := models.ParseRiskLevel(level.ValueString())
			if err != nil {
				return nil, fmt.Errorf("error parsing service principal risk level: %v", err)
			}
			riskLevels[i] = *riskLevel
		}
		conditions.SetServicePrincipalRiskLevels(riskLevels)
	}

	// Sign-in Risk Levels
	if len(data.SignInRiskLevels) > 0 {
		riskLevels := make([]models.RiskLevel, len(data.SignInRiskLevels))
		for i, level := range data.SignInRiskLevels {
			riskLevel, err := models.ParseRiskLevel(level.ValueString())
			if err != nil {
				return nil, fmt.Errorf("error parsing sign-in risk level: %v", err)
			}
			riskLevels[i] = *riskLevel
		}
		conditions.SetSignInRiskLevels(riskLevels)
	}

	// User Risk Levels
	if len(data.UserRiskLevels) > 0 {
		riskLevels := make([]models.RiskLevel, len(data.UserRiskLevels))
		for i, level := range data.UserRiskLevels {
			riskLevel, err := models.ParseRiskLevel(level.ValueString())
			if err != nil {
				return nil, fmt.Errorf("error parsing user risk level: %v", err)
			}
			riskLevels[i] = *riskLevel
		}
		conditions.SetUserRiskLevels(riskLevels)
	}

	// Users
	if data.Users != nil {
		users, err := constructUsers(data.Users)
		if err != nil {
			return nil, fmt.Errorf("error constructing users: %v", err)
		}
		conditions.SetUsers(users)
	}

	return conditions, nil
}

func constructApplications(data *ConditionalAccessApplicationsModel) (models.ConditionalAccessApplicationsable, error) {
	if data == nil {
		return nil, nil
	}

	applications := models.NewConditionalAccessApplications()

	if len(data.IncludeApplications) > 0 {
		includeApps := make([]string, len(data.IncludeApplications))
		for i, app := range data.IncludeApplications {
			includeApps[i] = app.ValueString()
		}
		applications.SetIncludeApplications(includeApps)
	}

	if len(data.ExcludeApplications) > 0 {
		excludeApps := make([]string, len(data.ExcludeApplications))
		for i, app := range data.ExcludeApplications {
			excludeApps[i] = app.ValueString()
		}
		applications.SetExcludeApplications(excludeApps)
	}

	if len(data.IncludeUserActions) > 0 {
		userActions := make([]string, len(data.IncludeUserActions))
		for i, action := range data.IncludeUserActions {
			userActions[i] = action.ValueString()
		}
		applications.SetIncludeUserActions(userActions)
	}

	return applications, nil
}

func constructUsers(data *ConditionalAccessUsersModel) (models.ConditionalAccessUsersable, error) {
	if data == nil {
		return nil, nil
	}

	users := models.NewConditionalAccessUsers()

	if len(data.IncludeUsers) > 0 {
		includeUsers := make([]string, len(data.IncludeUsers))
		for i, user := range data.IncludeUsers {
			includeUsers[i] = user.ValueString()
		}
		users.SetIncludeUsers(includeUsers)
	}

	if len(data.ExcludeUsers) > 0 {
		excludeUsers := make([]string, len(data.ExcludeUsers))
		for i, user := range data.ExcludeUsers {
			excludeUsers[i] = user.ValueString()
		}
		users.SetExcludeUsers(excludeUsers)
	}

	if len(data.IncludeGroups) > 0 {
		includeGroups := make([]string, len(data.IncludeGroups))
		for i, group := range data.IncludeGroups {
			includeGroups[i] = group.ValueString()
		}
		users.SetIncludeGroups(includeGroups)
	}

	if len(data.ExcludeGroups) > 0 {
		excludeGroups := make([]string, len(data.ExcludeGroups))
		for i, group := range data.ExcludeGroups {
			excludeGroups[i] = group.ValueString()
		}
		users.SetExcludeGroups(excludeGroups)
	}

	if len(data.IncludeRoles) > 0 {
		includeRoles := make([]string, len(data.IncludeRoles))
		for i, role := range data.IncludeRoles {
			includeRoles[i] = role.ValueString()
		}
		users.SetIncludeRoles(includeRoles)
	}

	if len(data.ExcludeRoles) > 0 {
		excludeRoles := make([]string, len(data.ExcludeRoles))
		for i, role := range data.ExcludeRoles {
			excludeRoles[i] = role.ValueString()
		}
		users.SetExcludeRoles(excludeRoles)
	}

	if data.IncludeGuestsOrExternalUsers != nil {
		guestsOrExternalUsers, err := constructGuestsOrExternalUsers(data.IncludeGuestsOrExternalUsers)
		if err != nil {
			return nil, fmt.Errorf("error constructing include guests or external users: %v", err)
		}
		users.SetIncludeGuestsOrExternalUsers(guestsOrExternalUsers)
	}

	if data.ExcludeGuestsOrExternalUsers != nil {
		guestsOrExternalUsers, err := constructGuestsOrExternalUsers(data.ExcludeGuestsOrExternalUsers)
		if err != nil {
			return nil, fmt.Errorf("error constructing exclude guests or external users: %v", err)
		}
		users.SetExcludeGuestsOrExternalUsers(guestsOrExternalUsers)
	}

	return users, nil
}

func constructGuestsOrExternalUsers(data *ConditionalAccessGuestsOrExternalUsersModel) (models.ConditionalAccessGuestsOrExternalUsersable, error) {
	if data == nil {
		return nil, nil
	}

	guestsOrExternalUsers := models.NewConditionalAccessGuestsOrExternalUsers()

	if !data.GuestOrExternalUserTypes.IsNull() {
		guestOrExternalUserTypes, err := models.ParseConditionalAccessGuestOrExternalUserTypes(data.GuestOrExternalUserTypes.ValueString())
		if err != nil {
			return nil, fmt.Errorf("error parsing guest or external user types: %v", err)
		}
		guestsOrExternalUsers.SetGuestOrExternalUserTypes(guestOrExternalUserTypes)
	}

	if data.ExternalTenants != nil {
		externalTenants, err := constructConditionalAccessExternalTenants(data.ExternalTenants)
		if err != nil {
			return nil, fmt.Errorf("error constructing external tenants: %v", err)
		}
		guestsOrExternalUsers.SetExternalTenants(externalTenants)
	}

	return guestsOrExternalUsers, nil
}

func constructConditionalAccessExternalTenants(data *ConditionalAccessExternalTenantsModel) (models.ConditionalAccessExternalTenantsable, error) {
	if data == nil {
		return nil, nil
	}

	externalTenants := models.NewConditionalAccessExternalTenants()

	// The SDK doesn't have a MembershipKind property, so we'll skip that

	if len(data.TenantIds) > 0 {
		tenantIds := make([]string, len(data.TenantIds))
		for i, id := range data.TenantIds {
			tenantIds[i] = id.ValueString()
		}
		externalTenants.SetTenantIds(tenantIds)
	}

	return externalTenants, nil
}

func constructClientApplications(data *ConditionalAccessClientApplicationsModel) (models.ConditionalAccessClientApplicationsable, error) {
	if data == nil {
		return nil, nil
	}

	clientApps := models.NewConditionalAccessClientApplications()

	if len(data.IncludeServicePrincipals) > 0 {
		includeServicePrincipals := make([]string, len(data.IncludeServicePrincipals))
		for i, sp := range data.IncludeServicePrincipals {
			includeServicePrincipals[i] = sp.ValueString()
		}
		clientApps.SetIncludeServicePrincipals(includeServicePrincipals)
	}

	if len(data.ExcludeServicePrincipals) > 0 {
		excludeServicePrincipals := make([]string, len(data.ExcludeServicePrincipals))
		for i, sp := range data.ExcludeServicePrincipals {
			excludeServicePrincipals[i] = sp.ValueString()
		}
		clientApps.SetExcludeServicePrincipals(excludeServicePrincipals)
	}

	return clientApps, nil
}

// Implement similar functions for other nested objects:
// func constructDevices(data *ConditionalAccessDevicesModel) (models.ConditionalAccessDevicesable, error)
// func constructDeviceStates(data *ConditionalAccessDeviceStatesModel) (models.ConditionalAccessDeviceStatesable, error)
// func constructLocations(data *ConditionalAccessLocationsModel) (models.ConditionalAccessLocationsable, error)
// func constructPlatforms(data *ConditionalAccessPlatformsModel) (models.ConditionalAccessPlatformsable, error)
// func constructAuthenticationFlows(data *ConditionalAccessAuthenticationFlowsModel) (models.ConditionalAccessAuthenticationFlowsable, error)

func constructGrantControls(data *ConditionalAccessGrantControlsModel) (*models.ConditionalAccessGrantControls, error) {
	if data == nil {
		return nil, nil
	}

	grantControls := models.NewConditionalAccessGrantControls()

	if !data.Operator.IsNull() {
		operator, err := models.ParseConditionalAccessGrantControlOperator(data.Operator.ValueString())
		if err != nil {
			return nil, fmt.Errorf("error parsing grant control operator: %v", err)
		}
		grantControls.SetOperator(operator)
	}

	if len(data.BuiltInControls) > 0 {
		builtInControls := make([]models.ConditionalAccessGrantControl, len(data.BuiltInControls))
		for i, control := range data.BuiltInControls {
			builtInControl, err := models.ParseConditionalAccessGrantControl(control.ValueString())
			if err != nil {
				return nil, fmt.Errorf("error parsing built-in control: %v", err)
			}
			builtInControls[i] = *builtInControl
		}
		grantControls.SetBuiltInControls(builtInControls)
	}

	if len(data.CustomAuthenticationFactors) > 0 {
		customFactors := make([]string, len(data.CustomAuthenticationFactors))
		for i, factor := range data.CustomAuthenticationFactors {
			customFactors[i] = factor.ValueString()
		}
		grantControls.SetCustomAuthenticationFactors(customFactors)
	}

	if len(data.TermsOfUse) > 0 {
		termsOfUse := make([]string, len(data.TermsOfUse))
		for i, term := range data.TermsOfUse {
			termsOfUse[i] = term.ValueString()
		}
		grantControls.SetTermsOfUse(termsOfUse)
	}

	if data.AuthenticationStrength != nil {
		authStrength, err := constructAuthenticationStrength(data.AuthenticationStrength)
		if err != nil {
			return nil, fmt.Errorf("error constructing authentication strength: %v", err)
		}
		grantControls.SetAuthenticationStrength(authStrength)
	}

	return grantControls, nil
}

func constructAuthenticationStrength(data *AuthenticationStrengthPolicyModel) (*models.AuthenticationStrengthPolicy, error) {
	if data == nil {
		return nil, nil
	}

	authStrength := models.NewAuthenticationStrengthPolicy()

	if !data.DisplayName.IsNull() {
		displayName := data.DisplayName.ValueString()
		authStrength.SetDisplayName(&displayName)
	}

	if !data.Description.IsNull() {
		description := data.Description.ValueString()
		authStrength.SetDescription(&description)
	}

	if !data.PolicyType.IsNull() {
		policyType := data.PolicyType.ValueString()
		authStrength.SetPolicyType(&policyType)
	}

	if !data.RequirementsSatisfied.IsNull() {
		requirementsSatisfied := data.RequirementsSatisfied.ValueString()
		authStrength.SetRequirementsSatisfied(&requirementsSatisfied)
	}

	if len(data.AllowedCombinations) > 0 {
		allowedCombinations := make([]string, len(data.AllowedCombinations))
		for i, combination := range data.AllowedCombinations {
			allowedCombinations[i] = combination.ValueString()
		}
		authStrength.SetAllowedCombinations(allowedCombinations)
	}

	return authStrength, nil
}

func constructSessionControls(data *ConditionalAccessSessionControlsModel) (models.ConditionalAccessSessionControlsable, error) {
	if data == nil {
		return nil, nil
	}

	sessionControls := models.NewConditionalAccessSessionControls()

	if data.ApplicationEnforcedRestrictions != nil {
		appRestrictions := models.NewApplicationEnforcedRestrictionsSessionControl()
		isEnabled := data.ApplicationEnforcedRestrictions.IsEnabled.ValueBool()
		appRestrictions.SetIsEnabled(&isEnabled)
		sessionControls.SetApplicationEnforcedRestrictions(appRestrictions)
	}

	if data.ApplicationEnforcedRestrictions != nil {
    appRestrictions := models.NewApplicationEnforcedRestrictionsSessionControl()
    isEnabled := data.ApplicationEnforcedRestrictions.IsEnabled.ValueBool()
    appRestrictions.SetIsEnabled(&isEnabled)
    sessionControls.SetApplicationEnforcedRestrictions(appRestrictions)
}

if data.ContinuousAccessEvaluation != nil {
	continuousAccessEvaluation := models.NewContinuousAccessEvaluationSessionControl()

	if !data.ContinuousAccessEvaluation.Mode.IsNull() {
			mode, err := models.ParseContinuousAccessEvaluationMode(data.ContinuousAccessEvaluation.Mode.ValueString())
			if err != nil {
					return nil, fmt.Errorf("error parsing continuous access evaluation mode: %v", err)
			}
			continuousAccessEvaluation.SetMode(mode)
	}

	sessionControls.SetContinuousAccessEvaluation(continuousAccessEvaluation)
}

if data.PersistentBrowser != nil {
	persistentBrowser := models.NewPersistentBrowserSessionControl()
	
	// SetIsEnabled is inherited from ConditionalAccessSessionControl
	isEnabled := data.PersistentBrowser.IsEnabled.ValueBool()
	persistentBrowser.SetIsEnabled(&isEnabled)
	
	if !data.PersistentBrowser.Mode.IsNull() {
			mode, err := models.ParsePersistentBrowserSessionMode(data.PersistentBrowser.Mode.ValueString())
			if err != nil {
					return nil, fmt.Errorf("error parsing persistent browser session mode: %v", err)
			}
			persistentBrowser.SetMode(mode)
	}
	
	sessionControls.SetPersistentBrowser(persistentBrowser)
}

if data.SignInFrequency != nil {
	signInFrequency := models.NewSignInFrequencySessionControl()
	
	// SetIsEnabled is inherited from ConditionalAccessSessionControl
	isEnabled := data.SignInFrequency.IsEnabled.ValueBool()
	signInFrequency.SetIsEnabled(&isEnabled)
	
	if !data.SignInFrequency.Type.IsNull() {
			freqType, err := models.ParseSigninFrequencyType(data.SignInFrequency.Type.ValueString())
			if err != nil {
					return nil, fmt.Errorf("error parsing sign-in frequency type: %v", err)
			}
			signInFrequency.SetTypeEscaped(freqType)
	}
	
	if !data.SignInFrequency.Value.IsNull() {
			value := data.SignInFrequency.Value.ValueInt32()
			signInFrequency.SetValue(&value)
	}
	
	if !data.SignInFrequency.FrequencyInterval.IsNull() {
			freqInterval, err := models.ParseSignInFrequencyInterval(data.SignInFrequency.FrequencyInterval.ValueString())
			if err != nil {
					return nil, fmt.Errorf("error parsing sign-in frequency interval: %v", err)
			}
			signInFrequency.SetFrequencyInterval(freqInterval)
	}
	
	if !data.SignInFrequency.AuthenticationType.IsNull() {
			authType, err := models.ParseSignInFrequencyAuthenticationType(data.SignInFrequency.AuthenticationType.ValueString())
			if err != nil {
					return nil, fmt.Errorf("error parsing sign-in frequency authentication type: %v", err)
			}
			signInFrequency.SetAuthenticationType(authType)
	}
	
	sessionControls.SetSignInFrequency(signInFrequency)
}

	if data.SecureSignInSession != nil {
		secureSignInSession := models.NewSecureSignInSessionControl()
		secureSignInSession.SetIsEnabled(data.SecureSignInSession.IsEnabled.ValueBool())
		sessionControls.SetSecureSignInSession(secureSignInSession)
	}

	sessionControls.SetDisableResilienceDefaults(data.DisableResilienceDefaults.ValueBool())

	return sessionControls, nil
}
