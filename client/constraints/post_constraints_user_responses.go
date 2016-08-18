package constraints

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

// PostConstraintsUserReader is a Reader for the PostConstraintsUser structure.
type PostConstraintsUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostConstraintsUserReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostConstraintsUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostConstraintsUserOK creates a PostConstraintsUserOK with default headers values
func NewPostConstraintsUserOK() *PostConstraintsUserOK {
	return &PostConstraintsUserOK{}
}

/*PostConstraintsUserOK handles this case with default header values.

successful operation
*/
type PostConstraintsUserOK struct {
	Payload *models.ID
}

func (o *PostConstraintsUserOK) Error() string {
	return fmt.Sprintf("[POST /constraints/user][%d] postConstraintsUserOK  %+v", 200, o.Payload)
}

func (o *PostConstraintsUserOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
