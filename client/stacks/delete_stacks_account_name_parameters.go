package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteStacksAccountNameParams creates a new DeleteStacksAccountNameParams object
// with the default values initialized.
func NewDeleteStacksAccountNameParams() *DeleteStacksAccountNameParams {
	var (
		forcedDefault bool = bool(false)
	)
	return &DeleteStacksAccountNameParams{
		Forced: &forcedDefault,
	}
}

/*DeleteStacksAccountNameParams contains all the parameters to send to the API endpoint
for the delete stacks account name operation typically these are written to a http.Request
*/
type DeleteStacksAccountNameParams struct {

	/*Forced*/
	Forced *bool
	/*Name*/
	Name string
}

// WithForced adds the forced to the delete stacks account name params
func (o *DeleteStacksAccountNameParams) WithForced(forced *bool) *DeleteStacksAccountNameParams {
	o.Forced = forced
	return o
}

// WithName adds the name to the delete stacks account name params
func (o *DeleteStacksAccountNameParams) WithName(name string) *DeleteStacksAccountNameParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStacksAccountNameParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
