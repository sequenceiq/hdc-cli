package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/sequenceiq/hdc-cli/models"
)

// GetStackForAmbariReader is a Reader for the GetStackForAmbari structure.
type GetStackForAmbariReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStackForAmbariReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStackForAmbariOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStackForAmbariOK creates a GetStackForAmbariOK with default headers values
func NewGetStackForAmbariOK() *GetStackForAmbariOK {
	return &GetStackForAmbariOK{}
}

/*GetStackForAmbariOK handles this case with default header values.

successful operation
*/
type GetStackForAmbariOK struct {
	Payload *models.StackResponse
}

func (o *GetStackForAmbariOK) Error() string {
	return fmt.Sprintf("[POST /stacks/ambari][%d] getStackForAmbariOK  %+v", 200, o.Payload)
}

func (o *GetStackForAmbariOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
