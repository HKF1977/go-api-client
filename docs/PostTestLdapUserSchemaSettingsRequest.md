# PostTestLdapUserSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserBase** | **string** | Base DN from which to search for Users | 
**UserIdAttribute** | **string** | Attribute type for the Users | 
**UserListFilter** | Pointer to **string** | Search filter for Users | [optional] 
**Limit** | Pointer to **int32** | Number of records to return. Default &#x3D; 100 | [optional] [default to 100]

## Methods

### NewPostTestLdapUserSchemaSettingsRequest

`func NewPostTestLdapUserSchemaSettingsRequest(userBase string, userIdAttribute string, ) *PostTestLdapUserSchemaSettingsRequest`

NewPostTestLdapUserSchemaSettingsRequest instantiates a new PostTestLdapUserSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTestLdapUserSchemaSettingsRequestWithDefaults

`func NewPostTestLdapUserSchemaSettingsRequestWithDefaults() *PostTestLdapUserSchemaSettingsRequest`

NewPostTestLdapUserSchemaSettingsRequestWithDefaults instantiates a new PostTestLdapUserSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserBase

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *PostTestLdapUserSchemaSettingsRequest) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.


### GetUserIdAttribute

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *PostTestLdapUserSchemaSettingsRequest) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.


### GetUserListFilter

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *PostTestLdapUserSchemaSettingsRequest) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *PostTestLdapUserSchemaSettingsRequest) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *PostTestLdapUserSchemaSettingsRequest) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.

### GetLimit

`func (o *PostTestLdapUserSchemaSettingsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostTestLdapUserSchemaSettingsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostTestLdapUserSchemaSettingsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostTestLdapUserSchemaSettingsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


