// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeImageDistroXV1Reader is a Reader for the ChangeImageDistroXV1 structure.
type ChangeImageDistroXV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeImageDistroXV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewChangeImageDistroXV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewChangeImageDistroXV1Default creates a ChangeImageDistroXV1Default with default headers values
func NewChangeImageDistroXV1Default(code int) *ChangeImageDistroXV1Default {
	return &ChangeImageDistroXV1Default{
		_statusCode: code,
	}
}

/*ChangeImageDistroXV1Default handles this case with default header values.

successful operation
*/
type ChangeImageDistroXV1Default struct {
	_statusCode int
}

// Code gets the status code for the change image distro x v1 default response
func (o *ChangeImageDistroXV1Default) Code() int {
	return o._statusCode
}

func (o *ChangeImageDistroXV1Default) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/{name}/change_image][%d] changeImageDistroXV1 default ", o._statusCode)
}

func (o *ChangeImageDistroXV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}