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

// LdapGroupSettings struct for LdapGroupSettings
type LdapGroupSettings struct {
	// Require used of LDAP groups
	GroupRequired *bool `json:"group_required,omitempty"`
	// Base DN from which to search for Groups
	GroupBase *string `json:"group_base,omitempty"`
	// Attribute type for the Groups
	GroupIdAttribute *string `json:"group_id_attribute,omitempty"`
	// Search filter for Groups
	GroupListFilter *string `json:"group_list_filter,omitempty"`
	// Attribute used to search for a user within the Group
	GroupMemberAttribute *string `json:"group_member_attribute,omitempty"`
	// Format of the Group Member attribute
	GroupMemberAttrFormat *string `json:"group_member_attr_format,omitempty"`
	// Format of the Group Member attribute
	GroupSearchScope *string `json:"group_search_scope,omitempty"`
}

// NewLdapGroupSettings instantiates a new LdapGroupSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapGroupSettings() *LdapGroupSettings {
	this := LdapGroupSettings{}
	var groupMemberAttrFormat string = "UserDN"
	this.GroupMemberAttrFormat = &groupMemberAttrFormat
	var groupSearchScope string = "subtree"
	this.GroupSearchScope = &groupSearchScope
	return &this
}

// NewLdapGroupSettingsWithDefaults instantiates a new LdapGroupSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapGroupSettingsWithDefaults() *LdapGroupSettings {
	this := LdapGroupSettings{}
	var groupMemberAttrFormat string = "UserDN"
	this.GroupMemberAttrFormat = &groupMemberAttrFormat
	var groupSearchScope string = "subtree"
	this.GroupSearchScope = &groupSearchScope
	return &this
}

// GetGroupRequired returns the GroupRequired field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupRequired() bool {
	if o == nil || o.GroupRequired == nil {
		var ret bool
		return ret
	}
	return *o.GroupRequired
}

// GetGroupRequiredOk returns a tuple with the GroupRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupRequiredOk() (*bool, bool) {
	if o == nil || o.GroupRequired == nil {
		return nil, false
	}
	return o.GroupRequired, true
}

// HasGroupRequired returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupRequired() bool {
	if o != nil && o.GroupRequired != nil {
		return true
	}

	return false
}

// SetGroupRequired gets a reference to the given bool and assigns it to the GroupRequired field.
func (o *LdapGroupSettings) SetGroupRequired(v bool) {
	o.GroupRequired = &v
}

// GetGroupBase returns the GroupBase field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupBase() string {
	if o == nil || o.GroupBase == nil {
		var ret string
		return ret
	}
	return *o.GroupBase
}

// GetGroupBaseOk returns a tuple with the GroupBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupBaseOk() (*string, bool) {
	if o == nil || o.GroupBase == nil {
		return nil, false
	}
	return o.GroupBase, true
}

// HasGroupBase returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupBase() bool {
	if o != nil && o.GroupBase != nil {
		return true
	}

	return false
}

// SetGroupBase gets a reference to the given string and assigns it to the GroupBase field.
func (o *LdapGroupSettings) SetGroupBase(v string) {
	o.GroupBase = &v
}

// GetGroupIdAttribute returns the GroupIdAttribute field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupIdAttribute() string {
	if o == nil || o.GroupIdAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupIdAttribute
}

// GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupIdAttributeOk() (*string, bool) {
	if o == nil || o.GroupIdAttribute == nil {
		return nil, false
	}
	return o.GroupIdAttribute, true
}

// HasGroupIdAttribute returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupIdAttribute() bool {
	if o != nil && o.GroupIdAttribute != nil {
		return true
	}

	return false
}

// SetGroupIdAttribute gets a reference to the given string and assigns it to the GroupIdAttribute field.
func (o *LdapGroupSettings) SetGroupIdAttribute(v string) {
	o.GroupIdAttribute = &v
}

// GetGroupListFilter returns the GroupListFilter field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupListFilter() string {
	if o == nil || o.GroupListFilter == nil {
		var ret string
		return ret
	}
	return *o.GroupListFilter
}

// GetGroupListFilterOk returns a tuple with the GroupListFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupListFilterOk() (*string, bool) {
	if o == nil || o.GroupListFilter == nil {
		return nil, false
	}
	return o.GroupListFilter, true
}

