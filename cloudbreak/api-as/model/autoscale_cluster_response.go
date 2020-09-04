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

// AutoscaleClusterResponse autoscale cluster response
// swagger:model AutoscaleClusterResponse
type AutoscaleClusterResponse struct {

	// Indicate that the Autoscaling feature set is Enabled or Disabled
	AutoscalingEnabled *bool `json:"autoscalingEnabled,omitempty"`

	// Ambari server host address
	Host string `json:"host,omitempty"`

	// Id of the cluster
	ID int64 `json:"id,omitempty"`

	// Metric based alerts of the cluster
	MetricAlerts []*MetricAlertResponse `json:"metricAlerts"`

	// Ambari server port
	Port string `json:"port,omitempty"`

	// Prometheus based alerts of the cluster
	PrometheusAlerts []*PrometheusAlertResponse `json:"prometheusAlerts"`

	// Scaling configuration for the cluster
	ScalingConfiguration *ScalingConfiguration `json:"scalingConfiguration,omitempty"`

	// Id of the stack in Cloudbreak
	// Required: true
	StackID *int64 `json:"stackId"`

	// State of the cluster
	State string `json:"state,omitempty"`

	// Time based alerts of the cluster
	TimeAlerts []*TimeAlertResponse `json:"timeAlerts"`

	// Ambari server username
	User string `json:"user,omitempty"`
}

// Validate validates this autoscale cluster response
func (m *AutoscaleClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrometheusAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalingConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeAlerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscaleClusterResponse) validateMetricAlerts(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricAlerts); i++ {
		if swag.IsZero(m.MetricAlerts[i]) { // not required
			continue
		}

		if m.MetricAlerts[i] != nil {
			if err := m.MetricAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metricAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoscaleClusterResponse) validatePrometheusAlerts(formats strfmt.Registry) error {

	if swag.IsZero(m.PrometheusAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.PrometheusAlerts); i++ {
		if swag.IsZero(m.PrometheusAlerts[i]) { // not required
			continue
		}

		if m.PrometheusAlerts[i] != nil {
			if err := m.PrometheusAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prometheusAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoscaleClusterResponse) validateScalingConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.ScalingConfiguration) { // not required
		return nil
	}

	if m.ScalingConfiguration != nil {
		if err := m.ScalingConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scalingConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateStackID(formats strfmt.Registry) error {

	if err := validate.Required("stackId", "body", m.StackID); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateTimeAlerts(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.TimeAlerts); i++ {
		if swag.IsZero(m.TimeAlerts[i]) { // not required
			continue
		}

		if m.TimeAlerts[i] != nil {
			if err := m.TimeAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoscaleClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscaleClusterResponse) UnmarshalBinary(b []byte) error {
	var res AutoscaleClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
