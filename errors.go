package customerio

import "fmt"

// APIError is returned by any method that fails at the API level
type APIError struct {
	body   []byte
	status int
	url    string
}

// Error is used to display the error message
func (a *APIError) Error() string {
	return fmt.Sprintf("%v: %v %v", a.status, a.url, string(a.body))
}

// ParamError is an error returned if a parameter to the track API is invalid.
type ParamError struct {
	Param string // Param is the name of the parameter.
}

// Error is used to display the error message
func (p ParamError) Error() string { return p.Param + ": missing" }
