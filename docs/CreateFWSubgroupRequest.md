# CreateFWSubgroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **string** | Chained firewall rules seperated by \\n. Rule should be preceded by group name | [optional] 
**Name** | Pointer to **string** | &#39;name of the subgroup. Must be valid chain that begins with one of the following: PRE_C_, PST_C_, FWD_C_, INP_C_, OUT_C_.&#39;   | [optional] 
**Position** | Pointer to **int32** | Position which the chain will be inserted in the list of Firewall rules.  Default is 0, which is first in the ruleset  | [optional] 
**Flush** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateFWSubgroupRequest

`func NewCreateFWSubgroupRequest() *CreateFWSubgroupRequest`

NewCreateFWSubgroupRequest instantiates a new CreateFWSubgroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFWSubgroupRequestWithDefaults

`func NewCreateFWSubgroupRequestWithDefaults() *CreateFWSubgroupRequest`

NewCreateFWSubgroupRequestWithDefaults instantiates a new CreateFWSubgroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *CreateFWSubgroupRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateFWSubgroupRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateFWSubgroupRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateFWSubgroupRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetName

`func (o *CreateFWSubgroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFWSubgroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFWSubgroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFWSubgroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *CreateFWSubgroupRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateFWSubgroupRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateFWSubgroupRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateFWSubgroupRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetFlush

`func (o *CreateFWSubgroupRequest) GetFlush() bool`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *CreateFWSubgroupRequest) GetFlushOk() (*bool, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *CreateFWSubgroupRequest) SetFlush(v bool)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *CreateFWSubgroupRequest) HasFlush() bool`

HasFlush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


