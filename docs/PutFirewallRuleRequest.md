# PutFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | **string** | New-line delimited string of firewall rules. | 

## Methods

### NewPutFirewallRuleRequest

`func NewPutFirewallRuleRequest(rules string, ) *PutFirewallRuleRequest`

NewPutFirewallRuleRequest instantiates a new PutFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFirewallRuleRequestWithDefaults

`func NewPutFirewallRuleRequestWithDefaults() *PutFirewallRuleRequest`

NewPutFirewallRuleRequestWithDefaults instantiates a new PutFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PutFirewallRuleRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PutFirewallRuleRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PutFirewallRuleRequest) SetRules(v string)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


