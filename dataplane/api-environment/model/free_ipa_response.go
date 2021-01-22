// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FreeIpaResponse free ipa response
// swagger:model FreeIpaResponse
type FreeIpaResponse struct {

	// Aws specific FreeIpa parameters
	Aws *AttachedFreeIpaRequestAwsParameters `json:"aws,omitempty"`

	// The number of FreeIPA instances to create per group when creating FreeIPA in environment
	InstanceCountByGroup int32 `json:"instanceCountByGroup,omitempty"`
}

// Validate validates this free ipa response
func (m *FreeIpaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FreeIpaResponse) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FreeIpaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeIpaResponse) UnmarshalBinary(b []byte) error {
	var res FreeIpaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
