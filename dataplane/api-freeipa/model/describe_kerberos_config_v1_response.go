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

// DescribeKerberosConfigV1Response describe kerberos config v1 response
// swagger:model DescribeKerberosConfigV1Response
type DescribeKerberosConfigV1Response struct {

	// kerberos admin server URL
	AdminURL string `json:"adminUrl,omitempty"`

	// container dn
	ContainerDn string `json:"containerDn,omitempty"`

	// crn
	Crn string `json:"crn,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Ambari kerberos descriptor
	Descriptor *SecretResponse `json:"descriptor,omitempty"`

	// cluster instances will set this as the domain part of their hostname
	Domain string `json:"domain,omitempty"`

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Ambari kerberos krb5.conf template
	Krb5Conf *SecretResponse `json:"krb5Conf,omitempty"`

	// ldap Url
	LdapURL string `json:"ldapUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// comma separated list of nameservers' IP address which will be used by cluster instances
	// Pattern: (^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(,((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))*$)
	NameServers string `json:"nameServers,omitempty"`

	// kerberos admin password
	Password *SecretResponse `json:"password,omitempty"`

	// kerberos principal
	Principal *SecretResponse `json:"principal,omitempty"`

	// realm
	Realm string `json:"realm,omitempty"`

	// tcp allowed
	TCPAllowed bool `json:"tcpAllowed,omitempty"`

	// type
	// Required: true
	// Enum: [ACTIVE_DIRECTORY MIT FREEIPA]
	Type *string `json:"type"`

	// kerberos KDC server URL
	URL string `json:"url,omitempty"`

	// Allows to select either a trusting SSL connection or a validating (non-trusting) SSL connection to KDC
	VerifyKdcTrust bool `json:"verifyKdcTrust,omitempty"`
}

// Validate validates this describe kerberos config v1 response
func (m *DescribeKerberosConfigV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKrb5Conf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipal(formats); err != nil {
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

func (m *DescribeKerberosConfigV1Response) validateDescription(formats strfmt.Registry) error {

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

func (m *DescribeKerberosConfigV1Response) validateDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.Descriptor) { // not required
		return nil
	}

	if m.Descriptor != nil {
		if err := m.Descriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeKerberosConfigV1Response) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *DescribeKerberosConfigV1Response) validateKrb5Conf(formats strfmt.Registry) error {

	if swag.IsZero(m.Krb5Conf) { // not required
		return nil
	}

	if m.Krb5Conf != nil {
		if err := m.Krb5Conf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("krb5Conf")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeKerberosConfigV1Response) validateNameServers(formats strfmt.Registry) error {

	if swag.IsZero(m.NameServers) { // not required
		return nil
	}

	if err := validate.Pattern("nameServers", "body", string(m.NameServers), `(^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(,((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))*$)`); err != nil {
		return err
	}

	return nil
}

func (m *DescribeKerberosConfigV1Response) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if m.Password != nil {
		if err := m.Password.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeKerberosConfigV1Response) validatePrincipal(formats strfmt.Registry) error {

	if swag.IsZero(m.Principal) { // not required
		return nil
	}

	if m.Principal != nil {
		if err := m.Principal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

var describeKerberosConfigV1ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE_DIRECTORY","MIT","FREEIPA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		describeKerberosConfigV1ResponseTypeTypePropEnum = append(describeKerberosConfigV1ResponseTypeTypePropEnum, v)
	}
}

const (

	// DescribeKerberosConfigV1ResponseTypeACTIVEDIRECTORY captures enum value "ACTIVE_DIRECTORY"
	DescribeKerberosConfigV1ResponseTypeACTIVEDIRECTORY string = "ACTIVE_DIRECTORY"

	// DescribeKerberosConfigV1ResponseTypeMIT captures enum value "MIT"
	DescribeKerberosConfigV1ResponseTypeMIT string = "MIT"

	// DescribeKerberosConfigV1ResponseTypeFREEIPA captures enum value "FREEIPA"
	DescribeKerberosConfigV1ResponseTypeFREEIPA string = "FREEIPA"
)

// prop value enum
func (m *DescribeKerberosConfigV1Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, describeKerberosConfigV1ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DescribeKerberosConfigV1Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeKerberosConfigV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeKerberosConfigV1Response) UnmarshalBinary(b []byte) error {
	var res DescribeKerberosConfigV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}