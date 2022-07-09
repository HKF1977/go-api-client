# FirewallFWSetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to **string** | Begins with name and is followed by Firewall rules. This is an indexed rule that allows for speedy matching on IPs | [optional] 

## Methods

### NewFirewallFWSetData

`func NewFirewallFWSetData() *FirewallFWSetData`

NewFirewallFWSetData instantiates a new FirewallFWSetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFWSetDataWithDefaults

`func NewFirewallFWSetDataWithDefaults() *FirewallFWSetData`

NewFirewallFWSetDataWithDefaults instantiates a new FirewallFWSetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FirewallFWSetData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallFWSetData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallFWSetData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallFWSetData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRules

`func (o *FirewallFWSetData) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallFWSetData) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallFWSetData) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallFWSetData) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


