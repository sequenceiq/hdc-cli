// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewChangeImageStackInWorkspaceV4InternalParams creates a new ChangeImageStackInWorkspaceV4InternalParams object
// with the default values initialized.
func NewChangeImageStackInWorkspaceV4InternalParams() *ChangeImageStackInWorkspaceV4InternalParams {
	var ()
	return &ChangeImageStackInWorkspaceV4InternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeImageStackInWorkspaceV4InternalParamsWithTimeout creates a new ChangeImageStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeImageStackInWorkspaceV4InternalParamsWithTimeout(timeout time.Duration) *ChangeImageStackInWorkspaceV4InternalParams {
	var ()
	return &ChangeImageStackInWorkspaceV4InternalParams{

		timeout: timeout,
	}
}

// NewChangeImageStackInWorkspaceV4InternalParamsWithContext creates a new ChangeImageStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeImageStackInWorkspaceV4InternalParamsWithContext(ctx context.Context) *ChangeImageStackInWorkspaceV4InternalParams {
	var ()
	return &ChangeImageStackInWorkspaceV4InternalParams{

		Context: ctx,
	}
}

// NewChangeImageStackInWorkspaceV4InternalParamsWithHTTPClient creates a new ChangeImageStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeImageStackInWorkspaceV4InternalParamsWithHTTPClient(client *http.Client) *ChangeImageStackInWorkspaceV4InternalParams {
	var ()
	return &ChangeImageStackInWorkspaceV4InternalParams{
		HTTPClient: client,
	}
}

/*ChangeImageStackInWorkspaceV4InternalParams contains all the parameters to send to the API endpoint
for the change image stack in workspace v4 internal operation typically these are written to a http.Request
*/
type ChangeImageStackInWorkspaceV4InternalParams struct {

	/*Body*/
	Body *model.StackImageChangeV4Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithTimeout(timeout time.Duration) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithContext(ctx context.Context) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithHTTPClient(client *http.Client) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithBody(body *model.StackImageChangeV4Request) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetBody(body *model.StackImageChangeV4Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithName(name string) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) WithWorkspaceID(workspaceID int64) *ChangeImageStackInWorkspaceV4InternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the change image stack in workspace v4 internal params
func (o *ChangeImageStackInWorkspaceV4InternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeImageStackInWorkspaceV4InternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
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