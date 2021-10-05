package models

type ResponseEntity struct {
	Body            interface{} `json:"body,omitempty"`
	StatusCode      string      `json:"statusCode,omitempty"`
	StatusCodeValue int32       `json:"statusCodeValue,omitempty"`
}
