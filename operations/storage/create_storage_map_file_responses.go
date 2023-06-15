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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateStorageMapFileCreatedCode is the HTTP code returned for type CreateStorageMapFileCreated
const CreateStorageMapFileCreatedCode int = 201

/*
CreateStorageMapFileCreated Map file created with its entries

swagger:response createStorageMapFileCreated
*/
type CreateStorageMapFileCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Map `json:"body,omitempty"`
}

// NewCreateStorageMapFileCreated creates CreateStorageMapFileCreated with default headers values
func NewCreateStorageMapFileCreated() *CreateStorageMapFileCreated {

	return &CreateStorageMapFileCreated{}
}

// WithPayload adds the payload to the create storage map file created response
func (o *CreateStorageMapFileCreated) WithPayload(payload *models.Map) *CreateStorageMapFileCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage map file created response
func (o *CreateStorageMapFileCreated) SetPayload(payload *models.Map) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageMapFileCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageMapFileBadRequestCode is the HTTP code returned for type CreateStorageMapFileBadRequest
const CreateStorageMapFileBadRequestCode int = 400

/*
CreateStorageMapFileBadRequest Bad request

swagger:response createStorageMapFileBadRequest
*/
type CreateStorageMapFileBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageMapFileBadRequest creates CreateStorageMapFileBadRequest with default headers values
func NewCreateStorageMapFileBadRequest() *CreateStorageMapFileBadRequest {

	return &CreateStorageMapFileBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create storage map file bad request response
func (o *CreateStorageMapFileBadRequest) WithConfigurationVersion(configurationVersion string) *CreateStorageMapFileBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage map file bad request response
func (o *CreateStorageMapFileBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage map file bad request response
func (o *CreateStorageMapFileBadRequest) WithPayload(payload *models.Error) *CreateStorageMapFileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage map file bad request response
func (o *CreateStorageMapFileBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageMapFileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateStorageMapFileConflictCode is the HTTP code returned for type CreateStorageMapFileConflict
const CreateStorageMapFileConflictCode int = 409

/*
CreateStorageMapFileConflict The specified resource already exists

swagger:response createStorageMapFileConflict
*/
type CreateStorageMapFileConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageMapFileConflict creates CreateStorageMapFileConflict with default headers values
func NewCreateStorageMapFileConflict() *CreateStorageMapFileConflict {

	return &CreateStorageMapFileConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create storage map file conflict response
func (o *CreateStorageMapFileConflict) WithConfigurationVersion(configurationVersion string) *CreateStorageMapFileConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage map file conflict response
func (o *CreateStorageMapFileConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage map file conflict response
func (o *CreateStorageMapFileConflict) WithPayload(payload *models.Error) *CreateStorageMapFileConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage map file conflict response
func (o *CreateStorageMapFileConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageMapFileConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateStorageMapFileDefault General Error

swagger:response createStorageMapFileDefault
*/
type CreateStorageMapFileDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageMapFileDefault creates CreateStorageMapFileDefault with default headers values
func NewCreateStorageMapFileDefault(code int) *CreateStorageMapFileDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateStorageMapFileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create storage map file default response
func (o *CreateStorageMapFileDefault) WithStatusCode(code int) *CreateStorageMapFileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create storage map file default response
func (o *CreateStorageMapFileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create storage map file default response
func (o *CreateStorageMapFileDefault) WithConfigurationVersion(configurationVersion string) *CreateStorageMapFileDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage map file default response
func (o *CreateStorageMapFileDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage map file default response
func (o *CreateStorageMapFileDefault) WithPayload(payload *models.Error) *CreateStorageMapFileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage map file default response
func (o *CreateStorageMapFileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageMapFileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
