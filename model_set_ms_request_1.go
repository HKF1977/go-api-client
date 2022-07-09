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

// SetMSRequest1 struct for SetMSRequest1
type SetMSRequest1 struct {
	// VNS3 Management system endpoint IP address
	Ip string `json:"ip"`
}

// NewSetMSRequest1 instantiates a new SetMSRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMSRequest1(ip string) *SetMSRequest1 {
	this := SetMSRequest1{}
	this.Ip = ip
	return &this
}

// NewSetMSRequest1WithDefaults instantiates a new SetMSRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMSRequest1WithDefaults() *SetMSRequest1 {
	this := SetMSRequest1{}
	return &this
}

// GetIp returns the Ip field value
func (o *SetMSRequest1) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *SetMSRequest1) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *SetMSRequest1) SetIp(v string) {
	o.Ip = v
}

func (o SetMSRequest1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableSetMSRequest1 struct {
	value *SetMSRequest1
	isSet bool
}

func (v NullableSetMSRequest1) Get() *SetMSRequest1 {
	return v.value
}

func (v *NullableSetMSRequest1) Set(val *SetMSRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMSRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMSRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMSRequest1(val *SetMSRequest1) *NullableSetMSRequest1 {
	return &NullableSetMSRequest1{value: val, isSet: true}
}

func (v NullableSetMSRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMSRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


