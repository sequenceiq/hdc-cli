// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// CheckForClusterUpgradeByCrnReader is a Reader for the CheckForClusterUpgradeByCrn structure.
type CheckForClusterUpgradeByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckForClusterUpgradeByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckForClusterUpgradeByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckForClusterUpgradeByCrnOK creates a CheckForClusterUpgradeByCrnOK with default headers values
func NewCheckForClusterUpgradeByCrnOK() *CheckForClusterUpgradeByCrnOK {
	return &CheckForClusterUpgradeByCrnOK{}
}

/*CheckForClusterUpgradeByCrnOK handles this case with default header values.

successful operation
*/
type CheckForClusterUpgradeByCrnOK struct {
	Payload *model.UpgradeOptionsV4Response
}

func (o *CheckForClusterUpgradeByCrnOK) Error() string {
	return fmt.Sprintf("[GET /sdx/crn/{crn}/check_cluster_upgrade][%d] checkForClusterUpgradeByCrnOK  %+v", 200, o.Payload)
}

func (o *CheckForClusterUpgradeByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.UpgradeOptionsV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}