// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// CleanupV1Reader is a Reader for the CleanupV1 structure.
type CleanupV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CleanupV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCleanupV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCleanupV1OK creates a CleanupV1OK with default headers values
func NewCleanupV1OK() *CleanupV1OK {
	return &CleanupV1OK{}
}

/*CleanupV1OK handles this case with default header values.

successful operation
*/
type CleanupV1OK struct {
	Payload *model.OperationV1Status
}

func (o *CleanupV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/cleanup][%d] cleanupV1OK  %+v", 200, o.Payload)
}

func (o *CleanupV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}