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

// TaskTokenResponse struct for TaskTokenResponse
type TaskTokenResponse struct {
	Response *TaskToken `json:"response,omitempty"`
}

// NewTaskTokenResponse instantiates a new TaskTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTokenResponse() *TaskTokenResponse {
	this := TaskTokenResponse{}
	return &this
}

// NewTaskTokenResponseWithDefaults instantiates a new TaskTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTokenResponseWithDefaults() *TaskTokenResponse {
	this := TaskTokenResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TaskTokenResponse) GetResponse() TaskToken {
	if o == nil || o.Response == nil {
		var ret TaskToken
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTokenResponse) GetResponseOk() (*TaskToken, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TaskTokenResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given TaskToken and assigns it to the Response field.
func (o *TaskTokenResponse) SetResponse(v TaskToken) {
	o.Response = &v
}

func (o TaskTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableTaskTokenResponse struct {
	value *TaskTokenResponse
	isSet bool
}

func (v NullableTaskTokenResponse) Get() *TaskTokenResponse {
	return v.value
}

func (v *NullableTaskTokenResponse) Set(val *TaskTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTokenResponse(val *TaskTokenResponse) *NullableTaskTokenResponse {
	return &NullableTaskTokenResponse{value: val, isSet: true}
}

func (v NullableTaskTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


