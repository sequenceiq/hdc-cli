package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetDailyFlexUsageReader is a Reader for the GetDailyFlexUsage structure.
type GetDailyFlexUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDailyFlexUsageReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDailyFlexUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDailyFlexUsageOK creates a GetDailyFlexUsageOK with default headers values
func NewGetDailyFlexUsageOK() *GetDailyFlexUsageOK {
	return &GetDailyFlexUsageOK{}
}

/*GetDailyFlexUsageOK handles this case with default header values.

successful operation
*/
type GetDailyFlexUsageOK struct {
	Payload *models_cloudbreak.CloudbreakFlexUsage
}

func (o *GetDailyFlexUsageOK) Error() string {
	return fmt.Sprintf("[GET /usages/flex/daily][%d] getDailyFlexUsageOK  %+v", 200, o.Payload)
}

func (o *GetDailyFlexUsageOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.CloudbreakFlexUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
