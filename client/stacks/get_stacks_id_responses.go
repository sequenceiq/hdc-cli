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

// GetStacksIDReader is a Reader for the GetStacksID structure.
type GetStacksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStacksIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStacksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStacksIDOK creates a GetStacksIDOK with default headers values
func NewGetStacksIDOK() *GetStacksIDOK {
	return &GetStacksIDOK{}
}

/*GetStacksIDOK handles this case with default header values.

successful operation
*/
type GetStacksIDOK struct {
	Payload *models.StackResponse
}

func (o *GetStacksIDOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] getStacksIdOK  %+v", 200, o.Payload)
}

func (o *GetStacksIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
