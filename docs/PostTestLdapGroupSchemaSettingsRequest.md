# PostTestLdapGroupSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBase** | **string** | Base DN from which to search for Groups | 
**GroupIdAttribute** | **string** | Attribute type for the Groups | 
**GroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**GroupMemberAttribute** | Pointer to **string** | ttribute used to search for a user within the Group | [optional] 
**GroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] 
**GroupSearchScope** | Pointer to **string** | Default&#x3D;subtree | [optional] 
**Limit** | Pointer to **int32** | Number of records to return. Default &#x3D; 100 | [optional] [default to 100]

## Methods

### NewPostTestLdapGroupSchemaSettingsRequest

`func NewPostTestLdapGroupSchemaSettingsRequest(groupBase string, groupIdAttribute string, ) *PostTestLdapGroupSchemaSettingsRequest`

NewPostTestLdapGroupSchemaSettingsRequest instantiates a new PostTestLdapGroupSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTestLdapGroupSchemaSettingsRequestWithDefaults

`func NewPostTestLdapGroupSchemaSettingsRequestWithDefaults() *PostTestLdapGroupSchemaSettingsRequest`

NewPostTestLdapGroupSchemaSettingsRequestWithDefaults instantiates a new PostTestLdapGroupSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBase

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.


### GetGroupIdAttribute

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.


### GetGroupListFilter

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *PostTestLdapGroupSchemaSettingsRequest) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *PostTestLdapGroupSchemaSettingsRequest) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *PostTestLdapGroupSchemaSettingsRequest) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *PostTestLdapGroupSchemaSettingsRequest) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.

### GetLimit

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostTestLdapGroupSchemaSettingsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostTestLdapGroupSchemaSettingsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostTestLdapGroupSchemaSettingsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


