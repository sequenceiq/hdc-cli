// Code generated by go-swagger; DO NOT EDIT.

package v4user_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4user profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4user profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTerminatedClustersPreferences deletes user preference to show or hide terminated clusters

User preference whether to show or hide terminated instances and how old deleted instances should be shown.
*/
func (a *Client) DeleteTerminatedClustersPreferences(params *DeleteTerminatedClustersPreferencesParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTerminatedClustersPreferencesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTerminatedClustersPreferences",
		Method:             "DELETE",
		PathPattern:        "/v4/user_profiles/terminated_clusters_preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTerminatedClustersPreferencesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetTerminatedClustersPreferences gets user preference to show or hide terminated clusters

User preference whether to show or hide terminated instances and how old deleted instances should be shown.
*/
func (a *Client) GetTerminatedClustersPreferences(params *GetTerminatedClustersPreferencesParams) (*GetTerminatedClustersPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTerminatedClustersPreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTerminatedClustersPreferences",
		Method:             "GET",
		PathPattern:        "/v4/user_profiles/terminated_clusters_preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTerminatedClustersPreferencesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTerminatedClustersPreferencesOK), nil

}

/*
GetUserProfileInWorkspace users related profile

Users can be invited under an account by the administrator, and all resources
*/
func (a *Client) GetUserProfileInWorkspace(params *GetUserProfileInWorkspaceParams) (*GetUserProfileInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserProfileInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserProfileInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/user_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserProfileInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserProfileInWorkspaceOK), nil

}

/*
PutTerminatedClustersPreferences sets user preference to show or hide terminated clusters

User preference whether to show or hide terminated instances and how old deleted instances should be shown.
*/
func (a *Client) PutTerminatedClustersPreferences(params *PutTerminatedClustersPreferencesParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutTerminatedClustersPreferencesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putTerminatedClustersPreferences",
		Method:             "PUT",
		PathPattern:        "/v4/user_profiles/terminated_clusters_preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutTerminatedClustersPreferencesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}