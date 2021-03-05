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

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceHTTPRequestRuleOKCode is the HTTP code returned for type ReplaceHTTPRequestRuleOK
const ReplaceHTTPRequestRuleOKCode int = 200

/*ReplaceHTTPRequestRuleOK HTTP Request Rule replaced

swagger:response replaceHttpRequestRuleOK
*/
type ReplaceHTTPRequestRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPRequestRule `json:"body,omitempty"`
}

// NewReplaceHTTPRequestRuleOK creates ReplaceHTTPRequestRuleOK with default headers values
func NewReplaceHTTPRequestRuleOK() *ReplaceHTTPRequestRuleOK {

	return &ReplaceHTTPRequestRuleOK{}
}

// WithPayload adds the payload to the replace Http request rule o k response
func (o *ReplaceHTTPRequestRuleOK) WithPayload(payload *models.HTTPRequestRule) *ReplaceHTTPRequestRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http request rule o k response
func (o *ReplaceHTTPRequestRuleOK) SetPayload(payload *models.HTTPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPRequestRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPRequestRuleAcceptedCode is the HTTP code returned for type ReplaceHTTPRequestRuleAccepted
const ReplaceHTTPRequestRuleAcceptedCode int = 202

/*ReplaceHTTPRequestRuleAccepted Configuration change accepted and reload requested

swagger:response replaceHttpRequestRuleAccepted
*/
type ReplaceHTTPRequestRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.HTTPRequestRule `json:"body,omitempty"`
}

// NewReplaceHTTPRequestRuleAccepted creates ReplaceHTTPRequestRuleAccepted with default headers values
func NewReplaceHTTPRequestRuleAccepted() *ReplaceHTTPRequestRuleAccepted {

	return &ReplaceHTTPRequestRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace Http request rule accepted response
func (o *ReplaceHTTPRequestRuleAccepted) WithReloadID(reloadID string) *ReplaceHTTPRequestRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Http request rule accepted response
func (o *ReplaceHTTPRequestRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Http request rule accepted response
func (o *ReplaceHTTPRequestRuleAccepted) WithPayload(payload *models.HTTPRequestRule) *ReplaceHTTPRequestRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http request rule accepted response
func (o *ReplaceHTTPRequestRuleAccepted) SetPayload(payload *models.HTTPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPRequestRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPRequestRuleBadRequestCode is the HTTP code returned for type ReplaceHTTPRequestRuleBadRequest
const ReplaceHTTPRequestRuleBadRequestCode int = 400

/*ReplaceHTTPRequestRuleBadRequest Bad request

swagger:response replaceHttpRequestRuleBadRequest
*/
type ReplaceHTTPRequestRuleBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPRequestRuleBadRequest creates ReplaceHTTPRequestRuleBadRequest with default headers values
func NewReplaceHTTPRequestRuleBadRequest() *ReplaceHTTPRequestRuleBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceHTTPRequestRuleBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http request rule bad request response
func (o *ReplaceHTTPRequestRuleBadRequest) WithConfigurationVersion(configurationVersion int64) *ReplaceHTTPRequestRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http request rule bad request response
func (o *ReplaceHTTPRequestRuleBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http request rule bad request response
func (o *ReplaceHTTPRequestRuleBadRequest) WithPayload(payload *models.Error) *ReplaceHTTPRequestRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http request rule bad request response
func (o *ReplaceHTTPRequestRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPRequestRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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

// ReplaceHTTPRequestRuleNotFoundCode is the HTTP code returned for type ReplaceHTTPRequestRuleNotFound
const ReplaceHTTPRequestRuleNotFoundCode int = 404

/*ReplaceHTTPRequestRuleNotFound The specified resource was not found

swagger:response replaceHttpRequestRuleNotFound
*/
type ReplaceHTTPRequestRuleNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPRequestRuleNotFound creates ReplaceHTTPRequestRuleNotFound with default headers values
func NewReplaceHTTPRequestRuleNotFound() *ReplaceHTTPRequestRuleNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceHTTPRequestRuleNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http request rule not found response
func (o *ReplaceHTTPRequestRuleNotFound) WithConfigurationVersion(configurationVersion int64) *ReplaceHTTPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http request rule not found response
func (o *ReplaceHTTPRequestRuleNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http request rule not found response
func (o *ReplaceHTTPRequestRuleNotFound) WithPayload(payload *models.Error) *ReplaceHTTPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http request rule not found response
func (o *ReplaceHTTPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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

/*ReplaceHTTPRequestRuleDefault General Error

swagger:response replaceHttpRequestRuleDefault
*/
type ReplaceHTTPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPRequestRuleDefault creates ReplaceHTTPRequestRuleDefault with default headers values
func NewReplaceHTTPRequestRuleDefault(code int) *ReplaceHTTPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceHTTPRequestRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) WithStatusCode(code int) *ReplaceHTTPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) WithConfigurationVersion(configurationVersion int64) *ReplaceHTTPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) WithPayload(payload *models.Error) *ReplaceHTTPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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
