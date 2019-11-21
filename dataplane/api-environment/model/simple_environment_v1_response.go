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

// SimpleEnvironmentV1Response simple environment v1 response
// swagger:model SimpleEnvironmentV1Response
type SimpleEnvironmentV1Response struct {

	// Name of the admin group to be used for all the services.
	AdminGroupName string `json:"adminGroupName,omitempty"`

	// SSH key for accessing cluster node instances.
	Authentication *EnvironmentAuthenticationV1Response `json:"authentication,omitempty"`

	// AWS Specific parameters.
	Aws *AwsEnvironmentV1Parameters `json:"aws,omitempty"`

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Create freeipa in environment
	CreateFreeIpa bool `json:"createFreeIpa,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// credential
	Credential *CredentialViewV1Response `json:"credential,omitempty"`

	// id of the resource
	Crn string `json:"crn,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// Status of the environment.
	// Enum: [CREATION_INITIATED DELETE_INITIATED UPDATE_INITIATED ENVIRONMENT_VALIDATION_IN_PROGRESS NETWORK_CREATION_IN_PROGRESS NETWORK_DELETE_IN_PROGRESS RDBMS_DELETE_IN_PROGRESS FREEIPA_CREATION_IN_PROGRESS FREEIPA_DELETE_IN_PROGRESS CLUSTER_DEFINITION_CLEANUP_PROGRESS UMS_RESOURCE_DELETE_IN_PROGRESS IDBROKER_MAPPINGS_DELETE_IN_PROGRESS S3GUARD_TABLE_DELETE_IN_PROGRESS DATAHUB_CLUSTERS_DELETE_IN_PROGRESS DATALAKE_CLUSTERS_DELETE_IN_PROGRESS PUBLICKEY_CREATE_IN_PROGRESS PUBLICKEY_DELETE_IN_PROGRESS AVAILABLE ARCHIVED CREATE_FAILED DELETE_FAILED UPDATE_FAILED STOP_DATAHUB_STARTED STOP_DATAHUB_FAILED STOP_DATALAKE_STARTED STOP_DATALAKE_FAILED STOP_FREEIPA_STARTED STOP_FREEIPA_FAILED ENV_STOPPED START_DATAHUB_STARTED START_DATAHUB_FAILED START_DATALAKE_STARTED START_DATALAKE_FAILED START_FREEIPA_STARTED START_FREEIPA_FAILED]
	EnvironmentStatus string `json:"environmentStatus,omitempty"`

	// IDBroker mapping source.
	// Enum: [NONE MOCK IDBMMS]
	IDBrokerMappingSource string `json:"idBrokerMappingSource,omitempty"`

	// Location of the environment.
	Location *LocationV1Response `json:"location,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV1Response `json:"network,omitempty"`

	// Regions of the environment.
	Regions *CompactRegionV1Response `json:"regions,omitempty"`

	// Security control for FreeIPA and Datalake deployment.
	SecurityAccess *SecurityAccessV1Response `json:"securityAccess,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// Telemetry related specifics of the environment.
	Telemetry *TelemetryResponse `json:"telemetry,omitempty"`

	// Configuration that the connection going directly or with ccm.
	// Enum: [DIRECT CCM]
	Tunnel string `json:"tunnel,omitempty"`
}

