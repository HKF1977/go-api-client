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

// RoutesListResponse struct for RoutesListResponse
type RoutesListResponse struct {
	Response *map[string]RoutesList `json:"response,omitempty"`
}

// NewRoutesListResponse instantiates a new RoutesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesListResponse() *RoutesListResponse {
	this := RoutesListResponse{}
	return &this
}

// NewRoutesListResponseWithDefaults instantiates a new RoutesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesListResponseWithDefaults() *RoutesListResponse {
	this := RoutesListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RoutesListResponse) GetResponse() map[string]RoutesList {
	if o == nil || o.Response == nil {
		var ret map[string]RoutesList
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesListResponse) GetResponseOk() (*map[string]RoutesList, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RoutesListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]RoutesList and assigns it to the Response field.
func (o *RoutesListResponse) SetResponse(v map[string]RoutesList) {
	o.Response = &v
}

func (o RoutesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesListResponse struct {
	value *RoutesListResponse
	isSet bool
}

func (v NullableRoutesListResponse) Get() *RoutesListResponse {
	return v.value
}

func (v *NullableRoutesListResponse) Set(val *RoutesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesListResponse(val *RoutesListResponse) *NullableRoutesListResponse {
	return &NullableRoutesListResponse{value: val, isSet: true}
}

func (v NullableRoutesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

