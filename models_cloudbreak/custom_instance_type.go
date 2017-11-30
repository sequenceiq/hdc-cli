// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustomInstanceType custom instance type
// swagger:model CustomInstanceType

type CustomInstanceType struct {

	// cpus
	Cpus int32 `json:"cpus,omitempty"`

	// memory
	Memory int32 `json:"memory,omitempty"`
}

/* polymorph CustomInstanceType cpus false */

/* polymorph CustomInstanceType memory false */

// Validate validates this custom instance type
func (m *CustomInstanceType) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CustomInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomInstanceType) UnmarshalBinary(b []byte) error {
	var res CustomInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
