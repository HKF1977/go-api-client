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

// PasswordResetResponseResponse struct for PasswordResetResponseResponse
type PasswordResetResponseResponse struct {
	PasswordReset *string `json:"password_reset,omitempty"`
}

// NewPasswordResetResponseResponse instantiates a new PasswordResetResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordResetResponseResponse() *PasswordResetResponseResponse {
	this := PasswordResetResponseResponse{}
	return &this
}

// NewPasswordResetResponseResponseWithDefaults instantiates a new PasswordResetResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordResetResponseResponseWithDefaults() *PasswordResetResponseResponse {
	this := PasswordResetResponseResponse{}
	return &this
}

// GetPasswordReset returns the PasswordReset field value if set, zero value otherwise.
func (o *PasswordResetResponseResponse) GetPasswordReset() string {
	if o == nil || o.PasswordReset == nil {
		var ret string
		return ret
	}
	return *o.PasswordReset
}

// GetPasswordResetOk returns a tuple with the PasswordReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordResetResponseResponse) GetPasswordResetOk() (*string, bool) {
	if o == nil || o.PasswordReset == nil {
		return nil, false
	}
	return o.PasswordReset, true
}

// HasPasswordReset returns a boolean if a field has been set.
func (o *PasswordResetResponseResponse) HasPasswordReset() bool {
	if o != nil && o.PasswordReset != nil {
		return true
	}

	return false
}

// SetPasswordReset gets a reference to the given string and assigns it to the PasswordReset field.
func (o *PasswordResetResponseResponse) SetPasswordReset(v string) {
	o.PasswordReset = &v
}

func (o PasswordResetResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordReset != nil {
		toSerialize["password_reset"] = o.PasswordReset
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordResetResponseResponse struct {
	value *PasswordResetResponseResponse
	isSet bool
}

func (v NullablePasswordResetResponseResponse) Get() *PasswordResetResponseResponse {
	return v.value
}

func (v *NullablePasswordResetResponseResponse) Set(val *PasswordResetResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordResetResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordResetResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordResetResponseResponse(val *PasswordResetResponseResponse) *NullablePasswordResetResponseResponse {
	return &NullablePasswordResetResponseResponse{value: val, isSet: true}
}

func (v NullablePasswordResetResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordResetResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


