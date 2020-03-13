// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualMachinesResponse virtual machines response
// swagger:model VirtualMachinesResponse
type VirtualMachinesResponse struct {

	// default virtual machines
	DefaultVirtualMachine *VMTypeResponse `json:"defaultVirtualMachine,omitempty"`

	// virtual machines
	// Unique: true
	VirtualMachines []*VMTypeResponse `json:"virtualMachines"`
}

// Validate validates this virtual machines response
func (m *VirtualMachinesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachinesResponse) validateDefaultVirtualMachine(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultVirtualMachine) { // not required
		return nil
	}

	if m.DefaultVirtualMachine != nil {
		if err := m.DefaultVirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultVirtualMachine")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachinesResponse) validateVirtualMachines(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachines) { // not required
		return nil
	}

	if err := validate.UniqueItems("virtualMachines", "body", m.VirtualMachines); err != nil {
		return err
	}

	for i := 0; i < len(m.VirtualMachines); i++ {
		if swag.IsZero(m.VirtualMachines[i]) { // not required
			continue
		}

		if m.VirtualMachines[i] != nil {
			if err := m.VirtualMachines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachinesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachinesResponse) UnmarshalBinary(b []byte) error {
	var res VirtualMachinesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}