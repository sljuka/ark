// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MarketHour v1 market hour
//
// swagger:model v1MarketHour
type V1MarketHour struct {

	// next end time
	// Format: date-time
	NextEndTime strfmt.DateTime `json:"nextEndTime,omitempty"`

	// next start time
	// Format: date-time
	NextStartTime strfmt.DateTime `json:"nextStartTime,omitempty"`

	// period
	Period string `json:"period,omitempty"`

	// round interval
	RoundInterval string `json:"roundInterval,omitempty"`
}

// Validate validates this v1 market hour
func (m *V1MarketHour) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MarketHour) validateNextEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.NextEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextEndTime", "body", "date-time", m.NextEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MarketHour) validateNextStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.NextStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextStartTime", "body", "date-time", m.NextStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 market hour based on context it is used
func (m *V1MarketHour) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MarketHour) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MarketHour) UnmarshalBinary(b []byte) error {
	var res V1MarketHour
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
