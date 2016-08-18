package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewHasResourcesParams creates a new HasResourcesParams object
// with the default values initialized.
func NewHasResourcesParams() *HasResourcesParams {
	var ()
	return &HasResourcesParams{}
}

/*HasResourcesParams contains all the parameters to send to the API endpoint
for the has resources operation typically these are written to a http.Request
*/
type HasResourcesParams struct {

	/*ID*/
	ID string
}

// WithID adds the id to the has resources params
func (o *HasResourcesParams) WithID(id string) *HasResourcesParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *HasResourcesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
