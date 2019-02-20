// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlatformDisksV4Response platform disks v4 response
// swagger:model PlatformDisksV4Response
type PlatformDisksV4Response struct {

	// default disks
	DefaultDisks map[string]string `json:"defaultDisks,omitempty"`

	// disk mappings
	DiskMappings map[string]map[string]string `json:"diskMappings,omitempty"`

	// disk types
	DiskTypes map[string][]string `json:"diskTypes,omitempty"`

	// disk displayNames
	DisplayNames map[string]map[string]string `json:"displayNames,omitempty"`
}

// Validate validates this platform disks v4 response
func (m *PlatformDisksV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlatformDisksV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformDisksV4Response) UnmarshalBinary(b []byte) error {
	var res PlatformDisksV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}