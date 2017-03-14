package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetSecuritygroupsAccountParams creates a new GetSecuritygroupsAccountParams object
// with the default values initialized.
func NewGetSecuritygroupsAccountParams() *GetSecuritygroupsAccountParams {

	return &GetSecuritygroupsAccountParams{}
}

/*GetSecuritygroupsAccountParams contains all the parameters to send to the API endpoint
for the get securitygroups account operation typically these are written to a http.Request
*/
type GetSecuritygroupsAccountParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecuritygroupsAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}