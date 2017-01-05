package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// MediaURL MediaUrl data
// swagger:model MediaUrl
type MediaURL struct {

	// Optional profile hint to the client to differentiate multiple MediaUrl objects from each other
	Profile string `json:"profile,omitempty"`

	// Url for the media
	URL string `json:"url,omitempty"`
}

// Validate validates this media Url
func (m *MediaURL) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
