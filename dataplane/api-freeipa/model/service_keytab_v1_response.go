// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceKeytabV1Response service keytab v1 response
// swagger:model ServiceKeytabV1Response
type ServiceKeytabV1Response struct {

	// Keytab that was requested
	Keytab *SecretResponse `json:"keytab,omitempty"`

	// Kerberos Principal Name
	ServicePrincipal *SecretResponse `json:"servicePrincipal,omitempty"`
}

// Validate validates this service keytab v1 response
func (m *ServiceKeytabV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeytab(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceKeytabV1Response) validateKeytab(formats strfmt.Registry) error {

	if swag.IsZero(m.Keytab) { // not required
		return nil
	}

	if m.Keytab != nil {
		if err := m.Keytab.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keytab")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceKeytabV1Response) validateServicePrincipal(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePrincipal) { // not required
		return nil
	}

	if m.ServicePrincipal != nil {
		if err := m.ServicePrincipal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePrincipal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceKeytabV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceKeytabV1Response) UnmarshalBinary(b []byte) error {
	var res ServiceKeytabV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
