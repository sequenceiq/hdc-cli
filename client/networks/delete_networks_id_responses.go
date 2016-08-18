package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteNetworksIDReader is a Reader for the DeleteNetworksID structure.
type DeleteNetworksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteNetworksIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteNetworksIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteNetworksIDDefault creates a DeleteNetworksIDDefault with default headers values
func NewDeleteNetworksIDDefault(code int) *DeleteNetworksIDDefault {
	return &DeleteNetworksIDDefault{
		_statusCode: code,
	}
}

/*DeleteNetworksIDDefault handles this case with default header values.

successful operation
*/
type DeleteNetworksIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete networks ID default response
func (o *DeleteNetworksIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{id}][%d] DeleteNetworksID default ", o._statusCode)
}

func (o *DeleteNetworksIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
