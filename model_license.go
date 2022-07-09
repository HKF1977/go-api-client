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

// License struct for License
type License struct {
	// Features available such as eBGP, CloudWAN etc.
	Capabilities []string `json:"capabilities,omitempty"`
	Finalized *bool `json:"finalized,omitempty"`
	MyManagerVip *string `json:"my_manager_vip,omitempty"`
	// State of license, accepted, in-progress, failed
	License *string `json:"license,omitempty"`
	LicensePresent *bool `json:"license_present,omitempty"`
	Sha1Checksum *string `json:"sha1_checksum,omitempty"`
	UploadedAt *string `json:"uploaded_at,omitempty"`
	CustomAddressing *bool `json:"custom_addressing,omitempty"`
	UploadedAtI *int32 `json:"uploaded_at_i,omitempty"`
	ContainerDetails *LicenseContainerDetails `json:"container_details,omitempty"`
	Topology *Topology `json:"topology,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *License) GetCapabilities() []string {
	if o == nil || o.Capabilities == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *License) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *License) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetFinalized returns the Finalized field value if set, zero value otherwise.
func (o *License) GetFinalized() bool {
	if o == nil || o.Finalized == nil {
		var ret bool
		return ret
	}
	return *o.Finalized
}

// GetFinalizedOk returns a tuple with the Finalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetFinalizedOk() (*bool, bool) {
	if o == nil || o.Finalized == nil {
		return nil, false
	}
	return o.Finalized, true
}

// HasFinalized returns a boolean if a field has been set.
func (o *License) HasFinalized() bool {
	if o != nil && o.Finalized != nil {
		return true
	}

	return false
}

// SetFinalized gets a reference to the given bool and assigns it to the Finalized field.
func (o *License) SetFinalized(v bool) {
	o.Finalized = &v
}

// GetMyManagerVip returns the MyManagerVip field value if set, zero value otherwise.
func (o *License) GetMyManagerVip() string {
	if o == nil || o.MyManagerVip == nil {
		var ret string
		return ret
	}
	return *o.MyManagerVip
}

// GetMyManagerVipOk returns a tuple with the MyManagerVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMyManagerVipOk() (*string, bool) {
	if o == nil || o.MyManagerVip == nil {
		return nil, false
	}
	return o.MyManagerVip, true
}

// HasMyManagerVip returns a boolean if a field has been set.
func (o *License) HasMyManagerVip() bool {
	if o != nil && o.MyManagerVip != nil {
		return true
	}

	return false
}

// SetMyManagerVip gets a reference to the given string and assigns it to the MyManagerVip field.
func (o *License) SetMyManagerVip(v string) {
	o.MyManagerVip = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *License) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *License) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *License) SetLicense(v string) {
	o.License = &v
}

// GetLicensePresent returns the LicensePresent field value if set, zero value otherwise.
func (o *License) GetLicensePresent() bool {
	if o == nil || o.LicensePresent == nil {
		var ret bool
		return ret
	}
	return *o.LicensePresent
}

// GetLicensePresentOk returns a tuple with the LicensePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicensePresentOk() (*bool, bool) {
	if o == nil || o.LicensePresent == nil {
		return nil, false
	}
	return o.LicensePresent, true
}

// HasLicensePresent returns a boolean if a field has been set.
func (o *License) HasLicensePresent() bool {
	if o != nil && o.LicensePresent != nil {
		return true
	}

	return false
}

// SetLicensePresent gets a reference to the given bool and assigns it to the LicensePresent field.
func (o *License) SetLicensePresent(v bool) {
	o.LicensePresent = &v
}

// GetSha1Checksum returns the Sha1Checksum field value if set, zero value otherwise.
func (o *License) GetSha1Checksum() string {
	if o == nil || o.Sha1Checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha1Checksum
}

// GetSha1ChecksumOk returns a tuple with the Sha1Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetSha1ChecksumOk() (*string, bool) {
	if o == nil || o.Sha1Checksum == nil {
		return nil, false
	}
	return o.Sha1Checksum, true
}

// HasSha1Checksum returns a boolean if a field has been set.
func (o *License) HasSha1Checksum() bool {
	if o != nil && o.Sha1Checksum != nil {
		return true
	}

	return false
}

// SetSha1Checksum gets a reference to the given string and assigns it to the Sha1Checksum field.
func (o *License) SetSha1Checksum(v string) {
	o.Sha1Checksum = &v
}

// GetUploadedAt returns the UploadedAt field value if set, zero value otherwise.
func (o *License) GetUploadedAt() string {
	if o == nil || o.UploadedAt == nil {
		var ret string
		return ret
	}
	return *o.UploadedAt
}

// GetUploadedAtOk returns a tuple with the UploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUploadedAtOk() (*string, bool) {
	if o == nil || o.UploadedAt == nil {
		return nil, false
	}
	return o.UploadedAt, true
}

// HasUploadedAt returns a boolean if a field has been set.
func (o *License) HasUploadedAt() bool {
	if o != nil && o.UploadedAt != nil {
		return true
	}

	return false
}

// SetUploadedAt gets a reference to the given string and assigns it to the UploadedAt field.
func (o *License) SetUploadedAt(v string) {
	o.UploadedAt = &v
}

// GetCustomAddressing returns the CustomAddressing field value if set, zero value otherwise.
func (o *License) GetCustomAddressing() bool {
	if o == nil || o.CustomAddressing == nil {
		var ret bool
		return ret
	}
	return *o.CustomAddressing
}

// GetCustomAddressingOk returns a tuple with the CustomAddressing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetCustomAddressingOk() (*bool, bool) {
	if o == nil || o.CustomAddressing == nil {
		return nil, false
	}
	return o.CustomAddressing, true
}

// HasCustomAddressing returns a boolean if a field has been set.
func (o *License) HasCustomAddressing() bool {
	if o != nil && o.CustomAddressing != nil {
		return true
	}

	return false
}

// SetCustomAddressing gets a reference to the given bool and assigns it to the CustomAddressing field.
func (o *License) SetCustomAddressing(v bool) {
	o.CustomAddressing = &v
}

// GetUploadedAtI returns the UploadedAtI field value if set, zero value otherwise.
func (o *License) GetUploadedAtI() int32 {
	if o == nil || o.UploadedAtI == nil {
		var ret int32
		return ret
	}
	return *o.UploadedAtI
}

// GetUploadedAtIOk returns a tuple with the UploadedAtI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUploadedAtIOk() (*int32, bool) {
	if o == nil || o.UploadedAtI == nil {
		return nil, false
	}
	return o.UploadedAtI, true
}

// HasUploadedAtI returns a boolean if a field has been set.
func (o *License) HasUploadedAtI() bool {
	if o != nil && o.UploadedAtI != nil {
		return true
	}

	return false
}

// SetUploadedAtI gets a reference to the given int32 and assigns it to the UploadedAtI field.
func (o *License) SetUploadedAtI(v int32) {
	o.UploadedAtI = &v
}

// GetContainerDetails returns the ContainerDetails field value if set, zero value otherwise.
func (o *License) GetContainerDetails() LicenseContainerDetails {
	if o == nil || o.ContainerDetails == nil {
		var ret LicenseContainerDetails
		return ret
	}
	return *o.ContainerDetails
}

// GetContainerDetailsOk returns a tuple with the ContainerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetContainerDetailsOk() (*LicenseContainerDetails, bool) {
	if o == nil || o.ContainerDetails == nil {
		return nil, false
	}
	return o.ContainerDetails, true
}

// HasContainerDetails returns a boolean if a field has been set.
func (o *License) HasContainerDetails() bool {
	if o != nil && o.ContainerDetails != nil {
		return true
	}

	return false
}

// SetContainerDetails gets a reference to the given LicenseContainerDetails and assigns it to the ContainerDetails field.
func (o *License) SetContainerDetails(v LicenseContainerDetails) {
	o.ContainerDetails = &v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *License) GetTopology() Topology {
	if o == nil || o.Topology == nil {
		var ret Topology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetTopologyOk() (*Topology, bool) {
	if o == nil || o.Topology == nil {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *License) HasTopology() bool {
	if o != nil && o.Topology != nil {
		return true
	}

	return false
}

// SetTopology gets a reference to the given Topology and assigns it to the Topology field.
func (o *License) SetTopology(v Topology) {
	o.Topology = &v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Finalized != nil {
		toSerialize["finalized"] = o.Finalized
	}
	if o.MyManagerVip != nil {
		toSerialize["my_manager_vip"] = o.MyManagerVip
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.LicensePresent != nil {
		toSerialize["license_present"] = o.LicensePresent
	}
	if o.Sha1Checksum != nil {
		toSerialize["sha1_checksum"] = o.Sha1Checksum
	}
	if o.UploadedAt != nil {
		toSerialize["uploaded_at"] = o.UploadedAt
	}
	if o.CustomAddressing != nil {
		toSerialize["custom_addressing"] = o.CustomAddressing
	}
	if o.UploadedAtI != nil {
		toSerialize["uploaded_at_i"] = o.UploadedAtI
	}
	if o.ContainerDetails != nil {
		toSerialize["container_details"] = o.ContainerDetails
	}
	if o.Topology != nil {
		toSerialize["topology"] = o.Topology
	}
	return json.Marshal(toSerialize)
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


