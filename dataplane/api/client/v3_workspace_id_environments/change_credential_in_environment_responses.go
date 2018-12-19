// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ChangeCredentialInEnvironmentReader is a Reader for the ChangeCredentialInEnvironment structure.
type ChangeCredentialInEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeCredentialInEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeCredentialInEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeCredentialInEnvironmentOK creates a ChangeCredentialInEnvironmentOK with default headers values
func NewChangeCredentialInEnvironmentOK() *ChangeCredentialInEnvironmentOK {
	return &ChangeCredentialInEnvironmentOK{}
}

/*ChangeCredentialInEnvironmentOK handles this case with default header values.

successful operation
*/
type ChangeCredentialInEnvironmentOK struct {
	Payload *model.DetailedEnvironmentResponse
}

func (o *ChangeCredentialInEnvironmentOK) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/environments/{name}/changeCredential][%d] changeCredentialInEnvironmentOK  %+v", 200, o.Payload)
}

func (o *ChangeCredentialInEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedEnvironmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}