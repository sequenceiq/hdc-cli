// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

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

// NewListRecipesByWorkspaceParams creates a new ListRecipesByWorkspaceParams object
// with the default values initialized.
func NewListRecipesByWorkspaceParams() *ListRecipesByWorkspaceParams {
	var ()
	return &ListRecipesByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRecipesByWorkspaceParamsWithTimeout creates a new ListRecipesByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRecipesByWorkspaceParamsWithTimeout(timeout time.Duration) *ListRecipesByWorkspaceParams {
	var ()
	return &ListRecipesByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListRecipesByWorkspaceParamsWithContext creates a new ListRecipesByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRecipesByWorkspaceParamsWithContext(ctx context.Context) *ListRecipesByWorkspaceParams {
	var ()
	return &ListRecipesByWorkspaceParams{

		Context: ctx,
	}
}

// NewListRecipesByWorkspaceParamsWithHTTPClient creates a new ListRecipesByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRecipesByWorkspaceParamsWithHTTPClient(client *http.Client) *ListRecipesByWorkspaceParams {
	var ()
	return &ListRecipesByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListRecipesByWorkspaceParams contains all the parameters to send to the API endpoint
for the list recipes by workspace operation typically these are written to a http.Request
*/
type ListRecipesByWorkspaceParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) WithTimeout(timeout time.Duration) *ListRecipesByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) WithContext(ctx context.Context) *ListRecipesByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) WithHTTPClient(client *http.Client) *ListRecipesByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListRecipesByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list recipes by workspace params
func (o *ListRecipesByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRecipesByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
