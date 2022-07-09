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

// StopContainerDetailResponse struct for StopContainerDetailResponse
type StopContainerDetailResponse struct {
	Response *StopContainerDetail `json:"response,omitempty"`
}

// NewStopContainerDetailResponse instantiates a new StopContainerDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopContainerDetailResponse() *StopContainerDetailResponse {
	this := StopContainerDetailResponse{}
	return &this
}

// NewStopContainerDetailResponseWithDefaults instantiates a new StopContainerDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopContainerDetailResponseWithDefaults() *StopContainerDetailResponse {
	this := StopContainerDetailResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *StopContainerDetailResponse) GetResponse() StopContainerDetail {
	if o == nil || o.Response == nil {
		var ret StopContainerDetail
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopContainerDetailResponse) GetResponseOk() (*StopContainerDetail, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *StopContainerDetailResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given StopContainerDetail and assigns it to the Response field.
func (o *StopContainerDetailResponse) SetResponse(v StopContainerDetail) {
	o.Response = &v
}

func (o StopContainerDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableStopContainerDetailResponse struct {
	value *StopContainerDetailResponse
	isSet bool
}

func (v NullableStopContainerDetailResponse) Get() *StopContainerDetailResponse {
	return v.value
}

func (v *NullableStopContainerDetailResponse) Set(val *StopContainerDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStopContainerDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStopContainerDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopContainerDetailResponse(val *StopContainerDetailResponse) *NullableStopContainerDetailResponse {
	return &NullableStopContainerDetailResponse{value: val, isSet: true}
}

func (v NullableStopContainerDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopContainerDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


