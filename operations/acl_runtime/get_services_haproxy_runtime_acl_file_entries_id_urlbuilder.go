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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetServicesHaproxyRuntimeACLFileEntriesIDURL generates an URL for the get services haproxy runtime ACL file entries ID operation
type GetServicesHaproxyRuntimeACLFileEntriesIDURL struct {
	ID string

	ACLID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) WithBasePath(bp string) *GetServicesHaproxyRuntimeACLFileEntriesIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/services/haproxy/runtime/acl_file_entries/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetServicesHaproxyRuntimeACLFileEntriesIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v3"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	aCLIDQ := o.ACLID
	if aCLIDQ != "" {
		qs.Set("acl_id", aCLIDQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetServicesHaproxyRuntimeACLFileEntriesIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetServicesHaproxyRuntimeACLFileEntriesIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
