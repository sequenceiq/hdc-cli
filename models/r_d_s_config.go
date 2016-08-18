package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*RDSConfig r d s config

swagger:model RDSConfig
*/
type RDSConfig struct {

	/* Password to use for the jdbc connection

	Required: true
	*/
	ConnectionPassword string `json:"connectionPassword"`

	/* JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>

	Required: true
	*/
	ConnectionURL string `json:"connectionURL"`

	/* Username to use for the jdbc connection

	Required: true
	*/
	ConnectionUserName string `json:"connectionUserName"`

	/* Type of the external database (allowed values: MYSQL, POSTGRES)

	Required: true
	*/
	DatabaseType string `json:"databaseType"`

	/* HDP version for the RDS configuration

	Required: true
	*/
	HdpVersion string `json:"hdpVersion"`

	/* Name of the RDS configuration resource

	Required: true
	*/
	Name string `json:"name"`

	/* If true, then the RDS configuration will be validated
	 */
	Validated *bool `json:"validated,omitempty"`
}

// Validate validates this r d s config
func (m *RDSConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHdpVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSConfig) validateConnectionPassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionPassword", "body", string(m.ConnectionPassword)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfig) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionURL", "body", string(m.ConnectionURL)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfig) validateConnectionUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionUserName", "body", string(m.ConnectionUserName)); err != nil {
		return err
	}

	return nil
}

var rDSConfigTypeDatabaseTypePropEnum []interface{}

func (m *RDSConfig) validateDatabaseTypeEnum(path, location string, value string) error {
	if rDSConfigTypeDatabaseTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["POSTGRES","MYSQL"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			rDSConfigTypeDatabaseTypePropEnum = append(rDSConfigTypeDatabaseTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, rDSConfigTypeDatabaseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RDSConfig) validateDatabaseType(formats strfmt.Registry) error {

	if err := validate.RequiredString("databaseType", "body", string(m.DatabaseType)); err != nil {
		return err
	}

	if err := m.validateDatabaseTypeEnum("databaseType", "body", m.DatabaseType); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfig) validateHdpVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("hdpVersion", "body", string(m.HdpVersion)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfig) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}
