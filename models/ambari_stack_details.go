package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*AmbariStackDetails ambari stack details

swagger:model AmbariStackDetails
*/
type AmbariStackDetails struct {

	/* operating system for the stack, like redhat6

	Required: true
	*/
	Os string `json:"os"`

	/* name of the stack, like HDP

	Required: true
	*/
	Stack string `json:"stack"`

	/* url of the stack repository

	Required: true
	*/
	StackBaseURL string `json:"stackBaseURL"`

	/* id of the stack repository

	Required: true
	*/
	StackRepoID string `json:"stackRepoId"`

	/* url of the stack utils repository

	Required: true
	*/
	UtilsBaseURL string `json:"utilsBaseURL"`

	/* id of the stack utils repository

	Required: true
	*/
	UtilsRepoID string `json:"utilsRepoId"`

	/* whether to verify or not the repo url

	Required: true
	*/
	Verify bool `json:"verify"`

	/* version of the stack

	Required: true
	*/
	Version string `json:"version"`
}

// Validate validates this ambari stack details
func (m *AmbariStackDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStack(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackBaseURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackRepoID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUtilsBaseURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUtilsRepoID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmbariStackDetails) validateOs(formats strfmt.Registry) error {

	if err := validate.RequiredString("os", "body", string(m.Os)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateStack(formats strfmt.Registry) error {

	if err := validate.RequiredString("stack", "body", string(m.Stack)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateStackBaseURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("stackBaseURL", "body", string(m.StackBaseURL)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateStackRepoID(formats strfmt.Registry) error {

	if err := validate.RequiredString("stackRepoId", "body", string(m.StackRepoID)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateUtilsBaseURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("utilsBaseURL", "body", string(m.UtilsBaseURL)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateUtilsRepoID(formats strfmt.Registry) error {

	if err := validate.RequiredString("utilsRepoId", "body", string(m.UtilsRepoID)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateVerify(formats strfmt.Registry) error {

	if err := validate.Required("verify", "body", bool(m.Verify)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("version", "body", string(m.Version)); err != nil {
		return err
	}

	return nil
}
