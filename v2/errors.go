package yacontest

import (
	"fmt"
)

// APIError Define API error when response status is 4xx or 5xx
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

type ValidateError struct {
	Message string
}

func newError() *ValidateError {
	return &ValidateError{}
}

func (e ValidateError) Errorf(message string, args ...interface{}) error {
	return &ValidateError{
		Message: fmt.Sprintf(message, args...),
	}
}

func (e ValidateError) Required(fieldName string) error {
	return e.Errorf("%s can not be empty", fieldName)
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("<ValidateError> msg=%s", e.Message)
}

// Error Return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d msg=%s", e.Code, e.Message)
}

// IsAPIError Check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}