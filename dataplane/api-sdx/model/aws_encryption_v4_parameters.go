// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AwsEncryptionV4Parameters aws encryption v4 parameters
// swagger:model AwsEncryptionV4Parameters
type AwsEncryptionV4Parameters struct {

	// encryption key for vm
	Key string `json:"key,omitempty"`

	// encryption type for vm (DEFAULT|CUSTOM|NONE)
	// Enum: [DEFAULT NONE CUSTOM]
	Type string `json:"type,omitempty"`
}

// Validate validates this aws encryption v4 parameters
func (m *AwsEncryptionV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsEncryptionV4ParametersTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","NONE","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsEncryptionV4ParametersTypeTypePropEnum = append(awsEncryptionV4ParametersTypeTypePropEnum, v)
	}
}

const (

	// AwsEncryptionV4ParametersTypeDEFAULT captures enum value "DEFAULT"
	AwsEncryptionV4ParametersTypeDEFAULT string = "DEFAULT"

	// AwsEncryptionV4ParametersTypeNONE captures enum value "NONE"
	AwsEncryptionV4ParametersTypeNONE string = "NONE"

	// AwsEncryptionV4ParametersTypeCUSTOM captures enum value "CUSTOM"
	AwsEncryptionV4ParametersTypeCUSTOM string = "CUSTOM"
)

// prop value enum
func (m *AwsEncryptionV4Parameters) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, awsEncryptionV4ParametersTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AwsEncryptionV4Parameters) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsEncryptionV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEncryptionV4Parameters) UnmarshalBinary(b []byte) error {
	var res AwsEncryptionV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}