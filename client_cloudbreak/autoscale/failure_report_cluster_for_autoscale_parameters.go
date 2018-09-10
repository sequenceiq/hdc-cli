// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewFailureReportClusterForAutoscaleParams creates a new FailureReportClusterForAutoscaleParams object
// with the default values initialized.
func NewFailureReportClusterForAutoscaleParams() *FailureReportClusterForAutoscaleParams {
	var ()
	return &FailureReportClusterForAutoscaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFailureReportClusterForAutoscaleParamsWithTimeout creates a new FailureReportClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFailureReportClusterForAutoscaleParamsWithTimeout(timeout time.Duration) *FailureReportClusterForAutoscaleParams {
	var ()
	return &FailureReportClusterForAutoscaleParams{

		timeout: timeout,
	}
}

// NewFailureReportClusterForAutoscaleParamsWithContext creates a new FailureReportClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewFailureReportClusterForAutoscaleParamsWithContext(ctx context.Context) *FailureReportClusterForAutoscaleParams {
	var ()
	return &FailureReportClusterForAutoscaleParams{

		Context: ctx,
	}
}

// NewFailureReportClusterForAutoscaleParamsWithHTTPClient creates a new FailureReportClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFailureReportClusterForAutoscaleParamsWithHTTPClient(client *http.Client) *FailureReportClusterForAutoscaleParams {
	var ()
	return &FailureReportClusterForAutoscaleParams{
		HTTPClient: client,
	}
}

/*FailureReportClusterForAutoscaleParams contains all the parameters to send to the API endpoint
for the failure report cluster for autoscale operation typically these are written to a http.Request
*/
type FailureReportClusterForAutoscaleParams struct {

	/*Body*/
	Body *models_cloudbreak.FailureReport
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) WithTimeout(timeout time.Duration) *FailureReportClusterForAutoscaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) WithContext(ctx context.Context) *FailureReportClusterForAutoscaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) WithHTTPClient(client *http.Client) *FailureReportClusterForAutoscaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) WithBody(body *models_cloudbreak.FailureReport) *FailureReportClusterForAutoscaleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) SetBody(body *models_cloudbreak.FailureReport) {
	o.Body = body
}

// WithID adds the id to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) WithID(id int64) *FailureReportClusterForAutoscaleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the failure report cluster for autoscale params
func (o *FailureReportClusterForAutoscaleParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FailureReportClusterForAutoscaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.FailureReport)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
