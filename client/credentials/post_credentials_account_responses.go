package credentials

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

// PostCredentialsAccountReader is a Reader for the PostCredentialsAccount structure.
type PostCredentialsAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostCredentialsAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCredentialsAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCredentialsAccountOK creates a PostCredentialsAccountOK with default headers values
func NewPostCredentialsAccountOK() *PostCredentialsAccountOK {
	return &PostCredentialsAccountOK{}
}

/*PostCredentialsAccountOK handles this case with default header values.

successful operation
*/
type PostCredentialsAccountOK struct {
	Payload *models.ID
}

func (o *PostCredentialsAccountOK) Error() string {
	return fmt.Sprintf("[POST /credentials/account][%d] postCredentialsAccountOK  %+v", 200, o.Payload)
}

func (o *PostCredentialsAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
