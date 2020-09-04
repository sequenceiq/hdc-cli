// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimeAlertRequest time alert request
// swagger:model TimeAlertRequest
type TimeAlertRequest struct {

	// Name of the alert
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*$)
	AlertName string `json:"alertName,omitempty"`

	// Cron expression of the time alert
	Cron string `json:"cron,omitempty"`

	// Description of the alert
	Description string `json:"description,omitempty"`

	// Id of the scaling ploicy
	ScalingPolicy *ScalingPolicyRequest `json:"scalingPolicy,omitempty"`

	// Timezone of the time alert
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this time alert request
func (m *TimeAlertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeAlertRequest) validateAlertName(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertName) { // not required
		return nil
	}

	if err := validate.Pattern("alertName", "body", string(m.AlertName), `(^[a-zA-Z][-a-zA-Z0-9]*$)`); err != nil {
		return err
	}

	return nil
}

func (m *TimeAlertRequest) validateScalingPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScalingPolicy) { // not required
		return nil
	}

	if m.ScalingPolicy != nil {
		if err := m.ScalingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scalingPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeAlertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeAlertRequest) UnmarshalBinary(b []byte) error {
	var res TimeAlertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
