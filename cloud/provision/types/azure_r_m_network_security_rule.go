// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureRMNetworkSecurityRule azure r m network security rule
// swagger:model AzureRMNetworkSecurityRule
type AzureRMNetworkSecurityRule struct {

	// access
	Access string `json:"access,omitempty"`

	// destination address prefix
	DestinationAddressPrefix string `json:"destination_address_prefix,omitempty"`

	// destination port range
	DestinationPortRange string `json:"destination_port_range,omitempty"`

	// direction
	Direction string `json:"direction,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// source address prefix
	SourceAddressPrefix string `json:"source_address_prefix,omitempty"`

	// source port range
	SourcePortRange string `json:"source_port_range,omitempty"`
}

// Validate validates this azure r m network security rule
func (m *AzureRMNetworkSecurityRule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureRMNetworkSecurityRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureRMNetworkSecurityRule) UnmarshalBinary(b []byte) error {
	var res AzureRMNetworkSecurityRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
