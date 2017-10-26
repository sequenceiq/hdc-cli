// Code generated by go-swagger; DO NOT EDIT.

package v1constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicConstraintReader is a Reader for the GetPublicConstraint structure.
type GetPublicConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicConstraintOK creates a GetPublicConstraintOK with default headers values
func NewGetPublicConstraintOK() *GetPublicConstraintOK {
	return &GetPublicConstraintOK{}
}

/*GetPublicConstraintOK handles this case with default header values.

successful operation
*/
type GetPublicConstraintOK struct {
	Payload *models_cloudbreak.ConstraintTemplateResponse
}

func (o *GetPublicConstraintOK) Error() string {
	return fmt.Sprintf("[GET /v1/constraints/account/{name}][%d] getPublicConstraintOK  %+v", 200, o.Payload)
}

func (o *GetPublicConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ConstraintTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}