// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenericError GenericError is a generic error message return by a server
//
// swagger:model GenericError
type GenericError struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this generic error
func (m *GenericError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generic error based on context it is used
func (m *GenericError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericError) UnmarshalBinary(b []byte) error {
	var res GenericError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
