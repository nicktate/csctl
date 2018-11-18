// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginCompatibility A Containership plugin compatibility definition
// swagger:model PluginCompatibility
type PluginCompatibility struct {

	// The definition for Kubernetes compatibility
	// Required: true
	Kubernetes *PluginKubernetesCompatibility `json:"kubernetes"`

	// The list of valid upgrade paths for the plugin
	// Required: true
	Upgrades []string `json:"upgrades"`
}

// Validate validates this plugin compatibility
func (m *PluginCompatibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrades(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginCompatibility) validateKubernetes(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes", "body", m.Kubernetes); err != nil {
		return err
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *PluginCompatibility) validateUpgrades(formats strfmt.Registry) error {

	if err := validate.Required("upgrades", "body", m.Upgrades); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginCompatibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginCompatibility) UnmarshalBinary(b []byte) error {
	var res PluginCompatibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}