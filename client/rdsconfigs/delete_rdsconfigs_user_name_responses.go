package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteRdsconfigsUserNameReader is a Reader for the DeleteRdsconfigsUserName structure.
type DeleteRdsconfigsUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteRdsconfigsUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteRdsconfigsUserNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteRdsconfigsUserNameDefault creates a DeleteRdsconfigsUserNameDefault with default headers values
func NewDeleteRdsconfigsUserNameDefault(code int) *DeleteRdsconfigsUserNameDefault {
	return &DeleteRdsconfigsUserNameDefault{
		_statusCode: code,
	}
}

/*DeleteRdsconfigsUserNameDefault handles this case with default header values.

successful operation
*/
type DeleteRdsconfigsUserNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete rdsconfigs user name default response
func (o *DeleteRdsconfigsUserNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRdsconfigsUserNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /rdsconfigs/user/{name}][%d] DeleteRdsconfigsUserName default ", o._statusCode)
}

func (o *DeleteRdsconfigsUserNameDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
