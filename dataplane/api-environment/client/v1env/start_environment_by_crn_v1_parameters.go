// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

// NewStartEnvironmentByCrnV1Params creates a new StartEnvironmentByCrnV1Params object
// with the default values initialized.
func NewStartEnvironmentByCrnV1Params() *StartEnvironmentByCrnV1Params {
	var (
		dataHubStartActionDefault = string("START_ALL")
	)
	return &StartEnvironmentByCrnV1Params{
		DataHubStartAction: &dataHubStartActionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStartEnvironmentByCrnV1ParamsWithTimeout creates a new StartEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartEnvironmentByCrnV1ParamsWithTimeout(timeout time.Duration) *StartEnvironmentByCrnV1Params {
	var (
		dataHubStartActionDefault = string("START_ALL")
	)
	return &StartEnvironmentByCrnV1Params{
		DataHubStartAction: &dataHubStartActionDefault,

		timeout: timeout,
	}
}

// NewStartEnvironmentByCrnV1ParamsWithContext creates a new StartEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewStartEnvironmentByCrnV1ParamsWithContext(ctx context.Context) *StartEnvironmentByCrnV1Params {
	var (
		dataHubStartActionDefault = string("START_ALL")
	)
	return &StartEnvironmentByCrnV1Params{
		DataHubStartAction: &dataHubStartActionDefault,

		Context: ctx,
	}
}

// NewStartEnvironmentByCrnV1ParamsWithHTTPClient creates a new StartEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartEnvironmentByCrnV1ParamsWithHTTPClient(client *http.Client) *StartEnvironmentByCrnV1Params {
	var (
		dataHubStartActionDefault = string("START_ALL")
	)
	return &StartEnvironmentByCrnV1Params{
		DataHubStartAction: &dataHubStartActionDefault,
		HTTPClient:         client,
	}
}

/*StartEnvironmentByCrnV1Params contains all the parameters to send to the API endpoint
for the start environment by crn v1 operation typically these are written to a http.Request
*/
type StartEnvironmentByCrnV1Params struct {

	/*Crn*/
	Crn string
	/*DataHubStartAction*/
	DataHubStartAction *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) WithTimeout(timeout time.Duration) *StartEnvironmentByCrnV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) WithContext(ctx context.Context) *StartEnvironmentByCrnV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) WithHTTPClient(client *http.Client) *StartEnvironmentByCrnV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) WithCrn(crn string) *StartEnvironmentByCrnV1Params {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) SetCrn(crn string) {
	o.Crn = crn
}

// WithDataHubStartAction adds the dataHubStartAction to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) WithDataHubStartAction(dataHubStartAction *string) *StartEnvironmentByCrnV1Params {
	o.SetDataHubStartAction(dataHubStartAction)
	return o
}

// SetDataHubStartAction adds the dataHubStartAction to the start environment by crn v1 params
func (o *StartEnvironmentByCrnV1Params) SetDataHubStartAction(dataHubStartAction *string) {
	o.DataHubStartAction = dataHubStartAction
}

// WriteToRequest writes these params to a swagger request
func (o *StartEnvironmentByCrnV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if o.DataHubStartAction != nil {

		// query param dataHubStartAction
		var qrDataHubStartAction string
		if o.DataHubStartAction != nil {
			qrDataHubStartAction = *o.DataHubStartAction
		}
		qDataHubStartAction := qrDataHubStartAction
		if qDataHubStartAction != "" {
			if err := r.SetQueryParam("dataHubStartAction", qDataHubStartAction); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
