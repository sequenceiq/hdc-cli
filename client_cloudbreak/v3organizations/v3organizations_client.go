// Code generated by go-swagger; DO NOT EDIT.

package v3organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOrganizationUsers adds users to the given organization

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) AddOrganizationUsers(params *AddOrganizationUsersParams) (*AddOrganizationUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOrganizationUsers",
		Method:             "PUT",
		PathPattern:        "/v3/organizations/name/{name}/addUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOrganizationUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOrganizationUsersOK), nil

}

/*
ChangeOrganizationUsers changes users and their permissions in the organization

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) ChangeOrganizationUsers(params *ChangeOrganizationUsersParams) (*ChangeOrganizationUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeOrganizationUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeOrganizationUsers",
		Method:             "PUT",
		PathPattern:        "/v3/organizations/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeOrganizationUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeOrganizationUsersOK), nil

}

/*
CreateOrganization creates an organization

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) CreateOrganization(params *CreateOrganizationParams) (*CreateOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrganization",
		Method:             "POST",
		PathPattern:        "/v3/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrganizationOK), nil

}

/*
DeleteOrganizationByName deletes an organization by name

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) DeleteOrganizationByName(params *DeleteOrganizationByNameParams) (*DeleteOrganizationByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganizationByName",
		Method:             "DELETE",
		PathPattern:        "/v3/organizations/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationByNameOK), nil

}

/*
GetOrganizationByName retrieves an organization by name

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) GetOrganizationByName(params *GetOrganizationByNameParams) (*GetOrganizationByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationByName",
		Method:             "GET",
		PathPattern:        "/v3/organizations/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationByNameOK), nil

}

/*
GetOrganizations retrieves organizations

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) GetOrganizations(params *GetOrganizationsParams) (*GetOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizations",
		Method:             "GET",
		PathPattern:        "/v3/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationsOK), nil

}

/*
RemoveOrganizationUsers removes users from the given organization by their user ids

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) RemoveOrganizationUsers(params *RemoveOrganizationUsersParams) (*RemoveOrganizationUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOrganizationUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeOrganizationUsers",
		Method:             "PUT",
		PathPattern:        "/v3/organizations/name/{name}/removeUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOrganizationUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveOrganizationUsersOK), nil

}

/*
UpdateOrganizationUsers updates the users permissions in the given organization

Organizations are a way of grouping resources, organization owners can add users to their organizations with different permission sets
*/
func (a *Client) UpdateOrganizationUsers(params *UpdateOrganizationUsersParams) (*UpdateOrganizationUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationUsers",
		Method:             "PUT",
		PathPattern:        "/v3/organizations/name/{name}/updateUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrganizationUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOrganizationUsersOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}