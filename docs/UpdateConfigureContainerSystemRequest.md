# UpdateConfigureContainerSystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Subnet CIDR that will be used for the container network. | 

## Methods

### NewUpdateConfigureContainerSystemRequest

`func NewUpdateConfigureContainerSystemRequest(network string, ) *UpdateConfigureContainerSystemRequest`

NewUpdateConfigureContainerSystemRequest instantiates a new UpdateConfigureContainerSystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigureContainerSystemRequestWithDefaults

`func NewUpdateConfigureContainerSystemRequestWithDefaults() *UpdateConfigureContainerSystemRequest`

NewUpdateConfigureContainerSystemRequestWithDefaults instantiates a new UpdateConfigureContainerSystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *UpdateConfigureContainerSystemRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateConfigureContainerSystemRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateConfigureContainerSystemRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