// Validate validates this simple environment v1 response
func (m *SimpleEnvironmentV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDBrokerMappingSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelemetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

var simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATION_INITIATED","DELETE_INITIATED","UPDATE_INITIATED","ENVIRONMENT_VALIDATION_IN_PROGRESS","NETWORK_CREATION_IN_PROGRESS","NETWORK_DELETE_IN_PROGRESS","RDBMS_DELETE_IN_PROGRESS","FREEIPA_CREATION_IN_PROGRESS","FREEIPA_DELETE_IN_PROGRESS","CLUSTER_DEFINITION_CLEANUP_PROGRESS","UMS_RESOURCE_DELETE_IN_PROGRESS","IDBROKER_MAPPINGS_DELETE_IN_PROGRESS","S3GUARD_TABLE_DELETE_IN_PROGRESS","DATAHUB_CLUSTERS_DELETE_IN_PROGRESS","DATALAKE_CLUSTERS_DELETE_IN_PROGRESS","PUBLICKEY_CREATE_IN_PROGRESS","PUBLICKEY_DELETE_IN_PROGRESS","AVAILABLE","ARCHIVED","CREATE_FAILED","DELETE_FAILED","UPDATE_FAILED","STOP_DATAHUB_STARTED","STOP_DATAHUB_FAILED","STOP_DATALAKE_STARTED","STOP_DATALAKE_FAILED","STOP_FREEIPA_STARTED","STOP_FREEIPA_FAILED","ENV_STOPPED","START_DATAHUB_STARTED","START_DATAHUB_FAILED","START_DATALAKE_STARTED","START_DATALAKE_FAILED","START_FREEIPA_STARTED","START_FREEIPA_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum = append(simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum, v)
	}
}

const (

	// SimpleEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED captures enum value "CREATION_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED string = "CREATION_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED captures enum value "DELETE_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED string = "DELETE_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED captures enum value "UPDATE_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED string = "UPDATE_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusENVIRONMENTVALIDATIONINPROGRESS captures enum value "ENVIRONMENT_VALIDATION_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusENVIRONMENTVALIDATIONINPROGRESS string = "ENVIRONMENT_VALIDATION_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS captures enum value "NETWORK_CREATION_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS string = "NETWORK_CREATION_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS captures enum value "NETWORK_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS string = "NETWORK_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS captures enum value "RDBMS_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS string = "RDBMS_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS captures enum value "FREEIPA_CREATION_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS string = "FREEIPA_CREATION_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS captures enum value "FREEIPA_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS string = "FREEIPA_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusCLUSTERDEFINITIONCLEANUPPROGRESS captures enum value "CLUSTER_DEFINITION_CLEANUP_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusCLUSTERDEFINITIONCLEANUPPROGRESS string = "CLUSTER_DEFINITION_CLEANUP_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusUMSRESOURCEDELETEINPROGRESS captures enum value "UMS_RESOURCE_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusUMSRESOURCEDELETEINPROGRESS string = "UMS_RESOURCE_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusIDBROKERMAPPINGSDELETEINPROGRESS captures enum value "IDBROKER_MAPPINGS_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusIDBROKERMAPPINGSDELETEINPROGRESS string = "IDBROKER_MAPPINGS_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusS3GUARDTABLEDELETEINPROGRESS captures enum value "S3GUARD_TABLE_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusS3GUARDTABLEDELETEINPROGRESS string = "S3GUARD_TABLE_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDATAHUBCLUSTERSDELETEINPROGRESS captures enum value "DATAHUB_CLUSTERS_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusDATAHUBCLUSTERSDELETEINPROGRESS string = "DATAHUB_CLUSTERS_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDATALAKECLUSTERSDELETEINPROGRESS captures enum value "DATALAKE_CLUSTERS_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusDATALAKECLUSTERSDELETEINPROGRESS string = "DATALAKE_CLUSTERS_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusPUBLICKEYCREATEINPROGRESS captures enum value "PUBLICKEY_CREATE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusPUBLICKEYCREATEINPROGRESS string = "PUBLICKEY_CREATE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusPUBLICKEYDELETEINPROGRESS captures enum value "PUBLICKEY_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusPUBLICKEYDELETEINPROGRESS string = "PUBLICKEY_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusAVAILABLE captures enum value "AVAILABLE"
	SimpleEnvironmentV1ResponseEnvironmentStatusAVAILABLE string = "AVAILABLE"

	// SimpleEnvironmentV1ResponseEnvironmentStatusARCHIVED captures enum value "ARCHIVED"
	SimpleEnvironmentV1ResponseEnvironmentStatusARCHIVED string = "ARCHIVED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusCREATEFAILED captures enum value "CREATE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusCREATEFAILED string = "CREATE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusDELETEFAILED string = "DELETE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED string = "UPDATE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBSTARTED captures enum value "STOP_DATAHUB_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBSTARTED string = "STOP_DATAHUB_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBFAILED captures enum value "STOP_DATAHUB_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBFAILED string = "STOP_DATAHUB_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKESTARTED captures enum value "STOP_DATALAKE_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKESTARTED string = "STOP_DATALAKE_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKEFAILED captures enum value "STOP_DATALAKE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKEFAILED string = "STOP_DATALAKE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPASTARTED captures enum value "STOP_FREEIPA_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPASTARTED string = "STOP_FREEIPA_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPAFAILED captures enum value "STOP_FREEIPA_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPAFAILED string = "STOP_FREEIPA_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusENVSTOPPED captures enum value "ENV_STOPPED"
	SimpleEnvironmentV1ResponseEnvironmentStatusENVSTOPPED string = "ENV_STOPPED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBSTARTED captures enum value "START_DATAHUB_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBSTARTED string = "START_DATAHUB_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBFAILED captures enum value "START_DATAHUB_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBFAILED string = "START_DATAHUB_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKESTARTED captures enum value "START_DATALAKE_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKESTARTED string = "START_DATALAKE_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKEFAILED captures enum value "START_DATALAKE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKEFAILED string = "START_DATALAKE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPASTARTED captures enum value "START_FREEIPA_STARTED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPASTARTED string = "START_FREEIPA_STARTED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPAFAILED captures enum value "START_FREEIPA_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPAFAILED string = "START_FREEIPA_FAILED"
)

