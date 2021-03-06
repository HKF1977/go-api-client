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

// PostTestLdapGroupSchemaSettings200ResponseResponseInner struct for PostTestLdapGroupSchemaSettings200ResponseResponseInner
type PostTestLdapGroupSchemaSettings200ResponseResponseInner struct {
	LdapGroup *string `json:"ldap_group,omitempty"`
	LdapUser []string `json:"ldap_user,omitempty"`
}

// NewPostTestLdapGroupSchemaSettings200ResponseResponseInner instantiates a new PostTestLdapGroupSchemaSettings200ResponseResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTestLdapGroupSchemaSettings200ResponseResponseInner() *PostTestLdapGroupSchemaSettings200ResponseResponseInner {
	this := PostTestLdapGroupSchemaSettings200ResponseResponseInner{}
	return &this
}

// NewPostTestLdapGroupSchemaSettings200ResponseResponseInnerWithDefaults instantiates a new PostTestLdapGroupSchemaSettings200ResponseResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTestLdapGroupSchemaSettings200ResponseResponseInnerWithDefaults() *PostTestLdapGroupSchemaSettings200ResponseResponseInner {
	this := PostTestLdapGroupSchemaSettings200ResponseResponseInner{}
	return &this
}

// GetLdapGroup returns the LdapGroup field value if set, zero value otherwise.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) GetLdapGroup() string {
	if o == nil || o.LdapGroup == nil {
		var ret string
		return ret
	}
	return *o.LdapGroup
}

// GetLdapGroupOk returns a tuple with the LdapGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) GetLdapGroupOk() (*string, bool) {
	if o == nil || o.LdapGroup == nil {
		return nil, false
	}
	return o.LdapGroup, true
}

// HasLdapGroup returns a boolean if a field has been set.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) HasLdapGroup() bool {
	if o != nil && o.LdapGroup != nil {
		return true
	}

	return false
}

// SetLdapGroup gets a reference to the given string and assigns it to the LdapGroup field.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) SetLdapGroup(v string) {
	o.LdapGroup = &v
}

// GetLdapUser returns the LdapUser field value if set, zero value otherwise.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) GetLdapUser() []string {
	if o == nil || o.LdapUser == nil {
		var ret []string
		return ret
	}
	return o.LdapUser
}

// GetLdapUserOk returns a tuple with the LdapUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) GetLdapUserOk() ([]string, bool) {
	if o == nil || o.LdapUser == nil {
		return nil, false
	}
	return o.LdapUser, true
}

// HasLdapUser returns a boolean if a field has been set.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) HasLdapUser() bool {
	if o != nil && o.LdapUser != nil {
		return true
	}

	return false
}

// SetLdapUser gets a reference to the given []string and assigns it to the LdapUser field.
func (o *PostTestLdapGroupSchemaSettings200ResponseResponseInner) SetLdapUser(v []string) {
	o.LdapUser = v
}

func (o PostTestLdapGroupSchemaSettings200ResponseResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LdapGroup != nil {
		toSerialize["ldap_group"] = o.LdapGroup
	}
	if o.LdapUser != nil {
		toSerialize["ldap_user"] = o.LdapUser
	}
	return json.Marshal(toSerialize)
}

type NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner struct {
	value *PostTestLdapGroupSchemaSettings200ResponseResponseInner
	isSet bool
}

func (v NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) Get() *PostTestLdapGroupSchemaSettings200ResponseResponseInner {
	return v.value
}

func (v *NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) Set(val *PostTestLdapGroupSchemaSettings200ResponseResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTestLdapGroupSchemaSettings200ResponseResponseInner(val *PostTestLdapGroupSchemaSettings200ResponseResponseInner) *NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner {
	return &NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner{value: val, isSet: true}
}

func (v NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTestLdapGroupSchemaSettings200ResponseResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


