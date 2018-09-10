// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_rdsconfigs

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

// NewListRdsConfigsByWorkspaceParams creates a new ListRdsConfigsByWorkspaceParams object
// with the default values initialized.
func NewListRdsConfigsByWorkspaceParams() *ListRdsConfigsByWorkspaceParams {
	var ()
	return &ListRdsConfigsByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRdsConfigsByWorkspaceParamsWithTimeout creates a new ListRdsConfigsByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRdsConfigsByWorkspaceParamsWithTimeout(timeout time.Duration) *ListRdsConfigsByWorkspaceParams {
	var ()
	return &ListRdsConfigsByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListRdsConfigsByWorkspaceParamsWithContext creates a new ListRdsConfigsByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRdsConfigsByWorkspaceParamsWithContext(ctx context.Context) *ListRdsConfigsByWorkspaceParams {
	var ()
	return &ListRdsConfigsByWorkspaceParams{

		Context: ctx,
	}
}

// NewListRdsConfigsByWorkspaceParamsWithHTTPClient creates a new ListRdsConfigsByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRdsConfigsByWorkspaceParamsWithHTTPClient(client *http.Client) *ListRdsConfigsByWorkspaceParams {
	var ()
	return &ListRdsConfigsByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListRdsConfigsByWorkspaceParams contains all the parameters to send to the API endpoint
for the list rds configs by workspace operation typically these are written to a http.Request
*/
type ListRdsConfigsByWorkspaceParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) WithTimeout(timeout time.Duration) *ListRdsConfigsByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) WithContext(ctx context.Context) *ListRdsConfigsByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) WithHTTPClient(client *http.Client) *ListRdsConfigsByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListRdsConfigsByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list rds configs by workspace params
func (o *ListRdsConfigsByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRdsConfigsByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
