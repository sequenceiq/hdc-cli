// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_audits

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

// NewGetAuditEventsZipInWorkspaceParams creates a new GetAuditEventsZipInWorkspaceParams object
// with the default values initialized.
func NewGetAuditEventsZipInWorkspaceParams() *GetAuditEventsZipInWorkspaceParams {
	var ()
	return &GetAuditEventsZipInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditEventsZipInWorkspaceParamsWithTimeout creates a new GetAuditEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditEventsZipInWorkspaceParamsWithTimeout(timeout time.Duration) *GetAuditEventsZipInWorkspaceParams {
	var ()
	return &GetAuditEventsZipInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetAuditEventsZipInWorkspaceParamsWithContext creates a new GetAuditEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditEventsZipInWorkspaceParamsWithContext(ctx context.Context) *GetAuditEventsZipInWorkspaceParams {
	var ()
	return &GetAuditEventsZipInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetAuditEventsZipInWorkspaceParamsWithHTTPClient creates a new GetAuditEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditEventsZipInWorkspaceParamsWithHTTPClient(client *http.Client) *GetAuditEventsZipInWorkspaceParams {
	var ()
	return &GetAuditEventsZipInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetAuditEventsZipInWorkspaceParams contains all the parameters to send to the API endpoint
for the get audit events zip in workspace operation typically these are written to a http.Request
*/
type GetAuditEventsZipInWorkspaceParams struct {

	/*ResourceID*/
	ResourceID int64
	/*ResourceType*/
	ResourceType string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithTimeout(timeout time.Duration) *GetAuditEventsZipInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithContext(ctx context.Context) *GetAuditEventsZipInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithHTTPClient(client *http.Client) *GetAuditEventsZipInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithResourceID(resourceID int64) *GetAuditEventsZipInWorkspaceParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetResourceID(resourceID int64) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithResourceType(resourceType string) *GetAuditEventsZipInWorkspaceParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithWorkspaceID adds the workspaceID to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetAuditEventsZipInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get audit events zip in workspace params
func (o *GetAuditEventsZipInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditEventsZipInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", swag.FormatInt64(o.ResourceID)); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
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
