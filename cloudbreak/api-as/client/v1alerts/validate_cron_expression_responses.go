// Code generated by go-swagger; DO NOT EDIT.

package v1alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ValidateCronExpressionReader is a Reader for the ValidateCronExpression structure.
type ValidateCronExpressionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCronExpressionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateCronExpressionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCronExpressionOK creates a ValidateCronExpressionOK with default headers values
func NewValidateCronExpressionOK() *ValidateCronExpressionOK {
	return &ValidateCronExpressionOK{}
}

/*ValidateCronExpressionOK handles this case with default header values.

successful operation
*/
type ValidateCronExpressionOK struct {
	Payload bool
}

func (o *ValidateCronExpressionOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/alerts/time/validate][%d] validateCronExpressionOK  %+v", 200, o.Payload)
}

func (o *ValidateCronExpressionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
