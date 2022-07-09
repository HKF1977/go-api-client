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

// SnapshotImportStatusResponse struct for SnapshotImportStatusResponse
type SnapshotImportStatusResponse struct {
	Response *SnapshotImportStatus `json:"response,omitempty"`
}

// NewSnapshotImportStatusResponse instantiates a new SnapshotImportStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotImportStatusResponse() *SnapshotImportStatusResponse {
	this := SnapshotImportStatusResponse{}
	return &this
}

// NewSnapshotImportStatusResponseWithDefaults instantiates a new SnapshotImportStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotImportStatusResponseWithDefaults() *SnapshotImportStatusResponse {
	this := SnapshotImportStatusResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SnapshotImportStatusResponse) GetResponse() SnapshotImportStatus {
	if o == nil || o.Response == nil {
		var ret SnapshotImportStatus
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotImportStatusResponse) GetResponseOk() (*SnapshotImportStatus, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SnapshotImportStatusResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given SnapshotImportStatus and assigns it to the Response field.
func (o *SnapshotImportStatusResponse) SetResponse(v SnapshotImportStatus) {
	o.Response = &v
}

func (o SnapshotImportStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotImportStatusResponse struct {
	value *SnapshotImportStatusResponse
	isSet bool
}

func (v NullableSnapshotImportStatusResponse) Get() *SnapshotImportStatusResponse {
	return v.value
}

func (v *NullableSnapshotImportStatusResponse) Set(val *SnapshotImportStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotImportStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotImportStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotImportStatusResponse(val *SnapshotImportStatusResponse) *NullableSnapshotImportStatusResponse {
	return &NullableSnapshotImportStatusResponse{value: val, isSet: true}
}

func (v NullableSnapshotImportStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotImportStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

