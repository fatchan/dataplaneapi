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

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteCacheAcceptedCode is the HTTP code returned for type DeleteCacheAccepted
const DeleteCacheAcceptedCode int = 202

/*
DeleteCacheAccepted Configuration change accepted and reload requested

swagger:response deleteCacheAccepted
*/
type DeleteCacheAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteCacheAccepted creates DeleteCacheAccepted with default headers values
func NewDeleteCacheAccepted() *DeleteCacheAccepted {

	return &DeleteCacheAccepted{}
}

// WithReloadID adds the reloadId to the delete cache accepted response
func (o *DeleteCacheAccepted) WithReloadID(reloadID string) *DeleteCacheAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete cache accepted response
func (o *DeleteCacheAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteCacheAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteCacheNoContentCode is the HTTP code returned for type DeleteCacheNoContent
const DeleteCacheNoContentCode int = 204

/*
DeleteCacheNoContent Cache deleted

swagger:response deleteCacheNoContent
*/
type DeleteCacheNoContent struct {
}

// NewDeleteCacheNoContent creates DeleteCacheNoContent with default headers values
func NewDeleteCacheNoContent() *DeleteCacheNoContent {

	return &DeleteCacheNoContent{}
}

// WriteResponse to the client
func (o *DeleteCacheNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteCacheNotFoundCode is the HTTP code returned for type DeleteCacheNotFound
const DeleteCacheNotFoundCode int = 404

/*
DeleteCacheNotFound The specified resource was not found

swagger:response deleteCacheNotFound
*/
type DeleteCacheNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCacheNotFound creates DeleteCacheNotFound with default headers values
func NewDeleteCacheNotFound() *DeleteCacheNotFound {

	return &DeleteCacheNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete cache not found response
func (o *DeleteCacheNotFound) WithConfigurationVersion(configurationVersion string) *DeleteCacheNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete cache not found response
func (o *DeleteCacheNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete cache not found response
func (o *DeleteCacheNotFound) WithPayload(payload *models.Error) *DeleteCacheNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cache not found response
func (o *DeleteCacheNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCacheNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteCacheDefault General Error

swagger:response deleteCacheDefault
*/
type DeleteCacheDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCacheDefault creates DeleteCacheDefault with default headers values
func NewDeleteCacheDefault(code int) *DeleteCacheDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteCacheDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete cache default response
func (o *DeleteCacheDefault) WithStatusCode(code int) *DeleteCacheDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete cache default response
func (o *DeleteCacheDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete cache default response
func (o *DeleteCacheDefault) WithConfigurationVersion(configurationVersion string) *DeleteCacheDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete cache default response
func (o *DeleteCacheDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete cache default response
func (o *DeleteCacheDefault) WithPayload(payload *models.Error) *DeleteCacheDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cache default response
func (o *DeleteCacheDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCacheDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
