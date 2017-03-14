package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*TemplateResponse template response

swagger:model TemplateResponse
*/
type TemplateResponse struct {

	/* type of cloud provider

	Required: true
	*/
	CloudPlatform string `json:"cloudPlatform"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* type of the instance

	Required: true
	*/
	InstanceType string `json:"instanceType"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* cloud specific parameters for template
	 */
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	/* resource is visible in account
	 */
	Public *bool `json:"public,omitempty"`

	/* id of the topology the resource belongs to
	 */
	TopologyID *int64 `json:"topologyId,omitempty"`

	/* number of volumes

	Required: true
	*/
	VolumeCount int32 `json:"volumeCount"`

	/* size of volumes

	Required: true
	*/
	VolumeSize int32 `json:"volumeSize"`

	/* type of the volumes
	 */
	VolumeType *string `json:"volumeType,omitempty"`
}

// Validate validates this template response
func (m *TemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateResponse) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.RequiredString("cloudPlatform", "body", string(m.CloudPlatform)); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("instanceType", "body", string(m.InstanceType)); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateVolumeCount(formats strfmt.Registry) error {

	if err := validate.Required("volumeCount", "body", int32(m.VolumeCount)); err != nil {
		return err
	}

	return nil
}

func (m *TemplateResponse) validateVolumeSize(formats strfmt.Registry) error {

	if err := validate.Required("volumeSize", "body", int32(m.VolumeSize)); err != nil {
		return err
	}

	return nil
}