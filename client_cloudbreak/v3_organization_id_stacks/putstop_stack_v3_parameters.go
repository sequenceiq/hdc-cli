// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutstopStackV3Params creates a new PutstopStackV3Params object
// with the default values initialized.
func NewPutstopStackV3Params() *PutstopStackV3Params {
	var ()
	return &PutstopStackV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutstopStackV3ParamsWithTimeout creates a new PutstopStackV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutstopStackV3ParamsWithTimeout(timeout time.Duration) *PutstopStackV3Params {
	var ()
	return &PutstopStackV3Params{

		timeout: timeout,
	}
}

// NewPutstopStackV3ParamsWithContext creates a new PutstopStackV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutstopStackV3ParamsWithContext(ctx context.Context) *PutstopStackV3Params {
	var ()
	return &PutstopStackV3Params{

		Context: ctx,
	}
}

// NewPutstopStackV3ParamsWithHTTPClient creates a new PutstopStackV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutstopStackV3ParamsWithHTTPClient(client *http.Client) *PutstopStackV3Params {
	var ()
	return &PutstopStackV3Params{
		HTTPClient: client,
	}
}

/*PutstopStackV3Params contains all the parameters to send to the API endpoint
for the putstop stack v3 operation typically these are written to a http.Request
*/
type PutstopStackV3Params struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the putstop stack v3 params
func (o *PutstopStackV3Params) WithTimeout(timeout time.Duration) *PutstopStackV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the putstop stack v3 params
func (o *PutstopStackV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the putstop stack v3 params
func (o *PutstopStackV3Params) WithContext(ctx context.Context) *PutstopStackV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the putstop stack v3 params
func (o *PutstopStackV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the putstop stack v3 params
func (o *PutstopStackV3Params) WithHTTPClient(client *http.Client) *PutstopStackV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the putstop stack v3 params
func (o *PutstopStackV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the putstop stack v3 params
func (o *PutstopStackV3Params) WithName(name string) *PutstopStackV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the putstop stack v3 params
func (o *PutstopStackV3Params) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the putstop stack v3 params
func (o *PutstopStackV3Params) WithOrganizationID(organizationID int64) *PutstopStackV3Params {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the putstop stack v3 params
func (o *PutstopStackV3Params) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutstopStackV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}