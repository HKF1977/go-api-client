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

// AlertEventTypesListResponse struct for AlertEventTypesListResponse
type AlertEventTypesListResponse struct {
	Response []string `json:"response,omitempty"`
}

// NewAlertEventTypesListResponse instantiates a new AlertEventTypesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertEventTypesListResponse() *AlertEventTypesListResponse {
	this := AlertEventTypesListResponse{}
	return &this
}

// NewAlertEventTypesListResponseWithDefaults instantiates a new AlertEventTypesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertEventTypesListResponseWithDefaults() *AlertEventTypesListResponse {
	this := AlertEventTypesListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AlertEventTypesListResponse) GetResponse() []string {
	if o == nil || o.Response == nil {
		var ret []string
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventTypesListResponse) GetResponseOk() ([]string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AlertEventTypesListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []string and assigns it to the Response field.
func (o *AlertEventTypesListResponse) SetResponse(v []string) {
	o.Response = v
}

func (o AlertEventTypesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAlertEventTypesListResponse struct {
	value *AlertEventTypesListResponse
	isSet bool
}

func (v NullableAlertEventTypesListResponse) Get() *AlertEventTypesListResponse {
	return v.value
}

func (v *NullableAlertEventTypesListResponse) Set(val *AlertEventTypesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertEventTypesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertEventTypesListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertEventTypesListResponse(val *AlertEventTypesListResponse) *NullableAlertEventTypesListResponse {
	return &NullableAlertEventTypesListResponse{value: val, isSet: true}
}

func (v NullableAlertEventTypesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertEventTypesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