// HasGroupListFilter returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupListFilter() bool {
	if o != nil && o.GroupListFilter != nil {
		return true
	}

	return false
}

// SetGroupListFilter gets a reference to the given string and assigns it to the GroupListFilter field.
func (o *LdapGroupSettings) SetGroupListFilter(v string) {
	o.GroupListFilter = &v
}

// GetGroupMemberAttribute returns the GroupMemberAttribute field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupMemberAttribute() string {
	if o == nil || o.GroupMemberAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupMemberAttribute
}

// GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupMemberAttributeOk() (*string, bool) {
	if o == nil || o.GroupMemberAttribute == nil {
		return nil, false
	}
	return o.GroupMemberAttribute, true
}

// HasGroupMemberAttribute returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupMemberAttribute() bool {
	if o != nil && o.GroupMemberAttribute != nil {
		return true
	}

	return false
}

// SetGroupMemberAttribute gets a reference to the given string and assigns it to the GroupMemberAttribute field.
func (o *LdapGroupSettings) SetGroupMemberAttribute(v string) {
	o.GroupMemberAttribute = &v
}

// GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupMemberAttrFormat() string {
	if o == nil || o.GroupMemberAttrFormat == nil {
		var ret string
		return ret
	}
	return *o.GroupMemberAttrFormat
}

// GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupMemberAttrFormatOk() (*string, bool) {
	if o == nil || o.GroupMemberAttrFormat == nil {
		return nil, false
	}
	return o.GroupMemberAttrFormat, true
}

// HasGroupMemberAttrFormat returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupMemberAttrFormat() bool {
	if o != nil && o.GroupMemberAttrFormat != nil {
		return true
	}

	return false
}

// SetGroupMemberAttrFormat gets a reference to the given string and assigns it to the GroupMemberAttrFormat field.
func (o *LdapGroupSettings) SetGroupMemberAttrFormat(v string) {
	o.GroupMemberAttrFormat = &v
}

// GetGroupSearchScope returns the GroupSearchScope field value if set, zero value otherwise.
func (o *LdapGroupSettings) GetGroupSearchScope() string {
	if o == nil || o.GroupSearchScope == nil {
		var ret string
		return ret
	}
	return *o.GroupSearchScope
}

// GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapGroupSettings) GetGroupSearchScopeOk() (*string, bool) {
	if o == nil || o.GroupSearchScope == nil {
		return nil, false
	}
	return o.GroupSearchScope, true
}

// HasGroupSearchScope returns a boolean if a field has been set.
func (o *LdapGroupSettings) HasGroupSearchScope() bool {
	if o != nil && o.GroupSearchScope != nil {
		return true
	}

	return false
}

// SetGroupSearchScope gets a reference to the given string and assigns it to the GroupSearchScope field.
func (o *LdapGroupSettings) SetGroupSearchScope(v string) {
	o.GroupSearchScope = &v
}

func (o LdapGroupSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupRequired != nil {
		toSerialize["group_required"] = o.GroupRequired
	}
	if o.GroupBase != nil {
		toSerialize["group_base"] = o.GroupBase
	}
	if o.GroupIdAttribute != nil {
		toSerialize["group_id_attribute"] = o.GroupIdAttribute
	}
	if o.GroupListFilter != nil {
		toSerialize["group_list_filter"] = o.GroupListFilter
	}
	if o.GroupMemberAttribute != nil {
		toSerialize["group_member_attribute"] = o.GroupMemberAttribute
	}
	if o.GroupMemberAttrFormat != nil {
		toSerialize["group_member_attr_format"] = o.GroupMemberAttrFormat
	}
	if o.GroupSearchScope != nil {
		toSerialize["group_search_scope"] = o.GroupSearchScope
	}
	return json.Marshal(toSerialize)
}

type NullableLdapGroupSettings struct {
	value *LdapGroupSettings
	isSet bool
}

func (v NullableLdapGroupSettings) Get() *LdapGroupSettings {
	return v.value
}

func (v *NullableLdapGroupSettings) Set(val *LdapGroupSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapGroupSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapGroupSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapGroupSettings(val *LdapGroupSettings) *NullableLdapGroupSettings {
	return &NullableLdapGroupSettings{value: val, isSet: true}
}

func (v NullableLdapGroupSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapGroupSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

