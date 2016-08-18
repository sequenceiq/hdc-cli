package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteRdsconfigsIDReader is a Reader for the DeleteRdsconfigsID structure.
type DeleteRdsconfigsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteRdsconfigsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteRdsconfigsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteRdsconfigsIDDefault creates a DeleteRdsconfigsIDDefault with default headers values
func NewDeleteRdsconfigsIDDefault(code int) *DeleteRdsconfigsIDDefault {
	return &DeleteRdsconfigsIDDefault{
		_statusCode: code,
	}
}

/*DeleteRdsconfigsIDDefault handles this case with default header values.

successful operation
*/
type DeleteRdsconfigsIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete rdsconfigs ID default response
func (o *DeleteRdsconfigsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRdsconfigsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /rdsconfigs/{id}][%d] DeleteRdsconfigsID default ", o._statusCode)
}

func (o *DeleteRdsconfigsIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
