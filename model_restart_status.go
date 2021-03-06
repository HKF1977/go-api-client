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

// RestartStatus struct for RestartStatus
type RestartStatus struct {
	Response *RestartStatusResponse `json:"response,omitempty"`
}

// NewRestartStatus instantiates a new RestartStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartStatus() *RestartStatus {
	this := RestartStatus{}
	return &this
}

// NewRestartStatusWithDefaults instantiates a new RestartStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartStatusWithDefaults() *RestartStatus {
	this := RestartStatus{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RestartStatus) GetResponse() RestartStatusResponse {
	if o == nil || o.Response == nil {
		var ret RestartStatusResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartStatus) GetResponseOk() (*RestartStatusResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RestartStatus) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given RestartStatusResponse and assigns it to the Response field.
func (o *RestartStatus) SetResponse(v RestartStatusResponse) {
	o.Response = &v
}

func (o RestartStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableRestartStatus struct {
	value *RestartStatus
	isSet bool
}

func (v NullableRestartStatus) Get() *RestartStatus {
	return v.value
}

func (v *NullableRestartStatus) Set(val *RestartStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartStatus(val *RestartStatus) *NullableRestartStatus {
	return &NullableRestartStatus{value: val, isSet: true}
}

func (v NullableRestartStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


