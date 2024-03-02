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

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// CreateUserCreatedCode is the HTTP code returned for type CreateUserCreated
const CreateUserCreatedCode int = 201

/*
CreateUserCreated User created

swagger:response createUserCreated
*/
type CreateUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {

	return &CreateUserCreated{}
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *models.User) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserAcceptedCode is the HTTP code returned for type CreateUserAccepted
const CreateUserAcceptedCode int = 202

/*
CreateUserAccepted Configuration change accepted and reload requested

swagger:response createUserAccepted
*/
type CreateUserAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewCreateUserAccepted creates CreateUserAccepted with default headers values
func NewCreateUserAccepted() *CreateUserAccepted {

	return &CreateUserAccepted{}
}

// WithReloadID adds the reloadId to the create user accepted response
func (o *CreateUserAccepted) WithReloadID(reloadID string) *CreateUserAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create user accepted response
func (o *CreateUserAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create user accepted response
func (o *CreateUserAccepted) WithPayload(payload *models.User) *CreateUserAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user accepted response
func (o *CreateUserAccepted) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateUserBadRequestCode is the HTTP code returned for type CreateUserBadRequest
const CreateUserBadRequestCode int = 400

/*
CreateUserBadRequest Bad request

swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {

	return &CreateUserBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create user bad request response
func (o *CreateUserBadRequest) WithConfigurationVersion(configurationVersion string) *CreateUserBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create user bad request response
func (o *CreateUserBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create user bad request response
func (o *CreateUserBadRequest) WithPayload(payload *models.Error) *CreateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user bad request response
func (o *CreateUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateUserConflictCode is the HTTP code returned for type CreateUserConflict
const CreateUserConflictCode int = 409

/*
CreateUserConflict The specified resource already exists

swagger:response createUserConflict
*/
type CreateUserConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUserConflict creates CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {

	return &CreateUserConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create user conflict response
func (o *CreateUserConflict) WithConfigurationVersion(configurationVersion string) *CreateUserConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create user conflict response
func (o *CreateUserConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create user conflict response
func (o *CreateUserConflict) WithPayload(payload *models.Error) *CreateUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user conflict response
func (o *CreateUserConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
CreateUserDefault General Error

swagger:response createUserDefault
*/
type CreateUserDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUserDefault creates CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create user default response
func (o *CreateUserDefault) WithStatusCode(code int) *CreateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create user default response
func (o *CreateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create user default response
func (o *CreateUserDefault) WithConfigurationVersion(configurationVersion string) *CreateUserDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create user default response
func (o *CreateUserDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create user default response
func (o *CreateUserDefault) WithPayload(payload *models.Error) *CreateUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user default response
func (o *CreateUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
