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

// CreateImageDetail struct for CreateImageDetail
type CreateImageDetail struct {
	Status *string `json:"status,omitempty"`
	ImportUuid *string `json:"import_uuid,omitempty"`
}

// NewCreateImageDetail instantiates a new CreateImageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageDetail() *CreateImageDetail {
	this := CreateImageDetail{}
	return &this
}

// NewCreateImageDetailWithDefaults instantiates a new CreateImageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageDetailWithDefaults() *CreateImageDetail {
	this := CreateImageDetail{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateImageDetail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageDetail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateImageDetail) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateImageDetail) SetStatus(v string) {
	o.Status = &v
}

// GetImportUuid returns the ImportUuid field value if set, zero value otherwise.
func (o *CreateImageDetail) GetImportUuid() string {
	if o == nil || o.ImportUuid == nil {
		var ret string
		return ret
	}
	return *o.ImportUuid
}

// GetImportUuidOk returns a tuple with the ImportUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageDetail) GetImportUuidOk() (*string, bool) {
	if o == nil || o.ImportUuid == nil {
		return nil, false
	}
	return o.ImportUuid, true
}

// HasImportUuid returns a boolean if a field has been set.
func (o *CreateImageDetail) HasImportUuid() bool {
	if o != nil && o.ImportUuid != nil {
		return true
	}

	return false
}

// SetImportUuid gets a reference to the given string and assigns it to the ImportUuid field.
func (o *CreateImageDetail) SetImportUuid(v string) {
	o.ImportUuid = &v
}

func (o CreateImageDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ImportUuid != nil {
		toSerialize["import_uuid"] = o.ImportUuid
	}
	return json.Marshal(toSerialize)
}

type NullableCreateImageDetail struct {
	value *CreateImageDetail
	isSet bool
}

func (v NullableCreateImageDetail) Get() *CreateImageDetail {
	return v.value
}

func (v *NullableCreateImageDetail) Set(val *CreateImageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageDetail(val *CreateImageDetail) *NullableCreateImageDetail {
	return &NullableCreateImageDetail{value: val, isSet: true}
}

func (v NullableCreateImageDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

