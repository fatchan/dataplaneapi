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

package http_errors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteHTTPErrorsSectionAcceptedCode is the HTTP code returned for type DeleteHTTPErrorsSectionAccepted
const DeleteHTTPErrorsSectionAcceptedCode int = 202

/*
DeleteHTTPErrorsSectionAccepted Configuration change accepted and reload requested

swagger:response deleteHttpErrorsSectionAccepted
*/
type DeleteHTTPErrorsSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteHTTPErrorsSectionAccepted creates DeleteHTTPErrorsSectionAccepted with default headers values
func NewDeleteHTTPErrorsSectionAccepted() *DeleteHTTPErrorsSectionAccepted {

	return &DeleteHTTPErrorsSectionAccepted{}
}

// WithReloadID adds the reloadId to the delete Http errors section accepted response
func (o *DeleteHTTPErrorsSectionAccepted) WithReloadID(reloadID string) *DeleteHTTPErrorsSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Http errors section accepted response
func (o *DeleteHTTPErrorsSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteHTTPErrorsSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteHTTPErrorsSectionNoContentCode is the HTTP code returned for type DeleteHTTPErrorsSectionNoContent
const DeleteHTTPErrorsSectionNoContentCode int = 204

/*
DeleteHTTPErrorsSectionNoContent http-error section deleted

swagger:response deleteHttpErrorsSectionNoContent
*/
type DeleteHTTPErrorsSectionNoContent struct {
}

// NewDeleteHTTPErrorsSectionNoContent creates DeleteHTTPErrorsSectionNoContent with default headers values
func NewDeleteHTTPErrorsSectionNoContent() *DeleteHTTPErrorsSectionNoContent {

	return &DeleteHTTPErrorsSectionNoContent{}
}

// WriteResponse to the client
func (o *DeleteHTTPErrorsSectionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteHTTPErrorsSectionNotFoundCode is the HTTP code returned for type DeleteHTTPErrorsSectionNotFound
const DeleteHTTPErrorsSectionNotFoundCode int = 404

/*
DeleteHTTPErrorsSectionNotFound The specified resource was not found

swagger:response deleteHttpErrorsSectionNotFound
*/
type DeleteHTTPErrorsSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPErrorsSectionNotFound creates DeleteHTTPErrorsSectionNotFound with default headers values
func NewDeleteHTTPErrorsSectionNotFound() *DeleteHTTPErrorsSectionNotFound {

	return &DeleteHTTPErrorsSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Http errors section not found response
func (o *DeleteHTTPErrorsSectionNotFound) WithConfigurationVersion(configurationVersion string) *DeleteHTTPErrorsSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Http errors section not found response
func (o *DeleteHTTPErrorsSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Http errors section not found response
func (o *DeleteHTTPErrorsSectionNotFound) WithPayload(payload *models.Error) *DeleteHTTPErrorsSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Http errors section not found response
func (o *DeleteHTTPErrorsSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPErrorsSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteHTTPErrorsSectionDefault General Error

swagger:response deleteHttpErrorsSectionDefault
*/
type DeleteHTTPErrorsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPErrorsSectionDefault creates DeleteHTTPErrorsSectionDefault with default headers values
func NewDeleteHTTPErrorsSectionDefault(code int) *DeleteHTTPErrorsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteHTTPErrorsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) WithStatusCode(code int) *DeleteHTTPErrorsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) WithConfigurationVersion(configurationVersion string) *DeleteHTTPErrorsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) WithPayload(payload *models.Error) *DeleteHTTPErrorsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete HTTP errors section default response
func (o *DeleteHTTPErrorsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPErrorsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
