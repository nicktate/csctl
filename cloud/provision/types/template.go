// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Template template
// swagger:model Template
type Template struct {

	// Template configuration
	// Required: true
	Configuration *TemplateConfiguration `json:"configuration"`

	// Timestamp at which the cluster was created
	CreatedAt string `json:"created_at,omitempty"`

	// Description of this template
	// Required: true
	Description *string `json:"description"`

	// Engine to use for provisioning (deprecated - always containership_kubernetes_engine)
	// Required: true
	// Enum: [containership_kubernetes_engine]
	Engine *string `json:"engine"`

	// Cluster ID
	ID UUID `json:"id,omitempty"`

	// Organization ID of the organization the cluster belongs to
	OrganizationID UUID `json:"organization_id,omitempty"`

	// Account ID of the cluster owner
	OwnerID UUID `json:"owner_id,omitempty"`

	// Cloud provider name
	// Required: true
	// Enum: [digital_ocean azure google amazon_web_services packet]
	ProviderName *string `json:"provider_name"`

	// Timestamp at which the cluster was updated
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this template
func (m *Template) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Template) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *Template) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var templateTypeEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["containership_kubernetes_engine"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateTypeEnginePropEnum = append(templateTypeEnginePropEnum, v)
	}
}

const (

	// TemplateEngineContainershipKubernetesEngine captures enum value "containership_kubernetes_engine"
	TemplateEngineContainershipKubernetesEngine string = "containership_kubernetes_engine"
)

// prop value enum
func (m *Template) validateEngineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateTypeEnginePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Template) validateEngine(formats strfmt.Registry) error {

	if err := validate.Required("engine", "body", m.Engine); err != nil {
		return err
	}

	// value enum
	if err := m.validateEngineEnum("engine", "body", *m.Engine); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Template) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := m.OrganizationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("organization_id")
		}
		return err
	}

	return nil
}

func (m *Template) validateOwnerID(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerID) { // not required
		return nil
	}

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner_id")
		}
		return err
	}

	return nil
}

var templateTypeProviderNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["digital_ocean","azure","google","amazon_web_services","packet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateTypeProviderNamePropEnum = append(templateTypeProviderNamePropEnum, v)
	}
}

const (

	// TemplateProviderNameDigitalOcean captures enum value "digital_ocean"
	TemplateProviderNameDigitalOcean string = "digital_ocean"

	// TemplateProviderNameAzure captures enum value "azure"
	TemplateProviderNameAzure string = "azure"

	// TemplateProviderNameGoogle captures enum value "google"
	TemplateProviderNameGoogle string = "google"

	// TemplateProviderNameAmazonWebServices captures enum value "amazon_web_services"
	TemplateProviderNameAmazonWebServices string = "amazon_web_services"

	// TemplateProviderNamePacket captures enum value "packet"
	TemplateProviderNamePacket string = "packet"
)

// prop value enum
func (m *Template) validateProviderNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateTypeProviderNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Template) validateProviderName(formats strfmt.Registry) error {

	if err := validate.Required("provider_name", "body", m.ProviderName); err != nil {
		return err
	}

	// value enum
	if err := m.validateProviderNameEnum("provider_name", "body", *m.ProviderName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Template) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Template) UnmarshalBinary(b []byte) error {
	var res Template
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
