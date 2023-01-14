// Code generated by go-swagger; DO NOT EDIT.

package ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/case-management-suite/models"
)

// GetActionRecordsOKCode is the HTTP code returned for type GetActionRecordsOK
const GetActionRecordsOKCode int = 200

/*
GetActionRecordsOK CaseAction

swagger:response getActionRecordsOK
*/
type GetActionRecordsOK struct {

	/*
	  In: Body
	*/
	Payload []models.CaseAction `json:"body,omitempty"`
}

// NewGetActionRecordsOK creates GetActionRecordsOK with default headers values
func NewGetActionRecordsOK() *GetActionRecordsOK {

	return &GetActionRecordsOK{}
}

// WithPayload adds the payload to the get action records o k response
func (o *GetActionRecordsOK) WithPayload(payload []models.CaseAction) *GetActionRecordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action records o k response
func (o *GetActionRecordsOK) SetPayload(payload []models.CaseAction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionRecordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]models.CaseAction, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetActionRecordsDefault get action records default

swagger:response getActionRecordsDefault
*/
type GetActionRecordsDefault struct {
	_statusCode int
	/*

	 */
	Code int16 `json:"Code"`
	/*

	 */
	Message string `json:"Message"`
}

// NewGetActionRecordsDefault creates GetActionRecordsDefault with default headers values
func NewGetActionRecordsDefault(code int) *GetActionRecordsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetActionRecordsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get action records default response
func (o *GetActionRecordsDefault) WithStatusCode(code int) *GetActionRecordsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get action records default response
func (o *GetActionRecordsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithCode adds the code to the get action records default response
func (o *GetActionRecordsDefault) WithCode(code int16) *GetActionRecordsDefault {
	o.Code = code
	return o
}

// SetCode sets the code to the get action records default response
func (o *GetActionRecordsDefault) SetCode(code int16) {
	o.Code = code
}

// WithMessage adds the message to the get action records default response
func (o *GetActionRecordsDefault) WithMessage(message string) *GetActionRecordsDefault {
	o.Message = message
	return o
}

// SetMessage sets the message to the get action records default response
func (o *GetActionRecordsDefault) SetMessage(message string) {
	o.Message = message
}

// WriteResponse to the client
func (o *GetActionRecordsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Code

	code := swag.FormatInt16(o.Code)
	if code != "" {
		rw.Header().Set("Code", code)
	}

	// response header Message

	message := o.Message
	if message != "" {
		rw.Header().Set("Message", message)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
