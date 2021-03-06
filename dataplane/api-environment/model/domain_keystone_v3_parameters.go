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

// DomainKeystoneV3Parameters domain keystone v3 parameters
// swagger:model DomainKeystoneV3Parameters
type DomainKeystoneV3Parameters struct {
	KeystoneV3Base

	// domain name
	// Required: true
	DomainName *string `json:"domainName"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DomainKeystoneV3Parameters) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 KeystoneV3Base
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.KeystoneV3Base = aO0

	// AO1
	var dataAO1 struct {
		DomainName *string `json:"domainName"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DomainName = dataAO1.DomainName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DomainKeystoneV3Parameters) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.KeystoneV3Base)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DomainName *string `json:"domainName"`
	}

	dataAO1.DomainName = m.DomainName

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this domain keystone v3 parameters
func (m *DomainKeystoneV3Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with KeystoneV3Base
	if err := m.KeystoneV3Base.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainKeystoneV3Parameters) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domainName", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainKeystoneV3Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainKeystoneV3Parameters) UnmarshalBinary(b []byte) error {
	var res DomainKeystoneV3Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
