// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// DeleteUserByNameOKCode is the HTTP code returned for type DeleteUserByNameOK
const DeleteUserByNameOKCode int = 200

/*DeleteUserByNameOK Detailed User and User information.

swagger:response deleteUserByNameOK
*/
type DeleteUserByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewDeleteUserByNameOK creates DeleteUserByNameOK with default headers values
func NewDeleteUserByNameOK() *DeleteUserByNameOK {

	return &DeleteUserByNameOK{}
}

// WithPayload adds the payload to the delete user by name o k response
func (o *DeleteUserByNameOK) WithPayload(payload *models.User) *DeleteUserByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user by name o k response
func (o *DeleteUserByNameOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserByNameUnauthorizedCode is the HTTP code returned for type DeleteUserByNameUnauthorized
const DeleteUserByNameUnauthorizedCode int = 401

/*DeleteUserByNameUnauthorized Unauthorized

swagger:response deleteUserByNameUnauthorized
*/
type DeleteUserByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserByNameUnauthorized creates DeleteUserByNameUnauthorized with default headers values
func NewDeleteUserByNameUnauthorized() *DeleteUserByNameUnauthorized {

	return &DeleteUserByNameUnauthorized{}
}

// WithPayload adds the payload to the delete user by name unauthorized response
func (o *DeleteUserByNameUnauthorized) WithPayload(payload *models.Error) *DeleteUserByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user by name unauthorized response
func (o *DeleteUserByNameUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserByNameNotFoundCode is the HTTP code returned for type DeleteUserByNameNotFound
const DeleteUserByNameNotFoundCode int = 404

/*DeleteUserByNameNotFound Model with the given ID not found.

swagger:response deleteUserByNameNotFound
*/
type DeleteUserByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserByNameNotFound creates DeleteUserByNameNotFound with default headers values
func NewDeleteUserByNameNotFound() *DeleteUserByNameNotFound {

	return &DeleteUserByNameNotFound{}
}

// WithPayload adds the payload to the delete user by name not found response
func (o *DeleteUserByNameNotFound) WithPayload(payload *models.Error) *DeleteUserByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user by name not found response
func (o *DeleteUserByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}