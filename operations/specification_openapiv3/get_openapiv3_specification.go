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

package specification_openapiv3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetOpenapiv3SpecificationHandlerFunc turns a function with the right signature into a get openapiv3 specification handler
type GetOpenapiv3SpecificationHandlerFunc func(GetOpenapiv3SpecificationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOpenapiv3SpecificationHandlerFunc) Handle(params GetOpenapiv3SpecificationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetOpenapiv3SpecificationHandler interface for that can handle valid get openapiv3 specification params
type GetOpenapiv3SpecificationHandler interface {
	Handle(GetOpenapiv3SpecificationParams, interface{}) middleware.Responder
}

// NewGetOpenapiv3Specification creates a new http.Handler for the get openapiv3 specification operation
func NewGetOpenapiv3Specification(ctx *middleware.Context, handler GetOpenapiv3SpecificationHandler) *GetOpenapiv3Specification {
	return &GetOpenapiv3Specification{Context: ctx, Handler: handler}
}

/*GetOpenapiv3Specification swagger:route GET /specification_openapiv3 SpecificationOpenapiv3 getOpenapiv3Specification

Data Plane API v3 Specification

Return Data Plane API OpenAPI v3 specification

*/
type GetOpenapiv3Specification struct {
	Context *middleware.Context
	Handler GetOpenapiv3SpecificationHandler
}

func (o *GetOpenapiv3Specification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOpenapiv3SpecificationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}