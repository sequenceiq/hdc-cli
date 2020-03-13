// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStartSdxByCrnParams creates a new StartSdxByCrnParams object
// with the default values initialized.
func NewStartSdxByCrnParams() *StartSdxByCrnParams {
	var ()
	return &StartSdxByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartSdxByCrnParamsWithTimeout creates a new StartSdxByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartSdxByCrnParamsWithTimeout(timeout time.Duration) *StartSdxByCrnParams {
	var ()
	return &StartSdxByCrnParams{

		timeout: timeout,
	}
}

// NewStartSdxByCrnParamsWithContext creates a new StartSdxByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartSdxByCrnParamsWithContext(ctx context.Context) *StartSdxByCrnParams {
	var ()
	return &StartSdxByCrnParams{

		Context: ctx,
	}
}

// NewStartSdxByCrnParamsWithHTTPClient creates a new StartSdxByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartSdxByCrnParamsWithHTTPClient(client *http.Client) *StartSdxByCrnParams {
	var ()
	return &StartSdxByCrnParams{
		HTTPClient: client,
	}
}

/*StartSdxByCrnParams contains all the parameters to send to the API endpoint
for the start sdx by crn operation typically these are written to a http.Request
*/
type StartSdxByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start sdx by crn params
func (o *StartSdxByCrnParams) WithTimeout(timeout time.Duration) *StartSdxByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start sdx by crn params
func (o *StartSdxByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start sdx by crn params
func (o *StartSdxByCrnParams) WithContext(ctx context.Context) *StartSdxByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start sdx by crn params
func (o *StartSdxByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start sdx by crn params
func (o *StartSdxByCrnParams) WithHTTPClient(client *http.Client) *StartSdxByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start sdx by crn params
func (o *StartSdxByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the start sdx by crn params
func (o *StartSdxByCrnParams) WithCrn(crn string) *StartSdxByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the start sdx by crn params
func (o *StartSdxByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *StartSdxByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}