package topologies

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

// GetTopologiesAccountIDReader is a Reader for the GetTopologiesAccountID structure.
type GetTopologiesAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetTopologiesAccountIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTopologiesAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTopologiesAccountIDOK creates a GetTopologiesAccountIDOK with default headers values
func NewGetTopologiesAccountIDOK() *GetTopologiesAccountIDOK {
	return &GetTopologiesAccountIDOK{}
}

/*GetTopologiesAccountIDOK handles this case with default header values.

successful operation
*/
type GetTopologiesAccountIDOK struct {
	Payload *models.TopologyResponse
}

func (o *GetTopologiesAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /topologies/account/{id}][%d] getTopologiesAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetTopologiesAccountIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopologyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
