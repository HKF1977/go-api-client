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

// UpdateServerSSLRequest struct for UpdateServerSSLRequest
type UpdateServerSSLRequest struct {
	Cert string `json:"cert"`
	Key string `json:"key"`
}

// NewUpdateServerSSLRequest instantiates a new UpdateServerSSLRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerSSLRequest(cert string, key string) *UpdateServerSSLRequest {
	this := UpdateServerSSLRequest{}
	this.Cert = cert
	this.Key = key
	return &this
}

// NewUpdateServerSSLRequestWithDefaults instantiates a new UpdateServerSSLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerSSLRequestWithDefaults() *UpdateServerSSLRequest {
	this := UpdateServerSSLRequest{}
	return &this
}

// GetCert returns the Cert field value
func (o *UpdateServerSSLRequest) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *UpdateServerSSLRequest) GetCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *UpdateServerSSLRequest) SetCert(v string) {
	o.Cert = v
}

// GetKey returns the Key field value
func (o *UpdateServerSSLRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UpdateServerSSLRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UpdateServerSSLRequest) SetKey(v string) {
	o.Key = v
}

func (o UpdateServerSSLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert"] = o.Cert
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateServerSSLRequest struct {
	value *UpdateServerSSLRequest
	isSet bool
}

func (v NullableUpdateServerSSLRequest) Get() *UpdateServerSSLRequest {
	return v.value
}

func (v *NullableUpdateServerSSLRequest) Set(val *UpdateServerSSLRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerSSLRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerSSLRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerSSLRequest(val *UpdateServerSSLRequest) *NullableUpdateServerSSLRequest {
	return &NullableUpdateServerSSLRequest{value: val, isSet: true}
}

func (v NullableUpdateServerSSLRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerSSLRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


