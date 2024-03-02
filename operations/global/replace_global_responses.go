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

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceGlobalOKCode is the HTTP code returned for type ReplaceGlobalOK
const ReplaceGlobalOKCode int = 200

/*
ReplaceGlobalOK Global replaced

swagger:response replaceGlobalOK
*/
type ReplaceGlobalOK struct {

	/*
	  In: Body
	*/
	Payload *models.Global `json:"body,omitempty"`
}

// NewReplaceGlobalOK creates ReplaceGlobalOK with default headers values
func NewReplaceGlobalOK() *ReplaceGlobalOK {

	return &ReplaceGlobalOK{}
}

// WithPayload adds the payload to the replace global o k response
func (o *ReplaceGlobalOK) WithPayload(payload *models.Global) *ReplaceGlobalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace global o k response
func (o *ReplaceGlobalOK) SetPayload(payload *models.Global) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGlobalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceGlobalAcceptedCode is the HTTP code returned for type ReplaceGlobalAccepted
const ReplaceGlobalAcceptedCode int = 202

/*
ReplaceGlobalAccepted Configuration change accepted and reload requested

swagger:response replaceGlobalAccepted
*/
type ReplaceGlobalAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Global `json:"body,omitempty"`
}

// NewReplaceGlobalAccepted creates ReplaceGlobalAccepted with default headers values
func NewReplaceGlobalAccepted() *ReplaceGlobalAccepted {

	return &ReplaceGlobalAccepted{}
}

// WithReloadID adds the reloadId to the replace global accepted response
func (o *ReplaceGlobalAccepted) WithReloadID(reloadID string) *ReplaceGlobalAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace global accepted response
func (o *ReplaceGlobalAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace global accepted response
func (o *ReplaceGlobalAccepted) WithPayload(payload *models.Global) *ReplaceGlobalAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace global accepted response
func (o *ReplaceGlobalAccepted) SetPayload(payload *models.Global) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGlobalAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceGlobalBadRequestCode is the HTTP code returned for type ReplaceGlobalBadRequest
const ReplaceGlobalBadRequestCode int = 400

/*
ReplaceGlobalBadRequest Bad request

swagger:response replaceGlobalBadRequest
*/
type ReplaceGlobalBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceGlobalBadRequest creates ReplaceGlobalBadRequest with default headers values
func NewReplaceGlobalBadRequest() *ReplaceGlobalBadRequest {

	return &ReplaceGlobalBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace global bad request response
func (o *ReplaceGlobalBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceGlobalBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace global bad request response
func (o *ReplaceGlobalBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace global bad request response
func (o *ReplaceGlobalBadRequest) WithPayload(payload *models.Error) *ReplaceGlobalBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace global bad request response
func (o *ReplaceGlobalBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGlobalBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*
ReplaceGlobalDefault General Error

swagger:response replaceGlobalDefault
*/
type ReplaceGlobalDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceGlobalDefault creates ReplaceGlobalDefault with default headers values
func NewReplaceGlobalDefault(code int) *ReplaceGlobalDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceGlobalDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace global default response
func (o *ReplaceGlobalDefault) WithStatusCode(code int) *ReplaceGlobalDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace global default response
func (o *ReplaceGlobalDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace global default response
func (o *ReplaceGlobalDefault) WithConfigurationVersion(configurationVersion string) *ReplaceGlobalDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace global default response
func (o *ReplaceGlobalDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace global default response
func (o *ReplaceGlobalDefault) WithPayload(payload *models.Error) *ReplaceGlobalDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace global default response
func (o *ReplaceGlobalDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGlobalDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
