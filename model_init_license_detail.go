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

// InitLicenseDetail struct for InitLicenseDetail
type InitLicenseDetail struct {
	Response *LicenseInitial `json:"response,omitempty"`
}

// NewInitLicenseDetail instantiates a new InitLicenseDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitLicenseDetail() *InitLicenseDetail {
	this := InitLicenseDetail{}
	return &this
}

// NewInitLicenseDetailWithDefaults instantiates a new InitLicenseDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitLicenseDetailWithDefaults() *InitLicenseDetail {
	this := InitLicenseDetail{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *InitLicenseDetail) GetResponse() LicenseInitial {
	if o == nil || o.Response == nil {
		var ret LicenseInitial
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitLicenseDetail) GetResponseOk() (*LicenseInitial, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *InitLicenseDetail) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given LicenseInitial and assigns it to the Response field.
func (o *InitLicenseDetail) SetResponse(v LicenseInitial) {
	o.Response = &v
}

func (o InitLicenseDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableInitLicenseDetail struct {
	value *InitLicenseDetail
	isSet bool
}

func (v NullableInitLicenseDetail) Get() *InitLicenseDetail {
	return v.value
}

func (v *NullableInitLicenseDetail) Set(val *InitLicenseDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableInitLicenseDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableInitLicenseDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitLicenseDetail(val *InitLicenseDetail) *NullableInitLicenseDetail {
	return &NullableInitLicenseDetail{value: val, isSet: true}
}

func (v NullableInitLicenseDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitLicenseDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

