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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v5/models"
)

// GetSpoeAgentsHandlerFunc turns a function with the right signature into a get spoe agents handler
type GetSpoeAgentsHandlerFunc func(GetSpoeAgentsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpoeAgentsHandlerFunc) Handle(params GetSpoeAgentsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSpoeAgentsHandler interface for that can handle valid get spoe agents params
type GetSpoeAgentsHandler interface {
	Handle(GetSpoeAgentsParams, interface{}) middleware.Responder
}

// NewGetSpoeAgents creates a new http.Handler for the get spoe agents operation
func NewGetSpoeAgents(ctx *middleware.Context, handler GetSpoeAgentsHandler) *GetSpoeAgents {
	return &GetSpoeAgents{Context: ctx, Handler: handler}
}

/*
	GetSpoeAgents swagger:route GET /services/haproxy/spoe/spoe_agents Spoe getSpoeAgents

# Return an array of spoe agents in one scope

Returns an array of all configured spoe agents in one scope.
*/
type GetSpoeAgents struct {
	Context *middleware.Context
	Handler GetSpoeAgentsHandler
}

func (o *GetSpoeAgents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSpoeAgentsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetSpoeAgentsOKBody get spoe agents o k body
//
// swagger:model GetSpoeAgentsOKBody
type GetSpoeAgentsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeAgents `json:"data"`
}

// Validate validates this get spoe agents o k body
func (o *GetSpoeAgentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeAgentsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeAgentsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeAgentsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getSpoeAgentsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get spoe agents o k body based on the context it is used
func (o *GetSpoeAgentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeAgentsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeAgentsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getSpoeAgentsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeAgentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeAgentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeAgentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
