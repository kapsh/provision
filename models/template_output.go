package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TemplateOutput Template Output
// swagger:model template-output
type TemplateOutput struct {
	TemplateInput

	// errors
	Errors []string `json:"Errors"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TemplateOutput) UnmarshalJSON(raw []byte) error {

	var aO0 TemplateInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TemplateInput = aO0

	var data struct {
		Errors []string `json:"Errors,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Errors = data.Errors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TemplateOutput) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.TemplateInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Errors []string `json:"Errors,omitempty"`
	}

	data.Errors = m.Errors

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this template output
func (m *TemplateOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.TemplateInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateOutput) validateErrors(formats strfmt.Registry) error {

	return nil
}
