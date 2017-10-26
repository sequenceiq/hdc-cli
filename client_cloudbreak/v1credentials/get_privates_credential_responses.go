// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivatesCredentialReader is a Reader for the GetPrivatesCredential structure.
type GetPrivatesCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesCredentialOK creates a GetPrivatesCredentialOK with default headers values
func NewGetPrivatesCredentialOK() *GetPrivatesCredentialOK {
	return &GetPrivatesCredentialOK{}
}

/*GetPrivatesCredentialOK handles this case with default header values.

successful operation
*/
type GetPrivatesCredentialOK struct {
	Payload []*models_cloudbreak.CredentialResponse
}

func (o *GetPrivatesCredentialOK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/user][%d] getPrivatesCredentialOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}