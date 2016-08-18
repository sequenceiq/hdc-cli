package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetRdsconfigsAccountNameParams creates a new GetRdsconfigsAccountNameParams object
// with the default values initialized.
func NewGetRdsconfigsAccountNameParams() *GetRdsconfigsAccountNameParams {
	var ()
	return &GetRdsconfigsAccountNameParams{}
}

/*GetRdsconfigsAccountNameParams contains all the parameters to send to the API endpoint
for the get rdsconfigs account name operation typically these are written to a http.Request
*/
type GetRdsconfigsAccountNameParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the get rdsconfigs account name params
func (o *GetRdsconfigsAccountNameParams) WithName(name string) *GetRdsconfigsAccountNameParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRdsconfigsAccountNameParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
