# PutLdapGroupSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupRequired** | **bool** | Require use of LDAP groups | 
**GroupBase** | Pointer to **string** | Base DN from which to search for Groups | [optional] 
**GroupIdAttribute** | Pointer to **string** | Attribute type for the Groups | [optional] 
**GroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**GroupMemberAttribute** | Pointer to **string** | ttribute used to search for a user within the Group | [optional] 
**GroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] 
**GroupSearchScope** | Pointer to **string** | Default&#x3D;subtree | [optional] 

## Methods

### NewPutLdapGroupSchemaSettingsRequest

`func NewPutLdapGroupSchemaSettingsRequest(groupRequired bool, ) *PutLdapGroupSchemaSettingsRequest`

NewPutLdapGroupSchemaSettingsRequest instantiates a new PutLdapGroupSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLdapGroupSchemaSettingsRequestWithDefaults

`func NewPutLdapGroupSchemaSettingsRequestWithDefaults() *PutLdapGroupSchemaSettingsRequest`

NewPutLdapGroupSchemaSettingsRequestWithDefaults instantiates a new PutLdapGroupSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupRequired

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupRequired() bool`

GetGroupRequired returns the GroupRequired field if non-nil, zero value otherwise.

### GetGroupRequiredOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupRequiredOk() (*bool, bool)`

GetGroupRequiredOk returns a tuple with the GroupRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRequired

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupRequired(v bool)`

SetGroupRequired sets GroupRequired field to given value.


### GetGroupBase

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *PutLdapGroupSchemaSettingsRequest) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *PutLdapGroupSchemaSettingsRequest) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *PutLdapGroupSchemaSettingsRequest) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


