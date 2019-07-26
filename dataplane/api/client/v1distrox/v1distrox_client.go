// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1distrox API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1distrox API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteDistroXV1ByCrn deletes stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteDistroXV1ByCrn(params *DeleteDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDistroXV1ByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteDistroXV1ByName deletes an workspace by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteDistroXV1ByName(params *DeleteDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDistroXV1ByName",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteInstanceDistroXV1ByCrn deletes instance from the stack s cluster by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceDistroXV1ByCrn(params *DeleteInstanceDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstanceDistroXV1ByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/crn/{crn}/instance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteInstanceDistroXV1ByName deletes instance from the stack s cluster by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceDistroXV1ByName(params *DeleteInstanceDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstanceDistroXV1ByName",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/name/{name}/instance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteWithKerberosDistroXV1ByCrn deletes the stack with kerberos cluster by crn

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteWithKerberosDistroXV1ByCrn(params *DeleteWithKerberosDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWithKerberosDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWithKerberosDistroXV1ByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/crn/{crn}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteWithKerberosDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteWithKerberosDistroXV1ByName deletes the stack with kerberos cluster by name

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteWithKerberosDistroXV1ByName(params *DeleteWithKerberosDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWithKerberosDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWithKerberosDistroXV1ByName",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/name/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteWithKerberosDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetDistroXRequestV1ByCrn gets stack request by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXRequestV1ByCrn(params *GetDistroXRequestV1ByCrnParams) (*GetDistroXRequestV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXRequestV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXRequestV1ByCrn",
		Method:             "GET",
		PathPattern:        "/v1/distrox/crn/{crn}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXRequestV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXRequestV1ByCrnOK), nil

}

/*
GetDistroXRequestV1ByName gets stack request by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXRequestV1ByName(params *GetDistroXRequestV1ByNameParams) (*GetDistroXRequestV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXRequestV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXRequestV1ByName",
		Method:             "GET",
		PathPattern:        "/v1/distrox/name/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXRequestV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXRequestV1ByNameOK), nil

}

/*
GetDistroXV1ByCrn gets stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXV1ByCrn(params *GetDistroXV1ByCrnParams) (*GetDistroXV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXV1ByCrn",
		Method:             "GET",
		PathPattern:        "/v1/distrox/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXV1ByCrnOK), nil

}

/*
GetDistroXV1ByName gets stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXV1ByName(params *GetDistroXV1ByNameParams) (*GetDistroXV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXV1ByName",
		Method:             "GET",
		PathPattern:        "/v1/distrox/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXV1ByNameOK), nil

}

/*
ListDistroXV1 lists stacks by environment crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ListDistroXV1(params *ListDistroXV1Params) (*ListDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDistroXV1",
		Method:             "GET",
		PathPattern:        "/v1/distrox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDistroXV1OK), nil

}

/*
PostDistroXForBlueprintV1ByCrn posts stack for blueprint by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostDistroXForBlueprintV1ByCrn(params *PostDistroXForBlueprintV1ByCrnParams) (*PostDistroXForBlueprintV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDistroXForBlueprintV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDistroXForBlueprintV1ByCrn",
		Method:             "POST",
		PathPattern:        "/v1/distrox/crn/{crn}/blueprint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDistroXForBlueprintV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDistroXForBlueprintV1ByCrnOK), nil

}

/*
PostDistroXForBlueprintV1ByName posts stack for blueprint by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostDistroXForBlueprintV1ByName(params *PostDistroXForBlueprintV1ByNameParams) (*PostDistroXForBlueprintV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDistroXForBlueprintV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDistroXForBlueprintV1ByName",
		Method:             "POST",
		PathPattern:        "/v1/distrox/name/{name}/blueprint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDistroXForBlueprintV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDistroXForBlueprintV1ByNameOK), nil

}

/*
PostDistroXV1 creates stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostDistroXV1(params *PostDistroXV1Params) (*PostDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDistroXV1",
		Method:             "POST",
		PathPattern:        "/v1/distrox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDistroXV1OK), nil

}

/*
PutScalingDistroXV1ByCrn scales the stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutScalingDistroXV1ByCrn(params *PutScalingDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutScalingDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScalingDistroXV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/scaling",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutScalingDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutScalingDistroXV1ByName scales the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutScalingDistroXV1ByName(params *PutScalingDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutScalingDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScalingDistroXV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/scaling",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutScalingDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairDistroXV1ByCrn repairs the stack by crn

Removing the failed nodes and starting new nodes to substitute them.
*/
func (a *Client) RepairDistroXV1ByCrn(params *RepairDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairDistroXV1ByCrn",
		Method:             "POST",
		PathPattern:        "/v1/distrox/crn/{crn}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairDistroXV1ByName repairs the stack by name

Removing the failed nodes and starting new nodes to substitute them.
*/
func (a *Client) RepairDistroXV1ByName(params *RepairDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairDistroXV1ByName",
		Method:             "POST",
		PathPattern:        "/v1/distrox/name/{name}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RetryDistroXV1ByCrn retries the stack by crn

Failed or interrupted stack and cluster operations can be retried, after the cause of the failure was eliminated. The operations will continue at the state, where the previous process failed.
*/
func (a *Client) RetryDistroXV1ByCrn(params *RetryDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retryDistroXV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetryDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RetryDistroXV1ByName retries the stack by name

Failed or interrupted stack and cluster operations can be retried, after the cause of the failure was eliminated. The operations will continue at the state, where the previous process failed.
*/
func (a *Client) RetryDistroXV1ByName(params *RetryDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retryDistroXV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetryDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SetDistroXMaintenanceModeByCrn sets maintenance mode for the cluster by crn

Setting maintenance mode for the cluster in order to be able to update Ambari and/or the Hadoop stack.
*/
func (a *Client) SetDistroXMaintenanceModeByCrn(params *SetDistroXMaintenanceModeByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDistroXMaintenanceModeByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDistroXMaintenanceModeByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/maintenance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetDistroXMaintenanceModeByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SetDistroXMaintenanceModeByName sets maintenance mode for the cluster by name

Setting maintenance mode for the cluster in order to be able to update Ambari and/or the Hadoop stack.
*/
func (a *Client) SetDistroXMaintenanceModeByName(params *SetDistroXMaintenanceModeByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDistroXMaintenanceModeByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDistroXMaintenanceModeByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/maintenance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetDistroXMaintenanceModeByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartDistroXV1ByCrn starts the stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StartDistroXV1ByCrn(params *StartDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startDistroXV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartDistroXV1ByName starts the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StartDistroXV1ByName(params *StartDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startDistroXV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StatusDistroXV1ByCrn retrieves stack status by stack crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusDistroXV1ByCrn(params *StatusDistroXV1ByCrnParams) (*StatusDistroXV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusDistroXV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusDistroXV1ByCrn",
		Method:             "GET",
		PathPattern:        "/v1/distrox/crn/{crn}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusDistroXV1ByCrnOK), nil

}

/*
StatusDistroXV1ByName retrieves stack status by stack name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusDistroXV1ByName(params *StatusDistroXV1ByNameParams) (*StatusDistroXV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusDistroXV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusDistroXV1ByName",
		Method:             "GET",
		PathPattern:        "/v1/distrox/name/{name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusDistroXV1ByNameOK), nil

}

/*
StopDistroXV1ByCrn stops the stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StopDistroXV1ByCrn(params *StopDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopDistroXV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopDistroXV1ByName stops the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StopDistroXV1ByName(params *StopDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopDistroXV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopDistroXV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncDistroXV1ByCrn syncs the stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) SyncDistroXV1ByCrn(params *SyncDistroXV1ByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncDistroXV1ByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncDistroXV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/crn/{crn}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncDistroXV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncDistroXV1ByName syncs the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) SyncDistroXV1ByName(params *SyncDistroXV1ByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncDistroXV1ByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncDistroXV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/name/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncDistroXV1ByNameReader{formats: a.formats},
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
