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

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceHTTPErrorRulesOKCode is the HTTP code returned for type ReplaceHTTPErrorRulesOK
const ReplaceHTTPErrorRulesOKCode int = 200

/*
ReplaceHTTPErrorRulesOK All HTTP Error Rules lines replaced

swagger:response replaceHttpErrorRulesOK
*/
type ReplaceHTTPErrorRulesOK struct {

	/*
	  In: Body
	*/
	Payload models.HTTPErrorRules `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRulesOK creates ReplaceHTTPErrorRulesOK with default headers values
func NewReplaceHTTPErrorRulesOK() *ReplaceHTTPErrorRulesOK {

	return &ReplaceHTTPErrorRulesOK{}
}

// WithPayload adds the payload to the replace Http error rules o k response
func (o *ReplaceHTTPErrorRulesOK) WithPayload(payload models.HTTPErrorRules) *ReplaceHTTPErrorRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rules o k response
func (o *ReplaceHTTPErrorRulesOK) SetPayload(payload models.HTTPErrorRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPErrorRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceHTTPErrorRulesAcceptedCode is the HTTP code returned for type ReplaceHTTPErrorRulesAccepted
const ReplaceHTTPErrorRulesAcceptedCode int = 202

/*
ReplaceHTTPErrorRulesAccepted Configuration change accepted and reload requested

swagger:response replaceHttpErrorRulesAccepted
*/
type ReplaceHTTPErrorRulesAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.HTTPErrorRules `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRulesAccepted creates ReplaceHTTPErrorRulesAccepted with default headers values
func NewReplaceHTTPErrorRulesAccepted() *ReplaceHTTPErrorRulesAccepted {

	return &ReplaceHTTPErrorRulesAccepted{}
}

// WithReloadID adds the reloadId to the replace Http error rules accepted response
func (o *ReplaceHTTPErrorRulesAccepted) WithReloadID(reloadID string) *ReplaceHTTPErrorRulesAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Http error rules accepted response
func (o *ReplaceHTTPErrorRulesAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Http error rules accepted response
func (o *ReplaceHTTPErrorRulesAccepted) WithPayload(payload models.HTTPErrorRules) *ReplaceHTTPErrorRulesAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rules accepted response
func (o *ReplaceHTTPErrorRulesAccepted) SetPayload(payload models.HTTPErrorRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRulesAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPErrorRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceHTTPErrorRulesBadRequestCode is the HTTP code returned for type ReplaceHTTPErrorRulesBadRequest
const ReplaceHTTPErrorRulesBadRequestCode int = 400

/*
ReplaceHTTPErrorRulesBadRequest Bad request

swagger:response replaceHttpErrorRulesBadRequest
*/
type ReplaceHTTPErrorRulesBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRulesBadRequest creates ReplaceHTTPErrorRulesBadRequest with default headers values
func NewReplaceHTTPErrorRulesBadRequest() *ReplaceHTTPErrorRulesBadRequest {

	return &ReplaceHTTPErrorRulesBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http error rules bad request response
func (o *ReplaceHTTPErrorRulesBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorRulesBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http error rules bad request response
func (o *ReplaceHTTPErrorRulesBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http error rules bad request response
func (o *ReplaceHTTPErrorRulesBadRequest) WithPayload(payload *models.Error) *ReplaceHTTPErrorRulesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rules bad request response
func (o *ReplaceHTTPErrorRulesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRulesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceHTTPErrorRulesDefault General Error

swagger:response replaceHttpErrorRulesDefault
*/
type ReplaceHTTPErrorRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRulesDefault creates ReplaceHTTPErrorRulesDefault with default headers values
func NewReplaceHTTPErrorRulesDefault(code int) *ReplaceHTTPErrorRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceHTTPErrorRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) WithStatusCode(code int) *ReplaceHTTPErrorRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) WithPayload(payload *models.Error) *ReplaceHTTPErrorRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace HTTP error rules default response
func (o *ReplaceHTTPErrorRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
