# PutLdapVpnSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnAuthEnabled** | **bool** | Enable use of LDAP through VPN. If true, other params required. | 
**VpnGroupBase** | **string** | Base DN from which to search for Groups | 
**VpnGroupIdAttribute** | **string** | Attribute type for the Groups | 
**VpnGroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**VpnGroupMemberAttribute** | **string** | Attribute used to search for a user within the Group | 
**VpnGroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "UserDN"]
**VpnGroupSearchScope** | Pointer to **string** | Search scope for filter | [optional] [default to "subtree"]
**VpnGroupOtp** | Pointer to **bool** | Use Google authenticator (OTP)? | [optional] [default to false]

## Methods

### NewPutLdapVpnSchemaSettingsRequest

`func NewPutLdapVpnSchemaSettingsRequest(vpnAuthEnabled bool, vpnGroupBase string, vpnGroupIdAttribute string, vpnGroupMemberAttribute string, ) *PutLdapVpnSchemaSettingsRequest`

NewPutLdapVpnSchemaSettingsRequest instantiates a new PutLdapVpnSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLdapVpnSchemaSettingsRequestWithDefaults

`func NewPutLdapVpnSchemaSettingsRequestWithDefaults() *PutLdapVpnSchemaSettingsRequest`

NewPutLdapVpnSchemaSettingsRequestWithDefaults instantiates a new PutLdapVpnSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnAuthEnabled

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnAuthEnabled() bool`

GetVpnAuthEnabled returns the VpnAuthEnabled field if non-nil, zero value otherwise.

### GetVpnAuthEnabledOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnAuthEnabledOk() (*bool, bool)`

GetVpnAuthEnabledOk returns a tuple with the VpnAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthEnabled

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnAuthEnabled(v bool)`

SetVpnAuthEnabled sets VpnAuthEnabled field to given value.


### GetVpnGroupBase

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupBase() string`

GetVpnGroupBase returns the VpnGroupBase field if non-nil, zero value otherwise.

### GetVpnGroupBaseOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupBaseOk() (*string, bool)`

GetVpnGroupBaseOk returns a tuple with the VpnGroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupBase

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupBase(v string)`

SetVpnGroupBase sets VpnGroupBase field to given value.


### GetVpnGroupIdAttribute

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupIdAttribute() string`

GetVpnGroupIdAttribute returns the VpnGroupIdAttribute field if non-nil, zero value otherwise.

### GetVpnGroupIdAttributeOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupIdAttributeOk() (*string, bool)`

GetVpnGroupIdAttributeOk returns a tuple with the VpnGroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupIdAttribute

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupIdAttribute(v string)`

SetVpnGroupIdAttribute sets VpnGroupIdAttribute field to given value.


### GetVpnGroupListFilter

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupListFilter() string`

GetVpnGroupListFilter returns the VpnGroupListFilter field if non-nil, zero value otherwise.

### GetVpnGroupListFilterOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupListFilterOk() (*string, bool)`

GetVpnGroupListFilterOk returns a tuple with the VpnGroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupListFilter

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupListFilter(v string)`

SetVpnGroupListFilter sets VpnGroupListFilter field to given value.

### HasVpnGroupListFilter

`func (o *PutLdapVpnSchemaSettingsRequest) HasVpnGroupListFilter() bool`

HasVpnGroupListFilter returns a boolean if a field has been set.

### GetVpnGroupMemberAttribute

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttribute() string`

GetVpnGroupMemberAttribute returns the VpnGroupMemberAttribute field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttributeOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttributeOk() (*string, bool)`

GetVpnGroupMemberAttributeOk returns a tuple with the VpnGroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttribute

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupMemberAttribute(v string)`

SetVpnGroupMemberAttribute sets VpnGroupMemberAttribute field to given value.


### GetVpnGroupMemberAttrFormat

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttrFormat() string`

GetVpnGroupMemberAttrFormat returns the VpnGroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttrFormatOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttrFormatOk() (*string, bool)`

GetVpnGroupMemberAttrFormatOk returns a tuple with the VpnGroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttrFormat

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupMemberAttrFormat(v string)`

SetVpnGroupMemberAttrFormat sets VpnGroupMemberAttrFormat field to given value.

### HasVpnGroupMemberAttrFormat

`func (o *PutLdapVpnSchemaSettingsRequest) HasVpnGroupMemberAttrFormat() bool`

HasVpnGroupMemberAttrFormat returns a boolean if a field has been set.

### GetVpnGroupSearchScope

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupSearchScope() string`

GetVpnGroupSearchScope returns the VpnGroupSearchScope field if non-nil, zero value otherwise.

### GetVpnGroupSearchScopeOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupSearchScopeOk() (*string, bool)`

GetVpnGroupSearchScopeOk returns a tuple with the VpnGroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupSearchScope

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupSearchScope(v string)`

SetVpnGroupSearchScope sets VpnGroupSearchScope field to given value.

### HasVpnGroupSearchScope

`func (o *PutLdapVpnSchemaSettingsRequest) HasVpnGroupSearchScope() bool`

HasVpnGroupSearchScope returns a boolean if a field has been set.

### GetVpnGroupOtp

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupOtp() bool`

GetVpnGroupOtp returns the VpnGroupOtp field if non-nil, zero value otherwise.

### GetVpnGroupOtpOk

`func (o *PutLdapVpnSchemaSettingsRequest) GetVpnGroupOtpOk() (*bool, bool)`

GetVpnGroupOtpOk returns a tuple with the VpnGroupOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupOtp

`func (o *PutLdapVpnSchemaSettingsRequest) SetVpnGroupOtp(v bool)`

SetVpnGroupOtp sets VpnGroupOtp field to given value.

### HasVpnGroupOtp

`func (o *PutLdapVpnSchemaSettingsRequest) HasVpnGroupOtp() bool`

HasVpnGroupOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


