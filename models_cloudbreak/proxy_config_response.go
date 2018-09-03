// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProxyConfigResponse proxy config response
// swagger:model ProxyConfigResponse

type ProxyConfigResponse struct {

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// proxy configuration id for the cluster
	ID int64 `json:"id,omitempty"`

	// Name of the proxy configuration resource
	// Required: true
	// Max Length: 100
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// organization of the resource
	Organization *OrganizationResourceResponse `json:"organization,omitempty"`

	// determines the protocol (http or https)
	// Required: true
	// Pattern: ^http(s)?$
	Protocol *string `json:"protocol"`

	// host or IP address of proxy server
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Pattern: (^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)
	ServerHost *string `json:"serverHost"`

	// port of proxy server (typically: 3128 or 8080)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	ServerPort *int32 `json:"serverPort"`

	// Username to use for basic authentication
	UserName string `json:"userName,omitempty"`
}

/* polymorph ProxyConfigResponse description false */

/* polymorph ProxyConfigResponse id false */

/* polymorph ProxyConfigResponse name false */

/* polymorph ProxyConfigResponse organization false */

/* polymorph ProxyConfigResponse protocol false */

/* polymorph ProxyConfigResponse serverHost false */

/* polymorph ProxyConfigResponse serverPort false */

/* polymorph ProxyConfigResponse userName false */

// Validate validates this proxy config response
func (m *ProxyConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerPort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyConfigResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ProxyConfigResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyConfigResponse) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {

		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *ProxyConfigResponse) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if err := validate.Pattern("protocol", "body", string(*m.Protocol), `^http(s)?$`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyConfigResponse) validateServerHost(formats strfmt.Registry) error {

	if err := validate.Required("serverHost", "body", m.ServerHost); err != nil {
		return err
	}

	if err := validate.MinLength("serverHost", "body", string(*m.ServerHost), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("serverHost", "body", string(*m.ServerHost), 255); err != nil {
		return err
	}

	if err := validate.Pattern("serverHost", "body", string(*m.ServerHost), `(^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyConfigResponse) validateServerPort(formats strfmt.Registry) error {

	if err := validate.Required("serverPort", "body", m.ServerPort); err != nil {
		return err
	}

	if err := validate.MinimumInt("serverPort", "body", int64(*m.ServerPort), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("serverPort", "body", int64(*m.ServerPort), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyConfigResponse) UnmarshalBinary(b []byte) error {
	var res ProxyConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
