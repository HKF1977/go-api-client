# \FirewallApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFirewallFwSet**](FirewallApi.md#DeleteFirewallFwSet) | **Delete** /firewall/fwsets | Delete firewall FWSet
[**DeleteFirewallRuleByPosition**](FirewallApi.md#DeleteFirewallRuleByPosition) | **Delete** /firewall/rules/{position} | Delete firewall rule by position
[**DeleteFirewallRuleByRule**](FirewallApi.md#DeleteFirewallRuleByRule) | **Delete** /firewall/rules | Delete firewall rule
[**DeleteFirewallSubgroup**](FirewallApi.md#DeleteFirewallSubgroup) | **Delete** /firewall/rules/subgroup | Delete firewall subgroup
[**GetFirewallFwSets**](FirewallApi.md#GetFirewallFwSets) | **Get** /firewall/fwsets | Get firewall FWSets
[**GetFirewallRuleSubgroups**](FirewallApi.md#GetFirewallRuleSubgroups) | **Get** /firewall/rules/subgroup | Get firewall subgroups
[**GetFirewallRules**](FirewallApi.md#GetFirewallRules) | **Get** /firewall/rules | Get firewall rules
[**PostCreateFirewallFwSet**](FirewallApi.md#PostCreateFirewallFwSet) | **Post** /firewall/fwsets | Create firewall FWSet
[**PostCreateFirewallRule**](FirewallApi.md#PostCreateFirewallRule) | **Post** /firewall/rules | Create firewall rule
[**PostCreateFirewallSubgroup**](FirewallApi.md#PostCreateFirewallSubgroup) | **Post** /firewall/rules/subgroup | Create firewall subgroup
[**PutFirewall**](FirewallApi.md#PutFirewall) | **Put** /firewall/rules | Overwrite firewall
[**PutFirewallAction**](FirewallApi.md#PutFirewallAction) | **Put** /firewall/actions | Put firewall action
[**PutOverwriteFirewall**](FirewallApi.md#PutOverwriteFirewall) | **Put** /firewall | Put firewall
[**PutReinitializeFirewallSubgroups**](FirewallApi.md#PutReinitializeFirewallSubgroups) | **Put** /firewall/rules/subgroup | Reload firewall subgroups
[**PutReinitializeFwSets**](FirewallApi.md#PutReinitializeFwSets) | **Put** /firewall/fwsets | Reload all firewall FWsets



## DeleteFirewallFwSet

> Object DeleteFirewallFwSet(ctx).FirewallFWSetDeleteRequest(firewallFWSetDeleteRequest).Execute()

Delete firewall FWSet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    firewallFWSetDeleteRequest := *openapiclient.NewFirewallFWSetDeleteRequest() // FirewallFWSetDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallFwSet(context.Background()).FirewallFWSetDeleteRequest(firewallFWSetDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallFwSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallFwSet`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallFwSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallFwSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallFWSetDeleteRequest** | [**FirewallFWSetDeleteRequest**](FirewallFWSetDeleteRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRuleByPosition

> Object DeleteFirewallRuleByPosition(ctx, position).Execute()

Delete firewall rule by position



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    position := int32(56) // int32 | index position for firewall rule, 0 is first

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRuleByPosition(context.Background(), position).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRuleByPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRuleByPosition`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRuleByPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**position** | **int32** | index position for firewall rule, 0 is first | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleByPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRuleByRule

> Object DeleteFirewallRuleByRule(ctx).DeleteFirewallRuleRequest(deleteFirewallRuleRequest).Execute()

Delete firewall rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteFirewallRuleRequest := *openapiclient.NewDeleteFirewallRuleRequest("Rule_example") // DeleteFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRuleByRule(context.Background()).DeleteFirewallRuleRequest(deleteFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRuleByRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRuleByRule`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRuleByRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleByRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFirewallRuleRequest** | [**DeleteFirewallRuleRequest**](DeleteFirewallRuleRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallSubgroup

> Object DeleteFirewallSubgroup(ctx).FirewallSubgroupDeleteRequest(firewallSubgroupDeleteRequest).Execute()

Delete firewall subgroup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    firewallSubgroupDeleteRequest := *openapiclient.NewFirewallSubgroupDeleteRequest() // FirewallSubgroupDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallSubgroup(context.Background()).FirewallSubgroupDeleteRequest(firewallSubgroupDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallSubgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallSubgroup`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallSubgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallSubgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallSubgroupDeleteRequest** | [**FirewallSubgroupDeleteRequest**](FirewallSubgroupDeleteRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallFwSets

> FirewallFWSetListResponse GetFirewallFwSets(ctx).Name(name).Verbose(verbose).Execute()

Get firewall FWSets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | name of resource (optional)
    verbose := true // bool | True for verbose output (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallFwSets(context.Background()).Name(name).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallFwSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallFwSets`: FirewallFWSetListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallFwSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallFwSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of resource | 
 **verbose** | **bool** | True for verbose output | [default to true]

### Return type

[**FirewallFWSetListResponse**](FirewallFWSetListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleSubgroups

> FirewallSubgroupListResponse GetFirewallRuleSubgroups(ctx).Name(name).Verbose(verbose).Execute()

Get firewall subgroups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | name of resource (optional)
    verbose := true // bool | True for verbose output (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallRuleSubgroups(context.Background()).Name(name).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRuleSubgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRuleSubgroups`: FirewallSubgroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRuleSubgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleSubgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of resource | 
 **verbose** | **bool** | True for verbose output | [default to true]

### Return type

[**FirewallSubgroupListResponse**](FirewallSubgroupListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRules

> FirewallRuleListResponse GetFirewallRules(ctx).Execute()

Get firewall rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRules`: FirewallRuleListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRulesRequest struct via the builder pattern


### Return type

[**FirewallRuleListResponse**](FirewallRuleListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallFwSet

> PostCreateFirewallFwSet200Response PostCreateFirewallFwSet(ctx).CreateFWSetRequest(createFWSetRequest).Execute()

Create firewall FWSet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createFWSetRequest := *openapiclient.NewCreateFWSetRequest() // CreateFWSetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallFwSet(context.Background()).CreateFWSetRequest(createFWSetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallFwSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallFwSet`: PostCreateFirewallFwSet200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallFwSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallFwSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFWSetRequest** | [**CreateFWSetRequest**](CreateFWSetRequest.md) |  | 

### Return type

[**PostCreateFirewallFwSet200Response**](PostCreateFirewallFwSet200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallRule

> FirewallRuleOperationResponse PostCreateFirewallRule(ctx).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create firewall rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Rule_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallRule(context.Background()).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallRule`: FirewallRuleOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRuleOperationResponse**](FirewallRuleOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallSubgroup

> PostCreateFirewallSubgroup200Response PostCreateFirewallSubgroup(ctx).CreateFWSubgroupRequest(createFWSubgroupRequest).Execute()

Create firewall subgroup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createFWSubgroupRequest := openapiclient.CreateFWSubgroupRequest{Interface{}: new(interface{})} // CreateFWSubgroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallSubgroup(context.Background()).CreateFWSubgroupRequest(createFWSubgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallSubgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallSubgroup`: PostCreateFirewallSubgroup200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallSubgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallSubgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFWSubgroupRequest** | [**CreateFWSubgroupRequest**](CreateFWSubgroupRequest.md) |  | 

### Return type

[**PostCreateFirewallSubgroup200Response**](PostCreateFirewallSubgroup200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFirewall

> Object PutFirewall(ctx).PutFirewallRuleRequest(putFirewallRuleRequest).Execute()

Overwrite firewall



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    putFirewallRuleRequest := *openapiclient.NewPutFirewallRuleRequest("Rules_example") // PutFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutFirewall(context.Background()).PutFirewallRuleRequest(putFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFirewall`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putFirewallRuleRequest** | [**PutFirewallRuleRequest**](PutFirewallRuleRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFirewallAction

> Object PutFirewallAction(ctx).FirewallActionRequest(firewallActionRequest).Execute()

Put firewall action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    firewallActionRequest := *openapiclient.NewFirewallActionRequest("Action_example") // FirewallActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutFirewallAction(context.Background()).FirewallActionRequest(firewallActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutFirewallAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFirewallAction`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutFirewallAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallActionRequest** | [**FirewallActionRequest**](FirewallActionRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOverwriteFirewall

> TaskTokenResponse PutOverwriteFirewall(ctx).OverwriteRequest(overwriteRequest).Execute()

Put firewall



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    overwriteRequest := *openapiclient.NewOverwriteRequest("Rules_example") // OverwriteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutOverwriteFirewall(context.Background()).OverwriteRequest(overwriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutOverwriteFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOverwriteFirewall`: TaskTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutOverwriteFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutOverwriteFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overwriteRequest** | [**OverwriteRequest**](OverwriteRequest.md) |  | 

### Return type

[**TaskTokenResponse**](TaskTokenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReinitializeFirewallSubgroups

> PutReinitializeFirewallSubgroups(ctx).ReinitRequest(reinitRequest).Execute()

Reload firewall subgroups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reinitRequest := *openapiclient.NewReinitRequest() // ReinitRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutReinitializeFirewallSubgroups(context.Background()).ReinitRequest(reinitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutReinitializeFirewallSubgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReinitializeFirewallSubgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reinitRequest** | [**ReinitRequest**](ReinitRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReinitializeFwSets

> PutReinitializeFwSets(ctx).ReinitRequest(reinitRequest).Execute()

Reload all firewall FWsets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reinitRequest := *openapiclient.NewReinitRequest() // ReinitRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutReinitializeFwSets(context.Background()).ReinitRequest(reinitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutReinitializeFwSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReinitializeFwSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reinitRequest** | [**ReinitRequest**](ReinitRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

