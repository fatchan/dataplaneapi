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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateBackendSwitchingRuleHandlerFunc turns a function with the right signature into a create backend switching rule handler
type CreateBackendSwitchingRuleHandlerFunc func(CreateBackendSwitchingRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBackendSwitchingRuleHandlerFunc) Handle(params CreateBackendSwitchingRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateBackendSwitchingRuleHandler interface for that can handle valid create backend switching rule params
type CreateBackendSwitchingRuleHandler interface {
	Handle(CreateBackendSwitchingRuleParams, interface{}) middleware.Responder
}

// NewCreateBackendSwitchingRule creates a new http.Handler for the create backend switching rule operation
func NewCreateBackendSwitchingRule(ctx *middleware.Context, handler CreateBackendSwitchingRuleHandler) *CreateBackendSwitchingRule {
	return &CreateBackendSwitchingRule{Context: ctx, Handler: handler}
}

/*
	CreateBackendSwitchingRule swagger:route POST /services/haproxy/configuration/backend_switching_rules/{index} BackendSwitchingRule createBackendSwitchingRule

# Add a new Backend Switching Rule

Adds a new Backend Switching Rule of the specified type in the specified frontend.
*/
type CreateBackendSwitchingRule struct {
	Context *middleware.Context
	Handler CreateBackendSwitchingRuleHandler
}

func (o *CreateBackendSwitchingRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateBackendSwitchingRuleParams()
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
