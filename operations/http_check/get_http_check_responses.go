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

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetHTTPCheckOKCode is the HTTP code returned for type GetHTTPCheckOK
const GetHTTPCheckOKCode int = 200

/*GetHTTPCheckOK Successful operation

swagger:response getHttpCheckOK
*/
type GetHTTPCheckOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPCheckOKBody `json:"body,omitempty"`
}

// NewGetHTTPCheckOK creates GetHTTPCheckOK with default headers values
func NewGetHTTPCheckOK() *GetHTTPCheckOK {

	return &GetHTTPCheckOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http check o k response
func (o *GetHTTPCheckOK) WithConfigurationVersion(configurationVersion string) *GetHTTPCheckOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http check o k response
func (o *GetHTTPCheckOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http check o k response
func (o *GetHTTPCheckOK) WithPayload(payload *GetHTTPCheckOKBody) *GetHTTPCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http check o k response
func (o *GetHTTPCheckOK) SetPayload(payload *GetHTTPCheckOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHTTPCheckNotFoundCode is the HTTP code returned for type GetHTTPCheckNotFound
const GetHTTPCheckNotFoundCode int = 404

/*GetHTTPCheckNotFound The specified resource was not found

swagger:response getHttpCheckNotFound
*/
type GetHTTPCheckNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPCheckNotFound creates GetHTTPCheckNotFound with default headers values
func NewGetHTTPCheckNotFound() *GetHTTPCheckNotFound {

	return &GetHTTPCheckNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http check not found response
func (o *GetHTTPCheckNotFound) WithConfigurationVersion(configurationVersion string) *GetHTTPCheckNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http check not found response
func (o *GetHTTPCheckNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http check not found response
func (o *GetHTTPCheckNotFound) WithPayload(payload *models.Error) *GetHTTPCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http check not found response
func (o *GetHTTPCheckNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetHTTPCheckDefault General Error

swagger:response getHttpCheckDefault
*/
type GetHTTPCheckDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPCheckDefault creates GetHTTPCheckDefault with default headers values
func NewGetHTTPCheckDefault(code int) *GetHTTPCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP check default response
func (o *GetHTTPCheckDefault) WithStatusCode(code int) *GetHTTPCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP check default response
func (o *GetHTTPCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP check default response
func (o *GetHTTPCheckDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPCheckDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP check default response
func (o *GetHTTPCheckDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP check default response
func (o *GetHTTPCheckDefault) WithPayload(payload *models.Error) *GetHTTPCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP check default response
func (o *GetHTTPCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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