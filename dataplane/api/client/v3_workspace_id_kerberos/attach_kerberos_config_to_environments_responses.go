// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_kerberos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// AttachKerberosConfigToEnvironmentsReader is a Reader for the AttachKerberosConfigToEnvironments structure.
type AttachKerberosConfigToEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachKerberosConfigToEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttachKerberosConfigToEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAttachKerberosConfigToEnvironmentsOK creates a AttachKerberosConfigToEnvironmentsOK with default headers values
func NewAttachKerberosConfigToEnvironmentsOK() *AttachKerberosConfigToEnvironmentsOK {
	return &AttachKerberosConfigToEnvironmentsOK{}
}

/*AttachKerberosConfigToEnvironmentsOK handles this case with default header values.

successful operation
*/
type AttachKerberosConfigToEnvironmentsOK struct {
	Payload *model.KerberosResponse
}

func (o *AttachKerberosConfigToEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/kerberos/{name}/attach][%d] attachKerberosConfigToEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *AttachKerberosConfigToEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.KerberosResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}