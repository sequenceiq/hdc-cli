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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewEditEnvironmentParams creates a new EditEnvironmentParams object
// with the default values initialized.
func NewEditEnvironmentParams() *EditEnvironmentParams {
	var ()
	return &EditEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditEnvironmentParamsWithTimeout creates a new EditEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditEnvironmentParamsWithTimeout(timeout time.Duration) *EditEnvironmentParams {
	var ()
	return &EditEnvironmentParams{

		timeout: timeout,
	}
}

// NewEditEnvironmentParamsWithContext creates a new EditEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditEnvironmentParamsWithContext(ctx context.Context) *EditEnvironmentParams {
	var ()
	return &EditEnvironmentParams{

		Context: ctx,
	}
}

// NewEditEnvironmentParamsWithHTTPClient creates a new EditEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditEnvironmentParamsWithHTTPClient(client *http.Client) *EditEnvironmentParams {
	var ()
	return &EditEnvironmentParams{
		HTTPClient: client,
	}
}

/*EditEnvironmentParams contains all the parameters to send to the API endpoint
for the edit environment operation typically these are written to a http.Request
*/
type EditEnvironmentParams struct {

	/*Body*/
	Body *model.EnvironmentEditRequest
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit environment params
func (o *EditEnvironmentParams) WithTimeout(timeout time.Duration) *EditEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit environment params
func (o *EditEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit environment params
func (o *EditEnvironmentParams) WithContext(ctx context.Context) *EditEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit environment params
func (o *EditEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit environment params
func (o *EditEnvironmentParams) WithHTTPClient(client *http.Client) *EditEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit environment params
func (o *EditEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit environment params
func (o *EditEnvironmentParams) WithBody(body *model.EnvironmentEditRequest) *EditEnvironmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit environment params
func (o *EditEnvironmentParams) SetBody(body *model.EnvironmentEditRequest) {
	o.Body = body
}

// WithName adds the name to the edit environment params
func (o *EditEnvironmentParams) WithName(name string) *EditEnvironmentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edit environment params
func (o *EditEnvironmentParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the edit environment params
func (o *EditEnvironmentParams) WithWorkspaceID(workspaceID int64) *EditEnvironmentParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the edit environment params
func (o *EditEnvironmentParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *EditEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
