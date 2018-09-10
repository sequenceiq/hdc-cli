// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

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

// NewListStacksByWorkspaceParams creates a new ListStacksByWorkspaceParams object
// with the default values initialized.
func NewListStacksByWorkspaceParams() *ListStacksByWorkspaceParams {
	var ()
	return &ListStacksByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListStacksByWorkspaceParamsWithTimeout creates a new ListStacksByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListStacksByWorkspaceParamsWithTimeout(timeout time.Duration) *ListStacksByWorkspaceParams {
	var ()
	return &ListStacksByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListStacksByWorkspaceParamsWithContext creates a new ListStacksByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListStacksByWorkspaceParamsWithContext(ctx context.Context) *ListStacksByWorkspaceParams {
	var ()
	return &ListStacksByWorkspaceParams{

		Context: ctx,
	}
}

// NewListStacksByWorkspaceParamsWithHTTPClient creates a new ListStacksByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListStacksByWorkspaceParamsWithHTTPClient(client *http.Client) *ListStacksByWorkspaceParams {
	var ()
	return &ListStacksByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListStacksByWorkspaceParams contains all the parameters to send to the API endpoint
for the list stacks by workspace operation typically these are written to a http.Request
*/
type ListStacksByWorkspaceParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) WithTimeout(timeout time.Duration) *ListStacksByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) WithContext(ctx context.Context) *ListStacksByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) WithHTTPClient(client *http.Client) *ListStacksByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListStacksByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list stacks by workspace params
func (o *ListStacksByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListStacksByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
