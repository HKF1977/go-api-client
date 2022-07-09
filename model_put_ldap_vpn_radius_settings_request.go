/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 5.1.5
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PutLdapVpnRadiusSettingsRequest struct for PutLdapVpnRadiusSettingsRequest
type PutLdapVpnRadiusSettingsRequest struct {
	// Enable use of Radius through VPN. If true, other params required.
	VpnAuthEnabled bool `json:"vpn_auth_enabled"`
	// IP address or resolvable hostname
	VpnRadiusServer string `json:"vpn_radius_server"`
	// Authentication port
	VpnRadiusAuthPort *int32 `json:"vpn_radius_auth_port,omitempty"`
	// Accounting port
	VpnRadiusAccountingPort *int32 `json:"vpn_radius_accounting_port,omitempty"`
	// Shared password
	VpnRadiusPass string `json:"vpn_radius_pass"`
}

// NewPutLdapVpnRadiusSettingsRequest instantiates a new PutLdapVpnRadiusSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutLdapVpnRadiusSettingsRequest(vpnAuthEnabled bool, vpnRadiusServer string, vpnRadiusPass string) *PutLdapVpnRadiusSettingsRequest {
	this := PutLdapVpnRadiusSettingsRequest{}
	this.VpnAuthEnabled = vpnAuthEnabled
	this.VpnRadiusServer = vpnRadiusServer
	var vpnRadiusAuthPort int32 = 1812
	this.VpnRadiusAuthPort = &vpnRadiusAuthPort
	var vpnRadiusAccountingPort int32 = 1812
	this.VpnRadiusAccountingPort = &vpnRadiusAccountingPort
	this.VpnRadiusPass = vpnRadiusPass
	return &this
}

// NewPutLdapVpnRadiusSettingsRequestWithDefaults instantiates a new PutLdapVpnRadiusSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutLdapVpnRadiusSettingsRequestWithDefaults() *PutLdapVpnRadiusSettingsRequest {
	this := PutLdapVpnRadiusSettingsRequest{}
	var vpnRadiusAuthPort int32 = 1812
	this.VpnRadiusAuthPort = &vpnRadiusAuthPort
	var vpnRadiusAccountingPort int32 = 1812
	this.VpnRadiusAccountingPort = &vpnRadiusAccountingPort
	return &this
}

// GetVpnAuthEnabled returns the VpnAuthEnabled field value
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnAuthEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VpnAuthEnabled
}

// GetVpnAuthEnabledOk returns a tuple with the VpnAuthEnabled field value
// and a boolean to check if the value has been set.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnAuthEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnAuthEnabled, true
}

// SetVpnAuthEnabled sets field value
func (o *PutLdapVpnRadiusSettingsRequest) SetVpnAuthEnabled(v bool) {
	o.VpnAuthEnabled = v
}

// GetVpnRadiusServer returns the VpnRadiusServer field value
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnRadiusServer
}

// GetVpnRadiusServerOk returns a tuple with the VpnRadiusServer field value
// and a boolean to check if the value has been set.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnRadiusServer, true
}

// SetVpnRadiusServer sets field value
func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusServer(v string) {
	o.VpnRadiusServer = v
}

// GetVpnRadiusAuthPort returns the VpnRadiusAuthPort field value if set, zero value otherwise.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAuthPort() int32 {
	if o == nil || o.VpnRadiusAuthPort == nil {
		var ret int32
		return ret
	}
	return *o.VpnRadiusAuthPort
}

// GetVpnRadiusAuthPortOk returns a tuple with the VpnRadiusAuthPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAuthPortOk() (*int32, bool) {
	if o == nil || o.VpnRadiusAuthPort == nil {
		return nil, false
	}
	return o.VpnRadiusAuthPort, true
}

// HasVpnRadiusAuthPort returns a boolean if a field has been set.
func (o *PutLdapVpnRadiusSettingsRequest) HasVpnRadiusAuthPort() bool {
	if o != nil && o.VpnRadiusAuthPort != nil {
		return true
	}

	return false
}

// SetVpnRadiusAuthPort gets a reference to the given int32 and assigns it to the VpnRadiusAuthPort field.
func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusAuthPort(v int32) {
	o.VpnRadiusAuthPort = &v
}

// GetVpnRadiusAccountingPort returns the VpnRadiusAccountingPort field value if set, zero value otherwise.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAccountingPort() int32 {
	if o == nil || o.VpnRadiusAccountingPort == nil {
		var ret int32
		return ret
	}
	return *o.VpnRadiusAccountingPort
}

// GetVpnRadiusAccountingPortOk returns a tuple with the VpnRadiusAccountingPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAccountingPortOk() (*int32, bool) {
	if o == nil || o.VpnRadiusAccountingPort == nil {
		return nil, false
	}
	return o.VpnRadiusAccountingPort, true
}

// HasVpnRadiusAccountingPort returns a boolean if a field has been set.
func (o *PutLdapVpnRadiusSettingsRequest) HasVpnRadiusAccountingPort() bool {
	if o != nil && o.VpnRadiusAccountingPort != nil {
		return true
	}

	return false
}

// SetVpnRadiusAccountingPort gets a reference to the given int32 and assigns it to the VpnRadiusAccountingPort field.
func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusAccountingPort(v int32) {
	o.VpnRadiusAccountingPort = &v
}

// GetVpnRadiusPass returns the VpnRadiusPass field value
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusPass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnRadiusPass
}

// GetVpnRadiusPassOk returns a tuple with the VpnRadiusPass field value
// and a boolean to check if the value has been set.
func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnRadiusPass, true
}

// SetVpnRadiusPass sets field value
func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusPass(v string) {
	o.VpnRadiusPass = v
}

func (o PutLdapVpnRadiusSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vpn_auth_enabled"] = o.VpnAuthEnabled
	}
	if true {
		toSerialize["vpn_radius_server"] = o.VpnRadiusServer
	}
	if o.VpnRadiusAuthPort != nil {
		toSerialize["vpn_radius_auth_port"] = o.VpnRadiusAuthPort
	}
	if o.VpnRadiusAccountingPort != nil {
		toSerialize["vpn_radius_accounting_port"] = o.VpnRadiusAccountingPort
	}
	if true {
		toSerialize["vpn_radius_pass"] = o.VpnRadiusPass
	}
	return json.Marshal(toSerialize)
}

type NullablePutLdapVpnRadiusSettingsRequest struct {
	value *PutLdapVpnRadiusSettingsRequest
	isSet bool
}

func (v NullablePutLdapVpnRadiusSettingsRequest) Get() *PutLdapVpnRadiusSettingsRequest {
	return v.value
}

func (v *NullablePutLdapVpnRadiusSettingsRequest) Set(val *PutLdapVpnRadiusSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutLdapVpnRadiusSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutLdapVpnRadiusSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutLdapVpnRadiusSettingsRequest(val *PutLdapVpnRadiusSettingsRequest) *NullablePutLdapVpnRadiusSettingsRequest {
	return &NullablePutLdapVpnRadiusSettingsRequest{value: val, isSet: true}
}

func (v NullablePutLdapVpnRadiusSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutLdapVpnRadiusSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

