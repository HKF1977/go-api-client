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

// GCECloudInfo Metadata returned from GCE metadata call.
type GCECloudInfo struct {
	ProjectId *string `json:"projectId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GCECloudInfo GCECloudInfo

// NewGCECloudInfo instantiates a new GCECloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCECloudInfo() *GCECloudInfo {
	this := GCECloudInfo{}
	return &this
}

// NewGCECloudInfoWithDefaults instantiates a new GCECloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCECloudInfoWithDefaults() *GCECloudInfo {
	this := GCECloudInfo{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GCECloudInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCECloudInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GCECloudInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GCECloudInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o GCECloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GCECloudInfo) UnmarshalJSON(bytes []byte) (err error) {
	varGCECloudInfo := _GCECloudInfo{}

	if err = json.Unmarshal(bytes, &varGCECloudInfo); err == nil {
		*o = GCECloudInfo(varGCECloudInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGCECloudInfo struct {
	value *GCECloudInfo
	isSet bool
}

func (v NullableGCECloudInfo) Get() *GCECloudInfo {
	return v.value
}

func (v *NullableGCECloudInfo) Set(val *GCECloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGCECloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGCECloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCECloudInfo(val *GCECloudInfo) *NullableGCECloudInfo {
	return &NullableGCECloudInfo{value: val, isSet: true}
}

func (v NullableGCECloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCECloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


