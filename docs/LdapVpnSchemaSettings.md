# LdapVpnSchemaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnAuthEnabled** | Pointer to **bool** |  | [optional] 
**VpnAuthProvider** | Pointer to **string** |  | [optional] 
**VpnGroupBase** | Pointer to **string** | Base DN from which to search for Groups | [optional] 
**VpnGroupIdAttribute** | Pointer to **string** | Attribute type for the Groups | [optional] 
**VpnGroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**VpnGroupMemberAttribute** | Pointer to **string** | Attribute used to search for a user within the Group | [optional] 
**VpnGroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "UserDN"]
**VpnGroupSearchScope** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "subtree"]
**VpnGroupOtp** | Pointer to **bool** |  | [optional] 
**VpnGroupOtpUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewLdapVpnSchemaSettings

`func NewLdapVpnSchemaSettings() *LdapVpnSchemaSettings`

NewLdapVpnSchemaSettings instantiates a new LdapVpnSchemaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapVpnSchemaSettingsWithDefaults

`func NewLdapVpnSchemaSettingsWithDefaults() *LdapVpnSchemaSettings`

NewLdapVpnSchemaSettingsWithDefaults instantiates a new LdapVpnSchemaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnAuthEnabled

`func (o *LdapVpnSchemaSettings) GetVpnAuthEnabled() bool`

GetVpnAuthEnabled returns the VpnAuthEnabled field if non-nil, zero value otherwise.

### GetVpnAuthEnabledOk

`func (o *LdapVpnSchemaSettings) GetVpnAuthEnabledOk() (*bool, bool)`

GetVpnAuthEnabledOk returns a tuple with the VpnAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthEnabled

`func (o *LdapVpnSchemaSettings) SetVpnAuthEnabled(v bool)`

SetVpnAuthEnabled sets VpnAuthEnabled field to given value.

### HasVpnAuthEnabled

`func (o *LdapVpnSchemaSettings) HasVpnAuthEnabled() bool`

HasVpnAuthEnabled returns a boolean if a field has been set.

### GetVpnAuthProvider

`func (o *LdapVpnSchemaSettings) GetVpnAuthProvider() string`

GetVpnAuthProvider returns the VpnAuthProvider field if non-nil, zero value otherwise.

### GetVpnAuthProviderOk

`func (o *LdapVpnSchemaSettings) GetVpnAuthProviderOk() (*string, bool)`

GetVpnAuthProviderOk returns a tuple with the VpnAuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthProvider

`func (o *LdapVpnSchemaSettings) SetVpnAuthProvider(v string)`

SetVpnAuthProvider sets VpnAuthProvider field to given value.

### HasVpnAuthProvider

`func (o *LdapVpnSchemaSettings) HasVpnAuthProvider() bool`

HasVpnAuthProvider returns a boolean if a field has been set.

### GetVpnGroupBase

`func (o *LdapVpnSchemaSettings) GetVpnGroupBase() string`

GetVpnGroupBase returns the VpnGroupBase field if non-nil, zero value otherwise.

### GetVpnGroupBaseOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupBaseOk() (*string, bool)`

GetVpnGroupBaseOk returns a tuple with the VpnGroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupBase

`func (o *LdapVpnSchemaSettings) SetVpnGroupBase(v string)`

SetVpnGroupBase sets VpnGroupBase field to given value.

### HasVpnGroupBase

`func (o *LdapVpnSchemaSettings) HasVpnGroupBase() bool`

HasVpnGroupBase returns a boolean if a field has been set.

### GetVpnGroupIdAttribute

`func (o *LdapVpnSchemaSettings) GetVpnGroupIdAttribute() string`

GetVpnGroupIdAttribute returns the VpnGroupIdAttribute field if non-nil, zero value otherwise.

### GetVpnGroupIdAttributeOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupIdAttributeOk() (*string, bool)`

GetVpnGroupIdAttributeOk returns a tuple with the VpnGroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupIdAttribute

`func (o *LdapVpnSchemaSettings) SetVpnGroupIdAttribute(v string)`

SetVpnGroupIdAttribute sets VpnGroupIdAttribute field to given value.

### HasVpnGroupIdAttribute

`func (o *LdapVpnSchemaSettings) HasVpnGroupIdAttribute() bool`

HasVpnGroupIdAttribute returns a boolean if a field has been set.

### GetVpnGroupListFilter

`func (o *LdapVpnSchemaSettings) GetVpnGroupListFilter() string`

GetVpnGroupListFilter returns the VpnGroupListFilter field if non-nil, zero value otherwise.

### GetVpnGroupListFilterOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupListFilterOk() (*string, bool)`

GetVpnGroupListFilterOk returns a tuple with the VpnGroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupListFilter

`func (o *LdapVpnSchemaSettings) SetVpnGroupListFilter(v string)`

SetVpnGroupListFilter sets VpnGroupListFilter field to given value.

### HasVpnGroupListFilter

`func (o *LdapVpnSchemaSettings) HasVpnGroupListFilter() bool`

HasVpnGroupListFilter returns a boolean if a field has been set.

### GetVpnGroupMemberAttribute

`func (o *LdapVpnSchemaSettings) GetVpnGroupMemberAttribute() string`

GetVpnGroupMemberAttribute returns the VpnGroupMemberAttribute field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttributeOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupMemberAttributeOk() (*string, bool)`

GetVpnGroupMemberAttributeOk returns a tuple with the VpnGroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttribute

`func (o *LdapVpnSchemaSettings) SetVpnGroupMemberAttribute(v string)`

SetVpnGroupMemberAttribute sets VpnGroupMemberAttribute field to given value.

### HasVpnGroupMemberAttribute

`func (o *LdapVpnSchemaSettings) HasVpnGroupMemberAttribute() bool`

HasVpnGroupMemberAttribute returns a boolean if a field has been set.

### GetVpnGroupMemberAttrFormat

`func (o *LdapVpnSchemaSettings) GetVpnGroupMemberAttrFormat() string`

GetVpnGroupMemberAttrFormat returns the VpnGroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttrFormatOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupMemberAttrFormatOk() (*string, bool)`

GetVpnGroupMemberAttrFormatOk returns a tuple with the VpnGroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttrFormat

`func (o *LdapVpnSchemaSettings) SetVpnGroupMemberAttrFormat(v string)`

SetVpnGroupMemberAttrFormat sets VpnGroupMemberAttrFormat field to given value.

### HasVpnGroupMemberAttrFormat

`func (o *LdapVpnSchemaSettings) HasVpnGroupMemberAttrFormat() bool`

HasVpnGroupMemberAttrFormat returns a boolean if a field has been set.

### GetVpnGroupSearchScope

`func (o *LdapVpnSchemaSettings) GetVpnGroupSearchScope() string`

GetVpnGroupSearchScope returns the VpnGroupSearchScope field if non-nil, zero value otherwise.

### GetVpnGroupSearchScopeOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupSearchScopeOk() (*string, bool)`

GetVpnGroupSearchScopeOk returns a tuple with the VpnGroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupSearchScope

`func (o *LdapVpnSchemaSettings) SetVpnGroupSearchScope(v string)`

SetVpnGroupSearchScope sets VpnGroupSearchScope field to given value.

### HasVpnGroupSearchScope

`func (o *LdapVpnSchemaSettings) HasVpnGroupSearchScope() bool`

HasVpnGroupSearchScope returns a boolean if a field has been set.

### GetVpnGroupOtp

`func (o *LdapVpnSchemaSettings) GetVpnGroupOtp() bool`

GetVpnGroupOtp returns the VpnGroupOtp field if non-nil, zero value otherwise.

### GetVpnGroupOtpOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupOtpOk() (*bool, bool)`

GetVpnGroupOtpOk returns a tuple with the VpnGroupOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupOtp

`func (o *LdapVpnSchemaSettings) SetVpnGroupOtp(v bool)`

SetVpnGroupOtp sets VpnGroupOtp field to given value.

### HasVpnGroupOtp

`func (o *LdapVpnSchemaSettings) HasVpnGroupOtp() bool`

HasVpnGroupOtp returns a boolean if a field has been set.

### GetVpnGroupOtpUrl

`func (o *LdapVpnSchemaSettings) GetVpnGroupOtpUrl() string`

GetVpnGroupOtpUrl returns the VpnGroupOtpUrl field if non-nil, zero value otherwise.

### GetVpnGroupOtpUrlOk

`func (o *LdapVpnSchemaSettings) GetVpnGroupOtpUrlOk() (*string, bool)`

GetVpnGroupOtpUrlOk returns a tuple with the VpnGroupOtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupOtpUrl

`func (o *LdapVpnSchemaSettings) SetVpnGroupOtpUrl(v string)`

SetVpnGroupOtpUrl sets VpnGroupOtpUrl field to given value.

### HasVpnGroupOtpUrl

`func (o *LdapVpnSchemaSettings) HasVpnGroupOtpUrl() bool`

HasVpnGroupOtpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


