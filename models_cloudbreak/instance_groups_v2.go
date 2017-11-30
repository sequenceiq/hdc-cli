// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceGroupsV2 instance groups v2
// swagger:model instanceGroupsV2

type InstanceGroupsV2 struct {

	// name of the instance group
	// Required: true
	Group *string `json:"group"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// cloud specific parameters for instance group
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// referenced recipe names
	// Unique: true
	RecipeNames []string `json:"recipeNames"`

	// recovery mode of the hostgroup's nodes
	RecoveryMode string `json:"recoveryMode,omitempty"`

	// instancegroup related securitygroup
	SecurityGroup *SecurityGroupV2Request `json:"securityGroup,omitempty"`

	// instancegroup related template
	// Required: true
	Template *TemplateV2Request `json:"template"`

	// type of the instance group
	Type string `json:"type,omitempty"`
}

/* polymorph instanceGroupsV2 group false */

/* polymorph instanceGroupsV2 nodeCount false */

/* polymorph instanceGroupsV2 parameters false */

/* polymorph instanceGroupsV2 recipeNames false */

/* polymorph instanceGroupsV2 recoveryMode false */

/* polymorph instanceGroupsV2 securityGroup false */

/* polymorph instanceGroupsV2 template false */

/* polymorph instanceGroupsV2 type false */

// Validate validates this instance groups v2
func (m *InstanceGroupsV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecoveryMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
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

func (m *InstanceGroupsV2) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupsV2) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(*m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(*m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupsV2) validateRecipeNames(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipeNames", "body", m.RecipeNames); err != nil {
		return err
	}

	return nil
}

var instanceGroupsV2TypeRecoveryModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupsV2TypeRecoveryModePropEnum = append(instanceGroupsV2TypeRecoveryModePropEnum, v)
	}
}

const (
	// InstanceGroupsV2RecoveryModeMANUAL captures enum value "MANUAL"
	InstanceGroupsV2RecoveryModeMANUAL string = "MANUAL"
	// InstanceGroupsV2RecoveryModeAUTO captures enum value "AUTO"
	InstanceGroupsV2RecoveryModeAUTO string = "AUTO"
)

// prop value enum
func (m *InstanceGroupsV2) validateRecoveryModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupsV2TypeRecoveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupsV2) validateRecoveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryModeEnum("recoveryMode", "body", m.RecoveryMode); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupsV2) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {

		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupsV2) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {

		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

var instanceGroupsV2TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GATEWAY","CORE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupsV2TypeTypePropEnum = append(instanceGroupsV2TypeTypePropEnum, v)
	}
}

const (
	// InstanceGroupsV2TypeGATEWAY captures enum value "GATEWAY"
	InstanceGroupsV2TypeGATEWAY string = "GATEWAY"
	// InstanceGroupsV2TypeCORE captures enum value "CORE"
	InstanceGroupsV2TypeCORE string = "CORE"
)

// prop value enum
func (m *InstanceGroupsV2) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupsV2TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupsV2) validateType(formats strfmt.Registry) error {

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
func (m *InstanceGroupsV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupsV2) UnmarshalBinary(b []byte) error {
	var res InstanceGroupsV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
