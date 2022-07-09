# \AccessApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessUrl**](AccessApi.md#CreateAccessUrl) | **Post** /access/url | Create access URL
[**CreateApiToken**](AccessApi.md#CreateApiToken) | **Post** /access/token | Create API token
[**DeleteAccessUrl**](AccessApi.md#DeleteAccessUrl) | **Delete** /access/url/{access_url_id} | Delete access URL
[**DeleteAccessUrlBySearch**](AccessApi.md#DeleteAccessUrlBySearch) | **Delete** /access/url | Find and delete access URL
[**DeleteApiToken**](AccessApi.md#DeleteApiToken) | **Delete** /access/token/{token_id} | Delete API token
[**GetAccessUrl**](AccessApi.md#GetAccessUrl) | **Get** /access/url/{access_url_id} | Get access URL
[**GetAccessUrls**](AccessApi.md#GetAccessUrls) | **Get** /access/urls | Get access URLs
[**GetApiToken**](AccessApi.md#GetApiToken) | **Get** /access/token/{token_id} | Get API access token
[**GetApiTokens**](AccessApi.md#GetApiTokens) | **Get** /access/tokens | Get API access tokens
[**GetLdapGroupSchemaSettings**](AccessApi.md#GetLdapGroupSchemaSettings) | **Get** /settings/ldap/group_schema | Get LDAP group schema settings
[**GetLdapSettings**](AccessApi.md#GetLdapSettings) | **Get** /settings/ldap | Get LDAP settings
[**GetLdapUserSchemaSettings**](AccessApi.md#GetLdapUserSchemaSettings) | **Get** /settings/ldap/user_schema | Get LDAP user schema settings
[**GetLdapVpnRadiusSettings**](AccessApi.md#GetLdapVpnRadiusSettings) | **Get** /settings/ldap/vpn_radius | Get LDAP VPN Radius settings
[**GetLdapVpnSchemaSettings**](AccessApi.md#GetLdapVpnSchemaSettings) | **Get** /settings/ldap/vpn_schema | Get LDAP VPN schema settings
[**PostTestLdapGroupSchemaSettings**](AccessApi.md#PostTestLdapGroupSchemaSettings) | **Post** /settings/ldap/group_schema | Test LDAP group schema settings
[**PostTestLdapSettings**](AccessApi.md#PostTestLdapSettings) | **Post** /settings/ldap | Test LDAP settings
[**PostTestLdapUserSchemaSettings**](AccessApi.md#PostTestLdapUserSchemaSettings) | **Post** /settings/ldap/user_schema | Test LDAP user schema settings
[**PostTestLdapVpnSchemaSettings**](AccessApi.md#PostTestLdapVpnSchemaSettings) | **Post** /settings/ldap/vpn_schema | Test LDAP VPN schema settings
[**PutEnableLdap**](AccessApi.md#PutEnableLdap) | **Put** /settings/ldap/enabled | Enable/disable LDAP
[**PutExpireAccessUrl**](AccessApi.md#PutExpireAccessUrl) | **Put** /access/url/{access_url_id} | Expire access URL
[**PutExpireApiToken**](AccessApi.md#PutExpireApiToken) | **Put** /access/token/{token_id} | Expire API token
[**PutLdapGroupSchemaSettings**](AccessApi.md#PutLdapGroupSchemaSettings) | **Put** /settings/ldap/group_schema | Put LDAP group schema settings
[**PutLdapSettings**](AccessApi.md#PutLdapSettings) | **Put** /settings/ldap | Put LDAP settings
[**PutLdapUserSchemaSettings**](AccessApi.md#PutLdapUserSchemaSettings) | **Put** /settings/ldap/user_schema | Put LDAP user schema settings
[**PutLdapVpnRadiusSettings**](AccessApi.md#PutLdapVpnRadiusSettings) | **Put** /settings/ldap/vpn_radius | Put LDAP VPN Radius settings
[**PutLdapVpnSchemaSettings**](AccessApi.md#PutLdapVpnSchemaSettings) | **Put** /settings/ldap/vpn_schema | Put LDAP VPN schema settings
[**PutUploadLdapAuthCert**](AccessApi.md#PutUploadLdapAuthCert) | **Put** /settings/ldap/encrypt_auth_cert | Upload LDAP Auth Cert
[**PutUploadLdapAuthKey**](AccessApi.md#PutUploadLdapAuthKey) | **Put** /settings/ldap/encrypt_auth_key | Upload LDAP Auth Key
[**PutUploadLdapCaCert**](AccessApi.md#PutUploadLdapCaCert) | **Put** /settings/ldap/encrypt_ca_cert | Upload LDAP CA cert



## CreateAccessUrl

> AccessUrlDetail CreateAccessUrl(ctx).CreateAccessURLRequest(createAccessURLRequest).Execute()

Create access URL



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
    createAccessURLRequest := *openapiclient.NewCreateAccessURLRequest() // CreateAccessURLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.CreateAccessUrl(context.Background()).CreateAccessURLRequest(createAccessURLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.CreateAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessUrl`: AccessUrlDetail
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.CreateAccessUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessURLRequest** | [**CreateAccessURLRequest**](CreateAccessURLRequest.md) |  | 

### Return type

[**AccessUrlDetail**](AccessUrlDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiToken

> AccessTokenDetail CreateApiToken(ctx).CreateAPITokenRequest(createAPITokenRequest).Execute()

Create API token



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
    createAPITokenRequest := *openapiclient.NewCreateAPITokenRequest() // CreateAPITokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.CreateApiToken(context.Background()).CreateAPITokenRequest(createAPITokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.CreateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiToken`: AccessTokenDetail
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.CreateApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPITokenRequest** | [**CreateAPITokenRequest**](CreateAPITokenRequest.md) |  | 

### Return type

[**AccessTokenDetail**](AccessTokenDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessUrl

> Object DeleteAccessUrl(ctx, accessUrlId).Execute()

Delete access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteAccessUrl(context.Background(), accessUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessUrlRequest struct via the builder pattern


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


## DeleteAccessUrlBySearch

> Object DeleteAccessUrlBySearch(ctx).DeleteAccessURLRequest(deleteAccessURLRequest).Execute()

Find and delete access URL



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
    deleteAccessURLRequest := openapiclient.DeleteAccessURLRequest{Interface{}: new(interface{})} // DeleteAccessURLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteAccessUrlBySearch(context.Background()).DeleteAccessURLRequest(deleteAccessURLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteAccessUrlBySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessUrlBySearch`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteAccessUrlBySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessUrlBySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAccessURLRequest** | [**DeleteAccessURLRequest**](DeleteAccessURLRequest.md) |  | 

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


## DeleteApiToken

> SimpleStringResponse DeleteApiToken(ctx, tokenId).Execute()

Delete API token



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
    tokenId := int32(56) // int32 | Token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteApiToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiToken`: SimpleStringResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleStringResponse**](SimpleStringResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessUrl

> Object GetAccessUrl(ctx, accessUrlId).Execute()

Get access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.GetAccessUrl(context.Background(), accessUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessUrlRequest struct via the builder pattern


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


## GetAccessUrls

> AccessUrlListResponse GetAccessUrls(ctx).Execute()

Get access URLs



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
    resp, r, err := apiClient.AccessApi.GetAccessUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetAccessUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessUrls`: AccessUrlListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetAccessUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessUrlsRequest struct via the builder pattern


### Return type

[**AccessUrlListResponse**](AccessUrlListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiToken

> Object GetApiToken(ctx, tokenId).Execute()

Get API access token



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
    tokenId := int32(56) // int32 | Token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.GetApiToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiToken`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


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


## GetApiTokens

> AccessTokenListResponse GetApiTokens(ctx).Execute()

Get API access tokens



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
    resp, r, err := apiClient.AccessApi.GetApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTokens`: AccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokensRequest struct via the builder pattern


### Return type

[**AccessTokenListResponse**](AccessTokenListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapGroupSchemaSettings

> Object GetLdapGroupSchemaSettings(ctx).Execute()

Get LDAP group schema settings



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
    resp, r, err := apiClient.AccessApi.GetLdapGroupSchemaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetLdapGroupSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapGroupSchemaSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetLdapGroupSchemaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapGroupSchemaSettingsRequest struct via the builder pattern


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


## GetLdapSettings

> Object GetLdapSettings(ctx).Execute()

Get LDAP settings



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
    resp, r, err := apiClient.AccessApi.GetLdapSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetLdapSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetLdapSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapSettingsRequest struct via the builder pattern


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


## GetLdapUserSchemaSettings

> Object GetLdapUserSchemaSettings(ctx).Execute()

Get LDAP user schema settings



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
    resp, r, err := apiClient.AccessApi.GetLdapUserSchemaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetLdapUserSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapUserSchemaSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetLdapUserSchemaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapUserSchemaSettingsRequest struct via the builder pattern


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


## GetLdapVpnRadiusSettings

> Object GetLdapVpnRadiusSettings(ctx).Execute()

Get LDAP VPN Radius settings



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
    resp, r, err := apiClient.AccessApi.GetLdapVpnRadiusSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetLdapVpnRadiusSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapVpnRadiusSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetLdapVpnRadiusSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapVpnRadiusSettingsRequest struct via the builder pattern


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


## GetLdapVpnSchemaSettings

> Object GetLdapVpnSchemaSettings(ctx).Execute()

Get LDAP VPN schema settings



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
    resp, r, err := apiClient.AccessApi.GetLdapVpnSchemaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetLdapVpnSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapVpnSchemaSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetLdapVpnSchemaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapVpnSchemaSettingsRequest struct via the builder pattern


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


## PostTestLdapGroupSchemaSettings

> PostTestLdapGroupSchemaSettings200Response PostTestLdapGroupSchemaSettings(ctx).PostTestLdapGroupSchemaSettingsRequest(postTestLdapGroupSchemaSettingsRequest).Execute()

Test LDAP group schema settings



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
    postTestLdapGroupSchemaSettingsRequest := *openapiclient.NewPostTestLdapGroupSchemaSettingsRequest("GroupBase_example", "GroupIdAttribute_example") // PostTestLdapGroupSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestLdapGroupSchemaSettings(context.Background()).PostTestLdapGroupSchemaSettingsRequest(postTestLdapGroupSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestLdapGroupSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestLdapGroupSchemaSettings`: PostTestLdapGroupSchemaSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestLdapGroupSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestLdapGroupSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestLdapGroupSchemaSettingsRequest** | [**PostTestLdapGroupSchemaSettingsRequest**](PostTestLdapGroupSchemaSettingsRequest.md) |  | 

### Return type

[**PostTestLdapGroupSchemaSettings200Response**](PostTestLdapGroupSchemaSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTestLdapSettings

> PostTestLdapSettings200Response PostTestLdapSettings(ctx).PostTestLdapSettingsRequest(postTestLdapSettingsRequest).Execute()

Test LDAP settings



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
    postTestLdapSettingsRequest := *openapiclient.NewPostTestLdapSettingsRequest("Host_example") // PostTestLdapSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestLdapSettings(context.Background()).PostTestLdapSettingsRequest(postTestLdapSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestLdapSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestLdapSettings`: PostTestLdapSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestLdapSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestLdapSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestLdapSettingsRequest** | [**PostTestLdapSettingsRequest**](PostTestLdapSettingsRequest.md) |  | 

### Return type

[**PostTestLdapSettings200Response**](PostTestLdapSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTestLdapUserSchemaSettings

> PostTestLdapUserSchemaSettings200Response PostTestLdapUserSchemaSettings(ctx).PostTestLdapUserSchemaSettingsRequest(postTestLdapUserSchemaSettingsRequest).Execute()

Test LDAP user schema settings



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
    postTestLdapUserSchemaSettingsRequest := *openapiclient.NewPostTestLdapUserSchemaSettingsRequest("UserBase_example", "UserIdAttribute_example") // PostTestLdapUserSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestLdapUserSchemaSettings(context.Background()).PostTestLdapUserSchemaSettingsRequest(postTestLdapUserSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestLdapUserSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestLdapUserSchemaSettings`: PostTestLdapUserSchemaSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestLdapUserSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestLdapUserSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestLdapUserSchemaSettingsRequest** | [**PostTestLdapUserSchemaSettingsRequest**](PostTestLdapUserSchemaSettingsRequest.md) |  | 

### Return type

[**PostTestLdapUserSchemaSettings200Response**](PostTestLdapUserSchemaSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTestLdapVpnSchemaSettings

> PostTestLdapUserSchemaSettings200Response PostTestLdapVpnSchemaSettings(ctx).PostTestLdapVpnSchemaSettingsRequest(postTestLdapVpnSchemaSettingsRequest).Execute()

Test LDAP VPN schema settings



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
    postTestLdapVpnSchemaSettingsRequest := *openapiclient.NewPostTestLdapVpnSchemaSettingsRequest("VpnGroupBase_example", "VpnGroupIdAttribute_example", "VpnGroupMemberAttribute_example") // PostTestLdapVpnSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestLdapVpnSchemaSettings(context.Background()).PostTestLdapVpnSchemaSettingsRequest(postTestLdapVpnSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestLdapVpnSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestLdapVpnSchemaSettings`: PostTestLdapUserSchemaSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestLdapVpnSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestLdapVpnSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestLdapVpnSchemaSettingsRequest** | [**PostTestLdapVpnSchemaSettingsRequest**](PostTestLdapVpnSchemaSettingsRequest.md) |  | 

### Return type

[**PostTestLdapUserSchemaSettings200Response**](PostTestLdapUserSchemaSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEnableLdap

> SimpleEnabledResponse PutEnableLdap(ctx).PutEnableLdapRequest(putEnableLdapRequest).Execute()

Enable/disable LDAP



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
    putEnableLdapRequest := *openapiclient.NewPutEnableLdapRequest(false) // PutEnableLdapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutEnableLdap(context.Background()).PutEnableLdapRequest(putEnableLdapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutEnableLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutEnableLdap`: SimpleEnabledResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutEnableLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEnableLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putEnableLdapRequest** | [**PutEnableLdapRequest**](PutEnableLdapRequest.md) |  | 

### Return type

[**SimpleEnabledResponse**](SimpleEnabledResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExpireAccessUrl

> Object PutExpireAccessUrl(ctx, accessUrlId).ExpireRequest(expireRequest).Execute()

Expire access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID
    expireRequest := *openapiclient.NewExpireRequest() // ExpireRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutExpireAccessUrl(context.Background(), accessUrlId).ExpireRequest(expireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutExpireAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExpireAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutExpireAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExpireAccessUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireRequest** | [**ExpireRequest**](ExpireRequest.md) |  | 

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


## PutExpireApiToken

> Object PutExpireApiToken(ctx, tokenId).ExpireRequest(expireRequest).Execute()

Expire API token



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
    tokenId := int32(56) // int32 | Token ID
    expireRequest := *openapiclient.NewExpireRequest() // ExpireRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutExpireApiToken(context.Background(), tokenId).ExpireRequest(expireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutExpireApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExpireApiToken`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutExpireApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExpireApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireRequest** | [**ExpireRequest**](ExpireRequest.md) |  | 

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


## PutLdapGroupSchemaSettings

> LdapGroupSettingsResponse PutLdapGroupSchemaSettings(ctx).PutLdapGroupSchemaSettingsRequest(putLdapGroupSchemaSettingsRequest).Execute()

Put LDAP group schema settings



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
    putLdapGroupSchemaSettingsRequest := *openapiclient.NewPutLdapGroupSchemaSettingsRequest(false) // PutLdapGroupSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutLdapGroupSchemaSettings(context.Background()).PutLdapGroupSchemaSettingsRequest(putLdapGroupSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutLdapGroupSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLdapGroupSchemaSettings`: LdapGroupSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutLdapGroupSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLdapGroupSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLdapGroupSchemaSettingsRequest** | [**PutLdapGroupSchemaSettingsRequest**](PutLdapGroupSchemaSettingsRequest.md) |  | 

### Return type

[**LdapGroupSettingsResponse**](LdapGroupSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLdapSettings

> LdapSettingsResponse PutLdapSettings(ctx).PutLdapSettingsRequest(putLdapSettingsRequest).Execute()

Put LDAP settings



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
    putLdapSettingsRequest := *openapiclient.NewPutLdapSettingsRequest("Host_example") // PutLdapSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutLdapSettings(context.Background()).PutLdapSettingsRequest(putLdapSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutLdapSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLdapSettings`: LdapSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutLdapSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLdapSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLdapSettingsRequest** | [**PutLdapSettingsRequest**](PutLdapSettingsRequest.md) |  | 

### Return type

[**LdapSettingsResponse**](LdapSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLdapUserSchemaSettings

> LdapUserSettingsResponse PutLdapUserSchemaSettings(ctx).PutLdapUserSchemaSettingsRequest(putLdapUserSchemaSettingsRequest).Execute()

Put LDAP user schema settings



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
    putLdapUserSchemaSettingsRequest := *openapiclient.NewPutLdapUserSchemaSettingsRequest("UserBase_example", "UserIdAttribute_example") // PutLdapUserSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutLdapUserSchemaSettings(context.Background()).PutLdapUserSchemaSettingsRequest(putLdapUserSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutLdapUserSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLdapUserSchemaSettings`: LdapUserSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutLdapUserSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLdapUserSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLdapUserSchemaSettingsRequest** | [**PutLdapUserSchemaSettingsRequest**](PutLdapUserSchemaSettingsRequest.md) |  | 

### Return type

[**LdapUserSettingsResponse**](LdapUserSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLdapVpnRadiusSettings

> LdapVpnRadiusSettingsResponse PutLdapVpnRadiusSettings(ctx).PutLdapVpnRadiusSettingsRequest(putLdapVpnRadiusSettingsRequest).Execute()

Put LDAP VPN Radius settings



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
    putLdapVpnRadiusSettingsRequest := *openapiclient.NewPutLdapVpnRadiusSettingsRequest(false, "VpnRadiusServer_example", "VpnRadiusPass_example") // PutLdapVpnRadiusSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutLdapVpnRadiusSettings(context.Background()).PutLdapVpnRadiusSettingsRequest(putLdapVpnRadiusSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutLdapVpnRadiusSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLdapVpnRadiusSettings`: LdapVpnRadiusSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutLdapVpnRadiusSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLdapVpnRadiusSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLdapVpnRadiusSettingsRequest** | [**PutLdapVpnRadiusSettingsRequest**](PutLdapVpnRadiusSettingsRequest.md) |  | 

### Return type

[**LdapVpnRadiusSettingsResponse**](LdapVpnRadiusSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLdapVpnSchemaSettings

> LdapVpnSchemaSettingsResponse PutLdapVpnSchemaSettings(ctx).PutLdapVpnSchemaSettingsRequest(putLdapVpnSchemaSettingsRequest).Execute()

Put LDAP VPN schema settings



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
    putLdapVpnSchemaSettingsRequest := *openapiclient.NewPutLdapVpnSchemaSettingsRequest(false, "VpnGroupBase_example", "VpnGroupIdAttribute_example", "VpnGroupMemberAttribute_example") // PutLdapVpnSchemaSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutLdapVpnSchemaSettings(context.Background()).PutLdapVpnSchemaSettingsRequest(putLdapVpnSchemaSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutLdapVpnSchemaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLdapVpnSchemaSettings`: LdapVpnSchemaSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutLdapVpnSchemaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLdapVpnSchemaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLdapVpnSchemaSettingsRequest** | [**PutLdapVpnSchemaSettingsRequest**](PutLdapVpnSchemaSettingsRequest.md) |  | 

### Return type

[**LdapVpnSchemaSettingsResponse**](LdapVpnSchemaSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUploadLdapAuthCert

> PutUploadLdapAuthCert200Response PutUploadLdapAuthCert(ctx).Body(body).Execute()

Upload LDAP Auth Cert



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
    body := os.NewFile(1234, "some_file") // *os.File | Authentication cert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutUploadLdapAuthCert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutUploadLdapAuthCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUploadLdapAuthCert`: PutUploadLdapAuthCert200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutUploadLdapAuthCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUploadLdapAuthCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | Authentication cert | 

### Return type

[**PutUploadLdapAuthCert200Response**](PutUploadLdapAuthCert200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUploadLdapAuthKey

> PutUploadLdapAuthCert200Response PutUploadLdapAuthKey(ctx).Body(body).Execute()

Upload LDAP Auth Key



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
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutUploadLdapAuthKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutUploadLdapAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUploadLdapAuthKey`: PutUploadLdapAuthCert200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutUploadLdapAuthKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUploadLdapAuthKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**PutUploadLdapAuthCert200Response**](PutUploadLdapAuthCert200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUploadLdapCaCert

> PutUploadLdapAuthCert200Response PutUploadLdapCaCert(ctx).Body(body).Execute()

Upload LDAP CA cert



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
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutUploadLdapCaCert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutUploadLdapCaCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUploadLdapCaCert`: PutUploadLdapAuthCert200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutUploadLdapCaCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUploadLdapCaCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**PutUploadLdapAuthCert200Response**](PutUploadLdapAuthCert200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

