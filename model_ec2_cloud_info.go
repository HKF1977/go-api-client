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

// EC2CloudInfo Metadata returned from AWS instance metadata call. 
type EC2CloudInfo struct {
	AccountId *string `json:"accountId,omitempty"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	RamdiskId NullableString `json:"ramdiskId,omitempty"`
	KernelId NullableString `json:"kernelId,omitempty"`
	PendingTime *string `json:"pendingTime,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	PrivateIp *string `json:"privateIp,omitempty"`
	DevpayProductCodes NullableString `json:"devpayProductCodes,omitempty"`
	MarketplaceProductCodes NullableString `json:"marketplaceProductCodes,omitempty"`
	Version *string `json:"version,omitempty"`
	Region *string `json:"region,omitempty"`
	ImageId *string `json:"imageId,omitempty"`
	BillingProducts NullableString `json:"billingProducts,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EC2CloudInfo EC2CloudInfo

// NewEC2CloudInfo instantiates a new EC2CloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEC2CloudInfo() *EC2CloudInfo {
	this := EC2CloudInfo{}
	return &this
}

// NewEC2CloudInfoWithDefaults instantiates a new EC2CloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEC2CloudInfoWithDefaults() *EC2CloudInfo {
	this := EC2CloudInfo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *EC2CloudInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetAvailabilityZone() string {
	if o == nil || o.AvailabilityZone == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *EC2CloudInfo) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetRamdiskId returns the RamdiskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EC2CloudInfo) GetRamdiskId() string {
	if o == nil || o.RamdiskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RamdiskId.Get()
}

// GetRamdiskIdOk returns a tuple with the RamdiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EC2CloudInfo) GetRamdiskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RamdiskId.Get(), o.RamdiskId.IsSet()
}

// HasRamdiskId returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasRamdiskId() bool {
	if o != nil && o.RamdiskId.IsSet() {
		return true
	}

	return false
}

// SetRamdiskId gets a reference to the given NullableString and assigns it to the RamdiskId field.
func (o *EC2CloudInfo) SetRamdiskId(v string) {
	o.RamdiskId.Set(&v)
}
// SetRamdiskIdNil sets the value for RamdiskId to be an explicit nil
func (o *EC2CloudInfo) SetRamdiskIdNil() {
	o.RamdiskId.Set(nil)
}

// UnsetRamdiskId ensures that no value is present for RamdiskId, not even an explicit nil
func (o *EC2CloudInfo) UnsetRamdiskId() {
	o.RamdiskId.Unset()
}

// GetKernelId returns the KernelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EC2CloudInfo) GetKernelId() string {
	if o == nil || o.KernelId.Get() == nil {
		var ret string
		return ret
	}
	return *o.KernelId.Get()
}

// GetKernelIdOk returns a tuple with the KernelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EC2CloudInfo) GetKernelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KernelId.Get(), o.KernelId.IsSet()
}

// HasKernelId returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasKernelId() bool {
	if o != nil && o.KernelId.IsSet() {
		return true
	}

	return false
}

// SetKernelId gets a reference to the given NullableString and assigns it to the KernelId field.
func (o *EC2CloudInfo) SetKernelId(v string) {
	o.KernelId.Set(&v)
}
// SetKernelIdNil sets the value for KernelId to be an explicit nil
func (o *EC2CloudInfo) SetKernelIdNil() {
	o.KernelId.Set(nil)
}

// UnsetKernelId ensures that no value is present for KernelId, not even an explicit nil
func (o *EC2CloudInfo) UnsetKernelId() {
	o.KernelId.Unset()
}

// GetPendingTime returns the PendingTime field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetPendingTime() string {
	if o == nil || o.PendingTime == nil {
		var ret string
		return ret
	}
	return *o.PendingTime
}

// GetPendingTimeOk returns a tuple with the PendingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetPendingTimeOk() (*string, bool) {
	if o == nil || o.PendingTime == nil {
		return nil, false
	}
	return o.PendingTime, true
}

// HasPendingTime returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasPendingTime() bool {
	if o != nil && o.PendingTime != nil {
		return true
	}

	return false
}

// SetPendingTime gets a reference to the given string and assigns it to the PendingTime field.
func (o *EC2CloudInfo) SetPendingTime(v string) {
	o.PendingTime = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *EC2CloudInfo) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *EC2CloudInfo) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetDevpayProductCodes returns the DevpayProductCodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EC2CloudInfo) GetDevpayProductCodes() string {
	if o == nil || o.DevpayProductCodes.Get() == nil {
		var ret string
		return ret
	}
	return *o.DevpayProductCodes.Get()
}

// GetDevpayProductCodesOk returns a tuple with the DevpayProductCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EC2CloudInfo) GetDevpayProductCodesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevpayProductCodes.Get(), o.DevpayProductCodes.IsSet()
}

// HasDevpayProductCodes returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasDevpayProductCodes() bool {
	if o != nil && o.DevpayProductCodes.IsSet() {
		return true
	}

	return false
}

// SetDevpayProductCodes gets a reference to the given NullableString and assigns it to the DevpayProductCodes field.
func (o *EC2CloudInfo) SetDevpayProductCodes(v string) {
	o.DevpayProductCodes.Set(&v)
}
// SetDevpayProductCodesNil sets the value for DevpayProductCodes to be an explicit nil
func (o *EC2CloudInfo) SetDevpayProductCodesNil() {
	o.DevpayProductCodes.Set(nil)
}

// UnsetDevpayProductCodes ensures that no value is present for DevpayProductCodes, not even an explicit nil
func (o *EC2CloudInfo) UnsetDevpayProductCodes() {
	o.DevpayProductCodes.Unset()
}

// GetMarketplaceProductCodes returns the MarketplaceProductCodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EC2CloudInfo) GetMarketplaceProductCodes() string {
	if o == nil || o.MarketplaceProductCodes.Get() == nil {
		var ret string
		return ret
	}
	return *o.MarketplaceProductCodes.Get()
}

// GetMarketplaceProductCodesOk returns a tuple with the MarketplaceProductCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EC2CloudInfo) GetMarketplaceProductCodesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceProductCodes.Get(), o.MarketplaceProductCodes.IsSet()
}

// HasMarketplaceProductCodes returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasMarketplaceProductCodes() bool {
	if o != nil && o.MarketplaceProductCodes.IsSet() {
		return true
	}

	return false
}

// SetMarketplaceProductCodes gets a reference to the given NullableString and assigns it to the MarketplaceProductCodes field.
func (o *EC2CloudInfo) SetMarketplaceProductCodes(v string) {
	o.MarketplaceProductCodes.Set(&v)
}
// SetMarketplaceProductCodesNil sets the value for MarketplaceProductCodes to be an explicit nil
func (o *EC2CloudInfo) SetMarketplaceProductCodesNil() {
	o.MarketplaceProductCodes.Set(nil)
}

// UnsetMarketplaceProductCodes ensures that no value is present for MarketplaceProductCodes, not even an explicit nil
func (o *EC2CloudInfo) UnsetMarketplaceProductCodes() {
	o.MarketplaceProductCodes.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EC2CloudInfo) SetVersion(v string) {
	o.Version = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EC2CloudInfo) SetRegion(v string) {
	o.Region = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *EC2CloudInfo) SetImageId(v string) {
	o.ImageId = &v
}

// GetBillingProducts returns the BillingProducts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EC2CloudInfo) GetBillingProducts() string {
	if o == nil || o.BillingProducts.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingProducts.Get()
}

// GetBillingProductsOk returns a tuple with the BillingProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EC2CloudInfo) GetBillingProductsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingProducts.Get(), o.BillingProducts.IsSet()
}

// HasBillingProducts returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasBillingProducts() bool {
	if o != nil && o.BillingProducts.IsSet() {
		return true
	}

	return false
}

// SetBillingProducts gets a reference to the given NullableString and assigns it to the BillingProducts field.
func (o *EC2CloudInfo) SetBillingProducts(v string) {
	o.BillingProducts.Set(&v)
}
// SetBillingProductsNil sets the value for BillingProducts to be an explicit nil
func (o *EC2CloudInfo) SetBillingProductsNil() {
	o.BillingProducts.Set(nil)
}

// UnsetBillingProducts ensures that no value is present for BillingProducts, not even an explicit nil
func (o *EC2CloudInfo) UnsetBillingProducts() {
	o.BillingProducts.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *EC2CloudInfo) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *EC2CloudInfo) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2CloudInfo) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *EC2CloudInfo) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *EC2CloudInfo) SetInstanceType(v string) {
	o.InstanceType = &v
}

func (o EC2CloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if o.RamdiskId.IsSet() {
		toSerialize["ramdiskId"] = o.RamdiskId.Get()
	}
	if o.KernelId.IsSet() {
		toSerialize["kernelId"] = o.KernelId.Get()
	}
	if o.PendingTime != nil {
		toSerialize["pendingTime"] = o.PendingTime
	}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.PrivateIp != nil {
		toSerialize["privateIp"] = o.PrivateIp
	}
	if o.DevpayProductCodes.IsSet() {
		toSerialize["devpayProductCodes"] = o.DevpayProductCodes.Get()
	}
	if o.MarketplaceProductCodes.IsSet() {
		toSerialize["marketplaceProductCodes"] = o.MarketplaceProductCodes.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.ImageId != nil {
		toSerialize["imageId"] = o.ImageId
	}
	if o.BillingProducts.IsSet() {
		toSerialize["billingProducts"] = o.BillingProducts.Get()
	}
	if o.InstanceId != nil {
		toSerialize["instanceId"] = o.InstanceId
	}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EC2CloudInfo) UnmarshalJSON(bytes []byte) (err error) {
	varEC2CloudInfo := _EC2CloudInfo{}

	if err = json.Unmarshal(bytes, &varEC2CloudInfo); err == nil {
		*o = EC2CloudInfo(varEC2CloudInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "availabilityZone")
		delete(additionalProperties, "ramdiskId")
		delete(additionalProperties, "kernelId")
		delete(additionalProperties, "pendingTime")
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "privateIp")
		delete(additionalProperties, "devpayProductCodes")
		delete(additionalProperties, "marketplaceProductCodes")
		delete(additionalProperties, "version")
		delete(additionalProperties, "region")
		delete(additionalProperties, "imageId")
		delete(additionalProperties, "billingProducts")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "instanceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEC2CloudInfo struct {
	value *EC2CloudInfo
	isSet bool
}

func (v NullableEC2CloudInfo) Get() *EC2CloudInfo {
	return v.value
}

func (v *NullableEC2CloudInfo) Set(val *EC2CloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEC2CloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEC2CloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEC2CloudInfo(val *EC2CloudInfo) *NullableEC2CloudInfo {
	return &NullableEC2CloudInfo{value: val, isSet: true}
}

func (v NullableEC2CloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEC2CloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


