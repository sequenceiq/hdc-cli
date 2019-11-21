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

// FreeIpaServerV1Request free ipa server v1 request
// swagger:model FreeIpaServerV1Request
type FreeIpaServerV1Request struct {

	// Name of the admin group to be used for all the services.
	AdminGroupName string `json:"adminGroupName,omitempty"`

	// admin password
	// Required: true
	AdminPassword *string `json:"adminPassword"`

	// Domain name associated to the FreeIPA
	// Required: true
	// Pattern: (?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]
	Domain *string `json:"domain"`

	// Base hostname for FreeIPA servers
	// Required: true
	// Pattern: ^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$
	Hostname *string `json:"hostname"`
}

// Validate validates this free ipa server v1 request
func (m *FreeIpaServerV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FreeIpaServerV1Request) validateAdminPassword(formats strfmt.Registry) error {

	if err := validate.Required("adminPassword", "body", m.AdminPassword); err != nil {
		return err
	}

	return nil
}

func (m *FreeIpaServerV1Request) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	if err := validate.Pattern("domain", "body", string(*m.Domain), `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`); err != nil {
		return err
	}

	return nil
}

func (m *FreeIpaServerV1Request) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	if err := validate.Pattern("hostname", "body", string(*m.Hostname), `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FreeIpaServerV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeIpaServerV1Request) UnmarshalBinary(b []byte) error {
	var res FreeIpaServerV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
