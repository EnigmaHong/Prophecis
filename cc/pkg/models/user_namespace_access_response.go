// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserNamespaceAccessResponse user namespace access response
// swagger:model UserNamespaceAccessResponse
type UserNamespaceAccessResponse struct {

	// the accessible
	Accessible bool `json:"accessible,omitempty"`

	// the namespace
	Namespace string `json:"namespace,omitempty"`

	// the username
	Username string `json:"username,omitempty"`
}

// Validate validates this user namespace access response
func (m *UserNamespaceAccessResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserNamespaceAccessResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserNamespaceAccessResponse) UnmarshalBinary(b []byte) error {
	var res UserNamespaceAccessResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}