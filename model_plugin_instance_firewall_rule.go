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

// PluginInstanceFirewallRule struct for PluginInstanceFirewallRule
type PluginInstanceFirewallRule struct {
	Rule *string `json:"rule,omitempty"`
	Tags []string `json:"tags,omitempty"`
	PluginPort *int32 `json:"plugin_port,omitempty"`
	DestinationPort *int32 `json:"destination_port,omitempty"`
}

// NewPluginInstanceFirewallRule instantiates a new PluginInstanceFirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginInstanceFirewallRule() *PluginInstanceFirewallRule {
	this := PluginInstanceFirewallRule{}
	return &this
}

// NewPluginInstanceFirewallRuleWithDefaults instantiates a new PluginInstanceFirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginInstanceFirewallRuleWithDefaults() *PluginInstanceFirewallRule {
	this := PluginInstanceFirewallRule{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *PluginInstanceFirewallRule) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstanceFirewallRule) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *PluginInstanceFirewallRule) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *PluginInstanceFirewallRule) SetRule(v string) {
	o.Rule = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PluginInstanceFirewallRule) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstanceFirewallRule) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PluginInstanceFirewallRule) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PluginInstanceFirewallRule) SetTags(v []string) {
	o.Tags = v
}

// GetPluginPort returns the PluginPort field value if set, zero value otherwise.
func (o *PluginInstanceFirewallRule) GetPluginPort() int32 {
	if o == nil || o.PluginPort == nil {
		var ret int32
		return ret
	}
	return *o.PluginPort
}

// GetPluginPortOk returns a tuple with the PluginPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstanceFirewallRule) GetPluginPortOk() (*int32, bool) {
	if o == nil || o.PluginPort == nil {
		return nil, false
	}
	return o.PluginPort, true
}

// HasPluginPort returns a boolean if a field has been set.
func (o *PluginInstanceFirewallRule) HasPluginPort() bool {
	if o != nil && o.PluginPort != nil {
		return true
	}

	return false
}

// SetPluginPort gets a reference to the given int32 and assigns it to the PluginPort field.
func (o *PluginInstanceFirewallRule) SetPluginPort(v int32) {
	o.PluginPort = &v
}

// GetDestinationPort returns the DestinationPort field value if set, zero value otherwise.
func (o *PluginInstanceFirewallRule) GetDestinationPort() int32 {
	if o == nil || o.DestinationPort == nil {
		var ret int32
		return ret
	}
	return *o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstanceFirewallRule) GetDestinationPortOk() (*int32, bool) {
	if o == nil || o.DestinationPort == nil {
		return nil, false
	}
	return o.DestinationPort, true
}

// HasDestinationPort returns a boolean if a field has been set.
func (o *PluginInstanceFirewallRule) HasDestinationPort() bool {
	if o != nil && o.DestinationPort != nil {
		return true
	}

	return false
}

// SetDestinationPort gets a reference to the given int32 and assigns it to the DestinationPort field.
func (o *PluginInstanceFirewallRule) SetDestinationPort(v int32) {
	o.DestinationPort = &v
}

func (o PluginInstanceFirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.PluginPort != nil {
		toSerialize["plugin_port"] = o.PluginPort
	}
	if o.DestinationPort != nil {
		toSerialize["destination_port"] = o.DestinationPort
	}
	return json.Marshal(toSerialize)
}

type NullablePluginInstanceFirewallRule struct {
	value *PluginInstanceFirewallRule
	isSet bool
}

func (v NullablePluginInstanceFirewallRule) Get() *PluginInstanceFirewallRule {
	return v.value
}

func (v *NullablePluginInstanceFirewallRule) Set(val *PluginInstanceFirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginInstanceFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginInstanceFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginInstanceFirewallRule(val *PluginInstanceFirewallRule) *NullablePluginInstanceFirewallRule {
	return &NullablePluginInstanceFirewallRule{value: val, isSet: true}
}

func (v NullablePluginInstanceFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginInstanceFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


