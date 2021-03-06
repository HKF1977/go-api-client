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

// PluginManagerConfigConfigFilesResponse struct for PluginManagerConfigConfigFilesResponse
type PluginManagerConfigConfigFilesResponse struct {
	Response []PluginManagerConfigConfigFile `json:"response,omitempty"`
}

// NewPluginManagerConfigConfigFilesResponse instantiates a new PluginManagerConfigConfigFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerConfigConfigFilesResponse() *PluginManagerConfigConfigFilesResponse {
	this := PluginManagerConfigConfigFilesResponse{}
	return &this
}

// NewPluginManagerConfigConfigFilesResponseWithDefaults instantiates a new PluginManagerConfigConfigFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerConfigConfigFilesResponseWithDefaults() *PluginManagerConfigConfigFilesResponse {
	this := PluginManagerConfigConfigFilesResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFilesResponse) GetResponse() []PluginManagerConfigConfigFile {
	if o == nil || o.Response == nil {
		var ret []PluginManagerConfigConfigFile
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFilesResponse) GetResponseOk() ([]PluginManagerConfigConfigFile, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFilesResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []PluginManagerConfigConfigFile and assigns it to the Response field.
func (o *PluginManagerConfigConfigFilesResponse) SetResponse(v []PluginManagerConfigConfigFile) {
	o.Response = v
}

func (o PluginManagerConfigConfigFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerConfigConfigFilesResponse struct {
	value *PluginManagerConfigConfigFilesResponse
	isSet bool
}

func (v NullablePluginManagerConfigConfigFilesResponse) Get() *PluginManagerConfigConfigFilesResponse {
	return v.value
}

func (v *NullablePluginManagerConfigConfigFilesResponse) Set(val *PluginManagerConfigConfigFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerConfigConfigFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerConfigConfigFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerConfigConfigFilesResponse(val *PluginManagerConfigConfigFilesResponse) *NullablePluginManagerConfigConfigFilesResponse {
	return &NullablePluginManagerConfigConfigFilesResponse{value: val, isSet: true}
}

func (v NullablePluginManagerConfigConfigFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerConfigConfigFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


