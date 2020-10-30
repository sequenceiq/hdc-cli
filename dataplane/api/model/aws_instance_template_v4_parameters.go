// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AwsInstanceTemplateV4Parameters aws instance template v4 parameters
// swagger:model AwsInstanceTemplateV4Parameters
type AwsInstanceTemplateV4Parameters struct {

	// encryption for vm
	Encryption *AwsEncryptionV4Parameters `json:"encryption,omitempty"`

	// aws specific placement group parameter for vm
	PlacementGroup *AwsPlacementGroupV4Parameters `json:"placementGroup,omitempty"`

	// aws specific spot instance parameters for template
	Spot *AwsInstanceTemplateV4SpotParameters `json:"spot,omitempty"`
}

// Validate validates this aws instance template v4 parameters
func (m *AwsInstanceTemplateV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacementGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsInstanceTemplateV4Parameters) validateEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *AwsInstanceTemplateV4Parameters) validatePlacementGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.PlacementGroup) { // not required
		return nil
	}

	if m.PlacementGroup != nil {
		if err := m.PlacementGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placementGroup")
			}
			return err
		}
	}

	return nil
}

func (m *AwsInstanceTemplateV4Parameters) validateSpot(formats strfmt.Registry) error {

	if swag.IsZero(m.Spot) { // not required
		return nil
	}

	if m.Spot != nil {
		if err := m.Spot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsInstanceTemplateV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsInstanceTemplateV4Parameters) UnmarshalBinary(b []byte) error {
	var res AwsInstanceTemplateV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
