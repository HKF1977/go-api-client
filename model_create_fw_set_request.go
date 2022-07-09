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

// CreateFWSetRequest struct for CreateFWSetRequest
type CreateFWSetRequest struct {
	// Chained firewall rules seperated by \"\\n\"
	Rules *string `json:"rules,omitempty"`
	// 'name of the FWSet. Must be valid chain that begins with one of the following: NETS_, PORTS_, LIST_.'  
	Name *string `json:"name,omitempty"`
	Flush *bool `json:"flush,omitempty"`
}

// NewCreateFWSetRequest instantiates a new CreateFWSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFWSetRequest() *CreateFWSetRequest {
	this := CreateFWSetRequest{}
	var flush bool = true
	this.Flush = &flush
	return &this
}

// NewCreateFWSetRequestWithDefaults instantiates a new CreateFWSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFWSetRequestWithDefaults() *CreateFWSetRequest {
	this := CreateFWSetRequest{}
	var flush bool = true
	this.Flush = &flush
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *CreateFWSetRequest) GetRules() string {
	if o == nil || o.Rules == nil {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFWSetRequest) GetRulesOk() (*string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *CreateFWSetRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
func (o *CreateFWSetRequest) SetRules(v string) {
	o.Rules = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateFWSetRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFWSetRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateFWSetRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateFWSetRequest) SetName(v string) {
	o.Name = &v
}

// GetFlush returns the Flush field value if set, zero value otherwise.
func (o *CreateFWSetRequest) GetFlush() bool {
	if o == nil || o.Flush == nil {
		var ret bool
		return ret
	}
	return *o.Flush
}

// GetFlushOk returns a tuple with the Flush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFWSetRequest) GetFlushOk() (*bool, bool) {
	if o == nil || o.Flush == nil {
		return nil, false
	}
	return o.Flush, true
}

// HasFlush returns a boolean if a field has been set.
func (o *CreateFWSetRequest) HasFlush() bool {
	if o != nil && o.Flush != nil {
		return true
	}

	return false
}

// SetFlush gets a reference to the given bool and assigns it to the Flush field.
func (o *CreateFWSetRequest) SetFlush(v bool) {
	o.Flush = &v
}

func (o CreateFWSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Flush != nil {
		toSerialize["flush"] = o.Flush
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFWSetRequest struct {
	value *CreateFWSetRequest
	isSet bool
}

func (v NullableCreateFWSetRequest) Get() *CreateFWSetRequest {
	return v.value
}

func (v *NullableCreateFWSetRequest) Set(val *CreateFWSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFWSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFWSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFWSetRequest(val *CreateFWSetRequest) *NullableCreateFWSetRequest {
	return &NullableCreateFWSetRequest{value: val, isSet: true}
}

func (v NullableCreateFWSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFWSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

