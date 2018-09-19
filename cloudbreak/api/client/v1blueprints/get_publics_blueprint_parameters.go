// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

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

// NewGetPublicsBlueprintParams creates a new GetPublicsBlueprintParams object
// with the default values initialized.
func NewGetPublicsBlueprintParams() *GetPublicsBlueprintParams {

	return &GetPublicsBlueprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicsBlueprintParamsWithTimeout creates a new GetPublicsBlueprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicsBlueprintParamsWithTimeout(timeout time.Duration) *GetPublicsBlueprintParams {

	return &GetPublicsBlueprintParams{

		timeout: timeout,
	}
}

// NewGetPublicsBlueprintParamsWithContext creates a new GetPublicsBlueprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicsBlueprintParamsWithContext(ctx context.Context) *GetPublicsBlueprintParams {

	return &GetPublicsBlueprintParams{

		Context: ctx,
	}
}

// NewGetPublicsBlueprintParamsWithHTTPClient creates a new GetPublicsBlueprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicsBlueprintParamsWithHTTPClient(client *http.Client) *GetPublicsBlueprintParams {

	return &GetPublicsBlueprintParams{
		HTTPClient: client,
	}
}

/*GetPublicsBlueprintParams contains all the parameters to send to the API endpoint
for the get publics blueprint operation typically these are written to a http.Request
*/
type GetPublicsBlueprintParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get publics blueprint params
func (o *GetPublicsBlueprintParams) WithTimeout(timeout time.Duration) *GetPublicsBlueprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publics blueprint params
func (o *GetPublicsBlueprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publics blueprint params
func (o *GetPublicsBlueprintParams) WithContext(ctx context.Context) *GetPublicsBlueprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publics blueprint params
func (o *GetPublicsBlueprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publics blueprint params
func (o *GetPublicsBlueprintParams) WithHTTPClient(client *http.Client) *GetPublicsBlueprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publics blueprint params
func (o *GetPublicsBlueprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicsBlueprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}