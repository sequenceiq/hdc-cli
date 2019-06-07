// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TagSpecificationsResponse tag specifications response
// swagger:model TagSpecificationsResponse
type TagSpecificationsResponse struct {

	// tag specifications
	Specifications map[string]map[string]interface{} `json:"specifications,omitempty"`
}

// Validate validates this tag specifications response
func (m *TagSpecificationsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagSpecificationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagSpecificationsResponse) UnmarshalBinary(b []byte) error {
	var res TagSpecificationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}