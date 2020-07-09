// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// StartDistroXV1ByNameReader is a Reader for the StartDistroXV1ByName structure.
type StartDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartDistroXV1ByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartDistroXV1ByNameOK creates a StartDistroXV1ByNameOK with default headers values
func NewStartDistroXV1ByNameOK() *StartDistroXV1ByNameOK {
	return &StartDistroXV1ByNameOK{}
}

/*StartDistroXV1ByNameOK handles this case with default header values.

successful operation
*/
type StartDistroXV1ByNameOK struct {
	Payload *model.FlowIdentifier
}

func (o *StartDistroXV1ByNameOK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/name/{name}/start][%d] startDistroXV1ByNameOK  %+v", 200, o.Payload)
}

func (o *StartDistroXV1ByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
