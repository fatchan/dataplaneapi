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
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSpoeMessagesHandlerFunc turns a function with the right signature into a get spoe messages handler
type GetSpoeMessagesHandlerFunc func(GetSpoeMessagesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpoeMessagesHandlerFunc) Handle(params GetSpoeMessagesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSpoeMessagesHandler interface for that can handle valid get spoe messages params
type GetSpoeMessagesHandler interface {
	Handle(GetSpoeMessagesParams, interface{}) middleware.Responder
}

// NewGetSpoeMessages creates a new http.Handler for the get spoe messages operation
func NewGetSpoeMessages(ctx *middleware.Context, handler GetSpoeMessagesHandler) *GetSpoeMessages {
	return &GetSpoeMessages{Context: ctx, Handler: handler}
}

/*
	GetSpoeMessages swagger:route GET /services/haproxy/spoe/spoe_messages Spoe getSpoeMessages

# Return an array of spoe messages in one scope

Returns an array of all configured spoe messages in one scope.
*/
type GetSpoeMessages struct {
	Context *middleware.Context
	Handler GetSpoeMessagesHandler
}

func (o *GetSpoeMessages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSpoeMessagesParams()
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
