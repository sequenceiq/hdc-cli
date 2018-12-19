// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StructuredParameterQueryResponse structured parameter query response
// swagger:model StructuredParameterQueryResponse
type StructuredParameterQueryResponse struct {

	// default path
	DefaultPath string `json:"defaultPath,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// property display name
	PropertyDisplayName string `json:"propertyDisplayName,omitempty"`

	// property file
	PropertyFile string `json:"propertyFile,omitempty"`

	// property name
	PropertyName string `json:"propertyName,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// related service
	RelatedService string `json:"relatedService,omitempty"`
}

// Validate validates this structured parameter query response
func (m *StructuredParameterQueryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StructuredParameterQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StructuredParameterQueryResponse) UnmarshalBinary(b []byte) error {
	var res StructuredParameterQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}