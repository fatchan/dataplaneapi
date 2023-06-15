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

// ReplaceSiteOKCode is the HTTP code returned for type ReplaceSiteOK
const ReplaceSiteOKCode int = 200

/*
ReplaceSiteOK Site replaced

swagger:response replaceSiteOK
*/
type ReplaceSiteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Site `json:"body,omitempty"`
}

// NewReplaceSiteOK creates ReplaceSiteOK with default headers values
func NewReplaceSiteOK() *ReplaceSiteOK {

	return &ReplaceSiteOK{}
}

// WithPayload adds the payload to the replace site o k response
func (o *ReplaceSiteOK) WithPayload(payload *models.Site) *ReplaceSiteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace site o k response
func (o *ReplaceSiteOK) SetPayload(payload *models.Site) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSiteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceSiteAcceptedCode is the HTTP code returned for type ReplaceSiteAccepted
const ReplaceSiteAcceptedCode int = 202

/*
ReplaceSiteAccepted Configuration change accepted and reload requested

swagger:response replaceSiteAccepted
*/
type ReplaceSiteAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Site `json:"body,omitempty"`
}

// NewReplaceSiteAccepted creates ReplaceSiteAccepted with default headers values
func NewReplaceSiteAccepted() *ReplaceSiteAccepted {

	return &ReplaceSiteAccepted{}
}

// WithReloadID adds the reloadId to the replace site accepted response
func (o *ReplaceSiteAccepted) WithReloadID(reloadID string) *ReplaceSiteAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace site accepted response
func (o *ReplaceSiteAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace site accepted response
func (o *ReplaceSiteAccepted) WithPayload(payload *models.Site) *ReplaceSiteAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace site accepted response
func (o *ReplaceSiteAccepted) SetPayload(payload *models.Site) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSiteAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceSiteBadRequestCode is the HTTP code returned for type ReplaceSiteBadRequest
const ReplaceSiteBadRequestCode int = 400

/*
ReplaceSiteBadRequest Bad request

swagger:response replaceSiteBadRequest
*/
type ReplaceSiteBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceSiteBadRequest creates ReplaceSiteBadRequest with default headers values
func NewReplaceSiteBadRequest() *ReplaceSiteBadRequest {

	return &ReplaceSiteBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace site bad request response
func (o *ReplaceSiteBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceSiteBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace site bad request response
func (o *ReplaceSiteBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace site bad request response
func (o *ReplaceSiteBadRequest) WithPayload(payload *models.Error) *ReplaceSiteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace site bad request response
func (o *ReplaceSiteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSiteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceSiteNotFoundCode is the HTTP code returned for type ReplaceSiteNotFound
const ReplaceSiteNotFoundCode int = 404

/*
ReplaceSiteNotFound The specified resource was not found

swagger:response replaceSiteNotFound
*/
type ReplaceSiteNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceSiteNotFound creates ReplaceSiteNotFound with default headers values
func NewReplaceSiteNotFound() *ReplaceSiteNotFound {

	return &ReplaceSiteNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace site not found response
func (o *ReplaceSiteNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceSiteNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace site not found response
func (o *ReplaceSiteNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace site not found response
func (o *ReplaceSiteNotFound) WithPayload(payload *models.Error) *ReplaceSiteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace site not found response
func (o *ReplaceSiteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSiteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceSiteDefault General Error

swagger:response replaceSiteDefault
*/
type ReplaceSiteDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceSiteDefault creates ReplaceSiteDefault with default headers values
func NewReplaceSiteDefault(code int) *ReplaceSiteDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceSiteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace site default response
func (o *ReplaceSiteDefault) WithStatusCode(code int) *ReplaceSiteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace site default response
func (o *ReplaceSiteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace site default response
func (o *ReplaceSiteDefault) WithConfigurationVersion(configurationVersion string) *ReplaceSiteDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace site default response
func (o *ReplaceSiteDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace site default response
func (o *ReplaceSiteDefault) WithPayload(payload *models.Error) *ReplaceSiteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace site default response
func (o *ReplaceSiteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSiteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
