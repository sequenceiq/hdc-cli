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

// EnvironmentNetworkV1Request environment network v1 request
// swagger:model EnvironmentNetworkV1Request
type EnvironmentNetworkV1Request struct {

	// AWS-specific properties of the network
	Aws *EnvironmentNetworkAwsV1Params `json:"aws,omitempty"`

	// Azure-specific properties of the network
	Azure *EnvironmentNetworkAzureV1Params `json:"azure,omitempty"`

	// Subnet ids for the Public Endpoint Access Gateway. If provided, these are the subnets that will be used to create a public Knox endpoint for out-of-network UI/API access. If not provided, public subnets will be selected from the subnet list provided for environment creation. (Optional)
	// Unique: true
	EndpointGatewaySubnetIds []string `json:"endpointGatewaySubnetIds"`

	// GCP-specific properties of the network
	Gcp *EnvironmentNetworkGcpV1Params `json:"gcp,omitempty"`

	// Mock-specific properties of the network
	Mock *EnvironmentNetworkMockV1Params `json:"mock,omitempty"`

	// network cidr
	// Max Length: 255
	// Min Length: 0
	NetworkCidr *string `json:"networkCidr,omitempty"`

	// A flag to enable or disable the outbound internet traffic from the instances.
	// Enum: [ENABLED DISABLED]
	OutboundInternetTraffic string `json:"outboundInternetTraffic,omitempty"`

	// A flag to enable or disable the private subnet creation.
	// Enum: [ENABLED DISABLED]
	PrivateSubnetCreation string `json:"privateSubnetCreation,omitempty"`

	// A flag to enable the Public Endpoint Access Gateway, which provides public UI/API access to private data lakes and data hubs.
	// Enum: [ENABLED DISABLED]
	PublicEndpointAccessGateway string `json:"publicEndpointAccessGateway,omitempty"`

	// A flag to enable or disable the service endpoint creation.
	// Enum: [ENABLED DISABLED ENABLED_PRIVATE_ENDPOINT]
	ServiceEndpointCreation string `json:"serviceEndpointCreation,omitempty"`

	// Subnet ids of the specified networks
	// Unique: true
	SubnetIds []string `json:"subnetIds"`

	// Yarn-specific properties of the network
	Yarn *EnvironmentNetworkYarnV1Params `json:"yarn,omitempty"`
}

// Validate validates this environment network v1 request
func (m *EnvironmentNetworkV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointGatewaySubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundInternetTraffic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateSubnetCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicEndpointAccessGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceEndpointCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentNetworkV1Request) validateAws(formats strfmt.Registry) error {

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

func (m *EnvironmentNetworkV1Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateEndpointGatewaySubnetIds(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointGatewaySubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("endpointGatewaySubnetIds", "body", m.EndpointGatewaySubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateMock(formats strfmt.Registry) error {

	if swag.IsZero(m.Mock) { // not required
		return nil
	}

	if m.Mock != nil {
		if err := m.Mock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mock")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkCidr) { // not required
		return nil
	}

	if err := validate.MinLength("networkCidr", "body", string(*m.NetworkCidr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("networkCidr", "body", string(*m.NetworkCidr), 255); err != nil {
		return err
	}

	return nil
}

var environmentNetworkV1RequestTypeOutboundInternetTrafficPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentNetworkV1RequestTypeOutboundInternetTrafficPropEnum = append(environmentNetworkV1RequestTypeOutboundInternetTrafficPropEnum, v)
	}
}

const (

	// EnvironmentNetworkV1RequestOutboundInternetTrafficENABLED captures enum value "ENABLED"
	EnvironmentNetworkV1RequestOutboundInternetTrafficENABLED string = "ENABLED"

	// EnvironmentNetworkV1RequestOutboundInternetTrafficDISABLED captures enum value "DISABLED"
	EnvironmentNetworkV1RequestOutboundInternetTrafficDISABLED string = "DISABLED"
)

// prop value enum
func (m *EnvironmentNetworkV1Request) validateOutboundInternetTrafficEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentNetworkV1RequestTypeOutboundInternetTrafficPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentNetworkV1Request) validateOutboundInternetTraffic(formats strfmt.Registry) error {

	if swag.IsZero(m.OutboundInternetTraffic) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutboundInternetTrafficEnum("outboundInternetTraffic", "body", m.OutboundInternetTraffic); err != nil {
		return err
	}

	return nil
}

var environmentNetworkV1RequestTypePrivateSubnetCreationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentNetworkV1RequestTypePrivateSubnetCreationPropEnum = append(environmentNetworkV1RequestTypePrivateSubnetCreationPropEnum, v)
	}
}

const (

	// EnvironmentNetworkV1RequestPrivateSubnetCreationENABLED captures enum value "ENABLED"
	EnvironmentNetworkV1RequestPrivateSubnetCreationENABLED string = "ENABLED"

	// EnvironmentNetworkV1RequestPrivateSubnetCreationDISABLED captures enum value "DISABLED"
	EnvironmentNetworkV1RequestPrivateSubnetCreationDISABLED string = "DISABLED"
)

// prop value enum
func (m *EnvironmentNetworkV1Request) validatePrivateSubnetCreationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentNetworkV1RequestTypePrivateSubnetCreationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentNetworkV1Request) validatePrivateSubnetCreation(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateSubnetCreation) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivateSubnetCreationEnum("privateSubnetCreation", "body", m.PrivateSubnetCreation); err != nil {
		return err
	}

	return nil
}

