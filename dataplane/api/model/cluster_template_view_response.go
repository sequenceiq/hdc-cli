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

// ClusterTemplateViewResponse cluster template view response
// swagger:model ClusterTemplateViewResponse
type ClusterTemplateViewResponse struct {

	// cloudplatform which this template is compatible with
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// datalake required which this template is compatible with
	// Enum: [NONE OPTIONAL REQUIRED]
	DatalakeRequired string `json:"datalakeRequired,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// node count of the cluster template
	NodeCount int32 `json:"nodeCount,omitempty"`

	// stack type of cluster template
	StackType string `json:"stackType,omitempty"`

	// stack version of cluster template
	StackVersion string `json:"stackVersion,omitempty"`

	// status
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED OUTDATED]
	Status string `json:"status,omitempty"`

	// type of the cluster template
	// Enum: [SPARK HIVE DATASCIENCE EDW ETL OTHER]
	Type string `json:"type,omitempty"`
}

// Validate validates this cluster template view response
func (m *ClusterTemplateViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

var clusterTemplateViewResponseTypeDatalakeRequiredPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","OPTIONAL","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewResponseTypeDatalakeRequiredPropEnum = append(clusterTemplateViewResponseTypeDatalakeRequiredPropEnum, v)
	}
}

const (

	// ClusterTemplateViewResponseDatalakeRequiredNONE captures enum value "NONE"
	ClusterTemplateViewResponseDatalakeRequiredNONE string = "NONE"

	// ClusterTemplateViewResponseDatalakeRequiredOPTIONAL captures enum value "OPTIONAL"
	ClusterTemplateViewResponseDatalakeRequiredOPTIONAL string = "OPTIONAL"

	// ClusterTemplateViewResponseDatalakeRequiredREQUIRED captures enum value "REQUIRED"
	ClusterTemplateViewResponseDatalakeRequiredREQUIRED string = "REQUIRED"
)

// prop value enum
func (m *ClusterTemplateViewResponse) validateDatalakeRequiredEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewResponseTypeDatalakeRequiredPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewResponse) validateDatalakeRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeRequired) { // not required
		return nil
	}

	// value enum
	if err := m.validateDatalakeRequiredEnum("datalakeRequired", "body", m.DatalakeRequired); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateViewResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterTemplateViewResponse) validateName(formats strfmt.Registry) error {

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

var clusterTemplateViewResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED","OUTDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewResponseTypeStatusPropEnum = append(clusterTemplateViewResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterTemplateViewResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterTemplateViewResponseStatusDEFAULT string = "DEFAULT"

	// ClusterTemplateViewResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterTemplateViewResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterTemplateViewResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterTemplateViewResponseStatusUSERMANAGED string = "USER_MANAGED"

	// ClusterTemplateViewResponseStatusOUTDATED captures enum value "OUTDATED"
	ClusterTemplateViewResponseStatusOUTDATED string = "OUTDATED"
)

// prop value enum
func (m *ClusterTemplateViewResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var clusterTemplateViewResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SPARK","HIVE","DATASCIENCE","EDW","ETL","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewResponseTypeTypePropEnum = append(clusterTemplateViewResponseTypeTypePropEnum, v)
	}
}

const (

	// ClusterTemplateViewResponseTypeSPARK captures enum value "SPARK"
	ClusterTemplateViewResponseTypeSPARK string = "SPARK"

	// ClusterTemplateViewResponseTypeHIVE captures enum value "HIVE"
	ClusterTemplateViewResponseTypeHIVE string = "HIVE"

	// ClusterTemplateViewResponseTypeDATASCIENCE captures enum value "DATASCIENCE"
	ClusterTemplateViewResponseTypeDATASCIENCE string = "DATASCIENCE"

	// ClusterTemplateViewResponseTypeEDW captures enum value "EDW"
	ClusterTemplateViewResponseTypeEDW string = "EDW"

	// ClusterTemplateViewResponseTypeETL captures enum value "ETL"
	ClusterTemplateViewResponseTypeETL string = "ETL"

	// ClusterTemplateViewResponseTypeOTHER captures enum value "OTHER"
	ClusterTemplateViewResponseTypeOTHER string = "OTHER"
)

// prop value enum
func (m *ClusterTemplateViewResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewResponse) validateType(formats strfmt.Registry) error {

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
func (m *ClusterTemplateViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplateViewResponse) UnmarshalBinary(b []byte) error {
	var res ClusterTemplateViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}