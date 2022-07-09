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

// RunPluginInstanceExeCommandRequest struct for RunPluginInstanceExeCommandRequest
type RunPluginInstanceExeCommandRequest struct {
	// The command to run. (A key in the commands object)
	Command string `json:"command"`
	// Path to the executable. This is required if more  than 1 executable is defined. 
	ExecutablePath *string `json:"executable_path,omitempty"`
	// Number of seconds to wait before timing out.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewRunPluginInstanceExeCommandRequest instantiates a new RunPluginInstanceExeCommandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunPluginInstanceExeCommandRequest(command string) *RunPluginInstanceExeCommandRequest {
	this := RunPluginInstanceExeCommandRequest{}
	this.Command = command
	var timeout int32 = 30
	this.Timeout = &timeout
	return &this
}

// NewRunPluginInstanceExeCommandRequestWithDefaults instantiates a new RunPluginInstanceExeCommandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunPluginInstanceExeCommandRequestWithDefaults() *RunPluginInstanceExeCommandRequest {
	this := RunPluginInstanceExeCommandRequest{}
	var timeout int32 = 30
	this.Timeout = &timeout
	return &this
}

// GetCommand returns the Command field value
func (o *RunPluginInstanceExeCommandRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *RunPluginInstanceExeCommandRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *RunPluginInstanceExeCommandRequest) SetCommand(v string) {
	o.Command = v
}

// GetExecutablePath returns the ExecutablePath field value if set, zero value otherwise.
func (o *RunPluginInstanceExeCommandRequest) GetExecutablePath() string {
	if o == nil || o.ExecutablePath == nil {
		var ret string
		return ret
	}
	return *o.ExecutablePath
}

// GetExecutablePathOk returns a tuple with the ExecutablePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunPluginInstanceExeCommandRequest) GetExecutablePathOk() (*string, bool) {
	if o == nil || o.ExecutablePath == nil {
		return nil, false
	}
	return o.ExecutablePath, true
}

// HasExecutablePath returns a boolean if a field has been set.
func (o *RunPluginInstanceExeCommandRequest) HasExecutablePath() bool {
	if o != nil && o.ExecutablePath != nil {
		return true
	}

	return false
}

// SetExecutablePath gets a reference to the given string and assigns it to the ExecutablePath field.
func (o *RunPluginInstanceExeCommandRequest) SetExecutablePath(v string) {
	o.ExecutablePath = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *RunPluginInstanceExeCommandRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunPluginInstanceExeCommandRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *RunPluginInstanceExeCommandRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *RunPluginInstanceExeCommandRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o RunPluginInstanceExeCommandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["command"] = o.Command
	}
	if o.ExecutablePath != nil {
		toSerialize["executable_path"] = o.ExecutablePath
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableRunPluginInstanceExeCommandRequest struct {
	value *RunPluginInstanceExeCommandRequest
	isSet bool
}

func (v NullableRunPluginInstanceExeCommandRequest) Get() *RunPluginInstanceExeCommandRequest {
	return v.value
}

func (v *NullableRunPluginInstanceExeCommandRequest) Set(val *RunPluginInstanceExeCommandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunPluginInstanceExeCommandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunPluginInstanceExeCommandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunPluginInstanceExeCommandRequest(val *RunPluginInstanceExeCommandRequest) *NullableRunPluginInstanceExeCommandRequest {
	return &NullableRunPluginInstanceExeCommandRequest{value: val, isSet: true}
}

func (v NullableRunPluginInstanceExeCommandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunPluginInstanceExeCommandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

