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

	"github.com/haproxytech/client-native/v6/models"
)

// DeleteDefaultsSectionAcceptedCode is the HTTP code returned for type DeleteDefaultsSectionAccepted
const DeleteDefaultsSectionAcceptedCode int = 202

/*
DeleteDefaultsSectionAccepted Configuration change accepted and reload requested

swagger:response deleteDefaultsSectionAccepted
*/
type DeleteDefaultsSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteDefaultsSectionAccepted creates DeleteDefaultsSectionAccepted with default headers values
func NewDeleteDefaultsSectionAccepted() *DeleteDefaultsSectionAccepted {

	return &DeleteDefaultsSectionAccepted{}
}

// WithReloadID adds the reloadId to the delete defaults section accepted response
func (o *DeleteDefaultsSectionAccepted) WithReloadID(reloadID string) *DeleteDefaultsSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete defaults section accepted response
func (o *DeleteDefaultsSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteDefaultsSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteDefaultsSectionNoContentCode is the HTTP code returned for type DeleteDefaultsSectionNoContent
const DeleteDefaultsSectionNoContentCode int = 204

/*
DeleteDefaultsSectionNoContent Defaults section deleted

swagger:response deleteDefaultsSectionNoContent
*/
type DeleteDefaultsSectionNoContent struct {
}

// NewDeleteDefaultsSectionNoContent creates DeleteDefaultsSectionNoContent with default headers values
func NewDeleteDefaultsSectionNoContent() *DeleteDefaultsSectionNoContent {

	return &DeleteDefaultsSectionNoContent{}
}

// WriteResponse to the client
func (o *DeleteDefaultsSectionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteDefaultsSectionNotFoundCode is the HTTP code returned for type DeleteDefaultsSectionNotFound
const DeleteDefaultsSectionNotFoundCode int = 404

/*
DeleteDefaultsSectionNotFound The specified resource was not found

swagger:response deleteDefaultsSectionNotFound
*/
type DeleteDefaultsSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDefaultsSectionNotFound creates DeleteDefaultsSectionNotFound with default headers values
func NewDeleteDefaultsSectionNotFound() *DeleteDefaultsSectionNotFound {

	return &DeleteDefaultsSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete defaults section not found response
func (o *DeleteDefaultsSectionNotFound) WithConfigurationVersion(configurationVersion string) *DeleteDefaultsSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete defaults section not found response
func (o *DeleteDefaultsSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete defaults section not found response
func (o *DeleteDefaultsSectionNotFound) WithPayload(payload *models.Error) *DeleteDefaultsSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete defaults section not found response
func (o *DeleteDefaultsSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDefaultsSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteDefaultsSectionDefault General Error

swagger:response deleteDefaultsSectionDefault
*/
type DeleteDefaultsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDefaultsSectionDefault creates DeleteDefaultsSectionDefault with default headers values
func NewDeleteDefaultsSectionDefault(code int) *DeleteDefaultsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDefaultsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) WithStatusCode(code int) *DeleteDefaultsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) WithConfigurationVersion(configurationVersion string) *DeleteDefaultsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) WithPayload(payload *models.Error) *DeleteDefaultsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDefaultsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
