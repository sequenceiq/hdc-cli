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

// InstanceMetaDataV1Response instance meta data v1 response
// swagger:model InstanceMetaDataV1Response
type InstanceMetaDataV1Response struct {

	// the fully qualified domain name of the node in the service discovery cluster
	DiscoveryFQDN string `json:"discoveryFQDN,omitempty"`

	// name of the instance group
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// id of the instance
	InstanceID string `json:"instanceId,omitempty"`

	// status of the instance
	// Enum: [REQUESTED CREATED UNREGISTERED REGISTERED DECOMMISSIONED TERMINATED DELETED_ON_PROVIDER_SIDE DELETED_BY_PROVIDER FAILED STOPPED REBOOTING UNREACHABLE DELETE_REQUESTED]
	InstanceStatus string `json:"instanceStatus,omitempty"`

	// type of the instance
	// Enum: [GATEWAY GATEWAY_PRIMARY]
	InstanceType string `json:"instanceType,omitempty"`

	// life cycle
	// Enum: [NORMAL SPOT]
	LifeCycle string `json:"lifeCycle,omitempty"`

	// private ip of the insctance
	PrivateIP string `json:"privateIp,omitempty"`

	// public ip of the instance
	PublicIP string `json:"publicIp,omitempty"`

	// ssh port
	SSHPort int32 `json:"sshPort,omitempty"`

	// state of the host
	State string `json:"state,omitempty"`
}

// Validate validates this instance meta data v1 response
func (m *InstanceMetaDataV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifeCycle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var instanceMetaDataV1ResponseTypeInstanceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATED","UNREGISTERED","REGISTERED","DECOMMISSIONED","TERMINATED","DELETED_ON_PROVIDER_SIDE","DELETED_BY_PROVIDER","FAILED","STOPPED","REBOOTING","UNREACHABLE","DELETE_REQUESTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceMetaDataV1ResponseTypeInstanceStatusPropEnum = append(instanceMetaDataV1ResponseTypeInstanceStatusPropEnum, v)
	}
}

const (

	// InstanceMetaDataV1ResponseInstanceStatusREQUESTED captures enum value "REQUESTED"
	InstanceMetaDataV1ResponseInstanceStatusREQUESTED string = "REQUESTED"

	// InstanceMetaDataV1ResponseInstanceStatusCREATED captures enum value "CREATED"
	InstanceMetaDataV1ResponseInstanceStatusCREATED string = "CREATED"

	// InstanceMetaDataV1ResponseInstanceStatusUNREGISTERED captures enum value "UNREGISTERED"
	InstanceMetaDataV1ResponseInstanceStatusUNREGISTERED string = "UNREGISTERED"

	// InstanceMetaDataV1ResponseInstanceStatusREGISTERED captures enum value "REGISTERED"
	InstanceMetaDataV1ResponseInstanceStatusREGISTERED string = "REGISTERED"

	// InstanceMetaDataV1ResponseInstanceStatusDECOMMISSIONED captures enum value "DECOMMISSIONED"
	InstanceMetaDataV1ResponseInstanceStatusDECOMMISSIONED string = "DECOMMISSIONED"

	// InstanceMetaDataV1ResponseInstanceStatusTERMINATED captures enum value "TERMINATED"
	InstanceMetaDataV1ResponseInstanceStatusTERMINATED string = "TERMINATED"

	// InstanceMetaDataV1ResponseInstanceStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	InstanceMetaDataV1ResponseInstanceStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// InstanceMetaDataV1ResponseInstanceStatusDELETEDBYPROVIDER captures enum value "DELETED_BY_PROVIDER"
	InstanceMetaDataV1ResponseInstanceStatusDELETEDBYPROVIDER string = "DELETED_BY_PROVIDER"

	// InstanceMetaDataV1ResponseInstanceStatusFAILED captures enum value "FAILED"
	InstanceMetaDataV1ResponseInstanceStatusFAILED string = "FAILED"

	// InstanceMetaDataV1ResponseInstanceStatusSTOPPED captures enum value "STOPPED"
	InstanceMetaDataV1ResponseInstanceStatusSTOPPED string = "STOPPED"

	// InstanceMetaDataV1ResponseInstanceStatusREBOOTING captures enum value "REBOOTING"
	InstanceMetaDataV1ResponseInstanceStatusREBOOTING string = "REBOOTING"

	// InstanceMetaDataV1ResponseInstanceStatusUNREACHABLE captures enum value "UNREACHABLE"
	InstanceMetaDataV1ResponseInstanceStatusUNREACHABLE string = "UNREACHABLE"

	// InstanceMetaDataV1ResponseInstanceStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	InstanceMetaDataV1ResponseInstanceStatusDELETEREQUESTED string = "DELETE_REQUESTED"
)

// prop value enum
func (m *InstanceMetaDataV1Response) validateInstanceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceMetaDataV1ResponseTypeInstanceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceMetaDataV1Response) validateInstanceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateInstanceStatusEnum("instanceStatus", "body", m.InstanceStatus); err != nil {
		return err
	}

	return nil
}

var instanceMetaDataV1ResponseTypeInstanceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GATEWAY","GATEWAY_PRIMARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceMetaDataV1ResponseTypeInstanceTypePropEnum = append(instanceMetaDataV1ResponseTypeInstanceTypePropEnum, v)
	}
}

const (

	// InstanceMetaDataV1ResponseInstanceTypeGATEWAY captures enum value "GATEWAY"
	InstanceMetaDataV1ResponseInstanceTypeGATEWAY string = "GATEWAY"

	// InstanceMetaDataV1ResponseInstanceTypeGATEWAYPRIMARY captures enum value "GATEWAY_PRIMARY"
	InstanceMetaDataV1ResponseInstanceTypeGATEWAYPRIMARY string = "GATEWAY_PRIMARY"
)

// prop value enum
func (m *InstanceMetaDataV1Response) validateInstanceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceMetaDataV1ResponseTypeInstanceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceMetaDataV1Response) validateInstanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateInstanceTypeEnum("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

var instanceMetaDataV1ResponseTypeLifeCyclePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NORMAL","SPOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceMetaDataV1ResponseTypeLifeCyclePropEnum = append(instanceMetaDataV1ResponseTypeLifeCyclePropEnum, v)
	}
}

const (

	// InstanceMetaDataV1ResponseLifeCycleNORMAL captures enum value "NORMAL"
	InstanceMetaDataV1ResponseLifeCycleNORMAL string = "NORMAL"

	// InstanceMetaDataV1ResponseLifeCycleSPOT captures enum value "SPOT"
	InstanceMetaDataV1ResponseLifeCycleSPOT string = "SPOT"
)

// prop value enum
func (m *InstanceMetaDataV1Response) validateLifeCycleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceMetaDataV1ResponseTypeLifeCyclePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceMetaDataV1Response) validateLifeCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.LifeCycle) { // not required
		return nil
	}

	// value enum
	if err := m.validateLifeCycleEnum("lifeCycle", "body", m.LifeCycle); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceMetaDataV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceMetaDataV1Response) UnmarshalBinary(b []byte) error {
	var res InstanceMetaDataV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
