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

// ProxyViewResponse Cloudbreak allows you to save your existing proxy configuration information as an external source so that you can provide the proxy information to multiple clusters that you create with Cloudbreak
// swagger:model ProxyViewResponse
type ProxyViewResponse struct {

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// proxy configuration id for the cluster
	Crn string `json:"crn,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// host or IP address of proxy server
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Pattern: (^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)
	Host *string `json:"host"`

	// Name of the proxy configuration resource
	// Required: true
	// Max Length: 100
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// port of proxy server (typically: 3128 or 8080)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// determines the protocol (http or https)
	// Required: true
	// Pattern: ^http(s)?$
	Protocol *string `json:"protocol"`
}

// Validate validates this proxy view response
func (m *ProxyViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyViewResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ProxyViewResponse) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if err := validate.MinLength("host", "body", string(*m.Host), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("host", "body", string(*m.Host), 255); err != nil {
		return err
	}

	if err := validate.Pattern("host", "body", string(*m.Host), `(^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$)`); err != nil {
		return err
	}

	return nil
}

func (m *ProxyViewResponse) validateName(formats strfmt.Registry) error {

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

func (m *ProxyViewResponse) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *ProxyViewResponse) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if err := validate.Pattern("protocol", "body", string(*m.Protocol), `^http(s)?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyViewResponse) UnmarshalBinary(b []byte) error {
	var res ProxyViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
