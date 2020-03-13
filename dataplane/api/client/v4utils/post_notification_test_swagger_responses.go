// Code generated by go-swagger; DO NOT EDIT.

package v4utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PostNotificationTestReader is a Reader for the PostNotificationTest structure.
type PostNotificationTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNotificationTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNotificationTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNotificationTestOK creates a PostNotificationTestOK with default headers values
func NewPostNotificationTestOK() *PostNotificationTestOK {
	return &PostNotificationTestOK{}
}

/*PostNotificationTestOK handles this case with default header values.

successful operation
*/
type PostNotificationTestOK struct {
	Payload *model.ResourceEventResponse
}

func (o *PostNotificationTestOK) Error() string {
	return fmt.Sprintf("[POST /v4/utils/notification_test][%d] postNotificationTestOK  %+v", 200, o.Payload)
}

func (o *PostNotificationTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ResourceEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}