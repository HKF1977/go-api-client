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

// UpdateConfigClientpackRequest struct for UpdateConfigClientpackRequest
type UpdateConfigClientpackRequest struct {
	// Enable or disable clientpacks.
	Enabled *bool `json:"enabled,omitempty"`
	// Update whether clientpacks are checked out and thus unavailable
	CheckedOut *bool `json:"checked_out,omitempty"`
	// Turn compression on or off for all. Can be \"on\" or \"off\" currently.
	Compression *string `json:"compression,omitempty"`
}

// NewUpdateConfigClientpackRequest instantiates a new UpdateConfigClientpackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConfigClientpackRequest() *UpdateConfigClientpackRequest {
	this := UpdateConfigClientpackRequest{}
	return &this
}

// NewUpdateConfigClientpackRequestWithDefaults instantiates a new UpdateConfigClientpackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConfigClientpackRequestWithDefaults() *UpdateConfigClientpackRequest {
	this := UpdateConfigClientpackRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateConfigClientpackRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigClientpackRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateConfigClientpackRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateConfigClientpackRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCheckedOut returns the CheckedOut field value if set, zero value otherwise.
func (o *UpdateConfigClientpackRequest) GetCheckedOut() bool {
	if o == nil || o.CheckedOut == nil {
		var ret bool
		return ret
	}
	return *o.CheckedOut
}

// GetCheckedOutOk returns a tuple with the CheckedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigClientpackRequest) GetCheckedOutOk() (*bool, bool) {
	if o == nil || o.CheckedOut == nil {
		return nil, false
	}
	return o.CheckedOut, true
}

// HasCheckedOut returns a boolean if a field has been set.
func (o *UpdateConfigClientpackRequest) HasCheckedOut() bool {
	if o != nil && o.CheckedOut != nil {
		return true
	}

	return false
}

// SetCheckedOut gets a reference to the given bool and assigns it to the CheckedOut field.
func (o *UpdateConfigClientpackRequest) SetCheckedOut(v bool) {
	o.CheckedOut = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *UpdateConfigClientpackRequest) GetCompression() string {
	if o == nil || o.Compression == nil {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigClientpackRequest) GetCompressionOk() (*string, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *UpdateConfigClientpackRequest) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *UpdateConfigClientpackRequest) SetCompression(v string) {
	o.Compression = &v
}

func (o UpdateConfigClientpackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CheckedOut != nil {
		toSerialize["checked_out"] = o.CheckedOut
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConfigClientpackRequest struct {
	value *UpdateConfigClientpackRequest
	isSet bool
}

func (v NullableUpdateConfigClientpackRequest) Get() *UpdateConfigClientpackRequest {
	return v.value
}

func (v *NullableUpdateConfigClientpackRequest) Set(val *UpdateConfigClientpackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConfigClientpackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConfigClientpackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConfigClientpackRequest(val *UpdateConfigClientpackRequest) *NullableUpdateConfigClientpackRequest {
	return &NullableUpdateConfigClientpackRequest{value: val, isSet: true}
}

func (v NullableUpdateConfigClientpackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConfigClientpackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


