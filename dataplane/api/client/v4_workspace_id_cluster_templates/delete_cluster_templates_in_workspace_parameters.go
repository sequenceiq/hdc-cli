// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_cluster_templates

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

// NewDeleteClusterTemplatesInWorkspaceParams creates a new DeleteClusterTemplatesInWorkspaceParams object
// with the default values initialized.
func NewDeleteClusterTemplatesInWorkspaceParams() *DeleteClusterTemplatesInWorkspaceParams {
	var ()
	return &DeleteClusterTemplatesInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterTemplatesInWorkspaceParamsWithTimeout creates a new DeleteClusterTemplatesInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterTemplatesInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteClusterTemplatesInWorkspaceParams {
	var ()
	return &DeleteClusterTemplatesInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteClusterTemplatesInWorkspaceParamsWithContext creates a new DeleteClusterTemplatesInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterTemplatesInWorkspaceParamsWithContext(ctx context.Context) *DeleteClusterTemplatesInWorkspaceParams {
	var ()
	return &DeleteClusterTemplatesInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteClusterTemplatesInWorkspaceParamsWithHTTPClient creates a new DeleteClusterTemplatesInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterTemplatesInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteClusterTemplatesInWorkspaceParams {
	var ()
	return &DeleteClusterTemplatesInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteClusterTemplatesInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete cluster templates in workspace operation typically these are written to a http.Request
*/
type DeleteClusterTemplatesInWorkspaceParams struct {

	/*Body*/
	Body []string
	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*EnvironmentName*/
	EnvironmentName *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithContext(ctx context.Context) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithBody(body []string) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetBody(body []string) {
	o.Body = body
}

// WithEnvironmentCrn adds the environmentCrn to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithEnvironmentCrn(environmentCrn *string) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithEnvironmentName adds the environmentName to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithEnvironmentName(environmentName *string) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetEnvironmentName(environmentName)
	return o
}

// SetEnvironmentName adds the environmentName to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetEnvironmentName(environmentName *string) {
	o.EnvironmentName = environmentName
}

// WithWorkspaceID adds the workspaceID to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteClusterTemplatesInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete cluster templates in workspace params
func (o *DeleteClusterTemplatesInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterTemplatesInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
				return err
			}
		}

	}

	if o.EnvironmentName != nil {

		// query param environmentName
		var qrEnvironmentName string
		if o.EnvironmentName != nil {
			qrEnvironmentName = *o.EnvironmentName
		}
		qEnvironmentName := qrEnvironmentName
		if qEnvironmentName != "" {
			if err := r.SetQueryParam("environmentName", qEnvironmentName); err != nil {
				return err
			}
		}

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