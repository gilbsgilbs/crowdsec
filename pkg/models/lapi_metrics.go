// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LapiMetrics LapiMetrics
//
// swagger:model LapiMetrics
type LapiMetrics struct {
	BaseMetrics

	// console options
	ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LapiMetrics) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMetrics
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMetrics = aO0

	// AO1
	var dataAO1 struct {
		ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConsoleOptions = dataAO1.ConsoleOptions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LapiMetrics) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMetrics)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`
	}

	dataAO1.ConsoleOptions = m.ConsoleOptions

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this lapi metrics
func (m *LapiMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMetrics
	if err := m.BaseMetrics.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsoleOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LapiMetrics) validateConsoleOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsoleOptions) { // not required
		return nil
	}

	if err := m.ConsoleOptions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("console_options")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("console_options")
		}
		return err
	}

	return nil
}

// ContextValidate validate this lapi metrics based on the context it is used
func (m *LapiMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMetrics
	if err := m.BaseMetrics.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsoleOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LapiMetrics) contextValidateConsoleOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConsoleOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("console_options")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("console_options")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LapiMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LapiMetrics) UnmarshalBinary(b []byte) error {
	var res LapiMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}