// Code generated by go-swagger; DO NOT EDIT.

package oidc

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
)

// NewLoggedInUserInfoParams creates a new LoggedInUserInfoParams object
// with the default values initialized.
func NewLoggedInUserInfoParams() *LoggedInUserInfoParams {

	return &LoggedInUserInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoggedInUserInfoParamsWithTimeout creates a new LoggedInUserInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoggedInUserInfoParamsWithTimeout(timeout time.Duration) *LoggedInUserInfoParams {

	return &LoggedInUserInfoParams{

		timeout: timeout,
	}
}

// NewLoggedInUserInfoParamsWithContext creates a new LoggedInUserInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoggedInUserInfoParamsWithContext(ctx context.Context) *LoggedInUserInfoParams {

	return &LoggedInUserInfoParams{

		Context: ctx,
	}
}

// NewLoggedInUserInfoParamsWithHTTPClient creates a new LoggedInUserInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoggedInUserInfoParamsWithHTTPClient(client *http.Client) *LoggedInUserInfoParams {

	return &LoggedInUserInfoParams{
		HTTPClient: client,
	}
}

/*LoggedInUserInfoParams contains all the parameters to send to the API endpoint
for the logged in user info operation typically these are written to a http.Request
*/
type LoggedInUserInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logged in user info params
func (o *LoggedInUserInfoParams) WithTimeout(timeout time.Duration) *LoggedInUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logged in user info params
func (o *LoggedInUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logged in user info params
func (o *LoggedInUserInfoParams) WithContext(ctx context.Context) *LoggedInUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logged in user info params
func (o *LoggedInUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logged in user info params
func (o *LoggedInUserInfoParams) WithHTTPClient(client *http.Client) *LoggedInUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logged in user info params
func (o *LoggedInUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LoggedInUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}