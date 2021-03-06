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

// AdminUISettingsDetail struct for AdminUISettingsDetail
type AdminUISettingsDetail struct {
	Response *AdminUISettingsDetailResponse `json:"response,omitempty"`
}

// NewAdminUISettingsDetail instantiates a new AdminUISettingsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUISettingsDetail() *AdminUISettingsDetail {
	this := AdminUISettingsDetail{}
	return &this
}

// NewAdminUISettingsDetailWithDefaults instantiates a new AdminUISettingsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUISettingsDetailWithDefaults() *AdminUISettingsDetail {
	this := AdminUISettingsDetail{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AdminUISettingsDetail) GetResponse() AdminUISettingsDetailResponse {
	if o == nil || o.Response == nil {
		var ret AdminUISettingsDetailResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUISettingsDetail) GetResponseOk() (*AdminUISettingsDetailResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AdminUISettingsDetail) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given AdminUISettingsDetailResponse and assigns it to the Response field.
func (o *AdminUISettingsDetail) SetResponse(v AdminUISettingsDetailResponse) {
	o.Response = &v
}

func (o AdminUISettingsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAdminUISettingsDetail struct {
	value *AdminUISettingsDetail
	isSet bool
}

func (v NullableAdminUISettingsDetail) Get() *AdminUISettingsDetail {
	return v.value
}

func (v *NullableAdminUISettingsDetail) Set(val *AdminUISettingsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUISettingsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUISettingsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUISettingsDetail(val *AdminUISettingsDetail) *NullableAdminUISettingsDetail {
	return &NullableAdminUISettingsDetail{value: val, isSet: true}
}

func (v NullableAdminUISettingsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUISettingsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


