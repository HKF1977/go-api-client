# FirewallFWSetDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **string** | Chained firewall rules seperated by \&quot;\\n\&quot; | [optional] 
**Name** | Pointer to **string** | Name of the FWSet. Must be valid chain that begins with one of the following: NETS_, PORTS_, LIST_.  | [optional] 

## Methods

### NewFirewallFWSetDeleteRequest

`func NewFirewallFWSetDeleteRequest() *FirewallFWSetDeleteRequest`

NewFirewallFWSetDeleteRequest instantiates a new FirewallFWSetDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFWSetDeleteRequestWithDefaults

`func NewFirewallFWSetDeleteRequestWithDefaults() *FirewallFWSetDeleteRequest`

NewFirewallFWSetDeleteRequestWithDefaults instantiates a new FirewallFWSetDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FirewallFWSetDeleteRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallFWSetDeleteRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallFWSetDeleteRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallFWSetDeleteRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetName

`func (o *FirewallFWSetDeleteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallFWSetDeleteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallFWSetDeleteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallFWSetDeleteRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


