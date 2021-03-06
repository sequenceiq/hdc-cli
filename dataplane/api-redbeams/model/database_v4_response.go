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

// DatabaseV4Response Response containing information about a database that was acted upon, e.g., retrieved, deleted, listed
// swagger:model DatabaseV4Response
type DatabaseV4Response struct {

	// Name of the JDBC connection driver (for example: 'org.postgresql.Driver')
	// Required: true
	ConnectionDriver *string `json:"connectionDriver"`

	// Password to use for authentication
	ConnectionPassword *SecretResponse `json:"connectionPassword,omitempty"`

	// JDBC connection URL in the form of jdbc:<db-type>:<driver-specific-part>
	// Required: true
	ConnectionURL *string `json:"connectionURL"`

	// Username to use for authentication
	ConnectionUserName *SecretResponse `json:"connectionUserName,omitempty"`

	// Creation date / time of the database, in epoch milliseconds
	CreationDate int64 `json:"creationDate,omitempty"`

	// CRN of the database
	Crn string `json:"crn,omitempty"`

	// Name of the database vendor (MYSQL, POSTGRES...)
	// Required: true
	DatabaseEngine *string `json:"databaseEngine"`

	// Display name of the database vendor (MySQL, PostgreSQL, ...)
	// Required: true
	DatabaseEngineDisplayName *string `json:"databaseEngineDisplayName"`

	// Description of the database
	// Max Length: 1000000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// CRN of the environment of the database
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Name of the database
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Ownership status of the database
	// Enum: [UNKNOWN SERVICE_MANAGED USER_MANAGED]
	ResourceStatus string `json:"resourceStatus,omitempty"`

	// Type of database, i.e., the service name that will use the database (HIVE, DRUID, SUPERSET, RANGER, ...)
	// Required: true
	// Max Length: 56
	// Min Length: 3
	// Pattern: (^[a-zA-Z_][-a-zA-Z0-9_]*[a-zA-Z0-9_]$)
	Type *string `json:"type"`
}

// Validate validates this database v4 response
func (m *DatabaseV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseEngineDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceStatus(formats); err != nil {
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

func (m *DatabaseV4Response) validateConnectionDriver(formats strfmt.Registry) error {

	if err := validate.Required("connectionDriver", "body", m.ConnectionDriver); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionPassword) { // not required
		return nil
	}

	if m.ConnectionPassword != nil {
		if err := m.ConnectionPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionPassword")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionUserName) { // not required
		return nil
	}

	if m.ConnectionUserName != nil {
		if err := m.ConnectionUserName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionUserName")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseV4Response) validateDatabaseEngine(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngine", "body", m.DatabaseEngine); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateDatabaseEngineDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngineDisplayName", "body", m.DatabaseEngineDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000000); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateName(formats strfmt.Registry) error {

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

var databaseV4ResponseTypeResourceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","SERVICE_MANAGED","USER_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		databaseV4ResponseTypeResourceStatusPropEnum = append(databaseV4ResponseTypeResourceStatusPropEnum, v)
	}
}

const (

	// DatabaseV4ResponseResourceStatusUNKNOWN captures enum value "UNKNOWN"
	DatabaseV4ResponseResourceStatusUNKNOWN string = "UNKNOWN"

	// DatabaseV4ResponseResourceStatusSERVICEMANAGED captures enum value "SERVICE_MANAGED"
	DatabaseV4ResponseResourceStatusSERVICEMANAGED string = "SERVICE_MANAGED"

	// DatabaseV4ResponseResourceStatusUSERMANAGED captures enum value "USER_MANAGED"
	DatabaseV4ResponseResourceStatusUSERMANAGED string = "USER_MANAGED"
)

// prop value enum
func (m *DatabaseV4Response) validateResourceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, databaseV4ResponseTypeResourceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DatabaseV4Response) validateResourceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceStatusEnum("resourceStatus", "body", m.ResourceStatus); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 56); err != nil {
		return err
	}

	if err := validate.Pattern("type", "body", string(*m.Type), `(^[a-zA-Z_][-a-zA-Z0-9_]*[a-zA-Z0-9_]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseV4Response) UnmarshalBinary(b []byte) error {
	var res DatabaseV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
