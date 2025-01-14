// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SubmitRedeemTxRequest v1 submit redeem tx request
//
// swagger:model v1SubmitRedeemTxRequest
type V1SubmitRedeemTxRequest struct {

	// redeem tx
	RedeemTx string `json:"redeemTx,omitempty"`
}

// Validate validates this v1 submit redeem tx request
func (m *V1SubmitRedeemTxRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 submit redeem tx request based on context it is used
func (m *V1SubmitRedeemTxRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SubmitRedeemTxRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SubmitRedeemTxRequest) UnmarshalBinary(b []byte) error {
	var res V1SubmitRedeemTxRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
