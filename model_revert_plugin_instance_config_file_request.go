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

// RevertPluginInstanceConfigFileRequest struct for RevertPluginInstanceConfigFileRequest
type RevertPluginInstanceConfigFileRequest struct {
	// Version to revert to
	Version int32 `json:"version"`
}

// NewRevertPluginInstanceConfigFileRequest instantiates a new RevertPluginInstanceConfigFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevertPluginInstanceConfigFileRequest(version int32) *RevertPluginInstanceConfigFileRequest {
	this := RevertPluginInstanceConfigFileRequest{}
	this.Version = version
	return &this
}

// NewRevertPluginInstanceConfigFileRequestWithDefaults instantiates a new RevertPluginInstanceConfigFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevertPluginInstanceConfigFileRequestWithDefaults() *RevertPluginInstanceConfigFileRequest {
	this := RevertPluginInstanceConfigFileRequest{}
	return &this
}

// GetVersion returns the Version field value
func (o *RevertPluginInstanceConfigFileRequest) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RevertPluginInstanceConfigFileRequest) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RevertPluginInstanceConfigFileRequest) SetVersion(v int32) {
	o.Version = v
}

func (o RevertPluginInstanceConfigFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableRevertPluginInstanceConfigFileRequest struct {
	value *RevertPluginInstanceConfigFileRequest
	isSet bool
}

func (v NullableRevertPluginInstanceConfigFileRequest) Get() *RevertPluginInstanceConfigFileRequest {
	return v.value
}

func (v *NullableRevertPluginInstanceConfigFileRequest) Set(val *RevertPluginInstanceConfigFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRevertPluginInstanceConfigFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRevertPluginInstanceConfigFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevertPluginInstanceConfigFileRequest(val *RevertPluginInstanceConfigFileRequest) *NullableRevertPluginInstanceConfigFileRequest {
	return &NullableRevertPluginInstanceConfigFileRequest{value: val, isSet: true}
}

func (v NullableRevertPluginInstanceConfigFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevertPluginInstanceConfigFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


