package sssd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/sequenceiq/hdc-cli/models"
)

// NewPostSssdUserParams creates a new PostSssdUserParams object
// with the default values initialized.
func NewPostSssdUserParams() *PostSssdUserParams {
	var ()
	return &PostSssdUserParams{}
}

/*PostSssdUserParams contains all the parameters to send to the API endpoint
for the post sssd user operation typically these are written to a http.Request
*/
type PostSssdUserParams struct {

	/*Body*/
	Body *models.SssdConfigRequest
}

// WithBody adds the body to the post sssd user params
func (o *PostSssdUserParams) WithBody(body *models.SssdConfigRequest) *PostSssdUserParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostSssdUserParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.SssdConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
