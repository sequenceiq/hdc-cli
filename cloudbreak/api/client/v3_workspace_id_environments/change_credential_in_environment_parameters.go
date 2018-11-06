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

	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
)

// NewChangeCredentialInEnvironmentParams creates a new ChangeCredentialInEnvironmentParams object
// with the default values initialized.
func NewChangeCredentialInEnvironmentParams() *ChangeCredentialInEnvironmentParams {
	var ()
	return &ChangeCredentialInEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeCredentialInEnvironmentParamsWithTimeout creates a new ChangeCredentialInEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeCredentialInEnvironmentParamsWithTimeout(timeout time.Duration) *ChangeCredentialInEnvironmentParams {
	var ()
	return &ChangeCredentialInEnvironmentParams{

		timeout: timeout,
	}
}

// NewChangeCredentialInEnvironmentParamsWithContext creates a new ChangeCredentialInEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeCredentialInEnvironmentParamsWithContext(ctx context.Context) *ChangeCredentialInEnvironmentParams {
	var ()
	return &ChangeCredentialInEnvironmentParams{

		Context: ctx,
	}
}

// NewChangeCredentialInEnvironmentParamsWithHTTPClient creates a new ChangeCredentialInEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeCredentialInEnvironmentParamsWithHTTPClient(client *http.Client) *ChangeCredentialInEnvironmentParams {
	var ()
	return &ChangeCredentialInEnvironmentParams{
		HTTPClient: client,
	}
}

/*ChangeCredentialInEnvironmentParams contains all the parameters to send to the API endpoint
for the change credential in environment operation typically these are written to a http.Request
*/
type ChangeCredentialInEnvironmentParams struct {

	/*Body*/
	Body *model.EnvironmentChangeCredentialRequest
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithTimeout(timeout time.Duration) *ChangeCredentialInEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithContext(ctx context.Context) *ChangeCredentialInEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithHTTPClient(client *http.Client) *ChangeCredentialInEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithBody(body *model.EnvironmentChangeCredentialRequest) *ChangeCredentialInEnvironmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetBody(body *model.EnvironmentChangeCredentialRequest) {
	o.Body = body
}

// WithName adds the name to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithName(name string) *ChangeCredentialInEnvironmentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) WithWorkspaceID(workspaceID int64) *ChangeCredentialInEnvironmentParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the change credential in environment params
func (o *ChangeCredentialInEnvironmentParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeCredentialInEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(model.EnvironmentChangeCredentialRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

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
