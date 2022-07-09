# PostTestLdapVpnSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnGroupBase** | **string** | Base DN from which to search for Groups | 
**VpnGroupIdAttribute** | **string** | Attribute type for the Groups | 
**VpnGroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**VpnGroupMemberAttribute** | **string** | Attribute used to search for a user within the Group | 
**VpnGroupMemberAttrFormat** | Pointer to **string** | Format of the Group Member attribute | [optional] [default to "UserDN"]
**VpnGroupSearchScope** | Pointer to **string** | Search scope for filter | [optional] [default to "subtree"]
**VpnGroupOtp** | Pointer to **bool** | Use Google authenticator (OTP)? | [optional] [default to false]
**Limit** | Pointer to **int32** | Number of records to return. Default &#x3D; 100 | [optional] [default to 100]

## Methods

### NewPostTestLdapVpnSchemaSettingsRequest

`func NewPostTestLdapVpnSchemaSettingsRequest(vpnGroupBase string, vpnGroupIdAttribute string, vpnGroupMemberAttribute string, ) *PostTestLdapVpnSchemaSettingsRequest`

NewPostTestLdapVpnSchemaSettingsRequest instantiates a new PostTestLdapVpnSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTestLdapVpnSchemaSettingsRequestWithDefaults

`func NewPostTestLdapVpnSchemaSettingsRequestWithDefaults() *PostTestLdapVpnSchemaSettingsRequest`

NewPostTestLdapVpnSchemaSettingsRequestWithDefaults instantiates a new PostTestLdapVpnSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnGroupBase

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupBase() string`

GetVpnGroupBase returns the VpnGroupBase field if non-nil, zero value otherwise.

### GetVpnGroupBaseOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupBaseOk() (*string, bool)`

GetVpnGroupBaseOk returns a tuple with the VpnGroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupBase

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupBase(v string)`

SetVpnGroupBase sets VpnGroupBase field to given value.


### GetVpnGroupIdAttribute

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupIdAttribute() string`

GetVpnGroupIdAttribute returns the VpnGroupIdAttribute field if non-nil, zero value otherwise.

### GetVpnGroupIdAttributeOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupIdAttributeOk() (*string, bool)`

GetVpnGroupIdAttributeOk returns a tuple with the VpnGroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupIdAttribute

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupIdAttribute(v string)`

SetVpnGroupIdAttribute sets VpnGroupIdAttribute field to given value.


### GetVpnGroupListFilter

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupListFilter() string`

GetVpnGroupListFilter returns the VpnGroupListFilter field if non-nil, zero value otherwise.

### GetVpnGroupListFilterOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupListFilterOk() (*string, bool)`

GetVpnGroupListFilterOk returns a tuple with the VpnGroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupListFilter

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupListFilter(v string)`

SetVpnGroupListFilter sets VpnGroupListFilter field to given value.

### HasVpnGroupListFilter

`func (o *PostTestLdapVpnSchemaSettingsRequest) HasVpnGroupListFilter() bool`

HasVpnGroupListFilter returns a boolean if a field has been set.

### GetVpnGroupMemberAttribute

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttribute() string`

GetVpnGroupMemberAttribute returns the VpnGroupMemberAttribute field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttributeOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttributeOk() (*string, bool)`

GetVpnGroupMemberAttributeOk returns a tuple with the VpnGroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttribute

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupMemberAttribute(v string)`

SetVpnGroupMemberAttribute sets VpnGroupMemberAttribute field to given value.


### GetVpnGroupMemberAttrFormat

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttrFormat() string`

GetVpnGroupMemberAttrFormat returns the VpnGroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetVpnGroupMemberAttrFormatOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupMemberAttrFormatOk() (*string, bool)`

GetVpnGroupMemberAttrFormatOk returns a tuple with the VpnGroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupMemberAttrFormat

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupMemberAttrFormat(v string)`

SetVpnGroupMemberAttrFormat sets VpnGroupMemberAttrFormat field to given value.

### HasVpnGroupMemberAttrFormat

`func (o *PostTestLdapVpnSchemaSettingsRequest) HasVpnGroupMemberAttrFormat() bool`

HasVpnGroupMemberAttrFormat returns a boolean if a field has been set.

### GetVpnGroupSearchScope

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupSearchScope() string`

GetVpnGroupSearchScope returns the VpnGroupSearchScope field if non-nil, zero value otherwise.

### GetVpnGroupSearchScopeOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupSearchScopeOk() (*string, bool)`

GetVpnGroupSearchScopeOk returns a tuple with the VpnGroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupSearchScope

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupSearchScope(v string)`

SetVpnGroupSearchScope sets VpnGroupSearchScope field to given value.

### HasVpnGroupSearchScope

`func (o *PostTestLdapVpnSchemaSettingsRequest) HasVpnGroupSearchScope() bool`

HasVpnGroupSearchScope returns a boolean if a field has been set.

### GetVpnGroupOtp

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupOtp() bool`

GetVpnGroupOtp returns the VpnGroupOtp field if non-nil, zero value otherwise.

### GetVpnGroupOtpOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetVpnGroupOtpOk() (*bool, bool)`

GetVpnGroupOtpOk returns a tuple with the VpnGroupOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGroupOtp

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetVpnGroupOtp(v bool)`

SetVpnGroupOtp sets VpnGroupOtp field to given value.

### HasVpnGroupOtp

`func (o *PostTestLdapVpnSchemaSettingsRequest) HasVpnGroupOtp() bool`

HasVpnGroupOtp returns a boolean if a field has been set.

### GetLimit

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostTestLdapVpnSchemaSettingsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostTestLdapVpnSchemaSettingsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostTestLdapVpnSchemaSettingsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


