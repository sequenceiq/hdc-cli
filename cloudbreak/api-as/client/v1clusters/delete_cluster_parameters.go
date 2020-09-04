// Code generated by go-swagger; DO NOT EDIT.

package v1clusters

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
)

// NewDeleteClusterParams creates a new DeleteClusterParams object
// with the default values initialized.
func NewDeleteClusterParams() *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterParamsWithTimeout creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterParamsWithTimeout(timeout time.Duration) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		timeout: timeout,
	}
}

// NewDeleteClusterParamsWithContext creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterParamsWithContext(ctx context.Context) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{

		Context: ctx,
	}
}

// NewDeleteClusterParamsWithHTTPClient creates a new DeleteClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterParamsWithHTTPClient(client *http.Client) *DeleteClusterParams {
	var ()
	return &DeleteClusterParams{
		HTTPClient: client,
	}
}

/*DeleteClusterParams contains all the parameters to send to the API endpoint
for the delete cluster operation typically these are written to a http.Request
*/
type DeleteClusterParams struct {

	/*ClusterID*/
	ClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) WithTimeout(timeout time.Duration) *DeleteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster params
func (o *DeleteClusterParams) WithContext(ctx context.Context) *DeleteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster params
func (o *DeleteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) WithHTTPClient(client *http.Client) *DeleteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster params
func (o *DeleteClusterParams) WithClusterID(clusterID int64) *DeleteClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster params
func (o *DeleteClusterParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
