package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetAllSettingsReader is a Reader for the GetAllSettings structure.
type GetAllSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAllSettingsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllSettingsOK creates a GetAllSettingsOK with default headers values
func NewGetAllSettingsOK() *GetAllSettingsOK {
	return &GetAllSettingsOK{}
}

/*GetAllSettingsOK handles this case with default header values.

successful operation
*/
type GetAllSettingsOK struct {
	Payload GetAllSettingsOKBodyBody
}

func (o *GetAllSettingsOK) Error() string {
	return fmt.Sprintf("[GET /settings/all][%d] getAllSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAllSettingsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAllSettingsOKBodyBody get all settings o k body body

swagger:model GetAllSettingsOKBodyBody
*/
type GetAllSettingsOKBodyBody map[string]map[string]interface{}

// Validate validates this get all settings o k body body
func (o GetAllSettingsOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getAllSettingsOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}