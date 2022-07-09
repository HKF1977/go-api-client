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
	"time"
)

// PluginManagerConfigConfigFile struct for PluginManagerConfigConfigFile
type PluginManagerConfigConfigFile struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *int32 `json:"version,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	PreviousVersions []PluginManagerConfigConfigFilePreviousVersionsInner `json:"previous_versions,omitempty"`
}

// NewPluginManagerConfigConfigFile instantiates a new PluginManagerConfigConfigFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerConfigConfigFile() *PluginManagerConfigConfigFile {
	this := PluginManagerConfigConfigFile{}
	return &this
}

// NewPluginManagerConfigConfigFileWithDefaults instantiates a new PluginManagerConfigConfigFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerConfigConfigFileWithDefaults() *PluginManagerConfigConfigFile {
	this := PluginManagerConfigConfigFile{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginManagerConfigConfigFile) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PluginManagerConfigConfigFile) SetPath(v string) {
	o.Path = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluginManagerConfigConfigFile) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *PluginManagerConfigConfigFile) SetVersion(v int32) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PluginManagerConfigConfigFile) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPreviousVersions returns the PreviousVersions field value if set, zero value otherwise.
func (o *PluginManagerConfigConfigFile) GetPreviousVersions() []PluginManagerConfigConfigFilePreviousVersionsInner {
	if o == nil || o.PreviousVersions == nil {
		var ret []PluginManagerConfigConfigFilePreviousVersionsInner
		return ret
	}
	return o.PreviousVersions
}

// GetPreviousVersionsOk returns a tuple with the PreviousVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConfigConfigFile) GetPreviousVersionsOk() ([]PluginManagerConfigConfigFilePreviousVersionsInner, bool) {
	if o == nil || o.PreviousVersions == nil {
		return nil, false
	}
	return o.PreviousVersions, true
}

// HasPreviousVersions returns a boolean if a field has been set.
func (o *PluginManagerConfigConfigFile) HasPreviousVersions() bool {
	if o != nil && o.PreviousVersions != nil {
		return true
	}

	return false
}

// SetPreviousVersions gets a reference to the given []PluginManagerConfigConfigFilePreviousVersionsInner and assigns it to the PreviousVersions field.
func (o *PluginManagerConfigConfigFile) SetPreviousVersions(v []PluginManagerConfigConfigFilePreviousVersionsInner) {
	o.PreviousVersions = v
}

func (o PluginManagerConfigConfigFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.PreviousVersions != nil {
		toSerialize["previous_versions"] = o.PreviousVersions
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerConfigConfigFile struct {
	value *PluginManagerConfigConfigFile
	isSet bool
}

func (v NullablePluginManagerConfigConfigFile) Get() *PluginManagerConfigConfigFile {
	return v.value
}

func (v *NullablePluginManagerConfigConfigFile) Set(val *PluginManagerConfigConfigFile) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerConfigConfigFile) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerConfigConfigFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerConfigConfigFile(val *PluginManagerConfigConfigFile) *NullablePluginManagerConfigConfigFile {
	return &NullablePluginManagerConfigConfigFile{value: val, isSet: true}
}

func (v NullablePluginManagerConfigConfigFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerConfigConfigFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


