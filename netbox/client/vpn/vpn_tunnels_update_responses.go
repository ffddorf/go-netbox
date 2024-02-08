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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VpnTunnelsUpdateReader is a Reader for the VpnTunnelsUpdate structure.
type VpnTunnelsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVpnTunnelsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelsUpdateOK creates a VpnTunnelsUpdateOK with default headers values
func NewVpnTunnelsUpdateOK() *VpnTunnelsUpdateOK {
	return &VpnTunnelsUpdateOK{}
}

/*
VpnTunnelsUpdateOK describes a response with status code 200, with default header values.

VpnTunnelsUpdateOK vpn tunnels update o k
*/
type VpnTunnelsUpdateOK struct {
	Payload *models.Tunnel
}

// IsSuccess returns true when this vpn tunnels update o k response has a 2xx status code
func (o *VpnTunnelsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnels update o k response has a 3xx status code
func (o *VpnTunnelsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnels update o k response has a 4xx status code
func (o *VpnTunnelsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnels update o k response has a 5xx status code
func (o *VpnTunnelsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnels update o k response a status code equal to that given
func (o *VpnTunnelsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vpn tunnels update o k response
func (o *VpnTunnelsUpdateOK) Code() int {
	return 200
}

func (o *VpnTunnelsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpnTunnelsUpdateOK  %+v", 200, o.Payload)
}

func (o *VpnTunnelsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpnTunnelsUpdateOK  %+v", 200, o.Payload)
}

func (o *VpnTunnelsUpdateOK) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *VpnTunnelsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVpnTunnelsUpdateDefault creates a VpnTunnelsUpdateDefault with default headers values
func NewVpnTunnelsUpdateDefault(code int) *VpnTunnelsUpdateDefault {
	return &VpnTunnelsUpdateDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelsUpdateDefault describes a response with status code -1, with default header values.

VpnTunnelsUpdateDefault vpn tunnels update default
*/
type VpnTunnelsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnels update default response has a 2xx status code
func (o *VpnTunnelsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnels update default response has a 3xx status code
func (o *VpnTunnelsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnels update default response has a 4xx status code
func (o *VpnTunnelsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnels update default response has a 5xx status code
func (o *VpnTunnelsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnels update default response a status code equal to that given
func (o *VpnTunnelsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnels update default response
func (o *VpnTunnelsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpn_tunnels_update default  %+v", o._statusCode, o.Payload)
}

func (o *VpnTunnelsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpn_tunnels_update default  %+v", o._statusCode, o.Payload)
}

func (o *VpnTunnelsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
