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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// GetAWSRegionsOKCode is the HTTP code returned for type GetAWSRegionsOK
const GetAWSRegionsOKCode int = 200

/*
GetAWSRegionsOK Successful operation

swagger:response getAWSRegionsOK
*/
type GetAWSRegionsOK struct {

	/*
	  In: Body
	*/
	Payload models.AwsRegions `json:"body,omitempty"`
}

// NewGetAWSRegionsOK creates GetAWSRegionsOK with default headers values
func NewGetAWSRegionsOK() *GetAWSRegionsOK {

	return &GetAWSRegionsOK{}
}

// WithPayload adds the payload to the get a w s regions o k response
func (o *GetAWSRegionsOK) WithPayload(payload models.AwsRegions) *GetAWSRegionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions o k response
func (o *GetAWSRegionsOK) SetPayload(payload models.AwsRegions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.AwsRegions{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetAWSRegionsDefault General Error

swagger:response getAWSRegionsDefault
*/
type GetAWSRegionsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionsDefault creates GetAWSRegionsDefault with default headers values
func NewGetAWSRegionsDefault(code int) *GetAWSRegionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAWSRegionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get a w s regions default response
func (o *GetAWSRegionsDefault) WithStatusCode(code int) *GetAWSRegionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get a w s regions default response
func (o *GetAWSRegionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get a w s regions default response
func (o *GetAWSRegionsDefault) WithConfigurationVersion(configurationVersion string) *GetAWSRegionsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get a w s regions default response
func (o *GetAWSRegionsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get a w s regions default response
func (o *GetAWSRegionsDefault) WithPayload(payload *models.Error) *GetAWSRegionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions default response
func (o *GetAWSRegionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
