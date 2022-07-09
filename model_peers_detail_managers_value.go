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

// PeersDetailManagersValue struct for PeersDetailManagersValue
type PeersDetailManagersValue struct {
	Id *int32 `json:"id,omitempty"`
	NotSet *bool `json:"not_set,omitempty"`
	Self *bool `json:"self,omitempty"`
	Mtu *string `json:"mtu,omitempty"`
	Reachable *bool `json:"reachable,omitempty"`
	Address *string `json:"address,omitempty"`
	OverlayIpaddress *string `json:"overlay_ipaddress,omitempty"`
}

// NewPeersDetailManagersValue instantiates a new PeersDetailManagersValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeersDetailManagersValue() *PeersDetailManagersValue {
	this := PeersDetailManagersValue{}
	return &this
}

// NewPeersDetailManagersValueWithDefaults instantiates a new PeersDetailManagersValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeersDetailManagersValueWithDefaults() *PeersDetailManagersValue {
	this := PeersDetailManagersValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PeersDetailManagersValue) SetId(v int32) {
	o.Id = &v
}

// GetNotSet returns the NotSet field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetNotSet() bool {
	if o == nil || o.NotSet == nil {
		var ret bool
		return ret
	}
	return *o.NotSet
}

// GetNotSetOk returns a tuple with the NotSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetNotSetOk() (*bool, bool) {
	if o == nil || o.NotSet == nil {
		return nil, false
	}
	return o.NotSet, true
}

// HasNotSet returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasNotSet() bool {
	if o != nil && o.NotSet != nil {
		return true
	}

	return false
}

// SetNotSet gets a reference to the given bool and assigns it to the NotSet field.
func (o *PeersDetailManagersValue) SetNotSet(v bool) {
	o.NotSet = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetSelf() bool {
	if o == nil || o.Self == nil {
		var ret bool
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetSelfOk() (*bool, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given bool and assigns it to the Self field.
func (o *PeersDetailManagersValue) SetSelf(v bool) {
	o.Self = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetMtu() string {
	if o == nil || o.Mtu == nil {
		var ret string
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetMtuOk() (*string, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given string and assigns it to the Mtu field.
func (o *PeersDetailManagersValue) SetMtu(v string) {
	o.Mtu = &v
}

// GetReachable returns the Reachable field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetReachable() bool {
	if o == nil || o.Reachable == nil {
		var ret bool
		return ret
	}
	return *o.Reachable
}

// GetReachableOk returns a tuple with the Reachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetReachableOk() (*bool, bool) {
	if o == nil || o.Reachable == nil {
		return nil, false
	}
	return o.Reachable, true
}

// HasReachable returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasReachable() bool {
	if o != nil && o.Reachable != nil {
		return true
	}

	return false
}

// SetReachable gets a reference to the given bool and assigns it to the Reachable field.
func (o *PeersDetailManagersValue) SetReachable(v bool) {
	o.Reachable = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PeersDetailManagersValue) SetAddress(v string) {
	o.Address = &v
}

// GetOverlayIpaddress returns the OverlayIpaddress field value if set, zero value otherwise.
func (o *PeersDetailManagersValue) GetOverlayIpaddress() string {
	if o == nil || o.OverlayIpaddress == nil {
		var ret string
		return ret
	}
	return *o.OverlayIpaddress
}

// GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersDetailManagersValue) GetOverlayIpaddressOk() (*string, bool) {
	if o == nil || o.OverlayIpaddress == nil {
		return nil, false
	}
	return o.OverlayIpaddress, true
}

// HasOverlayIpaddress returns a boolean if a field has been set.
func (o *PeersDetailManagersValue) HasOverlayIpaddress() bool {
	if o != nil && o.OverlayIpaddress != nil {
		return true
	}

	return false
}

// SetOverlayIpaddress gets a reference to the given string and assigns it to the OverlayIpaddress field.
func (o *PeersDetailManagersValue) SetOverlayIpaddress(v string) {
	o.OverlayIpaddress = &v
}

func (o PeersDetailManagersValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NotSet != nil {
		toSerialize["not_set"] = o.NotSet
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Mtu != nil {
		toSerialize["mtu"] = o.Mtu
	}
	if o.Reachable != nil {
		toSerialize["reachable"] = o.Reachable
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.OverlayIpaddress != nil {
		toSerialize["overlay_ipaddress"] = o.OverlayIpaddress
	}
	return json.Marshal(toSerialize)
}

type NullablePeersDetailManagersValue struct {
	value *PeersDetailManagersValue
	isSet bool
}

func (v NullablePeersDetailManagersValue) Get() *PeersDetailManagersValue {
	return v.value
}

func (v *NullablePeersDetailManagersValue) Set(val *PeersDetailManagersValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePeersDetailManagersValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePeersDetailManagersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeersDetailManagersValue(val *PeersDetailManagersValue) *NullablePeersDetailManagersValue {
	return &NullablePeersDetailManagersValue{value: val, isSet: true}
}

func (v NullablePeersDetailManagersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeersDetailManagersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


