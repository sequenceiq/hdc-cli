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

// NewGetRecipeByCrnInWorkspaceParams creates a new GetRecipeByCrnInWorkspaceParams object
// with the default values initialized.
func NewGetRecipeByCrnInWorkspaceParams() *GetRecipeByCrnInWorkspaceParams {
	var ()
	return &GetRecipeByCrnInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeByCrnInWorkspaceParamsWithTimeout creates a new GetRecipeByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeByCrnInWorkspaceParamsWithTimeout(timeout time.Duration) *GetRecipeByCrnInWorkspaceParams {
	var ()
	return &GetRecipeByCrnInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetRecipeByCrnInWorkspaceParamsWithContext creates a new GetRecipeByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeByCrnInWorkspaceParamsWithContext(ctx context.Context) *GetRecipeByCrnInWorkspaceParams {
	var ()
	return &GetRecipeByCrnInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetRecipeByCrnInWorkspaceParamsWithHTTPClient creates a new GetRecipeByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeByCrnInWorkspaceParamsWithHTTPClient(client *http.Client) *GetRecipeByCrnInWorkspaceParams {
	var ()
	return &GetRecipeByCrnInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetRecipeByCrnInWorkspaceParams contains all the parameters to send to the API endpoint
for the get recipe by crn in workspace operation typically these are written to a http.Request
*/
type GetRecipeByCrnInWorkspaceParams struct {

	/*Crn*/
	Crn string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) WithTimeout(timeout time.Duration) *GetRecipeByCrnInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) WithContext(ctx context.Context) *GetRecipeByCrnInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) WithHTTPClient(client *http.Client) *GetRecipeByCrnInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) WithCrn(crn string) *GetRecipeByCrnInWorkspaceParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithWorkspaceID adds the workspaceID to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetRecipeByCrnInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get recipe by crn in workspace params
func (o *GetRecipeByCrnInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeByCrnInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
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