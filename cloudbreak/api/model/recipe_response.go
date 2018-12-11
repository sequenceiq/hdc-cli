// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecipeResponse recipe response
// swagger:model RecipeResponse

type RecipeResponse struct {

	// content of recipe
	Content string `json:"content,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name string `json:"name,omitempty"`

	// type of recipe
	// Required: true
	RecipeType *string `json:"recipeType"`

	// workspace of the resource
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

/* polymorph RecipeResponse content false */

/* polymorph RecipeResponse description false */

/* polymorph RecipeResponse id false */

/* polymorph RecipeResponse name false */

/* polymorph RecipeResponse recipeType false */

/* polymorph RecipeResponse workspace false */

// Validate validates this recipe response
func (m *RecipeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *RecipeResponse) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

var recipeResponseTypeRecipeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRE_AMBARI_START","PRE_TERMINATION","POST_AMBARI_START","POST_CLUSTER_INSTALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recipeResponseTypeRecipeTypePropEnum = append(recipeResponseTypeRecipeTypePropEnum, v)
	}
}

const (
	// RecipeResponseRecipeTypePREAMBARISTART captures enum value "PRE_AMBARI_START"
	RecipeResponseRecipeTypePREAMBARISTART string = "PRE_AMBARI_START"
	// RecipeResponseRecipeTypePRETERMINATION captures enum value "PRE_TERMINATION"
	RecipeResponseRecipeTypePRETERMINATION string = "PRE_TERMINATION"
	// RecipeResponseRecipeTypePOSTAMBARISTART captures enum value "POST_AMBARI_START"
	RecipeResponseRecipeTypePOSTAMBARISTART string = "POST_AMBARI_START"
	// RecipeResponseRecipeTypePOSTCLUSTERINSTALL captures enum value "POST_CLUSTER_INSTALL"
	RecipeResponseRecipeTypePOSTCLUSTERINSTALL string = "POST_CLUSTER_INSTALL"
)

// prop value enum
func (m *RecipeResponse) validateRecipeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recipeResponseTypeRecipeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RecipeResponse) validateRecipeType(formats strfmt.Registry) error {

	if err := validate.Required("recipeType", "body", m.RecipeType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipeTypeEnum("recipeType", "body", *m.RecipeType); err != nil {
		return err
	}

	return nil
}

func (m *RecipeResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {

		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecipeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeResponse) UnmarshalBinary(b []byte) error {
	var res RecipeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
