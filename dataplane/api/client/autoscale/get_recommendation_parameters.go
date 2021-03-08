// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewGetRecommendationParams creates a new GetRecommendationParams object
// with the default values initialized.
func NewGetRecommendationParams() *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecommendationParamsWithTimeout creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecommendationParamsWithTimeout(timeout time.Duration) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		timeout: timeout,
	}
}

// NewGetRecommendationParamsWithContext creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecommendationParamsWithContext(ctx context.Context) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		Context: ctx,
	}
}

// NewGetRecommendationParamsWithHTTPClient creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecommendationParamsWithHTTPClient(client *http.Client) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{
		HTTPClient: client,
	}
}

/*GetRecommendationParams contains all the parameters to send to the API endpoint
for the get recommendation operation typically these are written to a http.Request
*/
type GetRecommendationParams struct {

	/*BlueprintName*/
	BlueprintName *string
	/*WorkspaceID*/
	WorkspaceID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recommendation params
func (o *GetRecommendationParams) WithTimeout(timeout time.Duration) *GetRecommendationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recommendation params
func (o *GetRecommendationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recommendation params
func (o *GetRecommendationParams) WithContext(ctx context.Context) *GetRecommendationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recommendation params
func (o *GetRecommendationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recommendation params
func (o *GetRecommendationParams) WithHTTPClient(client *http.Client) *GetRecommendationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recommendation params
func (o *GetRecommendationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprintName adds the blueprintName to the get recommendation params
func (o *GetRecommendationParams) WithBlueprintName(blueprintName *string) *GetRecommendationParams {
	o.SetBlueprintName(blueprintName)
	return o
}

// SetBlueprintName adds the blueprintName to the get recommendation params
func (o *GetRecommendationParams) SetBlueprintName(blueprintName *string) {
	o.BlueprintName = blueprintName
}

// WithWorkspaceID adds the workspaceID to the get recommendation params
func (o *GetRecommendationParams) WithWorkspaceID(workspaceID *int64) *GetRecommendationParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get recommendation params
func (o *GetRecommendationParams) SetWorkspaceID(workspaceID *int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecommendationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BlueprintName != nil {

		// query param blueprintName
		var qrBlueprintName string
		if o.BlueprintName != nil {
			qrBlueprintName = *o.BlueprintName
		}
		qBlueprintName := qrBlueprintName
		if qBlueprintName != "" {
			if err := r.SetQueryParam("blueprintName", qBlueprintName); err != nil {
				return err
			}
		}

	}

	if o.WorkspaceID != nil {

		// query param workspaceId
		var qrWorkspaceID int64
		if o.WorkspaceID != nil {
			qrWorkspaceID = *o.WorkspaceID
		}
		qWorkspaceID := swag.FormatInt64(qrWorkspaceID)
		if qWorkspaceID != "" {
			if err := r.SetQueryParam("workspaceId", qWorkspaceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}