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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceAclsOKCode is the HTTP code returned for type ReplaceAclsOK
const ReplaceAclsOKCode int = 200

/*
ReplaceAclsOK All ACL lines replaced

swagger:response replaceAclsOK
*/
type ReplaceAclsOK struct {

	/*
	  In: Body
	*/
	Payload models.Acls `json:"body,omitempty"`
}

// NewReplaceAclsOK creates ReplaceAclsOK with default headers values
func NewReplaceAclsOK() *ReplaceAclsOK {

	return &ReplaceAclsOK{}
}

// WithPayload adds the payload to the replace acls o k response
func (o *ReplaceAclsOK) WithPayload(payload models.Acls) *ReplaceAclsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace acls o k response
func (o *ReplaceAclsOK) SetPayload(payload models.Acls) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAclsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Acls{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAclsAcceptedCode is the HTTP code returned for type ReplaceAclsAccepted
const ReplaceAclsAcceptedCode int = 202

/*
ReplaceAclsAccepted Configuration change accepted and reload requested

swagger:response replaceAclsAccepted
*/
type ReplaceAclsAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.Acls `json:"body,omitempty"`
}

// NewReplaceAclsAccepted creates ReplaceAclsAccepted with default headers values
func NewReplaceAclsAccepted() *ReplaceAclsAccepted {

	return &ReplaceAclsAccepted{}
}

// WithReloadID adds the reloadId to the replace acls accepted response
func (o *ReplaceAclsAccepted) WithReloadID(reloadID string) *ReplaceAclsAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace acls accepted response
func (o *ReplaceAclsAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace acls accepted response
func (o *ReplaceAclsAccepted) WithPayload(payload models.Acls) *ReplaceAclsAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace acls accepted response
func (o *ReplaceAclsAccepted) SetPayload(payload models.Acls) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAclsAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Acls{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAclsBadRequestCode is the HTTP code returned for type ReplaceAclsBadRequest
const ReplaceAclsBadRequestCode int = 400

/*
ReplaceAclsBadRequest Bad request

swagger:response replaceAclsBadRequest
*/
type ReplaceAclsBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAclsBadRequest creates ReplaceAclsBadRequest with default headers values
func NewReplaceAclsBadRequest() *ReplaceAclsBadRequest {

	return &ReplaceAclsBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace acls bad request response
func (o *ReplaceAclsBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceAclsBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace acls bad request response
func (o *ReplaceAclsBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace acls bad request response
func (o *ReplaceAclsBadRequest) WithPayload(payload *models.Error) *ReplaceAclsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace acls bad request response
func (o *ReplaceAclsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAclsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceAclsDefault General Error

swagger:response replaceAclsDefault
*/
type ReplaceAclsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAclsDefault creates ReplaceAclsDefault with default headers values
func NewReplaceAclsDefault(code int) *ReplaceAclsDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceAclsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace acls default response
func (o *ReplaceAclsDefault) WithStatusCode(code int) *ReplaceAclsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace acls default response
func (o *ReplaceAclsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace acls default response
func (o *ReplaceAclsDefault) WithConfigurationVersion(configurationVersion string) *ReplaceAclsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace acls default response
func (o *ReplaceAclsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace acls default response
func (o *ReplaceAclsDefault) WithPayload(payload *models.Error) *ReplaceAclsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace acls default response
func (o *ReplaceAclsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAclsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
