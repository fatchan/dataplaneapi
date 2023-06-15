// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package log_forward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateLogForwardCreatedCode is the HTTP code returned for type CreateLogForwardCreated
const CreateLogForwardCreatedCode int = 201

/*
CreateLogForwardCreated Log Forward created

swagger:response createLogForwardCreated
*/
type CreateLogForwardCreated struct {

	/*
	  In: Body
	*/
	Payload *models.LogForward `json:"body,omitempty"`
}

// NewCreateLogForwardCreated creates CreateLogForwardCreated with default headers values
func NewCreateLogForwardCreated() *CreateLogForwardCreated {

	return &CreateLogForwardCreated{}
}

// WithPayload adds the payload to the create log forward created response
func (o *CreateLogForwardCreated) WithPayload(payload *models.LogForward) *CreateLogForwardCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log forward created response
func (o *CreateLogForwardCreated) SetPayload(payload *models.LogForward) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogForwardCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogForwardAcceptedCode is the HTTP code returned for type CreateLogForwardAccepted
const CreateLogForwardAcceptedCode int = 202

/*
CreateLogForwardAccepted Configuration change accepted and reload requested

swagger:response createLogForwardAccepted
*/
type CreateLogForwardAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.LogForward `json:"body,omitempty"`
}

// NewCreateLogForwardAccepted creates CreateLogForwardAccepted with default headers values
func NewCreateLogForwardAccepted() *CreateLogForwardAccepted {

	return &CreateLogForwardAccepted{}
}

// WithReloadID adds the reloadId to the create log forward accepted response
func (o *CreateLogForwardAccepted) WithReloadID(reloadID string) *CreateLogForwardAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create log forward accepted response
func (o *CreateLogForwardAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create log forward accepted response
func (o *CreateLogForwardAccepted) WithPayload(payload *models.LogForward) *CreateLogForwardAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log forward accepted response
func (o *CreateLogForwardAccepted) SetPayload(payload *models.LogForward) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogForwardAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogForwardBadRequestCode is the HTTP code returned for type CreateLogForwardBadRequest
const CreateLogForwardBadRequestCode int = 400

/*
CreateLogForwardBadRequest Bad request

swagger:response createLogForwardBadRequest
*/
type CreateLogForwardBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogForwardBadRequest creates CreateLogForwardBadRequest with default headers values
func NewCreateLogForwardBadRequest() *CreateLogForwardBadRequest {

	return &CreateLogForwardBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create log forward bad request response
func (o *CreateLogForwardBadRequest) WithConfigurationVersion(configurationVersion string) *CreateLogForwardBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log forward bad request response
func (o *CreateLogForwardBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log forward bad request response
func (o *CreateLogForwardBadRequest) WithPayload(payload *models.Error) *CreateLogForwardBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log forward bad request response
func (o *CreateLogForwardBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogForwardBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogForwardConflictCode is the HTTP code returned for type CreateLogForwardConflict
const CreateLogForwardConflictCode int = 409

/*
CreateLogForwardConflict The specified resource already exists

swagger:response createLogForwardConflict
*/
type CreateLogForwardConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogForwardConflict creates CreateLogForwardConflict with default headers values
func NewCreateLogForwardConflict() *CreateLogForwardConflict {

	return &CreateLogForwardConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create log forward conflict response
func (o *CreateLogForwardConflict) WithConfigurationVersion(configurationVersion string) *CreateLogForwardConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log forward conflict response
func (o *CreateLogForwardConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log forward conflict response
func (o *CreateLogForwardConflict) WithPayload(payload *models.Error) *CreateLogForwardConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log forward conflict response
func (o *CreateLogForwardConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogForwardConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateLogForwardDefault General Error

swagger:response createLogForwardDefault
*/
type CreateLogForwardDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogForwardDefault creates CreateLogForwardDefault with default headers values
func NewCreateLogForwardDefault(code int) *CreateLogForwardDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateLogForwardDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create log forward default response
func (o *CreateLogForwardDefault) WithStatusCode(code int) *CreateLogForwardDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create log forward default response
func (o *CreateLogForwardDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create log forward default response
func (o *CreateLogForwardDefault) WithConfigurationVersion(configurationVersion string) *CreateLogForwardDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log forward default response
func (o *CreateLogForwardDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log forward default response
func (o *CreateLogForwardDefault) WithPayload(payload *models.Error) *CreateLogForwardDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log forward default response
func (o *CreateLogForwardDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogForwardDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
