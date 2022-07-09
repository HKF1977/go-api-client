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

// SnapshotsListSnapshotsValue struct for SnapshotsListSnapshotsValue
type SnapshotsListSnapshotsValue struct {
	Sha1Checksum *string `json:"sha1_checksum,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedAtI *int32 `json:"created_at_i,omitempty"`
	Size *int32 `json:"size,omitempty"`
	// Status of snapshot, pending, finished_ok, finished_failed or invalid
	Status *string `json:"status,omitempty"`
	// Token if snapshot task is still pending for polling
	Token *string `json:"token,omitempty"`
}

// NewSnapshotsListSnapshotsValue instantiates a new SnapshotsListSnapshotsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotsListSnapshotsValue() *SnapshotsListSnapshotsValue {
	this := SnapshotsListSnapshotsValue{}
	return &this
}

// NewSnapshotsListSnapshotsValueWithDefaults instantiates a new SnapshotsListSnapshotsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotsListSnapshotsValueWithDefaults() *SnapshotsListSnapshotsValue {
	this := SnapshotsListSnapshotsValue{}
	return &this
}

// GetSha1Checksum returns the Sha1Checksum field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetSha1Checksum() string {
	if o == nil || o.Sha1Checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha1Checksum
}

// GetSha1ChecksumOk returns a tuple with the Sha1Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetSha1ChecksumOk() (*string, bool) {
	if o == nil || o.Sha1Checksum == nil {
		return nil, false
	}
	return o.Sha1Checksum, true
}

// HasSha1Checksum returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasSha1Checksum() bool {
	if o != nil && o.Sha1Checksum != nil {
		return true
	}

	return false
}

// SetSha1Checksum gets a reference to the given string and assigns it to the Sha1Checksum field.
func (o *SnapshotsListSnapshotsValue) SetSha1Checksum(v string) {
	o.Sha1Checksum = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SnapshotsListSnapshotsValue) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedAtI returns the CreatedAtI field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetCreatedAtI() int32 {
	if o == nil || o.CreatedAtI == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAtI
}

// GetCreatedAtIOk returns a tuple with the CreatedAtI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetCreatedAtIOk() (*int32, bool) {
	if o == nil || o.CreatedAtI == nil {
		return nil, false
	}
	return o.CreatedAtI, true
}

// HasCreatedAtI returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasCreatedAtI() bool {
	if o != nil && o.CreatedAtI != nil {
		return true
	}

	return false
}

// SetCreatedAtI gets a reference to the given int32 and assigns it to the CreatedAtI field.
func (o *SnapshotsListSnapshotsValue) SetCreatedAtI(v int32) {
	o.CreatedAtI = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *SnapshotsListSnapshotsValue) SetSize(v int32) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SnapshotsListSnapshotsValue) SetStatus(v string) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SnapshotsListSnapshotsValue) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsListSnapshotsValue) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SnapshotsListSnapshotsValue) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SnapshotsListSnapshotsValue) SetToken(v string) {
	o.Token = &v
}

func (o SnapshotsListSnapshotsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sha1Checksum != nil {
		toSerialize["sha1_checksum"] = o.Sha1Checksum
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedAtI != nil {
		toSerialize["created_at_i"] = o.CreatedAtI
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotsListSnapshotsValue struct {
	value *SnapshotsListSnapshotsValue
	isSet bool
}

func (v NullableSnapshotsListSnapshotsValue) Get() *SnapshotsListSnapshotsValue {
	return v.value
}

func (v *NullableSnapshotsListSnapshotsValue) Set(val *SnapshotsListSnapshotsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotsListSnapshotsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotsListSnapshotsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotsListSnapshotsValue(val *SnapshotsListSnapshotsValue) *NullableSnapshotsListSnapshotsValue {
	return &NullableSnapshotsListSnapshotsValue{value: val, isSet: true}
}

func (v NullableSnapshotsListSnapshotsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotsListSnapshotsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


