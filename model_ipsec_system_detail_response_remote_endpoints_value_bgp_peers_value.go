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

// IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue struct for IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue
type IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue struct {
	Asn *int32 `json:"asn,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	// List of \"in permit CIDR\" and/or \"out permit CIDR\" statements in a string delimited by \"\\n\"
	AccessList *string `json:"access_list,omitempty"`
	Id *int32 `json:"id,omitempty"`
	BgpPassword *string `json:"bgp_password,omitempty"`
	AddNetworkDistance *bool `json:"add_network_distance,omitempty"`
	// in or out
	AddNetworkDistanceDirection *string `json:"add_network_distance_direction,omitempty"`
	AddNetworkDistanceHops *int32 `json:"add_network_distance_hops,omitempty"`
	ConnectionDetail *string `json:"connection_detail,omitempty"`
}

// NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue() *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue {
	this := IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue{}
	return &this
}

// NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValueWithDefaults instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValueWithDefaults() *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue {
	this := IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAsn() int32 {
	if o == nil || o.Asn == nil {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAsnOk() (*int32, bool) {
	if o == nil || o.Asn == nil {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAsn() bool {
	if o != nil && o.Asn != nil {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAsn(v int32) {
	o.Asn = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAccessList() string {
	if o == nil || o.AccessList == nil {
		var ret string
		return ret
	}
	return *o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAccessListOk() (*string, bool) {
	if o == nil || o.AccessList == nil {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAccessList() bool {
	if o != nil && o.AccessList != nil {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given string and assigns it to the AccessList field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAccessList(v string) {
	o.AccessList = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetId(v int32) {
	o.Id = &v
}

// GetBgpPassword returns the BgpPassword field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetBgpPassword() string {
	if o == nil || o.BgpPassword == nil {
		var ret string
		return ret
	}
	return *o.BgpPassword
}

// GetBgpPasswordOk returns a tuple with the BgpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetBgpPasswordOk() (*string, bool) {
	if o == nil || o.BgpPassword == nil {
		return nil, false
	}
	return o.BgpPassword, true
}

// HasBgpPassword returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasBgpPassword() bool {
	if o != nil && o.BgpPassword != nil {
		return true
	}

	return false
}

// SetBgpPassword gets a reference to the given string and assigns it to the BgpPassword field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetBgpPassword(v string) {
	o.BgpPassword = &v
}

// GetAddNetworkDistance returns the AddNetworkDistance field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistance() bool {
	if o == nil || o.AddNetworkDistance == nil {
		var ret bool
		return ret
	}
	return *o.AddNetworkDistance
}

// GetAddNetworkDistanceOk returns a tuple with the AddNetworkDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceOk() (*bool, bool) {
	if o == nil || o.AddNetworkDistance == nil {
		return nil, false
	}
	return o.AddNetworkDistance, true
}

// HasAddNetworkDistance returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistance() bool {
	if o != nil && o.AddNetworkDistance != nil {
		return true
	}

	return false
}

// SetAddNetworkDistance gets a reference to the given bool and assigns it to the AddNetworkDistance field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistance(v bool) {
	o.AddNetworkDistance = &v
}

// GetAddNetworkDistanceDirection returns the AddNetworkDistanceDirection field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceDirection() string {
	if o == nil || o.AddNetworkDistanceDirection == nil {
		var ret string
		return ret
	}
	return *o.AddNetworkDistanceDirection
}

// GetAddNetworkDistanceDirectionOk returns a tuple with the AddNetworkDistanceDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceDirectionOk() (*string, bool) {
	if o == nil || o.AddNetworkDistanceDirection == nil {
		return nil, false
	}
	return o.AddNetworkDistanceDirection, true
}

// HasAddNetworkDistanceDirection returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistanceDirection() bool {
	if o != nil && o.AddNetworkDistanceDirection != nil {
		return true
	}

	return false
}

// SetAddNetworkDistanceDirection gets a reference to the given string and assigns it to the AddNetworkDistanceDirection field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistanceDirection(v string) {
	o.AddNetworkDistanceDirection = &v
}

// GetAddNetworkDistanceHops returns the AddNetworkDistanceHops field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceHops() int32 {
	if o == nil || o.AddNetworkDistanceHops == nil {
		var ret int32
		return ret
	}
	return *o.AddNetworkDistanceHops
}

// GetAddNetworkDistanceHopsOk returns a tuple with the AddNetworkDistanceHops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceHopsOk() (*int32, bool) {
	if o == nil || o.AddNetworkDistanceHops == nil {
		return nil, false
	}
	return o.AddNetworkDistanceHops, true
}

// HasAddNetworkDistanceHops returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistanceHops() bool {
	if o != nil && o.AddNetworkDistanceHops != nil {
		return true
	}

	return false
}

// SetAddNetworkDistanceHops gets a reference to the given int32 and assigns it to the AddNetworkDistanceHops field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistanceHops(v int32) {
	o.AddNetworkDistanceHops = &v
}

// GetConnectionDetail returns the ConnectionDetail field value if set, zero value otherwise.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetConnectionDetail() string {
	if o == nil || o.ConnectionDetail == nil {
		var ret string
		return ret
	}
	return *o.ConnectionDetail
}

// GetConnectionDetailOk returns a tuple with the ConnectionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetConnectionDetailOk() (*string, bool) {
	if o == nil || o.ConnectionDetail == nil {
		return nil, false
	}
	return o.ConnectionDetail, true
}

// HasConnectionDetail returns a boolean if a field has been set.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasConnectionDetail() bool {
	if o != nil && o.ConnectionDetail != nil {
		return true
	}

	return false
}

// SetConnectionDetail gets a reference to the given string and assigns it to the ConnectionDetail field.
func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetConnectionDetail(v string) {
	o.ConnectionDetail = &v
}

func (o IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asn != nil {
		toSerialize["asn"] = o.Asn
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.AccessList != nil {
		toSerialize["access_list"] = o.AccessList
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BgpPassword != nil {
		toSerialize["bgp_password"] = o.BgpPassword
	}
	if o.AddNetworkDistance != nil {
		toSerialize["add_network_distance"] = o.AddNetworkDistance
	}
	if o.AddNetworkDistanceDirection != nil {
		toSerialize["add_network_distance_direction"] = o.AddNetworkDistanceDirection
	}
	if o.AddNetworkDistanceHops != nil {
		toSerialize["add_network_distance_hops"] = o.AddNetworkDistanceHops
	}
	if o.ConnectionDetail != nil {
		toSerialize["connection_detail"] = o.ConnectionDetail
	}
	return json.Marshal(toSerialize)
}

type NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue struct {
	value *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue
	isSet bool
}

func (v NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) Get() *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue {
	return v.value
}

func (v *NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) Set(val *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue(val *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) *NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue {
	return &NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue{value: val, isSet: true}
}

func (v NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


