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

	"github.com/haproxytech/client-native/v6/models"
)

// GetAllStorageSSLCertificatesOKCode is the HTTP code returned for type GetAllStorageSSLCertificatesOK
const GetAllStorageSSLCertificatesOKCode int = 200

/*
GetAllStorageSSLCertificatesOK Successful operation

swagger:response getAllStorageSSLCertificatesOK
*/
type GetAllStorageSSLCertificatesOK struct {

	/*
	  In: Body
	*/
	Payload models.SslCertificates `json:"body,omitempty"`
}

// NewGetAllStorageSSLCertificatesOK creates GetAllStorageSSLCertificatesOK with default headers values
func NewGetAllStorageSSLCertificatesOK() *GetAllStorageSSLCertificatesOK {

	return &GetAllStorageSSLCertificatesOK{}
}

// WithPayload adds the payload to the get all storage s s l certificates o k response
func (o *GetAllStorageSSLCertificatesOK) WithPayload(payload models.SslCertificates) *GetAllStorageSSLCertificatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage s s l certificates o k response
func (o *GetAllStorageSSLCertificatesOK) SetPayload(payload models.SslCertificates) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageSSLCertificatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.SslCertificates{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllStorageSSLCertificatesNotFoundCode is the HTTP code returned for type GetAllStorageSSLCertificatesNotFound
const GetAllStorageSSLCertificatesNotFoundCode int = 404

/*
GetAllStorageSSLCertificatesNotFound The specified resource was not found

swagger:response getAllStorageSSLCertificatesNotFound
*/
type GetAllStorageSSLCertificatesNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllStorageSSLCertificatesNotFound creates GetAllStorageSSLCertificatesNotFound with default headers values
func NewGetAllStorageSSLCertificatesNotFound() *GetAllStorageSSLCertificatesNotFound {

	return &GetAllStorageSSLCertificatesNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get all storage s s l certificates not found response
func (o *GetAllStorageSSLCertificatesNotFound) WithConfigurationVersion(configurationVersion string) *GetAllStorageSSLCertificatesNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all storage s s l certificates not found response
func (o *GetAllStorageSSLCertificatesNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all storage s s l certificates not found response
func (o *GetAllStorageSSLCertificatesNotFound) WithPayload(payload *models.Error) *GetAllStorageSSLCertificatesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage s s l certificates not found response
func (o *GetAllStorageSSLCertificatesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageSSLCertificatesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetAllStorageSSLCertificatesDefault General Error

swagger:response getAllStorageSSLCertificatesDefault
*/
type GetAllStorageSSLCertificatesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllStorageSSLCertificatesDefault creates GetAllStorageSSLCertificatesDefault with default headers values
func NewGetAllStorageSSLCertificatesDefault(code int) *GetAllStorageSSLCertificatesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllStorageSSLCertificatesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) WithStatusCode(code int) *GetAllStorageSSLCertificatesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) WithConfigurationVersion(configurationVersion string) *GetAllStorageSSLCertificatesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) WithPayload(payload *models.Error) *GetAllStorageSSLCertificatesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageSSLCertificatesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
