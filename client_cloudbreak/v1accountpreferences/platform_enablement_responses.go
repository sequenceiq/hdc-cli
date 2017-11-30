// Code generated by go-swagger; DO NOT EDIT.

package v1accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// PlatformEnablementReader is a Reader for the PlatformEnablement structure.
type PlatformEnablementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlatformEnablementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPlatformEnablementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPlatformEnablementOK creates a PlatformEnablementOK with default headers values
func NewPlatformEnablementOK() *PlatformEnablementOK {
	return &PlatformEnablementOK{}
}

/*PlatformEnablementOK handles this case with default header values.

successful operation
*/
type PlatformEnablementOK struct {
	Payload PlatformEnablementOKBody
}

func (o *PlatformEnablementOK) Error() string {
	return fmt.Sprintf("[GET /v1/accountpreferences/platformenabled][%d] platformEnablementOK  %+v", 200, o.Payload)
}

func (o *PlatformEnablementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PlatformEnablementOKBody platform enablement o k body
swagger:model PlatformEnablementOKBody
*/

type PlatformEnablementOKBody map[string]bool

// Validate validates this platform enablement o k body
func (o PlatformEnablementOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
