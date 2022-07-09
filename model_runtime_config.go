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

// RuntimeConfig struct for RuntimeConfig
type RuntimeConfig struct {
	// Autonomous system number for controller if peered
	Asn *int32 `json:"asn,omitempty"`
	TopologyName *string `json:"topology_name,omitempty"`
	ControllerName *string `json:"controller_name,omitempty"`
	TopologyChecksum *string `json:"topology_checksum,omitempty"`
	// This managers ID in peered topology
	ManagerId *int32 `json:"manager_id,omitempty"`
	// NTP host endpoints, whitespace delimited
	NtpHosts *string `json:"ntp_hosts,omitempty"`
	Vns3Version *string `json:"vns3_version,omitempty"`
	Licensed *bool `json:"licensed,omitempty"`
	// This managers overlay IP in peered topology
	OverlayIpaddress *string `json:"overlay_ipaddress,omitempty"`
	Peered *bool `json:"peered,omitempty"`
	PublicIpaddress *string `json:"public_ipaddress,omitempty"`
	PrivateIpaddress *string `json:"private_ipaddress,omitempty"`
	// Security token in peered topology
	SecurityToken *string `json:"security_token,omitempty"`
}

// NewRuntimeConfig instantiates a new RuntimeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeConfig() *RuntimeConfig {
	this := RuntimeConfig{}
	return &this
}

