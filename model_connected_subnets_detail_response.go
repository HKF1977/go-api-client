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

// ConnectedSubnetsDetailResponse struct for ConnectedSubnetsDetailResponse
type ConnectedSubnetsDetailResponse struct {
	Response *ConnectedSubnetsDetail `json:"response,omitempty"`
}

// NewConnectedSubnetsDetailResponse instantiates a new ConnectedSubnetsDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedSubnetsDetailResponse() *ConnectedSubnetsDetailResponse {
	this := ConnectedSubnetsDetailResponse{}
	return &this
}

// NewConnectedSubnetsDetailResponseWithDefaults instantiates a new ConnectedSubnetsDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedSubnetsDetailResponseWithDefaults() *ConnectedSubnetsDetailResponse {
	this := ConnectedSubnetsDetailResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ConnectedSubnetsDetailResponse) GetResponse() ConnectedSubnetsDetail {
	if o == nil || o.Response == nil {
		var ret ConnectedSubnetsDetail
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnetsDetailResponse) GetResponseOk() (*ConnectedSubnetsDetail, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ConnectedSubnetsDetailResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ConnectedSubnetsDetail and assigns it to the Response field.
func (o *ConnectedSubnetsDetailResponse) SetResponse(v ConnectedSubnetsDetail) {
	o.Response = &v
}

func (o ConnectedSubnetsDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedSubnetsDetailResponse struct {
	value *ConnectedSubnetsDetailResponse
	isSet bool
}

func (v NullableConnectedSubnetsDetailResponse) Get() *ConnectedSubnetsDetailResponse {
	return v.value
}

func (v *NullableConnectedSubnetsDetailResponse) Set(val *ConnectedSubnetsDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedSubnetsDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedSubnetsDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedSubnetsDetailResponse(val *ConnectedSubnetsDetailResponse) *NullableConnectedSubnetsDetailResponse {
	return &NullableConnectedSubnetsDetailResponse{value: val, isSet: true}
}

func (v NullableConnectedSubnetsDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedSubnetsDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


