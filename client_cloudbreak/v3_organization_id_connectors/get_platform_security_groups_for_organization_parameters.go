// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_connectors

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewGetPlatformSecurityGroupsForOrganizationParams creates a new GetPlatformSecurityGroupsForOrganizationParams object
// with the default values initialized.
func NewGetPlatformSecurityGroupsForOrganizationParams() *GetPlatformSecurityGroupsForOrganizationParams {
	var ()
	return &GetPlatformSecurityGroupsForOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformSecurityGroupsForOrganizationParamsWithTimeout creates a new GetPlatformSecurityGroupsForOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformSecurityGroupsForOrganizationParamsWithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsForOrganizationParams {
	var ()
	return &GetPlatformSecurityGroupsForOrganizationParams{

		timeout: timeout,
	}
}

// NewGetPlatformSecurityGroupsForOrganizationParamsWithContext creates a new GetPlatformSecurityGroupsForOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformSecurityGroupsForOrganizationParamsWithContext(ctx context.Context) *GetPlatformSecurityGroupsForOrganizationParams {
	var ()
	return &GetPlatformSecurityGroupsForOrganizationParams{

		Context: ctx,
	}
}

// NewGetPlatformSecurityGroupsForOrganizationParamsWithHTTPClient creates a new GetPlatformSecurityGroupsForOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformSecurityGroupsForOrganizationParamsWithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsForOrganizationParams {
	var ()
	return &GetPlatformSecurityGroupsForOrganizationParams{
		HTTPClient: client,
	}
}

/*GetPlatformSecurityGroupsForOrganizationParams contains all the parameters to send to the API endpoint
for the get platform security groups for organization operation typically these are written to a http.Request
*/
type GetPlatformSecurityGroupsForOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) WithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsForOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) WithContext(ctx context.Context) *GetPlatformSecurityGroupsForOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) WithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsForOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetPlatformSecurityGroupsForOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) WithOrganizationID(organizationID int64) *GetPlatformSecurityGroupsForOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get platform security groups for organization params
func (o *GetPlatformSecurityGroupsForOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformSecurityGroupsForOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