// NewRuntimeConfigWithDefaults instantiates a new RuntimeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeConfigWithDefaults() *RuntimeConfig {
	this := RuntimeConfig{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *RuntimeConfig) GetAsn() int32 {
	if o == nil || o.Asn == nil {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetAsnOk() (*int32, bool) {
	if o == nil || o.Asn == nil {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *RuntimeConfig) HasAsn() bool {
	if o != nil && o.Asn != nil {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *RuntimeConfig) SetAsn(v int32) {
	o.Asn = &v
}

// GetTopologyName returns the TopologyName field value if set, zero value otherwise.
func (o *RuntimeConfig) GetTopologyName() string {
	if o == nil || o.TopologyName == nil {
		var ret string
		return ret
	}
	return *o.TopologyName
}

// GetTopologyNameOk returns a tuple with the TopologyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetTopologyNameOk() (*string, bool) {
	if o == nil || o.TopologyName == nil {
		return nil, false
	}
	return o.TopologyName, true
}

// HasTopologyName returns a boolean if a field has been set.
func (o *RuntimeConfig) HasTopologyName() bool {
	if o != nil && o.TopologyName != nil {
		return true
	}

	return false
}

// SetTopologyName gets a reference to the given string and assigns it to the TopologyName field.
func (o *RuntimeConfig) SetTopologyName(v string) {
	o.TopologyName = &v
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *RuntimeConfig) GetControllerName() string {
	if o == nil || o.ControllerName == nil {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetControllerNameOk() (*string, bool) {
	if o == nil || o.ControllerName == nil {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *RuntimeConfig) HasControllerName() bool {
	if o != nil && o.ControllerName != nil {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *RuntimeConfig) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetTopologyChecksum returns the TopologyChecksum field value if set, zero value otherwise.
func (o *RuntimeConfig) GetTopologyChecksum() string {
	if o == nil || o.TopologyChecksum == nil {
		var ret string
		return ret
	}
	return *o.TopologyChecksum
}

// GetTopologyChecksumOk returns a tuple with the TopologyChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetTopologyChecksumOk() (*string, bool) {
	if o == nil || o.TopologyChecksum == nil {
		return nil, false
	}
	return o.TopologyChecksum, true
}

// HasTopologyChecksum returns a boolean if a field has been set.
func (o *RuntimeConfig) HasTopologyChecksum() bool {
	if o != nil && o.TopologyChecksum != nil {
		return true
	}

	return false
}

// SetTopologyChecksum gets a reference to the given string and assigns it to the TopologyChecksum field.
func (o *RuntimeConfig) SetTopologyChecksum(v string) {
	o.TopologyChecksum = &v
}

// GetManagerId returns the ManagerId field value if set, zero value otherwise.
func (o *RuntimeConfig) GetManagerId() int32 {
	if o == nil || o.ManagerId == nil {
		var ret int32
		return ret
	}
	return *o.ManagerId
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetManagerIdOk() (*int32, bool) {
	if o == nil || o.ManagerId == nil {
		return nil, false
	}
	return o.ManagerId, true
}

// HasManagerId returns a boolean if a field has been set.
func (o *RuntimeConfig) HasManagerId() bool {
	if o != nil && o.ManagerId != nil {
		return true
	}

	return false
}

// SetManagerId gets a reference to the given int32 and assigns it to the ManagerId field.
func (o *RuntimeConfig) SetManagerId(v int32) {
	o.ManagerId = &v
}

// GetNtpHosts returns the NtpHosts field value if set, zero value otherwise.
func (o *RuntimeConfig) GetNtpHosts() string {
	if o == nil || o.NtpHosts == nil {
		var ret string
		return ret
	}
	return *o.NtpHosts
}

// GetNtpHostsOk returns a tuple with the NtpHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetNtpHostsOk() (*string, bool) {
	if o == nil || o.NtpHosts == nil {
		return nil, false
	}
	return o.NtpHosts, true
}

// HasNtpHosts returns a boolean if a field has been set.
func (o *RuntimeConfig) HasNtpHosts() bool {
	if o != nil && o.NtpHosts != nil {
		return true
	}

	return false
}

// SetNtpHosts gets a reference to the given string and assigns it to the NtpHosts field.
func (o *RuntimeConfig) SetNtpHosts(v string) {
	o.NtpHosts = &v
}

// GetVns3Version returns the Vns3Version field value if set, zero value otherwise.
func (o *RuntimeConfig) GetVns3Version() string {
	if o == nil || o.Vns3Version == nil {
		var ret string
		return ret
	}
	return *o.Vns3Version
}

// GetVns3VersionOk returns a tuple with the Vns3Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetVns3VersionOk() (*string, bool) {
	if o == nil || o.Vns3Version == nil {
		return nil, false
	}
	return o.Vns3Version, true
}

// HasVns3Version returns a boolean if a field has been set.
func (o *RuntimeConfig) HasVns3Version() bool {
	if o != nil && o.Vns3Version != nil {
		return true
	}

	return false
}

// SetVns3Version gets a reference to the given string and assigns it to the Vns3Version field.
func (o *RuntimeConfig) SetVns3Version(v string) {
	o.Vns3Version = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RuntimeConfig) GetLicensed() bool {
	if o == nil || o.Licensed == nil {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetLicensedOk() (*bool, bool) {
	if o == nil || o.Licensed == nil {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RuntimeConfig) HasLicensed() bool {
	if o != nil && o.Licensed != nil {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RuntimeConfig) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetOverlayIpaddress returns the OverlayIpaddress field value if set, zero value otherwise.
func (o *RuntimeConfig) GetOverlayIpaddress() string {
	if o == nil || o.OverlayIpaddress == nil {
		var ret string
		return ret
	}
	return *o.OverlayIpaddress
}

// GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetOverlayIpaddressOk() (*string, bool) {
	if o == nil || o.OverlayIpaddress == nil {
		return nil, false
	}
	return o.OverlayIpaddress, true
}

// HasOverlayIpaddress returns a boolean if a field has been set.
func (o *RuntimeConfig) HasOverlayIpaddress() bool {
	if o != nil && o.OverlayIpaddress != nil {
		return true
	}

	return false
}

// SetOverlayIpaddress gets a reference to the given string and assigns it to the OverlayIpaddress field.
func (o *RuntimeConfig) SetOverlayIpaddress(v string) {
	o.OverlayIpaddress = &v
}

// GetPeered returns the Peered field value if set, zero value otherwise.
func (o *RuntimeConfig) GetPeered() bool {
	if o == nil || o.Peered == nil {
		var ret bool
		return ret
	}
	return *o.Peered
}

// GetPeeredOk returns a tuple with the Peered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetPeeredOk() (*bool, bool) {
	if o == nil || o.Peered == nil {
		return nil, false
	}
	return o.Peered, true
}

// HasPeered returns a boolean if a field has been set.
func (o *RuntimeConfig) HasPeered() bool {
	if o != nil && o.Peered != nil {
		return true
	}

	return false
}

// SetPeered gets a reference to the given bool and assigns it to the Peered field.
func (o *RuntimeConfig) SetPeered(v bool) {
	o.Peered = &v
}

// GetPublicIpaddress returns the PublicIpaddress field value if set, zero value otherwise.
func (o *RuntimeConfig) GetPublicIpaddress() string {
	if o == nil || o.PublicIpaddress == nil {
		var ret string
		return ret
	}
	return *o.PublicIpaddress
}

// GetPublicIpaddressOk returns a tuple with the PublicIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetPublicIpaddressOk() (*string, bool) {
	if o == nil || o.PublicIpaddress == nil {
		return nil, false
	}
	return o.PublicIpaddress, true
}

// HasPublicIpaddress returns a boolean if a field has been set.
func (o *RuntimeConfig) HasPublicIpaddress() bool {
	if o != nil && o.PublicIpaddress != nil {
		return true
	}

	return false
}

// SetPublicIpaddress gets a reference to the given string and assigns it to the PublicIpaddress field.
func (o *RuntimeConfig) SetPublicIpaddress(v string) {
	o.PublicIpaddress = &v
}

// GetPrivateIpaddress returns the PrivateIpaddress field value if set, zero value otherwise.
func (o *RuntimeConfig) GetPrivateIpaddress() string {
	if o == nil || o.PrivateIpaddress == nil {
		var ret string
		return ret
	}
	return *o.PrivateIpaddress
}

// GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetPrivateIpaddressOk() (*string, bool) {
	if o == nil || o.PrivateIpaddress == nil {
		return nil, false
	}
	return o.PrivateIpaddress, true
}

// HasPrivateIpaddress returns a boolean if a field has been set.
func (o *RuntimeConfig) HasPrivateIpaddress() bool {
	if o != nil && o.PrivateIpaddress != nil {
		return true
	}

	return false
}

// SetPrivateIpaddress gets a reference to the given string and assigns it to the PrivateIpaddress field.
func (o *RuntimeConfig) SetPrivateIpaddress(v string) {
	o.PrivateIpaddress = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *RuntimeConfig) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeConfig) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *RuntimeConfig) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *RuntimeConfig) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

func (o RuntimeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asn != nil {
		toSerialize["asn"] = o.Asn
	}
	if o.TopologyName != nil {
		toSerialize["topology_name"] = o.TopologyName
	}
	if o.ControllerName != nil {
		toSerialize["controller_name"] = o.ControllerName
	}
	if o.TopologyChecksum != nil {
		toSerialize["topology_checksum"] = o.TopologyChecksum
	}
	if o.ManagerId != nil {
		toSerialize["manager_id"] = o.ManagerId
	}
	if o.NtpHosts != nil {
		toSerialize["ntp_hosts"] = o.NtpHosts
	}
	if o.Vns3Version != nil {
		toSerialize["vns3_version"] = o.Vns3Version
	}
	if o.Licensed != nil {
		toSerialize["licensed"] = o.Licensed
	}
	if o.OverlayIpaddress != nil {
		toSerialize["overlay_ipaddress"] = o.OverlayIpaddress
	}
	if o.Peered != nil {
		toSerialize["peered"] = o.Peered
	}
	if o.PublicIpaddress != nil {
		toSerialize["public_ipaddress"] = o.PublicIpaddress
	}
	if o.PrivateIpaddress != nil {
		toSerialize["private_ipaddress"] = o.PrivateIpaddress
	}
	if o.SecurityToken != nil {
		toSerialize["security_token"] = o.SecurityToken
	}
	return json.Marshal(toSerialize)
}

type NullableRuntimeConfig struct {
	value *RuntimeConfig
	isSet bool
}

func (v NullableRuntimeConfig) Get() *RuntimeConfig {
	return v.value
}

func (v *NullableRuntimeConfig) Set(val *RuntimeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeConfig(val *RuntimeConfig) *NullableRuntimeConfig {
	return &NullableRuntimeConfig{value: val, isSet: true}
}

func (v NullableRuntimeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

