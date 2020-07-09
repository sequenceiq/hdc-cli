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

// PostDistroXForBlueprintV1Reader is a Reader for the PostDistroXForBlueprintV1 structure.
type PostDistroXForBlueprintV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDistroXForBlueprintV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDistroXForBlueprintV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDistroXForBlueprintV1OK creates a PostDistroXForBlueprintV1OK with default headers values
func NewPostDistroXForBlueprintV1OK() *PostDistroXForBlueprintV1OK {
	return &PostDistroXForBlueprintV1OK{}
}

/*PostDistroXForBlueprintV1OK handles this case with default header values.

successful operation
*/
type PostDistroXForBlueprintV1OK struct {
	Payload *model.GeneratedBlueprintV4Response
}

func (o *PostDistroXForBlueprintV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/blueprint][%d] postDistroXForBlueprintV1OK  %+v", 200, o.Payload)
}

func (o *PostDistroXForBlueprintV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GeneratedBlueprintV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
