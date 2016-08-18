package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteRdsconfigsAccountNameReader is a Reader for the DeleteRdsconfigsAccountName structure.
type DeleteRdsconfigsAccountNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteRdsconfigsAccountNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteRdsconfigsAccountNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteRdsconfigsAccountNameDefault creates a DeleteRdsconfigsAccountNameDefault with default headers values
func NewDeleteRdsconfigsAccountNameDefault(code int) *DeleteRdsconfigsAccountNameDefault {
	return &DeleteRdsconfigsAccountNameDefault{
		_statusCode: code,
	}
}

/*DeleteRdsconfigsAccountNameDefault handles this case with default header values.

successful operation
*/
type DeleteRdsconfigsAccountNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete rdsconfigs account name default response
func (o *DeleteRdsconfigsAccountNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRdsconfigsAccountNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /rdsconfigs/account/{name}][%d] DeleteRdsconfigsAccountName default ", o._statusCode)
}

func (o *DeleteRdsconfigsAccountNameDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
