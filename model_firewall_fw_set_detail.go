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

// FirewallFWSetDetail struct for FirewallFWSetDetail
type FirewallFWSetDetail struct {
	Response *FirewallFWSetData `json:"response,omitempty"`
}

// NewFirewallFWSetDetail instantiates a new FirewallFWSetDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallFWSetDetail() *FirewallFWSetDetail {
	this := FirewallFWSetDetail{}
	return &this
}

// NewFirewallFWSetDetailWithDefaults instantiates a new FirewallFWSetDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallFWSetDetailWithDefaults() *FirewallFWSetDetail {
	this := FirewallFWSetDetail{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *FirewallFWSetDetail) GetResponse() FirewallFWSetData {
	if o == nil || o.Response == nil {
		var ret FirewallFWSetData
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFWSetDetail) GetResponseOk() (*FirewallFWSetData, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *FirewallFWSetDetail) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given FirewallFWSetData and assigns it to the Response field.
func (o *FirewallFWSetDetail) SetResponse(v FirewallFWSetData) {
	o.Response = &v
}

func (o FirewallFWSetDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallFWSetDetail struct {
	value *FirewallFWSetDetail
	isSet bool
}

func (v NullableFirewallFWSetDetail) Get() *FirewallFWSetDetail {
	return v.value
}

func (v *NullableFirewallFWSetDetail) Set(val *FirewallFWSetDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallFWSetDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallFWSetDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallFWSetDetail(val *FirewallFWSetDetail) *NullableFirewallFWSetDetail {
	return &NullableFirewallFWSetDetail{value: val, isSet: true}
}

func (v NullableFirewallFWSetDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallFWSetDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

