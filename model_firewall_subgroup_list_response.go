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

// FirewallSubgroupListResponse struct for FirewallSubgroupListResponse
type FirewallSubgroupListResponse struct {
	Response *FirewallSubgroupListResponseResponse `json:"response,omitempty"`
}

// NewFirewallSubgroupListResponse instantiates a new FirewallSubgroupListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSubgroupListResponse() *FirewallSubgroupListResponse {
	this := FirewallSubgroupListResponse{}
	return &this
}

// NewFirewallSubgroupListResponseWithDefaults instantiates a new FirewallSubgroupListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSubgroupListResponseWithDefaults() *FirewallSubgroupListResponse {
	this := FirewallSubgroupListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *FirewallSubgroupListResponse) GetResponse() FirewallSubgroupListResponseResponse {
	if o == nil || o.Response == nil {
		var ret FirewallSubgroupListResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubgroupListResponse) GetResponseOk() (*FirewallSubgroupListResponseResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *FirewallSubgroupListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given FirewallSubgroupListResponseResponse and assigns it to the Response field.
func (o *FirewallSubgroupListResponse) SetResponse(v FirewallSubgroupListResponseResponse) {
	o.Response = &v
}

func (o FirewallSubgroupListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallSubgroupListResponse struct {
	value *FirewallSubgroupListResponse
	isSet bool
}

func (v NullableFirewallSubgroupListResponse) Get() *FirewallSubgroupListResponse {
	return v.value
}

func (v *NullableFirewallSubgroupListResponse) Set(val *FirewallSubgroupListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSubgroupListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSubgroupListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSubgroupListResponse(val *FirewallSubgroupListResponse) *NullableFirewallSubgroupListResponse {
	return &NullableFirewallSubgroupListResponse{value: val, isSet: true}
}

func (v NullableFirewallSubgroupListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSubgroupListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


