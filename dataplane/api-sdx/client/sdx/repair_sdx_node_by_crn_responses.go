// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// RepairSdxNodeByCrnReader is a Reader for the RepairSdxNodeByCrn structure.
type RepairSdxNodeByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairSdxNodeByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepairSdxNodeByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepairSdxNodeByCrnOK creates a RepairSdxNodeByCrnOK with default headers values
func NewRepairSdxNodeByCrnOK() *RepairSdxNodeByCrnOK {
	return &RepairSdxNodeByCrnOK{}
}

/*RepairSdxNodeByCrnOK handles this case with default header values.

successful operation
*/
type RepairSdxNodeByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *RepairSdxNodeByCrnOK) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/manual_repair][%d] repairSdxNodeByCrnOK  %+v", 200, o.Payload)
}

func (o *RepairSdxNodeByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}