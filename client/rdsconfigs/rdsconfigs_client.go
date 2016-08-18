package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new rdsconfigs API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rdsconfigs API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteRdsconfigsAccountName deletes public owned or private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteRdsconfigsAccountName(params *DeleteRdsconfigsAccountNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRdsconfigsAccountNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRdsconfigsAccountName",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRdsconfigsAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteRdsconfigsID deletes r d s configuration by id

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteRdsconfigsID(params *DeleteRdsconfigsIDParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRdsconfigsIDParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRdsconfigsID",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRdsconfigsIDReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteRdsconfigsUserName deletes private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteRdsconfigsUserName(params *DeleteRdsconfigsUserNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRdsconfigsUserNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRdsconfigsUserName",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRdsconfigsUserNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetRdsconfigsAccount retrieves public and private owned r d s configurations

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsconfigsAccount(params *GetRdsconfigsAccountParams) (*GetRdsconfigsAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsconfigsAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRdsconfigsAccount",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRdsconfigsAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsconfigsAccountOK), nil
}

/*
GetRdsconfigsAccountName retrieves a public or private owned r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsconfigsAccountName(params *GetRdsconfigsAccountNameParams) (*GetRdsconfigsAccountNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsconfigsAccountNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRdsconfigsAccountName",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRdsconfigsAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsconfigsAccountNameOK), nil
}

/*
GetRdsconfigsID retrieves r d s configuration by id

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsconfigsID(params *GetRdsconfigsIDParams) (*GetRdsconfigsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsconfigsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRdsconfigsID",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRdsconfigsIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsconfigsIDOK), nil
}

/*
GetRdsconfigsUser retrieves private r d s configurations

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsconfigsUser(params *GetRdsconfigsUserParams) (*GetRdsconfigsUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsconfigsUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRdsconfigsUser",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRdsconfigsUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsconfigsUserOK), nil
}

/*
GetRdsconfigsUserName retrieves a private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsconfigsUserName(params *GetRdsconfigsUserNameParams) (*GetRdsconfigsUserNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsconfigsUserNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRdsconfigsUserName",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRdsconfigsUserNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsconfigsUserNameOK), nil
}

/*
PostRdsconfigsAccount creates r d s configuration as public resource

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) PostRdsconfigsAccount(params *PostRdsconfigsAccountParams) (*PostRdsconfigsAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRdsconfigsAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostRdsconfigsAccount",
		Method:             "POST",
		PathPattern:        "/rdsconfigs/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostRdsconfigsAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRdsconfigsAccountOK), nil
}

/*
PostRdsconfigsUser creates r d s configuration as private resource

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) PostRdsconfigsUser(params *PostRdsconfigsUserParams) (*PostRdsconfigsUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRdsconfigsUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostRdsconfigsUser",
		Method:             "POST",
		PathPattern:        "/rdsconfigs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostRdsconfigsUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRdsconfigsUserOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
