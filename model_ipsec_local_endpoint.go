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

// IpsecLocalEndpoint struct for IpsecLocalEndpoint
type IpsecLocalEndpoint struct {
	NatTraversal *bool `json:"nat_traversal,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	OverlaySubnet *string `json:"overlay_subnet,omitempty"`
	PrivateIpaddress *string `json:"private_ipaddress,omitempty"`
	IpsecLocalIpaddress *string `json:"ipsec_local_ipaddress,omitempty"`
	Asn *int32 `json:"asn,omitempty"`
}

// NewIpsecLocalEndpoint instantiates a new IpsecLocalEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpsecLocalEndpoint() *IpsecLocalEndpoint {
	this := IpsecLocalEndpoint{}
	return &this
}

// NewIpsecLocalEndpointWithDefaults instantiates a new IpsecLocalEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpsecLocalEndpointWithDefaults() *IpsecLocalEndpoint {
	this := IpsecLocalEndpoint{}
	return &this
}

// GetNatTraversal returns the NatTraversal field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetNatTraversal() bool {
	if o == nil || o.NatTraversal == nil {
		var ret bool
		return ret
	}
	return *o.NatTraversal
}

// GetNatTraversalOk returns a tuple with the NatTraversal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetNatTraversalOk() (*bool, bool) {
	if o == nil || o.NatTraversal == nil {
		return nil, false
	}
	return o.NatTraversal, true
}

// HasNatTraversal returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasNatTraversal() bool {
	if o != nil && o.NatTraversal != nil {
		return true
	}

	return false
}

// SetNatTraversal gets a reference to the given bool and assigns it to the NatTraversal field.
func (o *IpsecLocalEndpoint) SetNatTraversal(v bool) {
	o.NatTraversal = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *IpsecLocalEndpoint) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetOverlaySubnet returns the OverlaySubnet field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetOverlaySubnet() string {
	if o == nil || o.OverlaySubnet == nil {
		var ret string
		return ret
	}
	return *o.OverlaySubnet
}

// GetOverlaySubnetOk returns a tuple with the OverlaySubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetOverlaySubnetOk() (*string, bool) {
	if o == nil || o.OverlaySubnet == nil {
		return nil, false
	}
	return o.OverlaySubnet, true
}

// HasOverlaySubnet returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasOverlaySubnet() bool {
	if o != nil && o.OverlaySubnet != nil {
		return true
	}

	return false
}

// SetOverlaySubnet gets a reference to the given string and assigns it to the OverlaySubnet field.
func (o *IpsecLocalEndpoint) SetOverlaySubnet(v string) {
	o.OverlaySubnet = &v
}

// GetPrivateIpaddress returns the PrivateIpaddress field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetPrivateIpaddress() string {
	if o == nil || o.PrivateIpaddress == nil {
		var ret string
		return ret
	}
	return *o.PrivateIpaddress
}

// GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetPrivateIpaddressOk() (*string, bool) {
	if o == nil || o.PrivateIpaddress == nil {
		return nil, false
	}
	return o.PrivateIpaddress, true
}

// HasPrivateIpaddress returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasPrivateIpaddress() bool {
	if o != nil && o.PrivateIpaddress != nil {
		return true
	}

	return false
}

// SetPrivateIpaddress gets a reference to the given string and assigns it to the PrivateIpaddress field.
func (o *IpsecLocalEndpoint) SetPrivateIpaddress(v string) {
	o.PrivateIpaddress = &v
}

// GetIpsecLocalIpaddress returns the IpsecLocalIpaddress field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetIpsecLocalIpaddress() string {
	if o == nil || o.IpsecLocalIpaddress == nil {
		var ret string
		return ret
	}
	return *o.IpsecLocalIpaddress
}

// GetIpsecLocalIpaddressOk returns a tuple with the IpsecLocalIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetIpsecLocalIpaddressOk() (*string, bool) {
	if o == nil || o.IpsecLocalIpaddress == nil {
		return nil, false
	}
	return o.IpsecLocalIpaddress, true
}

// HasIpsecLocalIpaddress returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasIpsecLocalIpaddress() bool {
	if o != nil && o.IpsecLocalIpaddress != nil {
		return true
	}

	return false
}

// SetIpsecLocalIpaddress gets a reference to the given string and assigns it to the IpsecLocalIpaddress field.
func (o *IpsecLocalEndpoint) SetIpsecLocalIpaddress(v string) {
	o.IpsecLocalIpaddress = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *IpsecLocalEndpoint) GetAsn() int32 {
	if o == nil || o.Asn == nil {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpsecLocalEndpoint) GetAsnOk() (*int32, bool) {
	if o == nil || o.Asn == nil {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *IpsecLocalEndpoint) HasAsn() bool {
	if o != nil && o.Asn != nil {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *IpsecLocalEndpoint) SetAsn(v int32) {
	o.Asn = &v
}

func (o IpsecLocalEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatTraversal != nil {
		toSerialize["nat_traversal"] = o.NatTraversal
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.OverlaySubnet != nil {
		toSerialize["overlay_subnet"] = o.OverlaySubnet
	}
	if o.PrivateIpaddress != nil {
		toSerialize["private_ipaddress"] = o.PrivateIpaddress
	}
	if o.IpsecLocalIpaddress != nil {
		toSerialize["ipsec_local_ipaddress"] = o.IpsecLocalIpaddress
	}
	if o.Asn != nil {
		toSerialize["asn"] = o.Asn
	}
	return json.Marshal(toSerialize)
}

type NullableIpsecLocalEndpoint struct {
	value *IpsecLocalEndpoint
	isSet bool
}

func (v NullableIpsecLocalEndpoint) Get() *IpsecLocalEndpoint {
	return v.value
}

func (v *NullableIpsecLocalEndpoint) Set(val *IpsecLocalEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIpsecLocalEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIpsecLocalEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpsecLocalEndpoint(val *IpsecLocalEndpoint) *NullableIpsecLocalEndpoint {
	return &NullableIpsecLocalEndpoint{value: val, isSet: true}
}

func (v NullableIpsecLocalEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpsecLocalEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

