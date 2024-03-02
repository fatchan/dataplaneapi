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

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// GetTablesOKCode is the HTTP code returned for type GetTablesOK
const GetTablesOKCode int = 200

/*
GetTablesOK Successful operation

swagger:response getTablesOK
*/
type GetTablesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload models.Tables `json:"body,omitempty"`
}

// NewGetTablesOK creates GetTablesOK with default headers values
func NewGetTablesOK() *GetTablesOK {

	return &GetTablesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get tables o k response
func (o *GetTablesOK) WithConfigurationVersion(configurationVersion string) *GetTablesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get tables o k response
func (o *GetTablesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get tables o k response
func (o *GetTablesOK) WithPayload(payload models.Tables) *GetTablesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tables o k response
func (o *GetTablesOK) SetPayload(payload models.Tables) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTablesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Tables{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetTablesDefault General Error

swagger:response getTablesDefault
*/
type GetTablesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTablesDefault creates GetTablesDefault with default headers values
func NewGetTablesDefault(code int) *GetTablesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTablesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get tables default response
func (o *GetTablesDefault) WithStatusCode(code int) *GetTablesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get tables default response
func (o *GetTablesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get tables default response
func (o *GetTablesDefault) WithConfigurationVersion(configurationVersion string) *GetTablesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get tables default response
func (o *GetTablesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get tables default response
func (o *GetTablesDefault) WithPayload(payload *models.Error) *GetTablesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tables default response
func (o *GetTablesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTablesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
