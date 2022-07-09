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

// SystemInterface struct for SystemInterface
type SystemInterface struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// system or system_virtual
	InterfaceType *string `json:"interface_type,omitempty"`
	Description *string `json:"description,omitempty"`
	IpInternal NullableString `json:"ip_internal,omitempty"`
	Mtu *int32 `json:"mtu,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Availability of interface, Up or Down
	Status *string `json:"status,omitempty"`
	MaskBits NullableString `json:"mask_bits,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	SystemDefault *bool `json:"system_default,omitempty"`
	IpExternal NullableString `json:"ip_external,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewSystemInterface instantiates a new SystemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInterface() *SystemInterface {
	this := SystemInterface{}
	return &this
}

// NewSystemInterfaceWithDefaults instantiates a new SystemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInterfaceWithDefaults() *SystemInterface {
	this := SystemInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SystemInterface) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SystemInterface) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SystemInterface) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInterface) SetName(v string) {
	o.Name = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *SystemInterface) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *SystemInterface) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *SystemInterface) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SystemInterface) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SystemInterface) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SystemInterface) SetDescription(v string) {
	o.Description = &v
}

// GetIpInternal returns the IpInternal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemInterface) GetIpInternal() string {
	if o == nil || o.IpInternal.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpInternal.Get()
}

// GetIpInternalOk returns a tuple with the IpInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInterface) GetIpInternalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpInternal.Get(), o.IpInternal.IsSet()
}

// HasIpInternal returns a boolean if a field has been set.
func (o *SystemInterface) HasIpInternal() bool {
	if o != nil && o.IpInternal.IsSet() {
		return true
	}

	return false
}

// SetIpInternal gets a reference to the given NullableString and assigns it to the IpInternal field.
func (o *SystemInterface) SetIpInternal(v string) {
	o.IpInternal.Set(&v)
}
// SetIpInternalNil sets the value for IpInternal to be an explicit nil
func (o *SystemInterface) SetIpInternalNil() {
	o.IpInternal.Set(nil)
}

// UnsetIpInternal ensures that no value is present for IpInternal, not even an explicit nil
func (o *SystemInterface) UnsetIpInternal() {
	o.IpInternal.Unset()
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *SystemInterface) GetMtu() int32 {
	if o == nil || o.Mtu == nil {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetMtuOk() (*int32, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *SystemInterface) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *SystemInterface) SetMtu(v int32) {
	o.Mtu = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SystemInterface) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SystemInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SystemInterface) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SystemInterface) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SystemInterface) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SystemInterface) SetStatus(v string) {
	o.Status = &v
}

// GetMaskBits returns the MaskBits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemInterface) GetMaskBits() string {
	if o == nil || o.MaskBits.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaskBits.Get()
}

// GetMaskBitsOk returns a tuple with the MaskBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInterface) GetMaskBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskBits.Get(), o.MaskBits.IsSet()
}

// HasMaskBits returns a boolean if a field has been set.
func (o *SystemInterface) HasMaskBits() bool {
	if o != nil && o.MaskBits.IsSet() {
		return true
	}

	return false
}

// SetMaskBits gets a reference to the given NullableString and assigns it to the MaskBits field.
func (o *SystemInterface) SetMaskBits(v string) {
	o.MaskBits.Set(&v)
}
// SetMaskBitsNil sets the value for MaskBits to be an explicit nil
func (o *SystemInterface) SetMaskBitsNil() {
	o.MaskBits.Set(nil)
}

// UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil
func (o *SystemInterface) UnsetMaskBits() {
	o.MaskBits.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemInterface) GetGateway() string {
	if o == nil || o.Gateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInterface) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *SystemInterface) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *SystemInterface) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *SystemInterface) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *SystemInterface) UnsetGateway() {
	o.Gateway.Unset()
}

// GetSystemDefault returns the SystemDefault field value if set, zero value otherwise.
func (o *SystemInterface) GetSystemDefault() bool {
	if o == nil || o.SystemDefault == nil {
		var ret bool
		return ret
	}
	return *o.SystemDefault
}

// GetSystemDefaultOk returns a tuple with the SystemDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetSystemDefaultOk() (*bool, bool) {
	if o == nil || o.SystemDefault == nil {
		return nil, false
	}
	return o.SystemDefault, true
}

// HasSystemDefault returns a boolean if a field has been set.
func (o *SystemInterface) HasSystemDefault() bool {
	if o != nil && o.SystemDefault != nil {
		return true
	}

	return false
}

// SetSystemDefault gets a reference to the given bool and assigns it to the SystemDefault field.
func (o *SystemInterface) SetSystemDefault(v bool) {
	o.SystemDefault = &v
}

// GetIpExternal returns the IpExternal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemInterface) GetIpExternal() string {
	if o == nil || o.IpExternal.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpExternal.Get()
}

// GetIpExternalOk returns a tuple with the IpExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInterface) GetIpExternalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpExternal.Get(), o.IpExternal.IsSet()
}

// HasIpExternal returns a boolean if a field has been set.
func (o *SystemInterface) HasIpExternal() bool {
	if o != nil && o.IpExternal.IsSet() {
		return true
	}

	return false
}

// SetIpExternal gets a reference to the given NullableString and assigns it to the IpExternal field.
func (o *SystemInterface) SetIpExternal(v string) {
	o.IpExternal.Set(&v)
}
// SetIpExternalNil sets the value for IpExternal to be an explicit nil
func (o *SystemInterface) SetIpExternalNil() {
	o.IpExternal.Set(nil)
}

// UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
func (o *SystemInterface) UnsetIpExternal() {
	o.IpExternal.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SystemInterface) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterface) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SystemInterface) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SystemInterface) SetTags(v []string) {
	o.Tags = v
}

func (o SystemInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.InterfaceType != nil {
		toSerialize["interface_type"] = o.InterfaceType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IpInternal.IsSet() {
		toSerialize["ip_internal"] = o.IpInternal.Get()
	}
	if o.Mtu != nil {
		toSerialize["mtu"] = o.Mtu
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.MaskBits.IsSet() {
		toSerialize["mask_bits"] = o.MaskBits.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.SystemDefault != nil {
		toSerialize["system_default"] = o.SystemDefault
	}
	if o.IpExternal.IsSet() {
		toSerialize["ip_external"] = o.IpExternal.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInterface struct {
	value *SystemInterface
	isSet bool
}

func (v NullableSystemInterface) Get() *SystemInterface {
	return v.value
}

func (v *NullableSystemInterface) Set(val *SystemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInterface(val *SystemInterface) *NullableSystemInterface {
	return &NullableSystemInterface{value: val, isSet: true}
}

func (v NullableSystemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


