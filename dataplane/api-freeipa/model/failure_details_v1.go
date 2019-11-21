// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FailureDetailsV1 failure details v1
// swagger:model FailureDetailsV1
type FailureDetailsV1 struct {

	// additional details
	AdditionalDetails map[string]string `json:"additionalDetails,omitempty"`

	// environment
	Environment string `json:"environment,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this failure details v1
func (m *FailureDetailsV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FailureDetailsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailureDetailsV1) UnmarshalBinary(b []byte) error {
	var res FailureDetailsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
