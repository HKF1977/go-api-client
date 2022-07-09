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

// ConnectedSubnet struct for ConnectedSubnet
type ConnectedSubnet struct {
	Subnet *string `json:"subnet,omitempty"`
	Network *string `json:"network,omitempty"`
	CidrMask *string `json:"cidr_mask,omitempty"`
	Managerid *int32 `json:"managerid,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	// ipsec, local_no_encryption, remote_manager, or ebgp
	Origin *string `json:"origin,omitempty"`
}

// NewConnectedSubnet instantiates a new ConnectedSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedSubnet() *ConnectedSubnet {
	this := ConnectedSubnet{}
	return &this
}

// NewConnectedSubnetWithDefaults instantiates a new ConnectedSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedSubnetWithDefaults() *ConnectedSubnet {
	this := ConnectedSubnet{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *ConnectedSubnet) SetSubnet(v string) {
	o.Subnet = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ConnectedSubnet) SetNetwork(v string) {
	o.Network = &v
}

// GetCidrMask returns the CidrMask field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetCidrMask() string {
	if o == nil || o.CidrMask == nil {
		var ret string
		return ret
	}
	return *o.CidrMask
}

// GetCidrMaskOk returns a tuple with the CidrMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetCidrMaskOk() (*string, bool) {
	if o == nil || o.CidrMask == nil {
		return nil, false
	}
	return o.CidrMask, true
}

// HasCidrMask returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasCidrMask() bool {
	if o != nil && o.CidrMask != nil {
		return true
	}

	return false
}

// SetCidrMask gets a reference to the given string and assigns it to the CidrMask field.
func (o *ConnectedSubnet) SetCidrMask(v string) {
	o.CidrMask = &v
}

// GetManagerid returns the Managerid field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetManagerid() int32 {
	if o == nil || o.Managerid == nil {
		var ret int32
		return ret
	}
	return *o.Managerid
}

// GetManageridOk returns a tuple with the Managerid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetManageridOk() (*int32, bool) {
	if o == nil || o.Managerid == nil {
		return nil, false
	}
	return o.Managerid, true
}

// HasManagerid returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasManagerid() bool {
	if o != nil && o.Managerid != nil {
		return true
	}

	return false
}

// SetManagerid gets a reference to the given int32 and assigns it to the Managerid field.
func (o *ConnectedSubnet) SetManagerid(v int32) {
	o.Managerid = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *ConnectedSubnet) SetNetmask(v string) {
	o.Netmask = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ConnectedSubnet) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSubnet) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ConnectedSubnet) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *ConnectedSubnet) SetOrigin(v string) {
	o.Origin = &v
}

func (o ConnectedSubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.CidrMask != nil {
		toSerialize["cidr_mask"] = o.CidrMask
	}
	if o.Managerid != nil {
		toSerialize["managerid"] = o.Managerid
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedSubnet struct {
	value *ConnectedSubnet
	isSet bool
}

func (v NullableConnectedSubnet) Get() *ConnectedSubnet {
	return v.value
}

func (v *NullableConnectedSubnet) Set(val *ConnectedSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedSubnet(val *ConnectedSubnet) *NullableConnectedSubnet {
	return &NullableConnectedSubnet{value: val, isSet: true}
}

func (v NullableConnectedSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


