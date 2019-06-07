// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewEditEnvironmentV1Params creates a new EditEnvironmentV1Params object
// with the default values initialized.
func NewEditEnvironmentV1Params() *EditEnvironmentV1Params {
	var ()
	return &EditEnvironmentV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditEnvironmentV1ParamsWithTimeout creates a new EditEnvironmentV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditEnvironmentV1ParamsWithTimeout(timeout time.Duration) *EditEnvironmentV1Params {
	var ()
	return &EditEnvironmentV1Params{

		timeout: timeout,
	}
}

// NewEditEnvironmentV1ParamsWithContext creates a new EditEnvironmentV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewEditEnvironmentV1ParamsWithContext(ctx context.Context) *EditEnvironmentV1Params {
	var ()
	return &EditEnvironmentV1Params{

		Context: ctx,
	}
}

// NewEditEnvironmentV1ParamsWithHTTPClient creates a new EditEnvironmentV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditEnvironmentV1ParamsWithHTTPClient(client *http.Client) *EditEnvironmentV1Params {
	var ()
	return &EditEnvironmentV1Params{
		HTTPClient: client,
	}
}

/*EditEnvironmentV1Params contains all the parameters to send to the API endpoint
for the edit environment v1 operation typically these are written to a http.Request
*/
type EditEnvironmentV1Params struct {

	/*Body*/
	Body *model.EnvironmentEditV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit environment v1 params
func (o *EditEnvironmentV1Params) WithTimeout(timeout time.Duration) *EditEnvironmentV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit environment v1 params
func (o *EditEnvironmentV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit environment v1 params
func (o *EditEnvironmentV1Params) WithContext(ctx context.Context) *EditEnvironmentV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit environment v1 params
func (o *EditEnvironmentV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit environment v1 params
func (o *EditEnvironmentV1Params) WithHTTPClient(client *http.Client) *EditEnvironmentV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit environment v1 params
func (o *EditEnvironmentV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit environment v1 params
func (o *EditEnvironmentV1Params) WithBody(body *model.EnvironmentEditV1Request) *EditEnvironmentV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit environment v1 params
func (o *EditEnvironmentV1Params) SetBody(body *model.EnvironmentEditV1Request) {
	o.Body = body
}

// WithName adds the name to the edit environment v1 params
func (o *EditEnvironmentV1Params) WithName(name string) *EditEnvironmentV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the edit environment v1 params
func (o *EditEnvironmentV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EditEnvironmentV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}