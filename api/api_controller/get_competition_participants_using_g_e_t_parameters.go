// Code generated by go-swagger; DO NOT EDIT.

package api_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCompetitionParticipantsUsingGETParams creates a new GetCompetitionParticipantsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompetitionParticipantsUsingGETParams() *GetCompetitionParticipantsUsingGETParams {
	return &GetCompetitionParticipantsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompetitionParticipantsUsingGETParamsWithTimeout creates a new GetCompetitionParticipantsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetCompetitionParticipantsUsingGETParamsWithTimeout(timeout time.Duration) *GetCompetitionParticipantsUsingGETParams {
	return &GetCompetitionParticipantsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetCompetitionParticipantsUsingGETParamsWithContext creates a new GetCompetitionParticipantsUsingGETParams object
// with the ability to set a context for a request.
func NewGetCompetitionParticipantsUsingGETParamsWithContext(ctx context.Context) *GetCompetitionParticipantsUsingGETParams {
	return &GetCompetitionParticipantsUsingGETParams{
		Context: ctx,
	}
}

// NewGetCompetitionParticipantsUsingGETParamsWithHTTPClient creates a new GetCompetitionParticipantsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompetitionParticipantsUsingGETParamsWithHTTPClient(client *http.Client) *GetCompetitionParticipantsUsingGETParams {
	return &GetCompetitionParticipantsUsingGETParams{
		HTTPClient: client,
	}
}

/* GetCompetitionParticipantsUsingGETParams contains all the parameters to send to the API endpoint
   for the get competition participants using g e t operation.

   Typically these are written to a http.Request.
*/
type GetCompetitionParticipantsUsingGETParams struct {

	/* CompetitionID.

	   competitionId

	   Format: int64
	*/
	CompetitionID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get competition participants using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompetitionParticipantsUsingGETParams) WithDefaults() *GetCompetitionParticipantsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get competition participants using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompetitionParticipantsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) WithTimeout(timeout time.Duration) *GetCompetitionParticipantsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) WithContext(ctx context.Context) *GetCompetitionParticipantsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) WithHTTPClient(client *http.Client) *GetCompetitionParticipantsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompetitionID adds the competitionID to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) WithCompetitionID(competitionID int64) *GetCompetitionParticipantsUsingGETParams {
	o.SetCompetitionID(competitionID)
	return o
}

// SetCompetitionID adds the competitionId to the get competition participants using g e t params
func (o *GetCompetitionParticipantsUsingGETParams) SetCompetitionID(competitionID int64) {
	o.CompetitionID = competitionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompetitionParticipantsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param competitionId
	if err := r.SetPathParam("competitionId", swag.FormatInt64(o.CompetitionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
