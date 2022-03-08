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

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// ReplaceGroupOKCode is the HTTP code returned for type ReplaceGroupOK
const ReplaceGroupOKCode int = 200

/*ReplaceGroupOK Group replaced

swagger:response replaceGroupOK
*/
type ReplaceGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewReplaceGroupOK creates ReplaceGroupOK with default headers values
func NewReplaceGroupOK() *ReplaceGroupOK {

	return &ReplaceGroupOK{}
}

// WithPayload adds the payload to the replace group o k response
func (o *ReplaceGroupOK) WithPayload(payload *models.Group) *ReplaceGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace group o k response
func (o *ReplaceGroupOK) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceGroupAcceptedCode is the HTTP code returned for type ReplaceGroupAccepted
const ReplaceGroupAcceptedCode int = 202

/*ReplaceGroupAccepted Configuration change accepted and reload requested

swagger:response replaceGroupAccepted
*/
type ReplaceGroupAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewReplaceGroupAccepted creates ReplaceGroupAccepted with default headers values
func NewReplaceGroupAccepted() *ReplaceGroupAccepted {

	return &ReplaceGroupAccepted{}
}

// WithReloadID adds the reloadId to the replace group accepted response
func (o *ReplaceGroupAccepted) WithReloadID(reloadID string) *ReplaceGroupAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace group accepted response
func (o *ReplaceGroupAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace group accepted response
func (o *ReplaceGroupAccepted) WithPayload(payload *models.Group) *ReplaceGroupAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace group accepted response
func (o *ReplaceGroupAccepted) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGroupAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceGroupBadRequestCode is the HTTP code returned for type ReplaceGroupBadRequest
const ReplaceGroupBadRequestCode int = 400

/*ReplaceGroupBadRequest Bad request

swagger:response replaceGroupBadRequest
*/
type ReplaceGroupBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceGroupBadRequest creates ReplaceGroupBadRequest with default headers values
func NewReplaceGroupBadRequest() *ReplaceGroupBadRequest {

	return &ReplaceGroupBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace group bad request response
func (o *ReplaceGroupBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceGroupBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace group bad request response
func (o *ReplaceGroupBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace group bad request response
func (o *ReplaceGroupBadRequest) WithPayload(payload *models.Error) *ReplaceGroupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace group bad request response
func (o *ReplaceGroupBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGroupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceGroupNotFoundCode is the HTTP code returned for type ReplaceGroupNotFound
const ReplaceGroupNotFoundCode int = 404

/*ReplaceGroupNotFound The specified resource was not found

swagger:response replaceGroupNotFound
*/
type ReplaceGroupNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceGroupNotFound creates ReplaceGroupNotFound with default headers values
func NewReplaceGroupNotFound() *ReplaceGroupNotFound {

	return &ReplaceGroupNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace group not found response
func (o *ReplaceGroupNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceGroupNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace group not found response
func (o *ReplaceGroupNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace group not found response
func (o *ReplaceGroupNotFound) WithPayload(payload *models.Error) *ReplaceGroupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace group not found response
func (o *ReplaceGroupNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGroupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceGroupDefault General Error

swagger:response replaceGroupDefault
*/
type ReplaceGroupDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceGroupDefault creates ReplaceGroupDefault with default headers values
func NewReplaceGroupDefault(code int) *ReplaceGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace group default response
func (o *ReplaceGroupDefault) WithStatusCode(code int) *ReplaceGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace group default response
func (o *ReplaceGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace group default response
func (o *ReplaceGroupDefault) WithConfigurationVersion(configurationVersion string) *ReplaceGroupDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace group default response
func (o *ReplaceGroupDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace group default response
func (o *ReplaceGroupDefault) WithPayload(payload *models.Error) *ReplaceGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace group default response
func (o *ReplaceGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
