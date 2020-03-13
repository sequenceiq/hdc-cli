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

// DatabaseV4Request database v4 request
// swagger:model DatabaseV4Request
type DatabaseV4Request struct {

	// Password to use for the jdbc connection
	// Required: true
	ConnectionPassword *string `json:"connectionPassword"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	ConnectionURL *string `json:"connectionURL"`

	// Username to use for the jdbc connection
	// Required: true
	ConnectionUserName *string `json:"connectionUserName"`

	// URL that points to the jar of the connection driver(connector)
	ConnectorJarURL string `json:"connectorJarUrl,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Name of the RDS configuration resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Oracle specific properties
	Oracle *OracleParameters `json:"oracle,omitempty"`

	// Type of RDS, aka the service name that will use the RDS like HIVE, DRUID, SUPERSET, RANGER, etc.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this database v4 request
func (m *DatabaseV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseV4Request) validateConnectionPassword(formats strfmt.Registry) error {

	if err := validate.Required("connectionPassword", "body", m.ConnectionPassword); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Request) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Request) validateConnectionUserName(formats strfmt.Registry) error {

	if err := validate.Required("connectionUserName", "body", m.ConnectionUserName); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Request) validateDescription(formats strfmt.Registry) error {

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

func (m *DatabaseV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
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

func (m *DatabaseV4Request) validateOracle(formats strfmt.Registry) error {

	if swag.IsZero(m.Oracle) { // not required
		return nil
	}

	if m.Oracle != nil {
		if err := m.Oracle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseV4Request) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseV4Request) UnmarshalBinary(b []byte) error {
	var res DatabaseV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}