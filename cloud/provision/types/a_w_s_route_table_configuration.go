// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AWSRouteTableConfiguration AWS route table configuration
// swagger:model AWSRouteTableConfiguration
type AWSRouteTableConfiguration struct {

	// List of routes in this route table
	Route []*AWSRoute `json:"route"`

	// AWS tags
	Tags interface{} `json:"tags,omitempty"`

	// VPC this route table belongs to
	VpcID string `json:"vpc_id,omitempty"`
}

// Validate validates this a w s route table configuration
func (m *AWSRouteTableConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSRouteTableConfiguration) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	for i := 0; i < len(m.Route); i++ {
		if swag.IsZero(m.Route[i]) { // not required
			continue
		}

		if m.Route[i] != nil {
			if err := m.Route[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("route" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSRouteTableConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSRouteTableConfiguration) UnmarshalBinary(b []byte) error {
	var res AWSRouteTableConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
