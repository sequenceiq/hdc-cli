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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteEnvironmentsByCrnParams creates a new DeleteEnvironmentsByCrnParams object
// with the default values initialized.
func NewDeleteEnvironmentsByCrnParams() *DeleteEnvironmentsByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteEnvironmentsByCrnParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentsByCrnParamsWithTimeout creates a new DeleteEnvironmentsByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentsByCrnParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentsByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteEnvironmentsByCrnParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteEnvironmentsByCrnParamsWithContext creates a new DeleteEnvironmentsByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentsByCrnParamsWithContext(ctx context.Context) *DeleteEnvironmentsByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteEnvironmentsByCrnParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteEnvironmentsByCrnParamsWithHTTPClient creates a new DeleteEnvironmentsByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentsByCrnParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentsByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteEnvironmentsByCrnParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*DeleteEnvironmentsByCrnParams contains all the parameters to send to the API endpoint
for the delete environments by crn operation typically these are written to a http.Request
*/
type DeleteEnvironmentsByCrnParams struct {

	/*Body*/
	Body []string
	/*Forced*/
	Forced *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentsByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) WithContext(ctx context.Context) *DeleteEnvironmentsByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentsByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) WithBody(body []string) *DeleteEnvironmentsByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) SetBody(body []string) {
	o.Body = body
}

// WithForced adds the forced to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) WithForced(forced *bool) *DeleteEnvironmentsByCrnParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete environments by crn params
func (o *DeleteEnvironmentsByCrnParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentsByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}