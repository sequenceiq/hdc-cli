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

// PostStacksAccountReader is a Reader for the PostStacksAccount structure.
type PostStacksAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostStacksAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostStacksAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStacksAccountOK creates a PostStacksAccountOK with default headers values
func NewPostStacksAccountOK() *PostStacksAccountOK {
	return &PostStacksAccountOK{}
}

/*PostStacksAccountOK handles this case with default header values.

successful operation
*/
type PostStacksAccountOK struct {
	Payload *models.ID
}

func (o *PostStacksAccountOK) Error() string {
	return fmt.Sprintf("[POST /stacks/account][%d] postStacksAccountOK  %+v", 200, o.Payload)
}

func (o *PostStacksAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
