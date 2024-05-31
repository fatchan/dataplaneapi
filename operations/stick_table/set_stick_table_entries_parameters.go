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

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewSetStickTableEntriesParams creates a new SetStickTableEntriesParams object
//
// There are no default values defined in the spec.
func NewSetStickTableEntriesParams() SetStickTableEntriesParams {

	return SetStickTableEntriesParams{}
}

// SetStickTableEntriesParams contains all the bound params for the set stick table entries operation
// typically these are obtained from a http.Request
//
// swagger:parameters setStickTableEntries
type SetStickTableEntriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Stick table name
	  Required: true
	  In: query
	*/
	StickTable string
	/*Stick table entry
	  In: body
	*/
	StickTableEntry SetStickTableEntriesBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetStickTableEntriesParams() beforehand.
func (o *SetStickTableEntriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qStickTable, qhkStickTable, _ := qs.GetOK("stick_table")
	if err := o.bindStickTable(qStickTable, qhkStickTable, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body SetStickTableEntriesBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("stickTableEntry", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.StickTableEntry = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindStickTable binds and validates parameter StickTable from query.
func (o *SetStickTableEntriesParams) bindStickTable(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("stick_table", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("stick_table", "query", raw); err != nil {
		return err
	}
	o.StickTable = raw

	return nil
}
