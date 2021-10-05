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

// GetParticipantStatsUsingGETReader is a Reader for the GetParticipantStatsUsingGET structure.
type GetParticipantStatsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParticipantStatsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParticipantStatsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetParticipantStatsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParticipantStatsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetParticipantStatsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetParticipantStatsUsingGETOK creates a GetParticipantStatsUsingGETOK with default headers values
func NewGetParticipantStatsUsingGETOK() *GetParticipantStatsUsingGETOK {
	return &GetParticipantStatsUsingGETOK{}
}

/* GetParticipantStatsUsingGETOK describes a response with status code 200, with default header values.

Successfully received information about participant runs
*/
type GetParticipantStatsUsingGETOK struct {
	Payload *models.ParticipantStats
}

func (o *GetParticipantStatsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/participants/{participantId}/stats][%d] getParticipantStatsUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetParticipantStatsUsingGETOK) GetPayload() *models.ParticipantStats {
	return o.Payload
}

func (o *GetParticipantStatsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParticipantStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParticipantStatsUsingGETUnauthorized creates a GetParticipantStatsUsingGETUnauthorized with default headers values
func NewGetParticipantStatsUsingGETUnauthorized() *GetParticipantStatsUsingGETUnauthorized {
	return &GetParticipantStatsUsingGETUnauthorized{}
}

/* GetParticipantStatsUsingGETUnauthorized describes a response with status code 401, with default header values.

OAuth header is not declared or is wrong
*/
type GetParticipantStatsUsingGETUnauthorized struct {
}

func (o *GetParticipantStatsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/participants/{participantId}/stats][%d] getParticipantStatsUsingGETUnauthorized ", 401)
}

func (o *GetParticipantStatsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParticipantStatsUsingGETForbidden creates a GetParticipantStatsUsingGETForbidden with default headers values
func NewGetParticipantStatsUsingGETForbidden() *GetParticipantStatsUsingGETForbidden {
	return &GetParticipantStatsUsingGETForbidden{}
}

/* GetParticipantStatsUsingGETForbidden describes a response with status code 403, with default header values.

You don't have enough permissions
*/
type GetParticipantStatsUsingGETForbidden struct {
}

func (o *GetParticipantStatsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/participants/{participantId}/stats][%d] getParticipantStatsUsingGETForbidden ", 403)
}

func (o *GetParticipantStatsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParticipantStatsUsingGETNotFound creates a GetParticipantStatsUsingGETNotFound with default headers values
func NewGetParticipantStatsUsingGETNotFound() *GetParticipantStatsUsingGETNotFound {
	return &GetParticipantStatsUsingGETNotFound{}
}

/* GetParticipantStatsUsingGETNotFound describes a response with status code 404, with default header values.

Contest or your participation is not found
*/
type GetParticipantStatsUsingGETNotFound struct {
}

func (o *GetParticipantStatsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /contests/{contestId}/participants/{participantId}/stats][%d] getParticipantStatsUsingGETNotFound ", 404)
}

func (o *GetParticipantStatsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
