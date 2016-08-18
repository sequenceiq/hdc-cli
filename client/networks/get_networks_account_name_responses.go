package networks

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

// GetNetworksAccountNameReader is a Reader for the GetNetworksAccountName structure.
type GetNetworksAccountNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNetworksAccountNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworksAccountNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworksAccountNameOK creates a GetNetworksAccountNameOK with default headers values
func NewGetNetworksAccountNameOK() *GetNetworksAccountNameOK {
	return &GetNetworksAccountNameOK{}
}

/*GetNetworksAccountNameOK handles this case with default header values.

successful operation
*/
type GetNetworksAccountNameOK struct {
	Payload *models.NetworkJSON
}

func (o *GetNetworksAccountNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/account/{name}][%d] getNetworksAccountNameOK  %+v", 200, o.Payload)
}

func (o *GetNetworksAccountNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
