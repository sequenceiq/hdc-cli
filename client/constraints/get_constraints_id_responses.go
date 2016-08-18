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

// GetConstraintsIDReader is a Reader for the GetConstraintsID structure.
type GetConstraintsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetConstraintsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConstraintsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConstraintsIDOK creates a GetConstraintsIDOK with default headers values
func NewGetConstraintsIDOK() *GetConstraintsIDOK {
	return &GetConstraintsIDOK{}
}

/*GetConstraintsIDOK handles this case with default header values.

successful operation
*/
type GetConstraintsIDOK struct {
	Payload *models.ConstraintTemplateResponse
}

func (o *GetConstraintsIDOK) Error() string {
	return fmt.Sprintf("[GET /constraints/{id}][%d] getConstraintsIdOK  %+v", 200, o.Payload)
}

func (o *GetConstraintsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstraintTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
