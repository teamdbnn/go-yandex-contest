package common

import "fmt"

type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d msg=%s", e.Code, e.Message)
}

func IsApiError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
