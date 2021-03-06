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

// PluginManagerConfigProcessManager struct for PluginManagerConfigProcessManager
type PluginManagerConfigProcessManager struct {
	// name of process manager such as supervisor. Currently  we support commands for supervisor and service. 
	Name *string `json:"name,omitempty"`
	// Name of programs, services or units managed
	Subprocesses []string `json:"subprocesses,omitempty"`
}

// NewPluginManagerConfigProcessManager instantiates a new PluginManagerConfigProcessManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerConfigProcessManager() *PluginManagerConfigProcessManager {
	this := PluginManagerConfigProcessManager{}
	return &this
}

// NewPluginManagerConfigProcessManagerWithDefaults instantiates a new PluginManagerConfigProcessManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerConfigProcessManagerWithDefaults() *PluginManagerConfigProcessManager {
	this := PluginManagerConfigProcessManager{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginManagerConfigProcessManager) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigProcessManager) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginManagerConfigProcessManager) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginManagerConfigProcessManager) SetName(v string) {
	o.Name = &v
}

// GetSubprocesses returns the Subprocesses field value if set, zero value otherwise.
func (o *PluginManagerConfigProcessManager) GetSubprocesses() []string {
	if o == nil || o.Subprocesses == nil {
		var ret []string
		return ret
	}
	return o.Subprocesses
}

// GetSubprocessesOk returns a tuple with the Subprocesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigProcessManager) GetSubprocessesOk() ([]string, bool) {
	if o == nil || o.Subprocesses == nil {
		return nil, false
	}
	return o.Subprocesses, true
}

// HasSubprocesses returns a boolean if a field has been set.
func (o *PluginManagerConfigProcessManager) HasSubprocesses() bool {
	if o != nil && o.Subprocesses != nil {
		return true
	}

	return false
}

// SetSubprocesses gets a reference to the given []string and assigns it to the Subprocesses field.
func (o *PluginManagerConfigProcessManager) SetSubprocesses(v []string) {
	o.Subprocesses = v
}

func (o PluginManagerConfigProcessManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subprocesses != nil {
		toSerialize["subprocesses"] = o.Subprocesses
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerConfigProcessManager struct {
	value *PluginManagerConfigProcessManager
	isSet bool
}

func (v NullablePluginManagerConfigProcessManager) Get() *PluginManagerConfigProcessManager {
	return v.value
}

func (v *NullablePluginManagerConfigProcessManager) Set(val *PluginManagerConfigProcessManager) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerConfigProcessManager) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerConfigProcessManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerConfigProcessManager(val *PluginManagerConfigProcessManager) *NullablePluginManagerConfigProcessManager {
	return &NullablePluginManagerConfigProcessManager{value: val, isSet: true}
}

func (v NullablePluginManagerConfigProcessManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerConfigProcessManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


