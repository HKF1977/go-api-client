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

// AddClientpackRequest struct for AddClientpackRequest
type AddClientpackRequest struct {
	// CSV of IP addresses to be used for new clientpacks
	RequestedIps string `json:"requested_ips"`
}

// NewAddClientpackRequest instantiates a new AddClientpackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClientpackRequest(requestedIps string) *AddClientpackRequest {
	this := AddClientpackRequest{}
	this.RequestedIps = requestedIps
	return &this
}

// NewAddClientpackRequestWithDefaults instantiates a new AddClientpackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClientpackRequestWithDefaults() *AddClientpackRequest {
	this := AddClientpackRequest{}
	return &this
}

// GetRequestedIps returns the RequestedIps field value
func (o *AddClientpackRequest) GetRequestedIps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestedIps
}

// GetRequestedIpsOk returns a tuple with the RequestedIps field value
// and a boolean to check if the value has been set.
func (o *AddClientpackRequest) GetRequestedIpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedIps, true
}

// SetRequestedIps sets field value
func (o *AddClientpackRequest) SetRequestedIps(v string) {
	o.RequestedIps = v
}

func (o AddClientpackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requested_ips"] = o.RequestedIps
	}
	return json.Marshal(toSerialize)
}

type NullableAddClientpackRequest struct {
	value *AddClientpackRequest
	isSet bool
}

func (v NullableAddClientpackRequest) Get() *AddClientpackRequest {
	return v.value
}

func (v *NullableAddClientpackRequest) Set(val *AddClientpackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddClientpackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddClientpackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddClientpackRequest(val *AddClientpackRequest) *NullableAddClientpackRequest {
	return &NullableAddClientpackRequest{value: val, isSet: true}
}

func (v NullableAddClientpackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddClientpackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


