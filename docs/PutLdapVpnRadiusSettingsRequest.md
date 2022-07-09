# PutLdapVpnRadiusSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnAuthEnabled** | **bool** | Enable use of Radius through VPN. If true, other params required. | 
**VpnRadiusServer** | **string** | IP address or resolvable hostname | 
**VpnRadiusAuthPort** | Pointer to **int32** | Authentication port | [optional] [default to 1812]
**VpnRadiusAccountingPort** | Pointer to **int32** | Accounting port | [optional] [default to 1812]
**VpnRadiusPass** | **string** | Shared password | 

## Methods

### NewPutLdapVpnRadiusSettingsRequest

`func NewPutLdapVpnRadiusSettingsRequest(vpnAuthEnabled bool, vpnRadiusServer string, vpnRadiusPass string, ) *PutLdapVpnRadiusSettingsRequest`

NewPutLdapVpnRadiusSettingsRequest instantiates a new PutLdapVpnRadiusSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLdapVpnRadiusSettingsRequestWithDefaults

`func NewPutLdapVpnRadiusSettingsRequestWithDefaults() *PutLdapVpnRadiusSettingsRequest`

NewPutLdapVpnRadiusSettingsRequestWithDefaults instantiates a new PutLdapVpnRadiusSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnAuthEnabled

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnAuthEnabled() bool`

GetVpnAuthEnabled returns the VpnAuthEnabled field if non-nil, zero value otherwise.

### GetVpnAuthEnabledOk

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnAuthEnabledOk() (*bool, bool)`

GetVpnAuthEnabledOk returns a tuple with the VpnAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAuthEnabled

`func (o *PutLdapVpnRadiusSettingsRequest) SetVpnAuthEnabled(v bool)`

SetVpnAuthEnabled sets VpnAuthEnabled field to given value.


### GetVpnRadiusServer

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusServer() string`

GetVpnRadiusServer returns the VpnRadiusServer field if non-nil, zero value otherwise.

### GetVpnRadiusServerOk

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusServerOk() (*string, bool)`

GetVpnRadiusServerOk returns a tuple with the VpnRadiusServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusServer

`func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusServer(v string)`

SetVpnRadiusServer sets VpnRadiusServer field to given value.


### GetVpnRadiusAuthPort

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAuthPort() int32`

GetVpnRadiusAuthPort returns the VpnRadiusAuthPort field if non-nil, zero value otherwise.

### GetVpnRadiusAuthPortOk

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAuthPortOk() (*int32, bool)`

GetVpnRadiusAuthPortOk returns a tuple with the VpnRadiusAuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusAuthPort

`func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusAuthPort(v int32)`

SetVpnRadiusAuthPort sets VpnRadiusAuthPort field to given value.

### HasVpnRadiusAuthPort

`func (o *PutLdapVpnRadiusSettingsRequest) HasVpnRadiusAuthPort() bool`

HasVpnRadiusAuthPort returns a boolean if a field has been set.

### GetVpnRadiusAccountingPort

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAccountingPort() int32`

GetVpnRadiusAccountingPort returns the VpnRadiusAccountingPort field if non-nil, zero value otherwise.

### GetVpnRadiusAccountingPortOk

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusAccountingPortOk() (*int32, bool)`

GetVpnRadiusAccountingPortOk returns a tuple with the VpnRadiusAccountingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusAccountingPort

`func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusAccountingPort(v int32)`

SetVpnRadiusAccountingPort sets VpnRadiusAccountingPort field to given value.

### HasVpnRadiusAccountingPort

`func (o *PutLdapVpnRadiusSettingsRequest) HasVpnRadiusAccountingPort() bool`

HasVpnRadiusAccountingPort returns a boolean if a field has been set.

### GetVpnRadiusPass

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusPass() string`

GetVpnRadiusPass returns the VpnRadiusPass field if non-nil, zero value otherwise.

### GetVpnRadiusPassOk

`func (o *PutLdapVpnRadiusSettingsRequest) GetVpnRadiusPassOk() (*string, bool)`

GetVpnRadiusPassOk returns a tuple with the VpnRadiusPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnRadiusPass

`func (o *PutLdapVpnRadiusSettingsRequest) SetVpnRadiusPass(v string)`

SetVpnRadiusPass sets VpnRadiusPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


