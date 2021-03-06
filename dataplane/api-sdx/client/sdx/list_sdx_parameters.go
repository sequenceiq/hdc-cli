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

// NewListSdxParams creates a new ListSdxParams object
// with the default values initialized.
func NewListSdxParams() *ListSdxParams {
	var ()
	return &ListSdxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSdxParamsWithTimeout creates a new ListSdxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSdxParamsWithTimeout(timeout time.Duration) *ListSdxParams {
	var ()
	return &ListSdxParams{

		timeout: timeout,
	}
}

// NewListSdxParamsWithContext creates a new ListSdxParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSdxParamsWithContext(ctx context.Context) *ListSdxParams {
	var ()
	return &ListSdxParams{

		Context: ctx,
	}
}

// NewListSdxParamsWithHTTPClient creates a new ListSdxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSdxParamsWithHTTPClient(client *http.Client) *ListSdxParams {
	var ()
	return &ListSdxParams{
		HTTPClient: client,
	}
}

/*ListSdxParams contains all the parameters to send to the API endpoint
for the list sdx operation typically these are written to a http.Request
*/
type ListSdxParams struct {

	/*EnvName*/
	EnvName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sdx params
func (o *ListSdxParams) WithTimeout(timeout time.Duration) *ListSdxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sdx params
func (o *ListSdxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sdx params
func (o *ListSdxParams) WithContext(ctx context.Context) *ListSdxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sdx params
func (o *ListSdxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sdx params
func (o *ListSdxParams) WithHTTPClient(client *http.Client) *ListSdxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sdx params
func (o *ListSdxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvName adds the envName to the list sdx params
func (o *ListSdxParams) WithEnvName(envName *string) *ListSdxParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the list sdx params
func (o *ListSdxParams) SetEnvName(envName *string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *ListSdxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvName != nil {

		// query param envName
		var qrEnvName string
		if o.EnvName != nil {
			qrEnvName = *o.EnvName
		}
		qEnvName := qrEnvName
		if qEnvName != "" {
			if err := r.SetQueryParam("envName", qEnvName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
