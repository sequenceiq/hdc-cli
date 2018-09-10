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

// OperationDetails operation details
// swagger:model OperationDetails

type OperationDetails struct {

	// account
	Account string `json:"account,omitempty"`

	// cloudbreak Id
	CloudbreakID string `json:"cloudbreakId,omitempty"`

	// cloudbreak version
	CloudbreakVersion string `json:"cloudbreakVersion,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// resource Id
	ResourceID int64 `json:"resourceId,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// user Id v3
	UserIDV3 string `json:"userIdV3,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`

	// user name v3
	UserNameV3 string `json:"userNameV3,omitempty"`

	// workspace Id
	WorkspaceID int64 `json:"workspaceId,omitempty"`

	// zoned date time
	ZonedDateTime strfmt.DateTime `json:"zonedDateTime,omitempty"`
}

/* polymorph OperationDetails account false */

/* polymorph OperationDetails cloudbreakId false */

/* polymorph OperationDetails cloudbreakVersion false */

/* polymorph OperationDetails eventType false */

/* polymorph OperationDetails resourceId false */

/* polymorph OperationDetails resourceName false */

/* polymorph OperationDetails resourceType false */

/* polymorph OperationDetails timestamp false */

/* polymorph OperationDetails userId false */

/* polymorph OperationDetails userIdV3 false */

/* polymorph OperationDetails userName false */

/* polymorph OperationDetails userNameV3 false */

/* polymorph OperationDetails workspaceId false */

/* polymorph OperationDetails zonedDateTime false */

// Validate validates this operation details
func (m *OperationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var operationDetailsTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REST","FLOW","NOTIFICATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationDetailsTypeEventTypePropEnum = append(operationDetailsTypeEventTypePropEnum, v)
	}
}

const (
	// OperationDetailsEventTypeREST captures enum value "REST"
	OperationDetailsEventTypeREST string = "REST"
	// OperationDetailsEventTypeFLOW captures enum value "FLOW"
	OperationDetailsEventTypeFLOW string = "FLOW"
	// OperationDetailsEventTypeNOTIFICATION captures enum value "NOTIFICATION"
	OperationDetailsEventTypeNOTIFICATION string = "NOTIFICATION"
)

// prop value enum
func (m *OperationDetails) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, operationDetailsTypeEventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OperationDetails) validateEventType(formats strfmt.Registry) error {

	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventTypeEnum("eventType", "body", m.EventType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationDetails) UnmarshalBinary(b []byte) error {
	var res OperationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
