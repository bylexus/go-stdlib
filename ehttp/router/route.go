package router

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

// The Route type defines a set of rules to match a request against:
//
//   - the Method defines the HTTP method of the request to be matched.
//     It can be any of the supported HTTP methods (GET, POST etc), or `ANY` to match any HTTP method.
//
//   - Pattern is the URL pattern to be matched. A pattern can either be a simple string
//     (e.g. '/foo/bar') or a string containing placeholders (e.g. '/foo/{:param_name_123}').
//     The placeholder will match the regular expression `.+?`, which matches any string with
//     at least one character.
//     Placeholders also support custom regular expressions. For example, to match only numbers,
//     you can use `/foo/{:id|[\d]+}`: your custom regex is separated by a pipe (`|`).
//
//   - Handler is the handler to be called if the pattern matches the request. The matching route is injected
//     into the request's Context before the Handler is called (see Router for more details).
type Route struct {
	Handler      *http.Handler
	Pattern      string
	Method       string
	regexPattern *regexp.Regexp
}

// Create a new route by using `NewRoute` instead of the Struct constructor:
// Upon creation, the Pattern is parsed and a pre-compiled regular expression is generated.
func NewRoute(method string, pattern string, handler *http.Handler) (*Route, error) {
	re, err := createRegexRoutePattern(pattern)
	if err != nil {
		return nil, err
	}

	r := Route{
		Handler:      handler,
		Pattern:      pattern,
		Method:       method,
		regexPattern: re,
	}
	return &r, nil
}

// The MatchedRoute type contains the information of the matched request after
// a route has been found. It is injected to the request's Context, and contains the following information:
//
// - Route is the matched route
// - Params is the parameters of the matched route, if any, or an empty map
// - URL is the URL of the request that matched
// - Method is the matched HTTP method of the request
type MatchedRoute struct {
	Route  Route
	Params RouteParams
	URL    url.URL
	Method string
}

// Generates a regex patten of an URL pattern string
//
// An example URL pattern string:
//
//	/foo/bar.buz/{:param_name_123}/{:id|[\\d]+}
//
// we need to:
// - convert all '/' to '\/'
// - convert all '.' to '\.'
// - convert placeholders {:param_name_123} to '(?P<param_name_123>.+?)'
// - convert placeholders {:param_name_123|re} to '(?P<param_name_123>re)'
//
// first we need to extract the {:...} groups, replace them to a temporary
// value, replace all other regex meta chars to their escaped version,
// then insert the groups back as named regex group.
// so
//
//	/foo/bar.buz/{:param_name_123}/{:id|[\d]+}
//
// becomes
//
//	\/foo\/bar\.buz\/(?P<param_name_123>.+?)\/(?P<id>[\d]+)
func createRegexRoutePattern(rawPattern string) (*regexp.Regexp, error) {
	re, err := regexp.Compile(`(\{:.+?\})`)
	if err != nil {
		return nil, err
	}

	// extract param substrings:
	params := make([]string, 0)
	for _, param := range re.FindAllStringSubmatch(rawPattern, -1) {
		pattern, err := createParamRegexStrFromPattern(param[1])
		if err != nil {
			return nil, err
		}
		params = append(params, *pattern)
	}
	// extract non-param parts:
	nonParamParts := re.Split(rawPattern, -1)
	for i, part := range nonParamParts {
		nonParamParts[i] = regexp.QuoteMeta(part)
	}

	// mix'em together again:
	processed := "^"
	for i := 0; i < len(nonParamParts); i++ {
		processed += nonParamParts[i]
		if len(params) > i {
			processed += params[i]
		}
	}
	processed += "$"

	// compile the route to a regex:
	re, err = regexp.Compile(processed)
	if err != nil {
		return nil, err
	}

	return re, err
}

// input: e.g. {:param_name_123} or {:id|[\d]+}
// output: e.g. (?P<param_name_123>.+?) or (?P<id>[\d]+)
func createParamRegexStrFromPattern(pattern string) (*string, error) {
	re, err := regexp.Compile(`\{:(\w+)(\|(.+?))?\}`)
	if err != nil {
		return nil, err
	}
	res := re.FindStringSubmatch(pattern)
	if len(res) >= 1 {
		name := res[1]
		reStr := `.+?`
		if len(res) == 4 && len(res[3]) > 0 {
			reStr = res[3]
		}
		result := fmt.Sprintf("(?P<%s>%s)", name, reStr)
		return &result, nil
	} else {
		return nil, errors.New("no placeholder param pattern")
	}
}
