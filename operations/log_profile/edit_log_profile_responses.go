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

package log_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// EditLogProfileOKCode is the HTTP code returned for type EditLogProfileOK
const EditLogProfileOKCode int = 200

/*
EditLogProfileOK log_profile configuration updated

swagger:response editLogProfileOK
*/
type EditLogProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.LogProfile `json:"body,omitempty"`
}

// NewEditLogProfileOK creates EditLogProfileOK with default headers values
func NewEditLogProfileOK() *EditLogProfileOK {

	return &EditLogProfileOK{}
}

// WithPayload adds the payload to the edit log profile o k response
func (o *EditLogProfileOK) WithPayload(payload *models.LogProfile) *EditLogProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit log profile o k response
func (o *EditLogProfileOK) SetPayload(payload *models.LogProfile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditLogProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EditLogProfileAcceptedCode is the HTTP code returned for type EditLogProfileAccepted
const EditLogProfileAcceptedCode int = 202

/*
EditLogProfileAccepted Configuration change accepted and reload requested

swagger:response editLogProfileAccepted
*/
type EditLogProfileAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.LogProfile `json:"body,omitempty"`
}

// NewEditLogProfileAccepted creates EditLogProfileAccepted with default headers values
func NewEditLogProfileAccepted() *EditLogProfileAccepted {

	return &EditLogProfileAccepted{}
}

// WithReloadID adds the reloadId to the edit log profile accepted response
func (o *EditLogProfileAccepted) WithReloadID(reloadID string) *EditLogProfileAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the edit log profile accepted response
func (o *EditLogProfileAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the edit log profile accepted response
func (o *EditLogProfileAccepted) WithPayload(payload *models.LogProfile) *EditLogProfileAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit log profile accepted response
func (o *EditLogProfileAccepted) SetPayload(payload *models.LogProfile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditLogProfileAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// EditLogProfileBadRequestCode is the HTTP code returned for type EditLogProfileBadRequest
const EditLogProfileBadRequestCode int = 400

/*
EditLogProfileBadRequest Bad request

swagger:response editLogProfileBadRequest
*/
type EditLogProfileBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditLogProfileBadRequest creates EditLogProfileBadRequest with default headers values
func NewEditLogProfileBadRequest() *EditLogProfileBadRequest {

	return &EditLogProfileBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the edit log profile bad request response
func (o *EditLogProfileBadRequest) WithConfigurationVersion(configurationVersion string) *EditLogProfileBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit log profile bad request response
func (o *EditLogProfileBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit log profile bad request response
func (o *EditLogProfileBadRequest) WithPayload(payload *models.Error) *EditLogProfileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit log profile bad request response
func (o *EditLogProfileBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditLogProfileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// EditLogProfileNotFoundCode is the HTTP code returned for type EditLogProfileNotFound
const EditLogProfileNotFoundCode int = 404

/*
EditLogProfileNotFound The specified resource was not found

swagger:response editLogProfileNotFound
*/
type EditLogProfileNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditLogProfileNotFound creates EditLogProfileNotFound with default headers values
func NewEditLogProfileNotFound() *EditLogProfileNotFound {

	return &EditLogProfileNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the edit log profile not found response
func (o *EditLogProfileNotFound) WithConfigurationVersion(configurationVersion string) *EditLogProfileNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit log profile not found response
func (o *EditLogProfileNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit log profile not found response
func (o *EditLogProfileNotFound) WithPayload(payload *models.Error) *EditLogProfileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit log profile not found response
func (o *EditLogProfileNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditLogProfileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
EditLogProfileDefault General Error

swagger:response editLogProfileDefault
*/
type EditLogProfileDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditLogProfileDefault creates EditLogProfileDefault with default headers values
func NewEditLogProfileDefault(code int) *EditLogProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &EditLogProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the edit log profile default response
func (o *EditLogProfileDefault) WithStatusCode(code int) *EditLogProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the edit log profile default response
func (o *EditLogProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the edit log profile default response
func (o *EditLogProfileDefault) WithConfigurationVersion(configurationVersion string) *EditLogProfileDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit log profile default response
func (o *EditLogProfileDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit log profile default response
func (o *EditLogProfileDefault) WithPayload(payload *models.Error) *EditLogProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit log profile default response
func (o *EditLogProfileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditLogProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
