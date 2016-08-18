package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*RecipeRequest recipe request

swagger:model RecipeRequest
*/
type RecipeRequest struct {

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 1
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* list of consul plugins with execution types

	Required: true
	Unique: true
	*/
	Plugins []string `json:"plugins"`

	/* additional plugin properties
	 */
	Properties map[string]string `json:"properties,omitempty"`

	/* recipe timeout in minutes
	 */
	Timeout *int32 `json:"timeout,omitempty"`
}

// Validate validates this recipe request
func (m *RecipeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeRequest) validateDescription(formats strfmt.Registry) error {

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

func (m *RecipeRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
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

func (m *RecipeRequest) validatePlugins(formats strfmt.Registry) error {

	if err := validate.Required("plugins", "body", m.Plugins); err != nil {
		return err
	}

	if err := validate.UniqueItems("plugins", "body", m.Plugins); err != nil {
		return err
	}

	for i := 0; i < len(m.Plugins); i++ {

		if err := validate.RequiredString("plugins"+"."+strconv.Itoa(i), "body", string(m.Plugins[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *RecipeRequest) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if err := validate.Required("properties", "body", m.Properties); err != nil {
		return err
	}

	return nil
}
