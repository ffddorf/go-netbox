// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVpnTunnelsReadParams creates a new VpnTunnelsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVpnTunnelsReadParams() *VpnTunnelsReadParams {
	return &VpnTunnelsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVpnTunnelsReadParamsWithTimeout creates a new VpnTunnelsReadParams object
// with the ability to set a timeout on a request.
func NewVpnTunnelsReadParamsWithTimeout(timeout time.Duration) *VpnTunnelsReadParams {
	return &VpnTunnelsReadParams{
		timeout: timeout,
	}
}

// NewVpnTunnelsReadParamsWithContext creates a new VpnTunnelsReadParams object
// with the ability to set a context for a request.
func NewVpnTunnelsReadParamsWithContext(ctx context.Context) *VpnTunnelsReadParams {
	return &VpnTunnelsReadParams{
		Context: ctx,
	}
}

// NewVpnTunnelsReadParamsWithHTTPClient creates a new VpnTunnelsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewVpnTunnelsReadParamsWithHTTPClient(client *http.Client) *VpnTunnelsReadParams {
	return &VpnTunnelsReadParams{
		HTTPClient: client,
	}
}

/*
VpnTunnelsReadParams contains all the parameters to send to the API endpoint

	for the vpn tunnels read operation.

	Typically these are written to a http.Request.
*/
type VpnTunnelsReadParams struct {

	/* ID.

	   A unique integer value identifying this tunnel.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vpn tunnels read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsReadParams) WithDefaults() *VpnTunnelsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vpn tunnels read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vpn tunnels read params
func (o *VpnTunnelsReadParams) WithTimeout(timeout time.Duration) *VpnTunnelsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vpn tunnels read params
func (o *VpnTunnelsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vpn tunnels read params
func (o *VpnTunnelsReadParams) WithContext(ctx context.Context) *VpnTunnelsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vpn tunnels read params
func (o *VpnTunnelsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vpn tunnels read params
func (o *VpnTunnelsReadParams) WithHTTPClient(client *http.Client) *VpnTunnelsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vpn tunnels read params
func (o *VpnTunnelsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the vpn tunnels read params
func (o *VpnTunnelsReadParams) WithID(id int64) *VpnTunnelsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the vpn tunnels read params
func (o *VpnTunnelsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VpnTunnelsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
