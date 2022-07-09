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

// UpdatePluginInstanceConfigFileRequest struct for UpdatePluginInstanceConfigFileRequest
type UpdatePluginInstanceConfigFileRequest struct {
	// New config file contents
	Content string `json:"content"`
}

// NewUpdatePluginInstanceConfigFileRequest instantiates a new UpdatePluginInstanceConfigFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePluginInstanceConfigFileRequest(content string) *UpdatePluginInstanceConfigFileRequest {
	this := UpdatePluginInstanceConfigFileRequest{}
	this.Content = content
	return &this
}

// NewUpdatePluginInstanceConfigFileRequestWithDefaults instantiates a new UpdatePluginInstanceConfigFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePluginInstanceConfigFileRequestWithDefaults() *UpdatePluginInstanceConfigFileRequest {
	this := UpdatePluginInstanceConfigFileRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *UpdatePluginInstanceConfigFileRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *UpdatePluginInstanceConfigFileRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *UpdatePluginInstanceConfigFileRequest) SetContent(v string) {
	o.Content = v
}

func (o UpdatePluginInstanceConfigFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePluginInstanceConfigFileRequest struct {
	value *UpdatePluginInstanceConfigFileRequest
	isSet bool
}

func (v NullableUpdatePluginInstanceConfigFileRequest) Get() *UpdatePluginInstanceConfigFileRequest {
	return v.value
}

func (v *NullableUpdatePluginInstanceConfigFileRequest) Set(val *UpdatePluginInstanceConfigFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePluginInstanceConfigFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePluginInstanceConfigFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePluginInstanceConfigFileRequest(val *UpdatePluginInstanceConfigFileRequest) *NullableUpdatePluginInstanceConfigFileRequest {
	return &NullableUpdatePluginInstanceConfigFileRequest{value: val, isSet: true}
}

func (v NullableUpdatePluginInstanceConfigFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePluginInstanceConfigFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


