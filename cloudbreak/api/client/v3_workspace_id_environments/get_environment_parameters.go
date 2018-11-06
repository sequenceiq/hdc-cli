// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_environments

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

// NewGetEnvironmentParams creates a new GetEnvironmentParams object
// with the default values initialized.
func NewGetEnvironmentParams() *GetEnvironmentParams {
	var ()
	return &GetEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentParamsWithTimeout creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentParamsWithTimeout(timeout time.Duration) *GetEnvironmentParams {
	var ()
	return &GetEnvironmentParams{

		timeout: timeout,
	}
}

// NewGetEnvironmentParamsWithContext creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentParamsWithContext(ctx context.Context) *GetEnvironmentParams {
	var ()
	return &GetEnvironmentParams{

		Context: ctx,
	}
}

// NewGetEnvironmentParamsWithHTTPClient creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentParamsWithHTTPClient(client *http.Client) *GetEnvironmentParams {
	var ()
	return &GetEnvironmentParams{
		HTTPClient: client,
	}
}

/*GetEnvironmentParams contains all the parameters to send to the API endpoint
for the get environment operation typically these are written to a http.Request
*/
type GetEnvironmentParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) WithTimeout(timeout time.Duration) *GetEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment params
func (o *GetEnvironmentParams) WithContext(ctx context.Context) *GetEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment params
func (o *GetEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) WithHTTPClient(client *http.Client) *GetEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get environment params
func (o *GetEnvironmentParams) WithName(name string) *GetEnvironmentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get environment params
func (o *GetEnvironmentParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get environment params
func (o *GetEnvironmentParams) WithWorkspaceID(workspaceID int64) *GetEnvironmentParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get environment params
func (o *GetEnvironmentParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
