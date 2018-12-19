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

// SecurityRuleResponse security rule response
// swagger:model SecurityRuleResponse
type SecurityRuleResponse struct {

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// flag for making the rule modifiable
	Modifiable *bool `json:"modifiable,omitempty"`

	// comma separated list of accessible ports
	// Required: true
	// Pattern: ^[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1}(,[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1})*$
	Ports *string `json:"ports"`

	// protocol of the rule
	// Required: true
	Protocol *string `json:"protocol"`

	// definition of allowed subnet in CIDR format
	// Required: true
	// Pattern: ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$
	Subnet *string `json:"subnet"`
}

// Validate validates this security rule response
func (m *SecurityRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityRuleResponse) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	if err := validate.Pattern("ports", "body", string(*m.Ports), `^[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1}(,[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1})*$`); err != nil {
		return err
	}

	return nil
}

func (m *SecurityRuleResponse) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *SecurityRuleResponse) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	if err := validate.Pattern("subnet", "body", string(*m.Subnet), `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityRuleResponse) UnmarshalBinary(b []byte) error {
	var res SecurityRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}