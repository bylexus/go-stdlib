package router

// RouteParams keeps the parameters of a matched route
// with placeholders (e.g '/foo/{:entity}/{:id|[\d]+}').
// Each placeholder can only appear once in the URL,
// no multiple/array values are supported.
type RouteParams map[string]string

// This helper function creates a RouteParams from two keys/values
// array. keys and values must have the same length,
// then each key in keys is mapped to the corresponding value in values.
func RouteParamsFromKeyValues(keys []string, values []string) RouteParams {
	params := make(RouteParams)
	if len(keys) != len(values) {
		return params
	}
	for i := 0; i < len(keys); i++ {
		params[keys[i]] = values[i]
	}
	return params
}
