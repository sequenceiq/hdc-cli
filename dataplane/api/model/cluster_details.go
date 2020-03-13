// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterDetails cluster details
// swagger:model ClusterDetails
type ClusterDetails struct {

	// ambari version
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// cluster type
	ClusterType string `json:"clusterType,omitempty"`

	// cluster version
	ClusterVersion string `json:"clusterVersion,omitempty"`

	// database type
	DatabaseType string `json:"databaseType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// external database
	ExternalDatabase bool `json:"externalDatabase,omitempty"`

	// file system type
	FileSystemType string `json:"fileSystemType,omitempty"`

	// gateway enabled
	GatewayEnabled bool `json:"gatewayEnabled,omitempty"`

	// gateway type
	GatewayType string `json:"gatewayType,omitempty"`

	// host groups
	HostGroups []string `json:"hostGroups"`

	// id
	ID int64 `json:"id,omitempty"`

	// kerberos type
	KerberosType string `json:"kerberosType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sso type
	SsoType string `json:"ssoType,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this cluster details
func (m *ClusterDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDetails) UnmarshalBinary(b []byte) error {
	var res ClusterDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}