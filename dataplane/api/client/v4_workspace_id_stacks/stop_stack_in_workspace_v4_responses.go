// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StopStackInWorkspaceV4Reader is a Reader for the StopStackInWorkspaceV4 structure.
type StopStackInWorkspaceV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopStackInWorkspaceV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStopStackInWorkspaceV4Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStopStackInWorkspaceV4Default creates a StopStackInWorkspaceV4Default with default headers values
func NewStopStackInWorkspaceV4Default(code int) *StopStackInWorkspaceV4Default {
	return &StopStackInWorkspaceV4Default{
		_statusCode: code,
	}
}

/*StopStackInWorkspaceV4Default handles this case with default header values.

successful operation
*/
type StopStackInWorkspaceV4Default struct {
	_statusCode int
}

// Code gets the status code for the stop stack in workspace v4 default response
func (o *StopStackInWorkspaceV4Default) Code() int {
	return o._statusCode
}

func (o *StopStackInWorkspaceV4Default) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/{name}/stop][%d] stopStackInWorkspaceV4 default ", o._statusCode)
}

func (o *StopStackInWorkspaceV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}