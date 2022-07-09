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

// SystemStatusDetail struct for SystemStatusDetail
type SystemStatusDetail struct {
	Response *SystemStatus `json:"response,omitempty"`
}

// NewSystemStatusDetail instantiates a new SystemStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatusDetail() *SystemStatusDetail {
	this := SystemStatusDetail{}
	return &this
}

// NewSystemStatusDetailWithDefaults instantiates a new SystemStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatusDetailWithDefaults() *SystemStatusDetail {
	this := SystemStatusDetail{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SystemStatusDetail) GetResponse() SystemStatus {
	if o == nil || o.Response == nil {
		var ret SystemStatus
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatusDetail) GetResponseOk() (*SystemStatus, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SystemStatusDetail) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given SystemStatus and assigns it to the Response field.
func (o *SystemStatusDetail) SetResponse(v SystemStatus) {
	o.Response = &v
}

func (o SystemStatusDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableSystemStatusDetail struct {
	value *SystemStatusDetail
	isSet bool
}

func (v NullableSystemStatusDetail) Get() *SystemStatusDetail {
	return v.value
}

func (v *NullableSystemStatusDetail) Set(val *SystemStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatusDetail(val *SystemStatusDetail) *NullableSystemStatusDetail {
	return &NullableSystemStatusDetail{value: val, isSet: true}
}

func (v NullableSystemStatusDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


