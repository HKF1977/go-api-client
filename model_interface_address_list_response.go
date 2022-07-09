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

// InterfaceAddressListResponse struct for InterfaceAddressListResponse
type InterfaceAddressListResponse struct {
	Response []InterfaceAddress `json:"response,omitempty"`
}

// NewInterfaceAddressListResponse instantiates a new InterfaceAddressListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceAddressListResponse() *InterfaceAddressListResponse {
	this := InterfaceAddressListResponse{}
	return &this
}

// NewInterfaceAddressListResponseWithDefaults instantiates a new InterfaceAddressListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceAddressListResponseWithDefaults() *InterfaceAddressListResponse {
	this := InterfaceAddressListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *InterfaceAddressListResponse) GetResponse() []InterfaceAddress {
	if o == nil || o.Response == nil {
		var ret []InterfaceAddress
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAddressListResponse) GetResponseOk() ([]InterfaceAddress, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *InterfaceAddressListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []InterfaceAddress and assigns it to the Response field.
func (o *InterfaceAddressListResponse) SetResponse(v []InterfaceAddress) {
	o.Response = v
}

func (o InterfaceAddressListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceAddressListResponse struct {
	value *InterfaceAddressListResponse
	isSet bool
}

func (v NullableInterfaceAddressListResponse) Get() *InterfaceAddressListResponse {
	return v.value
}

func (v *NullableInterfaceAddressListResponse) Set(val *InterfaceAddressListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceAddressListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceAddressListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceAddressListResponse(val *InterfaceAddressListResponse) *NullableInterfaceAddressListResponse {
	return &NullableInterfaceAddressListResponse{value: val, isSet: true}
}

func (v NullableInterfaceAddressListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceAddressListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


