# UpdateConfigClientpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable clientpacks. | [optional] 
**CheckedOut** | Pointer to **bool** | Update whether clientpacks are checked out and thus unavailable | [optional] 
**Compression** | Pointer to **string** | Turn compression on or off for all. Can be \&quot;on\&quot; or \&quot;off\&quot; currently. | [optional] 

## Methods

### NewUpdateConfigClientpackRequest

`func NewUpdateConfigClientpackRequest() *UpdateConfigClientpackRequest`

NewUpdateConfigClientpackRequest instantiates a new UpdateConfigClientpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigClientpackRequestWithDefaults

`func NewUpdateConfigClientpackRequestWithDefaults() *UpdateConfigClientpackRequest`

NewUpdateConfigClientpackRequestWithDefaults instantiates a new UpdateConfigClientpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateConfigClientpackRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateConfigClientpackRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateConfigClientpackRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateConfigClientpackRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCheckedOut

`func (o *UpdateConfigClientpackRequest) GetCheckedOut() bool`

GetCheckedOut returns the CheckedOut field if non-nil, zero value otherwise.

### GetCheckedOutOk

`func (o *UpdateConfigClientpackRequest) GetCheckedOutOk() (*bool, bool)`

GetCheckedOutOk returns a tuple with the CheckedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedOut

`func (o *UpdateConfigClientpackRequest) SetCheckedOut(v bool)`

SetCheckedOut sets CheckedOut field to given value.

### HasCheckedOut

`func (o *UpdateConfigClientpackRequest) HasCheckedOut() bool`

HasCheckedOut returns a boolean if a field has been set.

### GetCompression

`func (o *UpdateConfigClientpackRequest) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *UpdateConfigClientpackRequest) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *UpdateConfigClientpackRequest) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *UpdateConfigClientpackRequest) HasCompression() bool`

HasCompression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


