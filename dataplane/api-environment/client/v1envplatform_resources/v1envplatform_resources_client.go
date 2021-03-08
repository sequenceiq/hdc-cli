// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1envplatform resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1envplatform resources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccessConfigsByEnv retrieves access configs with properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetAccessConfigsByEnv(params *GetAccessConfigsByEnvParams) (*GetAccessConfigsByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessConfigsByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessConfigsByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/access_configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessConfigsByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccessConfigsByEnvOK), nil

}

/*
GetEncryptionKeysByEnv retrieves encryption keys with properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetEncryptionKeysByEnv(params *GetEncryptionKeysByEnvParams) (*GetEncryptionKeysByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEncryptionKeysByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEncryptionKeysByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/encryption_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEncryptionKeysByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEncryptionKeysByEnvOK), nil

}

/*
GetGatewaysByEnv retrieves gateways with properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetGatewaysByEnv(params *GetGatewaysByEnvParams) (*GetGatewaysByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewaysByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGatewaysByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGatewaysByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGatewaysByEnvOK), nil

}

/*
GetIPPoolsByEnv retrieves ip pools with properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetIPPoolsByEnv(params *GetIPPoolsByEnvParams) (*GetIPPoolsByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPPoolsByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIpPoolsByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/ip_pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetIPPoolsByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPPoolsByEnvOK), nil

}

/*
GetNoSQLTablesByEnv retrieves nosql tables by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetNoSQLTablesByEnv(params *GetNoSQLTablesByEnvParams) (*GetNoSQLTablesByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNoSQLTablesByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNoSqlTablesByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/nosql_tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNoSQLTablesByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNoSQLTablesByEnvOK), nil

}

/*
GetPlatformNetworksByEnv retrieves network properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformNetworksByEnv(params *GetPlatformNetworksByEnvParams) (*GetPlatformNetworksByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformNetworksByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformNetworksByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformNetworksByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformNetworksByEnvOK), nil

}

/*
GetPlatformSShKeysByEnv retrieves sshkeys properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSShKeysByEnv(params *GetPlatformSShKeysByEnvParams) (*GetPlatformSShKeysByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSShKeysByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSShKeysByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/ssh_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSShKeysByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSShKeysByEnvOK), nil

}

/*
GetPlatformSecurityGroupsByEnv retrieves securitygroups properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSecurityGroupsByEnv(params *GetPlatformSecurityGroupsByEnvParams) (*GetPlatformSecurityGroupsByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSecurityGroupsByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSecurityGroupsByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/security_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSecurityGroupsByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSecurityGroupsByEnvOK), nil

}

/*
GetRegionsByEnv retrieves regions by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetRegionsByEnv(params *GetRegionsByEnvParams) (*GetRegionsByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegionsByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionsByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionsByEnvOK), nil

}

/*
GetResourceGroupsByEnv retrieves resource groups by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetResourceGroupsByEnv(params *GetResourceGroupsByEnvParams) (*GetResourceGroupsByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceGroupsByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResourceGroupsByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/resource_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetResourceGroupsByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResourceGroupsByEnvOK), nil

}

/*
GetVMTypesByCredentialByEnv retrieves vmtype properties by environment

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetVMTypesByCredentialByEnv(params *GetVMTypesByCredentialByEnvParams) (*GetVMTypesByCredentialByEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMTypesByCredentialByEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVmTypesByCredentialByEnv",
		Method:             "GET",
		PathPattern:        "/v1/env/platform_resources/machine_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVMTypesByCredentialByEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVMTypesByCredentialByEnvOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}