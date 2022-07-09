# LdapVpnRadiusSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnAuthEnabled** | Pointer to **bool** |  | [optional] 
**VpnAuthProvider** | Pointer to **string** |  | [optional] 
**VpnRadiusServer** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**VpnRadiusAuthPort** | Pointer to **int32** | Authentication port | [optional] 
**VpnRadiusAccountingPort** | Pointer to **int32** | Accounting port | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewLdapVpnRadiusSettings

`func NewLdapVpnRadiusSettings() *LdapVpnRadiusSettings`

NewLdapVpnRadiusSettings instantiates a new LdapVpnRadiusSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapVpnRadiusSettingsWithDefaults

`func NewLdapVpnRadiusSettingsWithDefaults() *LdapVpnRadiusSettings`

NewLdapVpnRadiusSettingsWithDefaults instantiates a new LdapVpnRadiusSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnAuthEnabled

`func (o *LdapVpnRadiusSettings) GetVpnAuthEnabled() bool`

GetVpnAuthEnabled returns the VpnAuthEnabled field if non-nil, zero value otherwise.

### GetVpnAuthEnabledOk

`func (o *LdapVpnRadiusSettings) GetVpnAuthEnabledOk() (*bool, bool)`

GetVpnAuthEnabledOk returns a tuple with the VpnAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthEnabled

`func (o *LdapVpnRadiusSettings) SetVpnAuthEnabled(v bool)`

SetVpnAuthEnabled sets VpnAuthEnabled field to given value.

### HasVpnAuthEnabled

`func (o *LdapVpnRadiusSettings) HasVpnAuthEnabled() bool`

HasVpnAuthEnabled returns a boolean if a field has been set.

### GetVpnAuthProvider

`func (o *LdapVpnRadiusSettings) GetVpnAuthProvider() string`

GetVpnAuthProvider returns the VpnAuthProvider field if non-nil, zero value otherwise.

### GetVpnAuthProviderOk

`func (o *LdapVpnRadiusSettings) GetVpnAuthProviderOk() (*string, bool)`

GetVpnAuthProviderOk returns a tuple with the VpnAuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthProvider

`func (o *LdapVpnRadiusSettings) SetVpnAuthProvider(v string)`

SetVpnAuthProvider sets VpnAuthProvider field to given value.

### HasVpnAuthProvider

`func (o *LdapVpnRadiusSettings) HasVpnAuthProvider() bool`

HasVpnAuthProvider returns a boolean if a field has been set.

### GetVpnRadiusServer

`func (o *LdapVpnRadiusSettings) GetVpnRadiusServer() string`

GetVpnRadiusServer returns the VpnRadiusServer field if non-nil, zero value otherwise.

### GetVpnRadiusServerOk

`func (o *LdapVpnRadiusSettings) GetVpnRadiusServerOk() (*string, bool)`

GetVpnRadiusServerOk returns a tuple with the VpnRadiusServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusServer

`func (o *LdapVpnRadiusSettings) SetVpnRadiusServer(v string)`

SetVpnRadiusServer sets VpnRadiusServer field to given value.

### HasVpnRadiusServer

`func (o *LdapVpnRadiusSettings) HasVpnRadiusServer() bool`

HasVpnRadiusServer returns a boolean if a field has been set.

### GetVpnRadiusAuthPort

`func (o *LdapVpnRadiusSettings) GetVpnRadiusAuthPort() int32`

GetVpnRadiusAuthPort returns the VpnRadiusAuthPort field if non-nil, zero value otherwise.

### GetVpnRadiusAuthPortOk

`func (o *LdapVpnRadiusSettings) GetVpnRadiusAuthPortOk() (*int32, bool)`

GetVpnRadiusAuthPortOk returns a tuple with the VpnRadiusAuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusAuthPort

`func (o *LdapVpnRadiusSettings) SetVpnRadiusAuthPort(v int32)`

SetVpnRadiusAuthPort sets VpnRadiusAuthPort field to given value.

### HasVpnRadiusAuthPort

`func (o *LdapVpnRadiusSettings) HasVpnRadiusAuthPort() bool`

HasVpnRadiusAuthPort returns a boolean if a field has been set.

### GetVpnRadiusAccountingPort

`func (o *LdapVpnRadiusSettings) GetVpnRadiusAccountingPort() int32`

GetVpnRadiusAccountingPort returns the VpnRadiusAccountingPort field if non-nil, zero value otherwise.

### GetVpnRadiusAccountingPortOk

`func (o *LdapVpnRadiusSettings) GetVpnRadiusAccountingPortOk() (*int32, bool)`

GetVpnRadiusAccountingPortOk returns a tuple with the VpnRadiusAccountingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusAccountingPort

`func (o *LdapVpnRadiusSettings) SetVpnRadiusAccountingPort(v int32)`

SetVpnRadiusAccountingPort sets VpnRadiusAccountingPort field to given value.

### HasVpnRadiusAccountingPort

`func (o *LdapVpnRadiusSettings) HasVpnRadiusAccountingPort() bool`

HasVpnRadiusAccountingPort returns a boolean if a field has been set.

### GetToken

`func (o *LdapVpnRadiusSettings) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LdapVpnRadiusSettings) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LdapVpnRadiusSettings) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LdapVpnRadiusSettings) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


