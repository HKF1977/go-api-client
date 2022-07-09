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

// PutLdapUserSchemaSettingsRequest struct for PutLdapUserSchemaSettingsRequest
type PutLdapUserSchemaSettingsRequest struct {
	// Base DN from which to search for Users
	UserBase string `json:"user_base"`
	// Attribute type for the Users
	UserIdAttribute string `json:"user_id_attribute"`
	// Search filter for Users
	UserListFilter *string `json:"user_list_filter,omitempty"`
}

// NewPutLdapUserSchemaSettingsRequest instantiates a new PutLdapUserSchemaSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutLdapUserSchemaSettingsRequest(userBase string, userIdAttribute string) *PutLdapUserSchemaSettingsRequest {
	this := PutLdapUserSchemaSettingsRequest{}
	this.UserBase = userBase
	this.UserIdAttribute = userIdAttribute
	return &this
}

// NewPutLdapUserSchemaSettingsRequestWithDefaults instantiates a new PutLdapUserSchemaSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutLdapUserSchemaSettingsRequestWithDefaults() *PutLdapUserSchemaSettingsRequest {
	this := PutLdapUserSchemaSettingsRequest{}
	return &this
}

// GetUserBase returns the UserBase field value
func (o *PutLdapUserSchemaSettingsRequest) GetUserBase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserBase
}

// GetUserBaseOk returns a tuple with the UserBase field value
// and a boolean to check if the value has been set.
func (o *PutLdapUserSchemaSettingsRequest) GetUserBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserBase, true
}

// SetUserBase sets field value
func (o *PutLdapUserSchemaSettingsRequest) SetUserBase(v string) {
	o.UserBase = v
}

// GetUserIdAttribute returns the UserIdAttribute field value
func (o *PutLdapUserSchemaSettingsRequest) GetUserIdAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIdAttribute
}

// GetUserIdAttributeOk returns a tuple with the UserIdAttribute field value
// and a boolean to check if the value has been set.
func (o *PutLdapUserSchemaSettingsRequest) GetUserIdAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIdAttribute, true
}

// SetUserIdAttribute sets field value
func (o *PutLdapUserSchemaSettingsRequest) SetUserIdAttribute(v string) {
	o.UserIdAttribute = v
}

// GetUserListFilter returns the UserListFilter field value if set, zero value otherwise.
func (o *PutLdapUserSchemaSettingsRequest) GetUserListFilter() string {
	if o == nil || o.UserListFilter == nil {
		var ret string
		return ret
	}
	return *o.UserListFilter
}

// GetUserListFilterOk returns a tuple with the UserListFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLdapUserSchemaSettingsRequest) GetUserListFilterOk() (*string, bool) {
	if o == nil || o.UserListFilter == nil {
		return nil, false
	}
	return o.UserListFilter, true
}

// HasUserListFilter returns a boolean if a field has been set.
func (o *PutLdapUserSchemaSettingsRequest) HasUserListFilter() bool {
	if o != nil && o.UserListFilter != nil {
		return true
	}

	return false
}

// SetUserListFilter gets a reference to the given string and assigns it to the UserListFilter field.
func (o *PutLdapUserSchemaSettingsRequest) SetUserListFilter(v string) {
	o.UserListFilter = &v
}

func (o PutLdapUserSchemaSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_base"] = o.UserBase
	}
	if true {
		toSerialize["user_id_attribute"] = o.UserIdAttribute
	}
	if o.UserListFilter != nil {
		toSerialize["user_list_filter"] = o.UserListFilter
	}
	return json.Marshal(toSerialize)
}

type NullablePutLdapUserSchemaSettingsRequest struct {
	value *PutLdapUserSchemaSettingsRequest
	isSet bool
}

func (v NullablePutLdapUserSchemaSettingsRequest) Get() *PutLdapUserSchemaSettingsRequest {
	return v.value
}

func (v *NullablePutLdapUserSchemaSettingsRequest) Set(val *PutLdapUserSchemaSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutLdapUserSchemaSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutLdapUserSchemaSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutLdapUserSchemaSettingsRequest(val *PutLdapUserSchemaSettingsRequest) *NullablePutLdapUserSchemaSettingsRequest {
	return &NullablePutLdapUserSchemaSettingsRequest{value: val, isSet: true}
}

func (v NullablePutLdapUserSchemaSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutLdapUserSchemaSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