// prop value enum
func (m *SimpleEnvironmentV1Response) validateEnvironmentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateEnvironmentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentStatusEnum("environmentStatus", "body", m.EnvironmentStatus); err != nil {
		return err
	}

	return nil
}

var simpleEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","MOCK","IDBMMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum = append(simpleEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum, v)
	}
}

const (

	// SimpleEnvironmentV1ResponseIDBrokerMappingSourceNONE captures enum value "NONE"
	SimpleEnvironmentV1ResponseIDBrokerMappingSourceNONE string = "NONE"

	// SimpleEnvironmentV1ResponseIDBrokerMappingSourceMOCK captures enum value "MOCK"
	SimpleEnvironmentV1ResponseIDBrokerMappingSourceMOCK string = "MOCK"

	// SimpleEnvironmentV1ResponseIDBrokerMappingSourceIDBMMS captures enum value "IDBMMS"
	SimpleEnvironmentV1ResponseIDBrokerMappingSourceIDBMMS string = "IDBMMS"
)

// prop value enum
func (m *SimpleEnvironmentV1Response) validateIDBrokerMappingSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simpleEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateIDBrokerMappingSource(formats strfmt.Registry) error {

	if swag.IsZero(m.IDBrokerMappingSource) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDBrokerMappingSourceEnum("idBrokerMappingSource", "body", m.IDBrokerMappingSource); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if m.Regions != nil {
		if err := m.Regions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regions")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateSecurityAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityAccess) { // not required
		return nil
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateTelemetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Telemetry) { // not required
		return nil
	}

	if m.Telemetry != nil {
		if err := m.Telemetry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("telemetry")
			}
			return err
		}
	}

	return nil
}

var simpleEnvironmentV1ResponseTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleEnvironmentV1ResponseTypeTunnelPropEnum = append(simpleEnvironmentV1ResponseTypeTunnelPropEnum, v)
	}
}

const (

	// SimpleEnvironmentV1ResponseTunnelDIRECT captures enum value "DIRECT"
	SimpleEnvironmentV1ResponseTunnelDIRECT string = "DIRECT"

	// SimpleEnvironmentV1ResponseTunnelCCM captures enum value "CCM"
	SimpleEnvironmentV1ResponseTunnelCCM string = "CCM"
)

// prop value enum
func (m *SimpleEnvironmentV1Response) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simpleEnvironmentV1ResponseTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateTunnel(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnel) { // not required
		return nil
	}

	// value enum
	if err := m.validateTunnelEnum("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleEnvironmentV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleEnvironmentV1Response) UnmarshalBinary(b []byte) error {
	var res SimpleEnvironmentV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
