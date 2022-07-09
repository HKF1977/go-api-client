# LdapSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Encrypt** | Pointer to **bool** | Use SSL | [optional] 
**EncryptLdaps** | Pointer to **bool** | Use LDAPS or start TLS (default)? | [optional] 
**EncryptAuth** | Pointer to **bool** | Use certificates to authenticate via encrypted connection | [optional] 
**EncryptAuthKey** | Pointer to **bool** |  | [optional] 
**EncryptAuthCert** | Pointer to **bool** |  | [optional] 
**EncryptVerifyCa** | Pointer to **bool** | Verify certicate using authority | [optional] 
**EncryptCaCert** | Pointer to **bool** |  | [optional] 
**Binddn** | Pointer to **string** | Bind username | [optional] 

## Methods

### NewLdapSettings

`func NewLdapSettings() *LdapSettings`

NewLdapSettings instantiates a new LdapSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapSettingsWithDefaults

`func NewLdapSettingsWithDefaults() *LdapSettings`

NewLdapSettingsWithDefaults instantiates a new LdapSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *LdapSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LdapSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LdapSettings) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LdapSettings) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *LdapSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapSettings) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LdapSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *LdapSettings) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *LdapSettings) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *LdapSettings) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *LdapSettings) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *LdapSettings) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *LdapSettings) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *LdapSettings) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *LdapSettings) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *LdapSettings) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *LdapSettings) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *LdapSettings) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *LdapSettings) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptAuthKey

`func (o *LdapSettings) GetEncryptAuthKey() bool`

GetEncryptAuthKey returns the EncryptAuthKey field if non-nil, zero value otherwise.

### GetEncryptAuthKeyOk

`func (o *LdapSettings) GetEncryptAuthKeyOk() (*bool, bool)`

GetEncryptAuthKeyOk returns a tuple with the EncryptAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKey

`func (o *LdapSettings) SetEncryptAuthKey(v bool)`

SetEncryptAuthKey sets EncryptAuthKey field to given value.

### HasEncryptAuthKey

`func (o *LdapSettings) HasEncryptAuthKey() bool`

HasEncryptAuthKey returns a boolean if a field has been set.

### GetEncryptAuthCert

`func (o *LdapSettings) GetEncryptAuthCert() bool`

GetEncryptAuthCert returns the EncryptAuthCert field if non-nil, zero value otherwise.

### GetEncryptAuthCertOk

`func (o *LdapSettings) GetEncryptAuthCertOk() (*bool, bool)`

GetEncryptAuthCertOk returns a tuple with the EncryptAuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCert

`func (o *LdapSettings) SetEncryptAuthCert(v bool)`

SetEncryptAuthCert sets EncryptAuthCert field to given value.

### HasEncryptAuthCert

`func (o *LdapSettings) HasEncryptAuthCert() bool`

HasEncryptAuthCert returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *LdapSettings) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *LdapSettings) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *LdapSettings) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *LdapSettings) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetEncryptCaCert

`func (o *LdapSettings) GetEncryptCaCert() bool`

GetEncryptCaCert returns the EncryptCaCert field if non-nil, zero value otherwise.

### GetEncryptCaCertOk

`func (o *LdapSettings) GetEncryptCaCertOk() (*bool, bool)`

GetEncryptCaCertOk returns a tuple with the EncryptCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCert

`func (o *LdapSettings) SetEncryptCaCert(v bool)`

SetEncryptCaCert sets EncryptCaCert field to given value.

### HasEncryptCaCert

`func (o *LdapSettings) HasEncryptCaCert() bool`

HasEncryptCaCert returns a boolean if a field has been set.

### GetBinddn

`func (o *LdapSettings) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *LdapSettings) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *LdapSettings) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *LdapSettings) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


