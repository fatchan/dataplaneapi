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

package dgram_bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// DeleteDgramBindAcceptedCode is the HTTP code returned for type DeleteDgramBindAccepted
const DeleteDgramBindAcceptedCode int = 202

/*
DeleteDgramBindAccepted Configuration change accepted and reload requested

swagger:response deleteDgramBindAccepted
*/
type DeleteDgramBindAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteDgramBindAccepted creates DeleteDgramBindAccepted with default headers values
func NewDeleteDgramBindAccepted() *DeleteDgramBindAccepted {

	return &DeleteDgramBindAccepted{}
}

// WithReloadID adds the reloadId to the delete dgram bind accepted response
func (o *DeleteDgramBindAccepted) WithReloadID(reloadID string) *DeleteDgramBindAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete dgram bind accepted response
func (o *DeleteDgramBindAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteDgramBindAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteDgramBindNoContentCode is the HTTP code returned for type DeleteDgramBindNoContent
const DeleteDgramBindNoContentCode int = 204

/*
DeleteDgramBindNoContent Bind deleted

swagger:response deleteDgramBindNoContent
*/
type DeleteDgramBindNoContent struct {
}

// NewDeleteDgramBindNoContent creates DeleteDgramBindNoContent with default headers values
func NewDeleteDgramBindNoContent() *DeleteDgramBindNoContent {

	return &DeleteDgramBindNoContent{}
}

// WriteResponse to the client
func (o *DeleteDgramBindNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteDgramBindNotFoundCode is the HTTP code returned for type DeleteDgramBindNotFound
const DeleteDgramBindNotFoundCode int = 404

/*
DeleteDgramBindNotFound The specified resource was not found

swagger:response deleteDgramBindNotFound
*/
type DeleteDgramBindNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDgramBindNotFound creates DeleteDgramBindNotFound with default headers values
func NewDeleteDgramBindNotFound() *DeleteDgramBindNotFound {

	return &DeleteDgramBindNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete dgram bind not found response
func (o *DeleteDgramBindNotFound) WithConfigurationVersion(configurationVersion string) *DeleteDgramBindNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete dgram bind not found response
func (o *DeleteDgramBindNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete dgram bind not found response
func (o *DeleteDgramBindNotFound) WithPayload(payload *models.Error) *DeleteDgramBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dgram bind not found response
func (o *DeleteDgramBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDgramBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteDgramBindDefault General Error

swagger:response deleteDgramBindDefault
*/
type DeleteDgramBindDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDgramBindDefault creates DeleteDgramBindDefault with default headers values
func NewDeleteDgramBindDefault(code int) *DeleteDgramBindDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDgramBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete dgram bind default response
func (o *DeleteDgramBindDefault) WithStatusCode(code int) *DeleteDgramBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete dgram bind default response
func (o *DeleteDgramBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete dgram bind default response
func (o *DeleteDgramBindDefault) WithConfigurationVersion(configurationVersion string) *DeleteDgramBindDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete dgram bind default response
func (o *DeleteDgramBindDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete dgram bind default response
func (o *DeleteDgramBindDefault) WithPayload(payload *models.Error) *DeleteDgramBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete dgram bind default response
func (o *DeleteDgramBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDgramBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
