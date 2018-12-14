// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

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

// NewGetPublicLdapParams creates a new GetPublicLdapParams object
// with the default values initialized.
func NewGetPublicLdapParams() *GetPublicLdapParams {
	var ()
	return &GetPublicLdapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicLdapParamsWithTimeout creates a new GetPublicLdapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicLdapParamsWithTimeout(timeout time.Duration) *GetPublicLdapParams {
	var ()
	return &GetPublicLdapParams{

		timeout: timeout,
	}
}

// NewGetPublicLdapParamsWithContext creates a new GetPublicLdapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicLdapParamsWithContext(ctx context.Context) *GetPublicLdapParams {
	var ()
	return &GetPublicLdapParams{

		Context: ctx,
	}
}

// NewGetPublicLdapParamsWithHTTPClient creates a new GetPublicLdapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicLdapParamsWithHTTPClient(client *http.Client) *GetPublicLdapParams {
	var ()
	return &GetPublicLdapParams{
		HTTPClient: client,
	}
}

/*GetPublicLdapParams contains all the parameters to send to the API endpoint
for the get public ldap operation typically these are written to a http.Request
*/
type GetPublicLdapParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public ldap params
func (o *GetPublicLdapParams) WithTimeout(timeout time.Duration) *GetPublicLdapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public ldap params
func (o *GetPublicLdapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public ldap params
func (o *GetPublicLdapParams) WithContext(ctx context.Context) *GetPublicLdapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public ldap params
func (o *GetPublicLdapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public ldap params
func (o *GetPublicLdapParams) WithHTTPClient(client *http.Client) *GetPublicLdapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public ldap params
func (o *GetPublicLdapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get public ldap params
func (o *GetPublicLdapParams) WithName(name string) *GetPublicLdapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public ldap params
func (o *GetPublicLdapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicLdapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}