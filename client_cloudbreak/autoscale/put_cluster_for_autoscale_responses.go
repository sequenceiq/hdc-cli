// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClusterForAutoscaleReader is a Reader for the PutClusterForAutoscale structure.
type PutClusterForAutoscaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterForAutoscaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutClusterForAutoscaleDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutClusterForAutoscaleDefault creates a PutClusterForAutoscaleDefault with default headers values
func NewPutClusterForAutoscaleDefault(code int) *PutClusterForAutoscaleDefault {
	return &PutClusterForAutoscaleDefault{
		_statusCode: code,
	}
}

/*PutClusterForAutoscaleDefault handles this case with default header values.

successful operation
*/
type PutClusterForAutoscaleDefault struct {
	_statusCode int
}

// Code gets the status code for the put cluster for autoscale default response
func (o *PutClusterForAutoscaleDefault) Code() int {
	return o._statusCode
}

func (o *PutClusterForAutoscaleDefault) Error() string {
	return fmt.Sprintf("[PUT /autoscale/stack/{id}/{owner}/cluster][%d] putClusterForAutoscale default ", o._statusCode)
}

func (o *PutClusterForAutoscaleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
