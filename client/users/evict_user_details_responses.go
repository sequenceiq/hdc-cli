package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// EvictUserDetailsReader is a Reader for the EvictUserDetails structure.
type EvictUserDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *EvictUserDetailsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEvictUserDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEvictUserDetailsOK creates a EvictUserDetailsOK with default headers values
func NewEvictUserDetailsOK() *EvictUserDetailsOK {
	return &EvictUserDetailsOK{}
}

/*EvictUserDetailsOK handles this case with default header values.

successful operation
*/
type EvictUserDetailsOK struct {
	Payload string
}

func (o *EvictUserDetailsOK) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] evictUserDetailsOK  %+v", 200, o.Payload)
}

func (o *EvictUserDetailsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
