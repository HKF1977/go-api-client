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

// PluginManagerConfigResponse struct for PluginManagerConfigResponse
type PluginManagerConfigResponse struct {
	Response *PluginManagerConfig `json:"response,omitempty"`
}

// NewPluginManagerConfigResponse instantiates a new PluginManagerConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerConfigResponse() *PluginManagerConfigResponse {
	this := PluginManagerConfigResponse{}
	return &this
}

// NewPluginManagerConfigResponseWithDefaults instantiates a new PluginManagerConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerConfigResponseWithDefaults() *PluginManagerConfigResponse {
	this := PluginManagerConfigResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PluginManagerConfigResponse) GetResponse() PluginManagerConfig {
	if o == nil || o.Response == nil {
		var ret PluginManagerConfig
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigResponse) GetResponseOk() (*PluginManagerConfig, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PluginManagerConfigResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given PluginManagerConfig and assigns it to the Response field.
func (o *PluginManagerConfigResponse) SetResponse(v PluginManagerConfig) {
	o.Response = &v
}

func (o PluginManagerConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerConfigResponse struct {
	value *PluginManagerConfigResponse
	isSet bool
}

func (v NullablePluginManagerConfigResponse) Get() *PluginManagerConfigResponse {
	return v.value
}

func (v *NullablePluginManagerConfigResponse) Set(val *PluginManagerConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerConfigResponse(val *PluginManagerConfigResponse) *NullablePluginManagerConfigResponse {
	return &NullablePluginManagerConfigResponse{value: val, isSet: true}
}

func (v NullablePluginManagerConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


