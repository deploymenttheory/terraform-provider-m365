package normalize

// odataTypeKey is the OData discriminator. Microsoft Graph requires it on write wherever a
// collection's element type is abstract, but echoes it back only in those same positions — where the
// element type is already concrete it is dropped from the response as redundant. ProjectOntoShape
// therefore treats it as the one key worth preserving from the shape when the response omits it.
const odataTypeKey = "@odata.type"

// ProjectOntoShape returns the subset of response whose key structure is present in shape.
//
// Microsoft Graph returns every property of a resource on read, including ones that were
// never set — a device configuration that declares a single property comes back with ~190
// keys, the remainder as false, null or []. Storing that whole response in a *_json attribute
// would diff against a small configuration on every plan. Projecting the response onto the
// shape of the prior configuration keeps only what the operator declared, so the diff reflects
// real changes.
//
// A consequence worth being aware of: because Graph reports an unset boolean as false, there is
// no way to distinguish "declared false" from "never declared". The shape is the only record of
// which properties are managed, which is why removing a property from a configuration stops
// tracking it rather than resetting it on the server.
//
// Rules:
//   - both objects: keys of shape that also exist in response, recursively. Keys only in shape
//     (the server dropped them) are omitted so the diff legitimately surfaces; keys only in
//     response (server defaults) are omitted, which is the point of the function. A key whose
//     response value is null is still present and keeps its null — that is the server
//     reporting the property as explicitly unset. The sole exception is "@odata.type", which is
//     carried over from the shape when the response omits it: Graph accepts the discriminator
//     everywhere it is required but echoes it back only where it is load-bearing, so its absence
//     is normalization rather than drift.
//   - both arrays: index-wise recursion over the overlapping prefix. The result has
//     len(response) elements; those beyond len(shape) are taken verbatim. Arrays are not
//     truncated to len(shape) because for declarations such as homeScreenPages the array
//     itself is the operator's intent, and dropping elements would hide real remote state.
//   - type mismatch, or a scalar/null node in shape: the response value verbatim. Projection
//     prunes breadth, never values.
//
// Neither argument is mutated.
func ProjectOntoShape(shape, response any) any {
	switch shapeV := shape.(type) {
	case map[string]any:
		responseV, ok := response.(map[string]any)
		if !ok {
			return response
		}

		projected := make(map[string]any, len(shapeV))
		for k, shapeChild := range shapeV {
			// Two-value form: a key holding a null value is present and must be kept.
			if responseChild, exists := responseV[k]; exists {
				projected[k] = ProjectOntoShape(shapeChild, responseChild)
				continue
			}

			// Deliberate exception to "keys only in shape are omitted": a discriminator the response
			// drops is normalization, not drift.
			//
			// Graph echoes @odata.type only where it is load-bearing — where the enclosing
			// collection's element type is abstract and the discriminator is needed to resolve which
			// concrete type an element is. Where the element type is already concrete the value is
			// redundant and Graph omits it, even though it accepted it on write.
			//
			// Observed in an iosDeviceFeaturesConfiguration home screen tree:
			//   icons          []iosHomeScreenItem     (abstract) -> @odata.type echoed
			//   homeScreenPages, pages, apps           (concrete) -> @odata.type dropped
			//
			// Omitting these from state would make every apply fail with "Provider produced
			// inconsistent result after apply", since the operator must supply them for Graph to
			// accept the request in the first place. Carrying the configured value forward keeps
			// state consistent with configuration without inventing anything.
			if k == odataTypeKey {
				projected[k] = shapeChild
			}
		}
		return projected

	case []any:
		responseV, ok := response.([]any)
		if !ok {
			return response
		}

		projected := make([]any, len(responseV))
		for i, responseChild := range responseV {
			if i < len(shapeV) {
				projected[i] = ProjectOntoShape(shapeV[i], responseChild)
			} else {
				projected[i] = responseChild
			}
		}
		return projected

	default:
		return response
	}
}