var environmentNetworkV1RequestTypePublicEndpointAccessGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentNetworkV1RequestTypePublicEndpointAccessGatewayPropEnum = append(environmentNetworkV1RequestTypePublicEndpointAccessGatewayPropEnum, v)
	}
}

const (

	// EnvironmentNetworkV1RequestPublicEndpointAccessGatewayENABLED captures enum value "ENABLED"
	EnvironmentNetworkV1RequestPublicEndpointAccessGatewayENABLED string = "ENABLED"

	// EnvironmentNetworkV1RequestPublicEndpointAccessGatewayDISABLED captures enum value "DISABLED"
	EnvironmentNetworkV1RequestPublicEndpointAccessGatewayDISABLED string = "DISABLED"
)

// prop value enum
func (m *EnvironmentNetworkV1Request) validatePublicEndpointAccessGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentNetworkV1RequestTypePublicEndpointAccessGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentNetworkV1Request) validatePublicEndpointAccessGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicEndpointAccessGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validatePublicEndpointAccessGatewayEnum("publicEndpointAccessGateway", "body", m.PublicEndpointAccessGateway); err != nil {
		return err
	}

	return nil
}

var environmentNetworkV1RequestTypeServiceEndpointCreationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","ENABLED_PRIVATE_ENDPOINT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentNetworkV1RequestTypeServiceEndpointCreationPropEnum = append(environmentNetworkV1RequestTypeServiceEndpointCreationPropEnum, v)
	}
}

const (

	// EnvironmentNetworkV1RequestServiceEndpointCreationENABLED captures enum value "ENABLED"
	EnvironmentNetworkV1RequestServiceEndpointCreationENABLED string = "ENABLED"

	// EnvironmentNetworkV1RequestServiceEndpointCreationDISABLED captures enum value "DISABLED"
	EnvironmentNetworkV1RequestServiceEndpointCreationDISABLED string = "DISABLED"

	// EnvironmentNetworkV1RequestServiceEndpointCreationENABLEDPRIVATEENDPOINT captures enum value "ENABLED_PRIVATE_ENDPOINT"
	EnvironmentNetworkV1RequestServiceEndpointCreationENABLEDPRIVATEENDPOINT string = "ENABLED_PRIVATE_ENDPOINT"
)

// prop value enum
func (m *EnvironmentNetworkV1Request) validateServiceEndpointCreationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentNetworkV1RequestTypeServiceEndpointCreationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentNetworkV1Request) validateServiceEndpointCreation(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEndpointCreation) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceEndpointCreationEnum("serviceEndpointCreation", "body", m.ServiceEndpointCreation); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateSubnetIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Request) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentNetworkV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentNetworkV1Request) UnmarshalBinary(b []byte) error {
	var res EnvironmentNetworkV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
