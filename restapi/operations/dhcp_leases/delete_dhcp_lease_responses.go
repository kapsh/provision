package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type DeleteDhcpLeaseNoContent
const DeleteDhcpLeaseNoContentCode int = 204

/*DeleteDhcpLeaseNoContent delete dhcp lease no content

swagger:response deleteDhcpLeaseNoContent
*/
type DeleteDhcpLeaseNoContent struct {
}

// NewDeleteDhcpLeaseNoContent creates DeleteDhcpLeaseNoContent with default headers values
func NewDeleteDhcpLeaseNoContent() *DeleteDhcpLeaseNoContent {
	return &DeleteDhcpLeaseNoContent{}
}

// WriteResponse to the client
func (o *DeleteDhcpLeaseNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// HTTP code for type DeleteDhcpLeaseUnauthorized
const DeleteDhcpLeaseUnauthorizedCode int = 401

/*DeleteDhcpLeaseUnauthorized delete dhcp lease unauthorized

swagger:response deleteDhcpLeaseUnauthorized
*/
type DeleteDhcpLeaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDhcpLeaseUnauthorized creates DeleteDhcpLeaseUnauthorized with default headers values
func NewDeleteDhcpLeaseUnauthorized() *DeleteDhcpLeaseUnauthorized {
	return &DeleteDhcpLeaseUnauthorized{}
}

// WithPayload adds the payload to the delete dhcp lease unauthorized response
func (o *DeleteDhcpLeaseUnauthorized) WithPayload(payload *models.Error) *DeleteDhcpLeaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dhcp lease unauthorized response
func (o *DeleteDhcpLeaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDhcpLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteDhcpLeaseNotFound
const DeleteDhcpLeaseNotFoundCode int = 404

/*DeleteDhcpLeaseNotFound delete dhcp lease not found

swagger:response deleteDhcpLeaseNotFound
*/
type DeleteDhcpLeaseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDhcpLeaseNotFound creates DeleteDhcpLeaseNotFound with default headers values
func NewDeleteDhcpLeaseNotFound() *DeleteDhcpLeaseNotFound {
	return &DeleteDhcpLeaseNotFound{}
}

// WithPayload adds the payload to the delete dhcp lease not found response
func (o *DeleteDhcpLeaseNotFound) WithPayload(payload *models.Error) *DeleteDhcpLeaseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dhcp lease not found response
func (o *DeleteDhcpLeaseNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDhcpLeaseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteDhcpLeaseConflict
const DeleteDhcpLeaseConflictCode int = 409

/*DeleteDhcpLeaseConflict delete dhcp lease conflict

swagger:response deleteDhcpLeaseConflict
*/
type DeleteDhcpLeaseConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDhcpLeaseConflict creates DeleteDhcpLeaseConflict with default headers values
func NewDeleteDhcpLeaseConflict() *DeleteDhcpLeaseConflict {
	return &DeleteDhcpLeaseConflict{}
}

// WithPayload adds the payload to the delete dhcp lease conflict response
func (o *DeleteDhcpLeaseConflict) WithPayload(payload *models.Error) *DeleteDhcpLeaseConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dhcp lease conflict response
func (o *DeleteDhcpLeaseConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDhcpLeaseConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteDhcpLeaseInternalServerError
const DeleteDhcpLeaseInternalServerErrorCode int = 500

/*DeleteDhcpLeaseInternalServerError delete dhcp lease internal server error

swagger:response deleteDhcpLeaseInternalServerError
*/
type DeleteDhcpLeaseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDhcpLeaseInternalServerError creates DeleteDhcpLeaseInternalServerError with default headers values
func NewDeleteDhcpLeaseInternalServerError() *DeleteDhcpLeaseInternalServerError {
	return &DeleteDhcpLeaseInternalServerError{}
}

// WithPayload adds the payload to the delete dhcp lease internal server error response
func (o *DeleteDhcpLeaseInternalServerError) WithPayload(payload *models.Error) *DeleteDhcpLeaseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dhcp lease internal server error response
func (o *DeleteDhcpLeaseInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDhcpLeaseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
