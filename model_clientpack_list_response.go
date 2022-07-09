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

// ClientpackListResponse struct for ClientpackListResponse
type ClientpackListResponse struct {
	Response *map[string]ClientpackList `json:"response,omitempty"`
}

// NewClientpackListResponse instantiates a new ClientpackListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackListResponse() *ClientpackListResponse {
	this := ClientpackListResponse{}
	return &this
}

// NewClientpackListResponseWithDefaults instantiates a new ClientpackListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackListResponseWithDefaults() *ClientpackListResponse {
	this := ClientpackListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ClientpackListResponse) GetResponse() map[string]ClientpackList {
	if o == nil || o.Response == nil {
		var ret map[string]ClientpackList
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackListResponse) GetResponseOk() (*map[string]ClientpackList, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ClientpackListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]ClientpackList and assigns it to the Response field.
func (o *ClientpackListResponse) SetResponse(v map[string]ClientpackList) {
	o.Response = &v
}

func (o ClientpackListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackListResponse struct {
	value *ClientpackListResponse
	isSet bool
}

func (v NullableClientpackListResponse) Get() *ClientpackListResponse {
	return v.value
}

func (v *NullableClientpackListResponse) Set(val *ClientpackListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackListResponse(val *ClientpackListResponse) *NullableClientpackListResponse {
	return &NullableClientpackListResponse{value: val, isSet: true}
}

func (v NullableClientpackListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


