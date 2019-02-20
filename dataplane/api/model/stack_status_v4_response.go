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

// StackStatusV4Response stack status v4 response
// swagger:model StackStatusV4Response
type StackStatusV4Response struct {

	// cluster status
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// cluster status reason
	ClusterStatusReason string `json:"clusterStatusReason,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// status
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this stack status v4 response
func (m *StackStatusV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stackStatusV4ResponseTypeClusterStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackStatusV4ResponseTypeClusterStatusPropEnum = append(stackStatusV4ResponseTypeClusterStatusPropEnum, v)
	}
}

const (

	// StackStatusV4ResponseClusterStatusREQUESTED captures enum value "REQUESTED"
	StackStatusV4ResponseClusterStatusREQUESTED string = "REQUESTED"

	// StackStatusV4ResponseClusterStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusAVAILABLE captures enum value "AVAILABLE"
	StackStatusV4ResponseClusterStatusAVAILABLE string = "AVAILABLE"

	// StackStatusV4ResponseClusterStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackStatusV4ResponseClusterStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// StackStatusV4ResponseClusterStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackStatusV4ResponseClusterStatusUPDATEFAILED string = "UPDATE_FAILED"

	// StackStatusV4ResponseClusterStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackStatusV4ResponseClusterStatusCREATEFAILED string = "CREATE_FAILED"

	// StackStatusV4ResponseClusterStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackStatusV4ResponseClusterStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// StackStatusV4ResponseClusterStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackStatusV4ResponseClusterStatusDELETEFAILED string = "DELETE_FAILED"

	// StackStatusV4ResponseClusterStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackStatusV4ResponseClusterStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// StackStatusV4ResponseClusterStatusSTOPPED captures enum value "STOPPED"
	StackStatusV4ResponseClusterStatusSTOPPED string = "STOPPED"

	// StackStatusV4ResponseClusterStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackStatusV4ResponseClusterStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// StackStatusV4ResponseClusterStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackStatusV4ResponseClusterStatusSTARTREQUESTED string = "START_REQUESTED"

	// StackStatusV4ResponseClusterStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackStatusV4ResponseClusterStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// StackStatusV4ResponseClusterStatusSTARTFAILED captures enum value "START_FAILED"
	StackStatusV4ResponseClusterStatusSTARTFAILED string = "START_FAILED"

	// StackStatusV4ResponseClusterStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackStatusV4ResponseClusterStatusSTOPFAILED string = "STOP_FAILED"

	// StackStatusV4ResponseClusterStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackStatusV4ResponseClusterStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// StackStatusV4ResponseClusterStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackStatusV4ResponseClusterStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *StackStatusV4Response) validateClusterStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackStatusV4ResponseTypeClusterStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackStatusV4Response) validateClusterStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterStatusEnum("clusterStatus", "body", m.ClusterStatus); err != nil {
		return err
	}

	return nil
}

var stackStatusV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackStatusV4ResponseTypeStatusPropEnum = append(stackStatusV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// StackStatusV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	StackStatusV4ResponseStatusREQUESTED string = "REQUESTED"

	// StackStatusV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackStatusV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// StackStatusV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackStatusV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// StackStatusV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackStatusV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// StackStatusV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackStatusV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// StackStatusV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackStatusV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// StackStatusV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackStatusV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// StackStatusV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackStatusV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// StackStatusV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackStatusV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// StackStatusV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackStatusV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// StackStatusV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackStatusV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// StackStatusV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackStatusV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// StackStatusV4ResponseStatusSTOPPED captures enum value "STOPPED"
	StackStatusV4ResponseStatusSTOPPED string = "STOPPED"

	// StackStatusV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackStatusV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// StackStatusV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackStatusV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// StackStatusV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackStatusV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// StackStatusV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackStatusV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// StackStatusV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackStatusV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// StackStatusV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackStatusV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// StackStatusV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackStatusV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// StackStatusV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackStatusV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *StackStatusV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackStatusV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackStatusV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackStatusV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackStatusV4Response) UnmarshalBinary(b []byte) error {
	var res StackStatusV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}