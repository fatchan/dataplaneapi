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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// GetSpoeTransactionOKCode is the HTTP code returned for type GetSpoeTransactionOK
const GetSpoeTransactionOKCode int = 200

/*
GetSpoeTransactionOK Successful operation

swagger:response getSpoeTransactionOK
*/
type GetSpoeTransactionOK struct {

	/*
	  In: Body
	*/
	Payload *models.SpoeTransaction `json:"body,omitempty"`
}

// NewGetSpoeTransactionOK creates GetSpoeTransactionOK with default headers values
func NewGetSpoeTransactionOK() *GetSpoeTransactionOK {

	return &GetSpoeTransactionOK{}
}

// WithPayload adds the payload to the get spoe transaction o k response
func (o *GetSpoeTransactionOK) WithPayload(payload *models.SpoeTransaction) *GetSpoeTransactionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe transaction o k response
func (o *GetSpoeTransactionOK) SetPayload(payload *models.SpoeTransaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeTransactionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSpoeTransactionNotFoundCode is the HTTP code returned for type GetSpoeTransactionNotFound
const GetSpoeTransactionNotFoundCode int = 404

/*
GetSpoeTransactionNotFound The specified resource was not found

swagger:response getSpoeTransactionNotFound
*/
type GetSpoeTransactionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeTransactionNotFound creates GetSpoeTransactionNotFound with default headers values
func NewGetSpoeTransactionNotFound() *GetSpoeTransactionNotFound {

	return &GetSpoeTransactionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe transaction not found response
func (o *GetSpoeTransactionNotFound) WithConfigurationVersion(configurationVersion string) *GetSpoeTransactionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe transaction not found response
func (o *GetSpoeTransactionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe transaction not found response
func (o *GetSpoeTransactionNotFound) WithPayload(payload *models.Error) *GetSpoeTransactionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe transaction not found response
func (o *GetSpoeTransactionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetSpoeTransactionDefault General Error

swagger:response getSpoeTransactionDefault
*/
type GetSpoeTransactionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeTransactionDefault creates GetSpoeTransactionDefault with default headers values
func NewGetSpoeTransactionDefault(code int) *GetSpoeTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) WithStatusCode(code int) *GetSpoeTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) WithPayload(payload *models.Error) *GetSpoeTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe transaction default response
func (o *GetSpoeTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
