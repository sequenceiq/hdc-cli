// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LocationV1Request location v1 request
// swagger:model LocationV1Request
type LocationV1Request struct {

	// Location latitude of the environment.
	Latitude float64 `json:"latitude,omitempty"`

	// Location longitude of the environment.
	Longitude float64 `json:"longitude,omitempty"`

	// Location of the environment.
	// Required: true
	// Max Length: 100
	// Min Length: 0
	Name *string `json:"name"`
}

// Validate validates this location v1 request
func (m *LocationV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationV1Request) UnmarshalBinary(b []byte) error {
	var res LocationV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}