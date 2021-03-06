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

// UpdateKeysetRequest struct for UpdateKeysetRequest
type UpdateKeysetRequest struct {
	// If provided, fetches keyset from source manager
	Source *string `json:"source,omitempty"`
	// Arbitrary key used to help randomize the checksum, it must be identical for each manager in a topology.
	Token string `json:"token"`
	// Name for the topology
	TopologyName *string `json:"topology_name,omitempty"`
}

// NewUpdateKeysetRequest instantiates a new UpdateKeysetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeysetRequest(token string) *UpdateKeysetRequest {
	this := UpdateKeysetRequest{}
	this.Token = token
	return &this
}

// NewUpdateKeysetRequestWithDefaults instantiates a new UpdateKeysetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeysetRequestWithDefaults() *UpdateKeysetRequest {
	this := UpdateKeysetRequest{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *UpdateKeysetRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeysetRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *UpdateKeysetRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *UpdateKeysetRequest) SetSource(v string) {
	o.Source = &v
}

// GetToken returns the Token field value
func (o *UpdateKeysetRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateKeysetRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateKeysetRequest) SetToken(v string) {
	o.Token = v
}

// GetTopologyName returns the TopologyName field value if set, zero value otherwise.
func (o *UpdateKeysetRequest) GetTopologyName() string {
	if o == nil || o.TopologyName == nil {
		var ret string
		return ret
	}
	return *o.TopologyName
}

// GetTopologyNameOk returns a tuple with the TopologyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeysetRequest) GetTopologyNameOk() (*string, bool) {
	if o == nil || o.TopologyName == nil {
		return nil, false
	}
	return o.TopologyName, true
}

// HasTopologyName returns a boolean if a field has been set.
func (o *UpdateKeysetRequest) HasTopologyName() bool {
	if o != nil && o.TopologyName != nil {
		return true
	}

	return false
}

// SetTopologyName gets a reference to the given string and assigns it to the TopologyName field.
func (o *UpdateKeysetRequest) SetTopologyName(v string) {
	o.TopologyName = &v
}

func (o UpdateKeysetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if o.TopologyName != nil {
		toSerialize["topology_name"] = o.TopologyName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateKeysetRequest struct {
	value *UpdateKeysetRequest
	isSet bool
}

func (v NullableUpdateKeysetRequest) Get() *UpdateKeysetRequest {
	return v.value
}

func (v *NullableUpdateKeysetRequest) Set(val *UpdateKeysetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeysetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeysetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeysetRequest(val *UpdateKeysetRequest) *NullableUpdateKeysetRequest {
	return &NullableUpdateKeysetRequest{value: val, isSet: true}
}

func (v NullableUpdateKeysetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeysetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


