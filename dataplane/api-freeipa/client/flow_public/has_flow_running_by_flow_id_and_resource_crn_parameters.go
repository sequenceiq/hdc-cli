// Code generated by go-swagger; DO NOT EDIT.

package flow_public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHasFlowRunningByFlowIDAndResourceCrnParams creates a new HasFlowRunningByFlowIDAndResourceCrnParams object
// with the default values initialized.
func NewHasFlowRunningByFlowIDAndResourceCrnParams() *HasFlowRunningByFlowIDAndResourceCrnParams {
	var ()
	return &HasFlowRunningByFlowIDAndResourceCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHasFlowRunningByFlowIDAndResourceCrnParamsWithTimeout creates a new HasFlowRunningByFlowIDAndResourceCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHasFlowRunningByFlowIDAndResourceCrnParamsWithTimeout(timeout time.Duration) *HasFlowRunningByFlowIDAndResourceCrnParams {
	var ()
	return &HasFlowRunningByFlowIDAndResourceCrnParams{

		timeout: timeout,
	}
}

// NewHasFlowRunningByFlowIDAndResourceCrnParamsWithContext creates a new HasFlowRunningByFlowIDAndResourceCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewHasFlowRunningByFlowIDAndResourceCrnParamsWithContext(ctx context.Context) *HasFlowRunningByFlowIDAndResourceCrnParams {
	var ()
	return &HasFlowRunningByFlowIDAndResourceCrnParams{

		Context: ctx,
	}
}

// NewHasFlowRunningByFlowIDAndResourceCrnParamsWithHTTPClient creates a new HasFlowRunningByFlowIDAndResourceCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHasFlowRunningByFlowIDAndResourceCrnParamsWithHTTPClient(client *http.Client) *HasFlowRunningByFlowIDAndResourceCrnParams {
	var ()
	return &HasFlowRunningByFlowIDAndResourceCrnParams{
		HTTPClient: client,
	}
}

/*HasFlowRunningByFlowIDAndResourceCrnParams contains all the parameters to send to the API endpoint
for the has flow running by flow Id and resource crn operation typically these are written to a http.Request
*/
type HasFlowRunningByFlowIDAndResourceCrnParams struct {

	/*FlowID*/
	FlowID string
	/*ResourceCrn*/
	ResourceCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WithTimeout(timeout time.Duration) *HasFlowRunningByFlowIDAndResourceCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WithContext(ctx context.Context) *HasFlowRunningByFlowIDAndResourceCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WithHTTPClient(client *http.Client) *HasFlowRunningByFlowIDAndResourceCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WithFlowID(flowID string) *HasFlowRunningByFlowIDAndResourceCrnParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WithResourceCrn adds the resourceCrn to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WithResourceCrn(resourceCrn *string) *HasFlowRunningByFlowIDAndResourceCrnParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the has flow running by flow Id and resource crn params
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) SetResourceCrn(resourceCrn *string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *HasFlowRunningByFlowIDAndResourceCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if o.ResourceCrn != nil {

		// query param resourceCrn
		var qrResourceCrn string
		if o.ResourceCrn != nil {
			qrResourceCrn = *o.ResourceCrn
		}
		qResourceCrn := qrResourceCrn
		if qResourceCrn != "" {
			if err := r.SetQueryParam("resourceCrn", qResourceCrn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}