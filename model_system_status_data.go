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

// SystemStatusData struct for SystemStatusData
type SystemStatusData struct {
	Sysstat [][]string `json:"sysstat,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemStatusData SystemStatusData

// NewSystemStatusData instantiates a new SystemStatusData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatusData() *SystemStatusData {
	this := SystemStatusData{}
	return &this
}

// NewSystemStatusDataWithDefaults instantiates a new SystemStatusData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatusDataWithDefaults() *SystemStatusData {
	this := SystemStatusData{}
	return &this
}

// GetSysstat returns the Sysstat field value if set, zero value otherwise.
func (o *SystemStatusData) GetSysstat() [][]string {
	if o == nil || o.Sysstat == nil {
		var ret [][]string
		return ret
	}
	return o.Sysstat
}

// GetSysstatOk returns a tuple with the Sysstat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatusData) GetSysstatOk() ([][]string, bool) {
	if o == nil || o.Sysstat == nil {
		return nil, false
	}
	return o.Sysstat, true
}

// HasSysstat returns a boolean if a field has been set.
func (o *SystemStatusData) HasSysstat() bool {
	if o != nil && o.Sysstat != nil {
		return true
	}

	return false
}

// SetSysstat gets a reference to the given [][]string and assigns it to the Sysstat field.
func (o *SystemStatusData) SetSysstat(v [][]string) {
	o.Sysstat = v
}

func (o SystemStatusData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sysstat != nil {
		toSerialize["sysstat"] = o.Sysstat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SystemStatusData) UnmarshalJSON(bytes []byte) (err error) {
	varSystemStatusData := _SystemStatusData{}

	if err = json.Unmarshal(bytes, &varSystemStatusData); err == nil {
		*o = SystemStatusData(varSystemStatusData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sysstat")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemStatusData struct {
	value *SystemStatusData
	isSet bool
}

func (v NullableSystemStatusData) Get() *SystemStatusData {
	return v.value
}

func (v *NullableSystemStatusData) Set(val *SystemStatusData) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatusData) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatusData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatusData(val *SystemStatusData) *NullableSystemStatusData {
	return &NullableSystemStatusData{value: val, isSet: true}
}

func (v NullableSystemStatusData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatusData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

