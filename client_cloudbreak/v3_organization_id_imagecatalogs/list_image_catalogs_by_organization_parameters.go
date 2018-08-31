// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_imagecatalogs

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

// NewListImageCatalogsByOrganizationParams creates a new ListImageCatalogsByOrganizationParams object
// with the default values initialized.
func NewListImageCatalogsByOrganizationParams() *ListImageCatalogsByOrganizationParams {
	var ()
	return &ListImageCatalogsByOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListImageCatalogsByOrganizationParamsWithTimeout creates a new ListImageCatalogsByOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListImageCatalogsByOrganizationParamsWithTimeout(timeout time.Duration) *ListImageCatalogsByOrganizationParams {
	var ()
	return &ListImageCatalogsByOrganizationParams{

		timeout: timeout,
	}
}

// NewListImageCatalogsByOrganizationParamsWithContext creates a new ListImageCatalogsByOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewListImageCatalogsByOrganizationParamsWithContext(ctx context.Context) *ListImageCatalogsByOrganizationParams {
	var ()
	return &ListImageCatalogsByOrganizationParams{

		Context: ctx,
	}
}

// NewListImageCatalogsByOrganizationParamsWithHTTPClient creates a new ListImageCatalogsByOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListImageCatalogsByOrganizationParamsWithHTTPClient(client *http.Client) *ListImageCatalogsByOrganizationParams {
	var ()
	return &ListImageCatalogsByOrganizationParams{
		HTTPClient: client,
	}
}

/*ListImageCatalogsByOrganizationParams contains all the parameters to send to the API endpoint
for the list image catalogs by organization operation typically these are written to a http.Request
*/
type ListImageCatalogsByOrganizationParams struct {

	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) WithTimeout(timeout time.Duration) *ListImageCatalogsByOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) WithContext(ctx context.Context) *ListImageCatalogsByOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) WithHTTPClient(client *http.Client) *ListImageCatalogsByOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) WithOrganizationID(organizationID int64) *ListImageCatalogsByOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list image catalogs by organization params
func (o *ListImageCatalogsByOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListImageCatalogsByOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}