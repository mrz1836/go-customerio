package customerio

import (
	"encoding/json"
	"fmt"
	"strings"
)

// APIError is returned by any method that fails at the API level
type APIError struct {
	body   []byte
	status int
	url    string
}

// Error is used to display the error message
// Escape the returned JSON string will make it easier to consume in other applications
func (a *APIError) Error() string {

	// Escape error if JSON
	str := string(a.body)
	if strings.Contains(str, "{") {
		b, err := json.Marshal(str)
		if err != nil {
			return fmt.Sprintf("%v: %v %v", a.status, a.url, str)
		}
		s := string(b)
		return fmt.Sprintf("%v: %v %v", a.status, a.url, s[1:len(s)-1])
	}

	return fmt.Sprintf("%v: %v %v", a.status, a.url, string(a.body))
}

// ParamError is an error returned if a parameter to the track API is invalid.
type ParamError struct {
	Param string // Param is the name of the parameter.
}

// Error is used to display the error message
func (p ParamError) Error() string { return p.Param + ": missing" }
