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

// PluginData struct for PluginData
type PluginData struct {
	Comment *string `json:"comment,omitempty"`
	ContainerConfig map[string]interface{} `json:"container_config,omitempty"`
}

// NewPluginData instantiates a new PluginData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginData() *PluginData {
	this := PluginData{}
	return &this
}

// NewPluginDataWithDefaults instantiates a new PluginData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginDataWithDefaults() *PluginData {
	this := PluginData{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PluginData) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginData) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PluginData) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *PluginData) SetComment(v string) {
	o.Comment = &v
}

// GetContainerConfig returns the ContainerConfig field value if set, zero value otherwise.
func (o *PluginData) GetContainerConfig() map[string]interface{} {
	if o == nil || o.ContainerConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ContainerConfig
}

// GetContainerConfigOk returns a tuple with the ContainerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginData) GetContainerConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.ContainerConfig == nil {
		return nil, false
	}
	return o.ContainerConfig, true
}

// HasContainerConfig returns a boolean if a field has been set.
func (o *PluginData) HasContainerConfig() bool {
	if o != nil && o.ContainerConfig != nil {
		return true
	}

	return false
}

// SetContainerConfig gets a reference to the given map[string]interface{} and assigns it to the ContainerConfig field.
func (o *PluginData) SetContainerConfig(v map[string]interface{}) {
	o.ContainerConfig = v
}

func (o PluginData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.ContainerConfig != nil {
		toSerialize["container_config"] = o.ContainerConfig
	}
	return json.Marshal(toSerialize)
}

type NullablePluginData struct {
	value *PluginData
	isSet bool
}

func (v NullablePluginData) Get() *PluginData {
	return v.value
}

func (v *NullablePluginData) Set(val *PluginData) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginData) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginData(val *PluginData) *NullablePluginData {
	return &NullablePluginData{value: val, isSet: true}
}

func (v NullablePluginData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


