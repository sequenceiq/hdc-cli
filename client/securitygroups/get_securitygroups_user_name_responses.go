package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/sequenceiq/hdc-cli/models"
)

// GetSecuritygroupsUserNameReader is a Reader for the GetSecuritygroupsUserName structure.
type GetSecuritygroupsUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSecuritygroupsUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecuritygroupsUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecuritygroupsUserNameOK creates a GetSecuritygroupsUserNameOK with default headers values
func NewGetSecuritygroupsUserNameOK() *GetSecuritygroupsUserNameOK {
	return &GetSecuritygroupsUserNameOK{}
}

/*GetSecuritygroupsUserNameOK handles this case with default header values.

successful operation
*/
type GetSecuritygroupsUserNameOK struct {
	Payload *models.SecurityGroupJSON
}

func (o *GetSecuritygroupsUserNameOK) Error() string {
	return fmt.Sprintf("[GET /securitygroups/user/{name}][%d] getSecuritygroupsUserNameOK  %+v", 200, o.Payload)
}

func (o *GetSecuritygroupsUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
