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

package resolver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceResolverOKCode is the HTTP code returned for type ReplaceResolverOK
const ReplaceResolverOKCode int = 200

/*
ReplaceResolverOK Resolver replaced

swagger:response replaceResolverOK
*/
type ReplaceResolverOK struct {

	/*
	  In: Body
	*/
	Payload *models.Resolver `json:"body,omitempty"`
}

// NewReplaceResolverOK creates ReplaceResolverOK with default headers values
func NewReplaceResolverOK() *ReplaceResolverOK {

	return &ReplaceResolverOK{}
}

// WithPayload adds the payload to the replace resolver o k response
func (o *ReplaceResolverOK) WithPayload(payload *models.Resolver) *ReplaceResolverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace resolver o k response
func (o *ReplaceResolverOK) SetPayload(payload *models.Resolver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceResolverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceResolverAcceptedCode is the HTTP code returned for type ReplaceResolverAccepted
const ReplaceResolverAcceptedCode int = 202

/*
ReplaceResolverAccepted Configuration change accepted and reload requested

swagger:response replaceResolverAccepted
*/
type ReplaceResolverAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Resolver `json:"body,omitempty"`
}

// NewReplaceResolverAccepted creates ReplaceResolverAccepted with default headers values
func NewReplaceResolverAccepted() *ReplaceResolverAccepted {

	return &ReplaceResolverAccepted{}
}

// WithReloadID adds the reloadId to the replace resolver accepted response
func (o *ReplaceResolverAccepted) WithReloadID(reloadID string) *ReplaceResolverAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace resolver accepted response
func (o *ReplaceResolverAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace resolver accepted response
func (o *ReplaceResolverAccepted) WithPayload(payload *models.Resolver) *ReplaceResolverAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace resolver accepted response
func (o *ReplaceResolverAccepted) SetPayload(payload *models.Resolver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceResolverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceResolverBadRequestCode is the HTTP code returned for type ReplaceResolverBadRequest
const ReplaceResolverBadRequestCode int = 400

/*
ReplaceResolverBadRequest Bad request

swagger:response replaceResolverBadRequest
*/
type ReplaceResolverBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceResolverBadRequest creates ReplaceResolverBadRequest with default headers values
func NewReplaceResolverBadRequest() *ReplaceResolverBadRequest {

	return &ReplaceResolverBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace resolver bad request response
func (o *ReplaceResolverBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceResolverBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace resolver bad request response
func (o *ReplaceResolverBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace resolver bad request response
func (o *ReplaceResolverBadRequest) WithPayload(payload *models.Error) *ReplaceResolverBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace resolver bad request response
func (o *ReplaceResolverBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceResolverBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceResolverNotFoundCode is the HTTP code returned for type ReplaceResolverNotFound
const ReplaceResolverNotFoundCode int = 404

/*
ReplaceResolverNotFound The specified resource was not found

swagger:response replaceResolverNotFound
*/
type ReplaceResolverNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceResolverNotFound creates ReplaceResolverNotFound with default headers values
func NewReplaceResolverNotFound() *ReplaceResolverNotFound {

	return &ReplaceResolverNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace resolver not found response
func (o *ReplaceResolverNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceResolverNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace resolver not found response
func (o *ReplaceResolverNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace resolver not found response
func (o *ReplaceResolverNotFound) WithPayload(payload *models.Error) *ReplaceResolverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace resolver not found response
func (o *ReplaceResolverNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceResolverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceResolverDefault General Error

swagger:response replaceResolverDefault
*/
type ReplaceResolverDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceResolverDefault creates ReplaceResolverDefault with default headers values
func NewReplaceResolverDefault(code int) *ReplaceResolverDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceResolverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace resolver default response
func (o *ReplaceResolverDefault) WithStatusCode(code int) *ReplaceResolverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace resolver default response
func (o *ReplaceResolverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace resolver default response
func (o *ReplaceResolverDefault) WithConfigurationVersion(configurationVersion string) *ReplaceResolverDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace resolver default response
func (o *ReplaceResolverDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace resolver default response
func (o *ReplaceResolverDefault) WithPayload(payload *models.Error) *ReplaceResolverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace resolver default response
func (o *ReplaceResolverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceResolverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
