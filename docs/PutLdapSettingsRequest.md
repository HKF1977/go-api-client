# PutLdapSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP address or resolvable hostname of LDAP server | 
**Port** | Pointer to **int32** | Port for LDAP | [optional] [default to 389]
**Encrypt** | Pointer to **bool** | Use SSL | [optional] [default to false]
**EncryptLdaps** | Pointer to **bool** | Use LDAPS or start TLS (default)? | [optional] [default to true]
**EncryptAuth** | Pointer to **bool** | Use certificates to authenticate via encrypted connection | [optional] [default to false]
**EncryptVerifyCa** | Pointer to **bool** | Verify certicate using authority | [optional] [default to false]
**Binddn** | Pointer to **string** | Bind Username | [optional] 
**Bindpw** | Pointer to **string** | Bind Password | [optional] 

## Methods

### NewPutLdapSettingsRequest

`func NewPutLdapSettingsRequest(host string, ) *PutLdapSettingsRequest`

NewPutLdapSettingsRequest instantiates a new PutLdapSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLdapSettingsRequestWithDefaults

`func NewPutLdapSettingsRequestWithDefaults() *PutLdapSettingsRequest`

NewPutLdapSettingsRequestWithDefaults instantiates a new PutLdapSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PutLdapSettingsRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PutLdapSettingsRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PutLdapSettingsRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PutLdapSettingsRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PutLdapSettingsRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PutLdapSettingsRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PutLdapSettingsRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *PutLdapSettingsRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *PutLdapSettingsRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *PutLdapSettingsRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *PutLdapSettingsRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *PutLdapSettingsRequest) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *PutLdapSettingsRequest) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *PutLdapSettingsRequest) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *PutLdapSettingsRequest) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *PutLdapSettingsRequest) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *PutLdapSettingsRequest) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *PutLdapSettingsRequest) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *PutLdapSettingsRequest) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *PutLdapSettingsRequest) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *PutLdapSettingsRequest) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *PutLdapSettingsRequest) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *PutLdapSettingsRequest) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetBinddn

`func (o *PutLdapSettingsRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *PutLdapSettingsRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *PutLdapSettingsRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *PutLdapSettingsRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *PutLdapSettingsRequest) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *PutLdapSettingsRequest) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *PutLdapSettingsRequest) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *PutLdapSettingsRequest) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


