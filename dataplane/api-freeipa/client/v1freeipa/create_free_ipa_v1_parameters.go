// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewCreateFreeIpaV1Params creates a new CreateFreeIpaV1Params object
// with the default values initialized.
func NewCreateFreeIpaV1Params() *CreateFreeIpaV1Params {
	var ()
	return &CreateFreeIpaV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFreeIpaV1ParamsWithTimeout creates a new CreateFreeIpaV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFreeIpaV1ParamsWithTimeout(timeout time.Duration) *CreateFreeIpaV1Params {
	var ()
	return &CreateFreeIpaV1Params{

		timeout: timeout,
	}
}

// NewCreateFreeIpaV1ParamsWithContext creates a new CreateFreeIpaV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFreeIpaV1ParamsWithContext(ctx context.Context) *CreateFreeIpaV1Params {
	var ()
	return &CreateFreeIpaV1Params{

		Context: ctx,
	}
}

// NewCreateFreeIpaV1ParamsWithHTTPClient creates a new CreateFreeIpaV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFreeIpaV1ParamsWithHTTPClient(client *http.Client) *CreateFreeIpaV1Params {
	var ()
	return &CreateFreeIpaV1Params{
		HTTPClient: client,
	}
}

/*CreateFreeIpaV1Params contains all the parameters to send to the API endpoint
for the create free ipa v1 operation typically these are written to a http.Request
*/
type CreateFreeIpaV1Params struct {

	/*Body*/
	Body *model.CreateFreeIpaV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) WithTimeout(timeout time.Duration) *CreateFreeIpaV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) WithContext(ctx context.Context) *CreateFreeIpaV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) WithHTTPClient(client *http.Client) *CreateFreeIpaV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) WithBody(body *model.CreateFreeIpaV1Request) *CreateFreeIpaV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create free ipa v1 params
func (o *CreateFreeIpaV1Params) SetBody(body *model.CreateFreeIpaV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFreeIpaV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
