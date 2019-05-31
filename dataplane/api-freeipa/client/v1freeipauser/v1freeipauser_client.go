// Code generated by go-swagger; DO NOT EDIT.

package v1freeipauser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1freeipauser API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1freeipauser API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUsersV1 creates groups and users in the free IP a servers

User synchronization and management operations
*/
func (a *Client) CreateUsersV1(params *CreateUsersV1Params) (*CreateUsersV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsersV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUsersV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/user/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUsersV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUsersV1OK), nil

}

/*
GetSyncUsersStatusV1 gets the status of synchronization operation

User synchronization and management operations
*/
func (a *Client) GetSyncUsersStatusV1(params *GetSyncUsersStatusV1Params) (*GetSyncUsersStatusV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncUsersStatusV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncUsersStatusV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/user/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSyncUsersStatusV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncUsersStatusV1OK), nil

}

/*
SetPasswordV1 sets the user s password in the free IP a servers

User synchronization and management operations
*/
func (a *Client) SetPasswordV1(params *SetPasswordV1Params) (*SetPasswordV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPasswordV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPasswordV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/user/{username}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetPasswordV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetPasswordV1OK), nil

}

/*
SynchronizeUsersV1 synchronizes groups and users to the free IP a servers

User synchronization and management operations
*/
func (a *Client) SynchronizeUsersV1(params *SynchronizeUsersV1Params) (*SynchronizeUsersV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSynchronizeUsersV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "synchronizeUsersV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/user/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SynchronizeUsersV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SynchronizeUsersV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
