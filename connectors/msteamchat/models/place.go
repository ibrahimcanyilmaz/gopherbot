package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Place Place (entity type: "https://schema.org/Place")
// swagger:model Place
type Place struct {

	// Address of the place (may be `string` or complex object of type `PostalAddress`)
	Address Object `json:"address,omitempty"`

	// Geo coordinates of the place (may be complex object of type `GeoCoordinates` or `GeoShape`)
	Geo Object `json:"geo,omitempty"`

	// Map to the place (may be `string` (URL) or complex object of type `Map`)
	HasMap Object `json:"hasMap,omitempty"`

	// The name of the thing
	Name string `json:"name,omitempty"`

	// The type of the thing
	Type string `json:"type,omitempty"`
}

// Validate validates this place
func (m *Place) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
