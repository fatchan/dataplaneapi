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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateBackendSwitchingRuleCreatedCode is the HTTP code returned for type CreateBackendSwitchingRuleCreated
const CreateBackendSwitchingRuleCreatedCode int = 201

/*
CreateBackendSwitchingRuleCreated Backend Switching Rule created

swagger:response createBackendSwitchingRuleCreated
*/
type CreateBackendSwitchingRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.BackendSwitchingRule `json:"body,omitempty"`
}

// NewCreateBackendSwitchingRuleCreated creates CreateBackendSwitchingRuleCreated with default headers values
func NewCreateBackendSwitchingRuleCreated() *CreateBackendSwitchingRuleCreated {

	return &CreateBackendSwitchingRuleCreated{}
}

// WithPayload adds the payload to the create backend switching rule created response
func (o *CreateBackendSwitchingRuleCreated) WithPayload(payload *models.BackendSwitchingRule) *CreateBackendSwitchingRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend switching rule created response
func (o *CreateBackendSwitchingRuleCreated) SetPayload(payload *models.BackendSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendSwitchingRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBackendSwitchingRuleAcceptedCode is the HTTP code returned for type CreateBackendSwitchingRuleAccepted
const CreateBackendSwitchingRuleAcceptedCode int = 202

/*
CreateBackendSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response createBackendSwitchingRuleAccepted
*/
type CreateBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.BackendSwitchingRule `json:"body,omitempty"`
}

// NewCreateBackendSwitchingRuleAccepted creates CreateBackendSwitchingRuleAccepted with default headers values
func NewCreateBackendSwitchingRuleAccepted() *CreateBackendSwitchingRuleAccepted {

	return &CreateBackendSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the create backend switching rule accepted response
func (o *CreateBackendSwitchingRuleAccepted) WithReloadID(reloadID string) *CreateBackendSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create backend switching rule accepted response
func (o *CreateBackendSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create backend switching rule accepted response
func (o *CreateBackendSwitchingRuleAccepted) WithPayload(payload *models.BackendSwitchingRule) *CreateBackendSwitchingRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend switching rule accepted response
func (o *CreateBackendSwitchingRuleAccepted) SetPayload(payload *models.BackendSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateBackendSwitchingRuleBadRequestCode is the HTTP code returned for type CreateBackendSwitchingRuleBadRequest
const CreateBackendSwitchingRuleBadRequestCode int = 400

/*
CreateBackendSwitchingRuleBadRequest Bad request

swagger:response createBackendSwitchingRuleBadRequest
*/
type CreateBackendSwitchingRuleBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendSwitchingRuleBadRequest creates CreateBackendSwitchingRuleBadRequest with default headers values
func NewCreateBackendSwitchingRuleBadRequest() *CreateBackendSwitchingRuleBadRequest {

	return &CreateBackendSwitchingRuleBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create backend switching rule bad request response
func (o *CreateBackendSwitchingRuleBadRequest) WithConfigurationVersion(configurationVersion string) *CreateBackendSwitchingRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend switching rule bad request response
func (o *CreateBackendSwitchingRuleBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend switching rule bad request response
func (o *CreateBackendSwitchingRuleBadRequest) WithPayload(payload *models.Error) *CreateBackendSwitchingRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend switching rule bad request response
func (o *CreateBackendSwitchingRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendSwitchingRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateBackendSwitchingRuleConflictCode is the HTTP code returned for type CreateBackendSwitchingRuleConflict
const CreateBackendSwitchingRuleConflictCode int = 409

/*
CreateBackendSwitchingRuleConflict The specified resource already exists

swagger:response createBackendSwitchingRuleConflict
*/
type CreateBackendSwitchingRuleConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendSwitchingRuleConflict creates CreateBackendSwitchingRuleConflict with default headers values
func NewCreateBackendSwitchingRuleConflict() *CreateBackendSwitchingRuleConflict {

	return &CreateBackendSwitchingRuleConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create backend switching rule conflict response
func (o *CreateBackendSwitchingRuleConflict) WithConfigurationVersion(configurationVersion string) *CreateBackendSwitchingRuleConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend switching rule conflict response
func (o *CreateBackendSwitchingRuleConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend switching rule conflict response
func (o *CreateBackendSwitchingRuleConflict) WithPayload(payload *models.Error) *CreateBackendSwitchingRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend switching rule conflict response
func (o *CreateBackendSwitchingRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendSwitchingRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
CreateBackendSwitchingRuleDefault General Error

swagger:response createBackendSwitchingRuleDefault
*/
type CreateBackendSwitchingRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendSwitchingRuleDefault creates CreateBackendSwitchingRuleDefault with default headers values
func NewCreateBackendSwitchingRuleDefault(code int) *CreateBackendSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) WithStatusCode(code int) *CreateBackendSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) WithConfigurationVersion(configurationVersion string) *CreateBackendSwitchingRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) WithPayload(payload *models.Error) *CreateBackendSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
