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

package quic_initial_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetQUICInitialRuleFrontendHandlerFunc turns a function with the right signature into a get QUIC initial rule frontend handler
type GetQUICInitialRuleFrontendHandlerFunc func(GetQUICInitialRuleFrontendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQUICInitialRuleFrontendHandlerFunc) Handle(params GetQUICInitialRuleFrontendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetQUICInitialRuleFrontendHandler interface for that can handle valid get QUIC initial rule frontend params
type GetQUICInitialRuleFrontendHandler interface {
	Handle(GetQUICInitialRuleFrontendParams, interface{}) middleware.Responder
}

// NewGetQUICInitialRuleFrontend creates a new http.Handler for the get QUIC initial rule frontend operation
func NewGetQUICInitialRuleFrontend(ctx *middleware.Context, handler GetQUICInitialRuleFrontendHandler) *GetQUICInitialRuleFrontend {
	return &GetQUICInitialRuleFrontend{Context: ctx, Handler: handler}
}

/*
	GetQUICInitialRuleFrontend swagger:route GET /services/haproxy/configuration/frontends/{parent_name}/quic_initial_rules/{index} QUICInitialRule getQuicInitialRuleFrontend

# Return one QUIC Initial Rule

Returns one QUIC Initial Rule configuration by it's index in the specified parent.
*/
type GetQUICInitialRuleFrontend struct {
	Context *middleware.Context
	Handler GetQUICInitialRuleFrontendHandler
}

func (o *GetQUICInitialRuleFrontend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetQUICInitialRuleFrontendParams()
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
