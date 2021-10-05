// Code generated by go-swagger; DO NOT EDIT.

package api_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamdbnn/go-yandex-contest/models"
)

// GetClarificationsUsingGETReader is a Reader for the GetClarificationsUsingGET structure.
type GetClarificationsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClarificationsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClarificationsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClarificationsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClarificationsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClarificationsUsingGETOK creates a GetClarificationsUsingGETOK with default headers values
func NewGetClarificationsUsingGETOK() *GetClarificationsUsingGETOK {
	return &GetClarificationsUsingGETOK{}
}

/* GetClarificationsUsingGETOK describes a response with status code 200, with default header values.

Successfully received clarifications
*/
type GetClarificationsUsingGETOK struct {
	Payload *models.Clarifications
}

func (o *GetClarificationsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/clarifications][%d] getClarificationsUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetClarificationsUsingGETOK) GetPayload() *models.Clarifications {
	return o.Payload
}

func (o *GetClarificationsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Clarifications)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClarificationsUsingGETUnauthorized creates a GetClarificationsUsingGETUnauthorized with default headers values
func NewGetClarificationsUsingGETUnauthorized() *GetClarificationsUsingGETUnauthorized {
	return &GetClarificationsUsingGETUnauthorized{}
}

/* GetClarificationsUsingGETUnauthorized describes a response with status code 401, with default header values.

OAuth header is not declared or is wrong
*/
type GetClarificationsUsingGETUnauthorized struct {
}

func (o *GetClarificationsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/clarifications][%d] getClarificationsUsingGETUnauthorized ", 401)
}

func (o *GetClarificationsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClarificationsUsingGETNotFound creates a GetClarificationsUsingGETNotFound with default headers values
func NewGetClarificationsUsingGETNotFound() *GetClarificationsUsingGETNotFound {
	return &GetClarificationsUsingGETNotFound{}
}

/* GetClarificationsUsingGETNotFound describes a response with status code 404, with default header values.

Contest not found
*/
type GetClarificationsUsingGETNotFound struct {
}

func (o *GetClarificationsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/clarifications][%d] getClarificationsUsingGETNotFound ", 404)
}

func (o *GetClarificationsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
