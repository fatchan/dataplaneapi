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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// ReplaceDefaultsSectionOKCode is the HTTP code returned for type ReplaceDefaultsSectionOK
const ReplaceDefaultsSectionOKCode int = 200

/*
ReplaceDefaultsSectionOK Defaults section replaced

swagger:response replaceDefaultsSectionOK
*/
type ReplaceDefaultsSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewReplaceDefaultsSectionOK creates ReplaceDefaultsSectionOK with default headers values
func NewReplaceDefaultsSectionOK() *ReplaceDefaultsSectionOK {

	return &ReplaceDefaultsSectionOK{}
}

// WithPayload adds the payload to the replace defaults section o k response
func (o *ReplaceDefaultsSectionOK) WithPayload(payload *models.Defaults) *ReplaceDefaultsSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults section o k response
func (o *ReplaceDefaultsSectionOK) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceDefaultsSectionAcceptedCode is the HTTP code returned for type ReplaceDefaultsSectionAccepted
const ReplaceDefaultsSectionAcceptedCode int = 202

/*
ReplaceDefaultsSectionAccepted Configuration change accepted and reload requested

swagger:response replaceDefaultsSectionAccepted
*/
type ReplaceDefaultsSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewReplaceDefaultsSectionAccepted creates ReplaceDefaultsSectionAccepted with default headers values
func NewReplaceDefaultsSectionAccepted() *ReplaceDefaultsSectionAccepted {

	return &ReplaceDefaultsSectionAccepted{}
}

// WithReloadID adds the reloadId to the replace defaults section accepted response
func (o *ReplaceDefaultsSectionAccepted) WithReloadID(reloadID string) *ReplaceDefaultsSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace defaults section accepted response
func (o *ReplaceDefaultsSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace defaults section accepted response
func (o *ReplaceDefaultsSectionAccepted) WithPayload(payload *models.Defaults) *ReplaceDefaultsSectionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults section accepted response
func (o *ReplaceDefaultsSectionAccepted) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceDefaultsSectionBadRequestCode is the HTTP code returned for type ReplaceDefaultsSectionBadRequest
const ReplaceDefaultsSectionBadRequestCode int = 400

/*
ReplaceDefaultsSectionBadRequest Bad request

swagger:response replaceDefaultsSectionBadRequest
*/
type ReplaceDefaultsSectionBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDefaultsSectionBadRequest creates ReplaceDefaultsSectionBadRequest with default headers values
func NewReplaceDefaultsSectionBadRequest() *ReplaceDefaultsSectionBadRequest {

	return &ReplaceDefaultsSectionBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace defaults section bad request response
func (o *ReplaceDefaultsSectionBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceDefaultsSectionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace defaults section bad request response
func (o *ReplaceDefaultsSectionBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace defaults section bad request response
func (o *ReplaceDefaultsSectionBadRequest) WithPayload(payload *models.Error) *ReplaceDefaultsSectionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults section bad request response
func (o *ReplaceDefaultsSectionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsSectionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceDefaultsSectionNotFoundCode is the HTTP code returned for type ReplaceDefaultsSectionNotFound
const ReplaceDefaultsSectionNotFoundCode int = 404

/*
ReplaceDefaultsSectionNotFound The specified resource was not found

swagger:response replaceDefaultsSectionNotFound
*/
type ReplaceDefaultsSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDefaultsSectionNotFound creates ReplaceDefaultsSectionNotFound with default headers values
func NewReplaceDefaultsSectionNotFound() *ReplaceDefaultsSectionNotFound {

	return &ReplaceDefaultsSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace defaults section not found response
func (o *ReplaceDefaultsSectionNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceDefaultsSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace defaults section not found response
func (o *ReplaceDefaultsSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace defaults section not found response
func (o *ReplaceDefaultsSectionNotFound) WithPayload(payload *models.Error) *ReplaceDefaultsSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults section not found response
func (o *ReplaceDefaultsSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceDefaultsSectionDefault General Error

swagger:response replaceDefaultsSectionDefault
*/
type ReplaceDefaultsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDefaultsSectionDefault creates ReplaceDefaultsSectionDefault with default headers values
func NewReplaceDefaultsSectionDefault(code int) *ReplaceDefaultsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceDefaultsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) WithStatusCode(code int) *ReplaceDefaultsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) WithConfigurationVersion(configurationVersion string) *ReplaceDefaultsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) WithPayload(payload *models.Error) *ReplaceDefaultsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults section default response
func (o *ReplaceDefaultsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
