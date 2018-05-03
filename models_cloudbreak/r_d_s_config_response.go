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

// RDSConfigResponse r d s config response
// swagger:model RDSConfigResponse

type RDSConfigResponse struct {

	// list of clusters which use config
	// Unique: true
	ClusterNames []string `json:"clusterNames"`

	// Name of the JDBC connection driver (for example: 'org.postgresql.Driver')
	// Required: true
	ConnectionDriver *string `json:"connectionDriver"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	// Pattern: ^jdbc:(postgresql|mysql|oracle):(thin:@|//)[-\w\.]*:\d{1,5}/?:?\w*
	ConnectionURL *string `json:"connectionURL"`

	// URL that points to the jar of the connection driver(connector)
	// Max Length: 150
	// Min Length: 0
	// Pattern: ^http[s]?://[\w-/?=+&:,#.]*
	ConnectorJarURL *string `json:"connectorJarUrl,omitempty"`

	// creation time of the resource in long
	CreationDate int64 `json:"creationDate,omitempty"`

	// Name of the external database engine (MYSQL, POSTGRES...)
	// Required: true
	DatabaseEngine *string `json:"databaseEngine"`

	// Display name of the external database engine (Mysql, Postges...)
	// Required: true
	DatabaseEngineDisplayName *string `json:"databaseEngineDisplayName"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Name of the RDS configuration resource
	// Required: true
	// Max Length: 50
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// resource is visible in account
	PublicInAccount *bool `json:"publicInAccount,omitempty"`

	// (HDP, HDF)Stack version for the RDS configuration
	StackVersion string `json:"stackVersion,omitempty"`

	// Type of RDS, aka the service name that will use the RDS like HIVE, DRUID, SUPERSET, RANGER, etc.
	// Required: true
	// Max Length: 12
	// Min Length: 3
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)
	Type *string `json:"type"`
}

/* polymorph RDSConfigResponse clusterNames false */

/* polymorph RDSConfigResponse connectionDriver false */

/* polymorph RDSConfigResponse connectionURL false */

/* polymorph RDSConfigResponse connectorJarUrl false */

/* polymorph RDSConfigResponse creationDate false */

/* polymorph RDSConfigResponse databaseEngine false */

/* polymorph RDSConfigResponse databaseEngineDisplayName false */

/* polymorph RDSConfigResponse id false */

/* polymorph RDSConfigResponse name false */

/* polymorph RDSConfigResponse publicInAccount false */

/* polymorph RDSConfigResponse stackVersion false */

/* polymorph RDSConfigResponse type false */

// Validate validates this r d s config response
func (m *RDSConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectorJarURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseEngine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseEngineDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSConfigResponse) validateClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterNames", "body", m.ClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionDriver(formats strfmt.Registry) error {

	if err := validate.Required("connectionDriver", "body", m.ConnectionDriver); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	if err := validate.Pattern("connectionURL", "body", string(*m.ConnectionURL), `^jdbc:(postgresql|mysql|oracle):(thin:@|//)[-\w\.]*:\d{1,5}/?:?\w*`); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectorJarURL(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorJarURL) { // not required
		return nil
	}

	if err := validate.MinLength("connectorJarUrl", "body", string(*m.ConnectorJarURL), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("connectorJarUrl", "body", string(*m.ConnectorJarURL), 150); err != nil {
		return err
	}

	if err := validate.Pattern("connectorJarUrl", "body", string(*m.ConnectorJarURL), `^http[s]?://[\w-/?=+&:,#.]*`); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateDatabaseEngine(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngine", "body", m.DatabaseEngine); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateDatabaseEngineDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngineDisplayName", "body", m.DatabaseEngineDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 12); err != nil {
		return err
	}

	if err := validate.Pattern("type", "body", string(*m.Type), `(^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RDSConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDSConfigResponse) UnmarshalBinary(b []byte) error {
	var res RDSConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
