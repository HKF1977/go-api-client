# LdapGroupSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupRequired** | Pointer to **bool** | Require used of LDAP groups | [optional] 
**GroupBase** | Pointer to **string** | Base DN from which to search for Groups | [optional] 
**GroupIdAttribute** | Pointer to **string** | Attribute type for the Groups | [optional] 
**GroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**GroupMemberAttribute** | Pointer to **string** | Attribute used to search for a user within the Group | [optional] 
**GroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "UserDN"]
**GroupSearchScope** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "subtree"]

## Methods

### NewLdapGroupSettings

`func NewLdapGroupSettings() *LdapGroupSettings`

NewLdapGroupSettings instantiates a new LdapGroupSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapGroupSettingsWithDefaults

`func NewLdapGroupSettingsWithDefaults() *LdapGroupSettings`

NewLdapGroupSettingsWithDefaults instantiates a new LdapGroupSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupRequired

`func (o *LdapGroupSettings) GetGroupRequired() bool`

GetGroupRequired returns the GroupRequired field if non-nil, zero value otherwise.

### GetGroupRequiredOk

`func (o *LdapGroupSettings) GetGroupRequiredOk() (*bool, bool)`

GetGroupRequiredOk returns a tuple with the GroupRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRequired

`func (o *LdapGroupSettings) SetGroupRequired(v bool)`

SetGroupRequired sets GroupRequired field to given value.

### HasGroupRequired

`func (o *LdapGroupSettings) HasGroupRequired() bool`

HasGroupRequired returns a boolean if a field has been set.

### GetGroupBase

`func (o *LdapGroupSettings) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *LdapGroupSettings) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *LdapGroupSettings) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *LdapGroupSettings) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *LdapGroupSettings) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *LdapGroupSettings) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *LdapGroupSettings) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *LdapGroupSettings) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *LdapGroupSettings) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *LdapGroupSettings) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *LdapGroupSettings) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *LdapGroupSettings) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *LdapGroupSettings) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *LdapGroupSettings) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *LdapGroupSettings) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *LdapGroupSettings) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *LdapGroupSettings) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *LdapGroupSettings) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *LdapGroupSettings) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *LdapGroupSettings) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *LdapGroupSettings) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *LdapGroupSettings) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *LdapGroupSettings) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *LdapGroupSettings) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


