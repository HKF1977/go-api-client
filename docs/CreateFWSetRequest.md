# CreateFWSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **string** | Chained firewall rules seperated by \&quot;\\n\&quot; | [optional] 
**Name** | Pointer to **string** | &#39;name of the FWSet. Must be valid chain that begins with one of the following: NETS_, PORTS_, LIST_.&#39;   | [optional] 
**Flush** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateFWSetRequest

`func NewCreateFWSetRequest() *CreateFWSetRequest`

NewCreateFWSetRequest instantiates a new CreateFWSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFWSetRequestWithDefaults

`func NewCreateFWSetRequestWithDefaults() *CreateFWSetRequest`

NewCreateFWSetRequestWithDefaults instantiates a new CreateFWSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *CreateFWSetRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateFWSetRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateFWSetRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateFWSetRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetName

`func (o *CreateFWSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFWSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFWSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFWSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlush

`func (o *CreateFWSetRequest) GetFlush() bool`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *CreateFWSetRequest) GetFlushOk() (*bool, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *CreateFWSetRequest) SetFlush(v bool)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *CreateFWSetRequest) HasFlush() bool`

HasFlush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


