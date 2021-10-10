package common

import (
	"fmt"

	"github.com/pkg/errors"
)

const RequiredMessage string = "%[1]s is required. Call %[1]s() method."

// APIError Define API error when response status is 4xx or 5xx
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func ValidationRequiredError(fieldName string) error {
	return errors.Errorf(RequiredMessage, fieldName)
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
