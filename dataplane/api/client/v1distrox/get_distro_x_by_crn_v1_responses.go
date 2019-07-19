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

// GetDistroXByCrnV1Reader is a Reader for the GetDistroXByCrnV1 structure.
type GetDistroXByCrnV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroXByCrnV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroXByCrnV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroXByCrnV1OK creates a GetDistroXByCrnV1OK with default headers values
func NewGetDistroXByCrnV1OK() *GetDistroXByCrnV1OK {
	return &GetDistroXByCrnV1OK{}
}

/*GetDistroXByCrnV1OK handles this case with default header values.

successful operation
*/
type GetDistroXByCrnV1OK struct {
	Payload *model.StackV4Response
}

func (o *GetDistroXByCrnV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/crn/{crn}][%d] getDistroXByCrnV1OK  %+v", 200, o.Payload)
}

func (o *GetDistroXByCrnV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}