package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ConstraintTemplateResponse constraint template response

swagger:model ConstraintTemplateResponse
*/
type ConstraintTemplateResponse struct {

	/* number of CPU cores needed for the Ambari node

	Required: true
	Read Only: true
	*/
	CPU float64 `json:"cpu"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* disk size needed for an Ambari node (GB)

	Required: true
	Read Only: true
	*/
	Disk float64 `json:"disk"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* memory needed for the Ambari container (GB)

	Required: true
	Read Only: true
	*/
	Memory float64 `json:"memory"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* type of orchestrator

	Required: true
	Read Only: true
	*/
	OrchestratorType string `json:"orchestratorType"`

	/* resource is visible in account

	Required: true
	Read Only: true
	*/
	PublicInAccount bool `json:"publicInAccount"`
}

// Validate validates this constraint template response
func (m *ConstraintTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrchestratorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicInAccount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConstraintTemplateResponse) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", float64(m.CPU)); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validateDisk(formats strfmt.Registry) error {

	if err := validate.Required("disk", "body", float64(m.Disk)); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", float64(m.Memory)); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validateOrchestratorType(formats strfmt.Registry) error {

	if err := validate.RequiredString("orchestratorType", "body", string(m.OrchestratorType)); err != nil {
		return err
	}

	return nil
}

func (m *ConstraintTemplateResponse) validatePublicInAccount(formats strfmt.Registry) error {

	if err := validate.Required("publicInAccount", "body", bool(m.PublicInAccount)); err != nil {
		return err
	}

	return nil
}
