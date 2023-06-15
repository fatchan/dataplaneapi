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

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v5/models"
)

// GetSiteOKCode is the HTTP code returned for type GetSiteOK
const GetSiteOKCode int = 200

/*
GetSiteOK Successful operation

swagger:response getSiteOK
*/
type GetSiteOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetSiteOKBody `json:"body,omitempty"`
}

// NewGetSiteOK creates GetSiteOK with default headers values
func NewGetSiteOK() *GetSiteOK {

	return &GetSiteOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get site o k response
func (o *GetSiteOK) WithConfigurationVersion(configurationVersion string) *GetSiteOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get site o k response
func (o *GetSiteOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get site o k response
func (o *GetSiteOK) WithPayload(payload *GetSiteOKBody) *GetSiteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get site o k response
func (o *GetSiteOK) SetPayload(payload *GetSiteOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSiteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetSiteNotFoundCode is the HTTP code returned for type GetSiteNotFound
const GetSiteNotFoundCode int = 404

/*
GetSiteNotFound The specified resource was not found

swagger:response getSiteNotFound
*/
type GetSiteNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSiteNotFound creates GetSiteNotFound with default headers values
func NewGetSiteNotFound() *GetSiteNotFound {

	return &GetSiteNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get site not found response
func (o *GetSiteNotFound) WithConfigurationVersion(configurationVersion string) *GetSiteNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get site not found response
func (o *GetSiteNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get site not found response
func (o *GetSiteNotFound) WithPayload(payload *models.Error) *GetSiteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get site not found response
func (o *GetSiteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSiteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetSiteDefault General Error

swagger:response getSiteDefault
*/
type GetSiteDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSiteDefault creates GetSiteDefault with default headers values
func NewGetSiteDefault(code int) *GetSiteDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSiteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get site default response
func (o *GetSiteDefault) WithStatusCode(code int) *GetSiteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get site default response
func (o *GetSiteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get site default response
func (o *GetSiteDefault) WithConfigurationVersion(configurationVersion string) *GetSiteDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get site default response
func (o *GetSiteDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get site default response
func (o *GetSiteDefault) WithPayload(payload *models.Error) *GetSiteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get site default response
func (o *GetSiteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSiteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
