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

// WebhookCustomProperty struct for WebhookCustomProperty
type WebhookCustomProperty struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewWebhookCustomProperty instantiates a new WebhookCustomProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookCustomProperty() *WebhookCustomProperty {
	this := WebhookCustomProperty{}
	return &this
}

// NewWebhookCustomPropertyWithDefaults instantiates a new WebhookCustomProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookCustomPropertyWithDefaults() *WebhookCustomProperty {
	this := WebhookCustomProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookCustomProperty) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookCustomProperty) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookCustomProperty) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookCustomProperty) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WebhookCustomProperty) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookCustomProperty) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WebhookCustomProperty) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WebhookCustomProperty) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WebhookCustomProperty) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookCustomProperty) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WebhookCustomProperty) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WebhookCustomProperty) SetDescription(v string) {
	o.Description = &v
}

func (o WebhookCustomProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookCustomProperty struct {
	value *WebhookCustomProperty
	isSet bool
}

func (v NullableWebhookCustomProperty) Get() *WebhookCustomProperty {
	return v.value
}

func (v *NullableWebhookCustomProperty) Set(val *WebhookCustomProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookCustomProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookCustomProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookCustomProperty(val *WebhookCustomProperty) *NullableWebhookCustomProperty {
	return &NullableWebhookCustomProperty{value: val, isSet: true}
}

func (v NullableWebhookCustomProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookCustomProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


