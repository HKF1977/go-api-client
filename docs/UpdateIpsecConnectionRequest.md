# UpdateIpsecConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the connection. | [optional] 
**Description** | Pointer to **string** | Description of this IPsec endpoint | [optional] 
**Ipaddress** | Pointer to **string** | IP of the remote gateway | [optional] 
**Secret** | Pointer to **string** | Pre-shared key | [optional] 
**Pfs** | Pointer to **bool** | Perfect Forward Secrecy if true, disables if false. | [optional] 
**IkeVersion** | Pointer to **int32** | Version for IKE algorithm | [optional] 
**NatTEnabled** | Pointer to **bool** | True if you want encapsulated IPsec protocol to this gateway | [optional] 
**ExtraConfig** | Pointer to **string** | Additional optionals for connection such as &#39;phase1&#x3D;aes256_gcm-sha2_256-dh14 phase2&#x3D;aes256_gcm&#39; | [optional] 
**PrivateIpaddress** | Pointer to **string** | Internal NAT address of the remote gateway | [optional] 
**Gre** | Pointer to **bool** | True if GRE is being used for the specific endpoint | [optional] 
**GreInterfaceAddress** | Pointer to **string** | Interface address for GRE | [optional] 
**VpnType** | Pointer to **string** | policy, gre, vti | [optional] 
**RouteBasedIntAddress** | Pointer to **string** |  | [optional] 
**RouteBasedLocal** | Pointer to **string** |  | [optional] 
**RouteBasedRemote** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateIpsecConnectionRequest

`func NewUpdateIpsecConnectionRequest() *UpdateIpsecConnectionRequest`

NewUpdateIpsecConnectionRequest instantiates a new UpdateIpsecConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpsecConnectionRequestWithDefaults

`func NewUpdateIpsecConnectionRequestWithDefaults() *UpdateIpsecConnectionRequest`

NewUpdateIpsecConnectionRequestWithDefaults instantiates a new UpdateIpsecConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIpsecConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIpsecConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIpsecConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIpsecConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIpsecConnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIpsecConnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIpsecConnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIpsecConnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpaddress

`func (o *UpdateIpsecConnectionRequest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *UpdateIpsecConnectionRequest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *UpdateIpsecConnectionRequest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *UpdateIpsecConnectionRequest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateIpsecConnectionRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateIpsecConnectionRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateIpsecConnectionRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateIpsecConnectionRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPfs

`func (o *UpdateIpsecConnectionRequest) GetPfs() bool`

GetPfs returns the Pfs field if non-nil, zero value otherwise.

### GetPfsOk

`func (o *UpdateIpsecConnectionRequest) GetPfsOk() (*bool, bool)`

GetPfsOk returns a tuple with the Pfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfs

`func (o *UpdateIpsecConnectionRequest) SetPfs(v bool)`

SetPfs sets Pfs field to given value.

### HasPfs

`func (o *UpdateIpsecConnectionRequest) HasPfs() bool`

HasPfs returns a boolean if a field has been set.

### GetIkeVersion

`func (o *UpdateIpsecConnectionRequest) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *UpdateIpsecConnectionRequest) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *UpdateIpsecConnectionRequest) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *UpdateIpsecConnectionRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNatTEnabled

`func (o *UpdateIpsecConnectionRequest) GetNatTEnabled() bool`

GetNatTEnabled returns the NatTEnabled field if non-nil, zero value otherwise.

### GetNatTEnabledOk

`func (o *UpdateIpsecConnectionRequest) GetNatTEnabledOk() (*bool, bool)`

GetNatTEnabledOk returns a tuple with the NatTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTEnabled

`func (o *UpdateIpsecConnectionRequest) SetNatTEnabled(v bool)`

SetNatTEnabled sets NatTEnabled field to given value.

### HasNatTEnabled

`func (o *UpdateIpsecConnectionRequest) HasNatTEnabled() bool`

HasNatTEnabled returns a boolean if a field has been set.

### GetExtraConfig

`func (o *UpdateIpsecConnectionRequest) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *UpdateIpsecConnectionRequest) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *UpdateIpsecConnectionRequest) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *UpdateIpsecConnectionRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *UpdateIpsecConnectionRequest) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *UpdateIpsecConnectionRequest) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *UpdateIpsecConnectionRequest) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *UpdateIpsecConnectionRequest) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetGre

`func (o *UpdateIpsecConnectionRequest) GetGre() bool`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *UpdateIpsecConnectionRequest) GetGreOk() (*bool, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *UpdateIpsecConnectionRequest) SetGre(v bool)`

SetGre sets Gre field to given value.

### HasGre

`func (o *UpdateIpsecConnectionRequest) HasGre() bool`

HasGre returns a boolean if a field has been set.

### GetGreInterfaceAddress

`func (o *UpdateIpsecConnectionRequest) GetGreInterfaceAddress() string`

GetGreInterfaceAddress returns the GreInterfaceAddress field if non-nil, zero value otherwise.

### GetGreInterfaceAddressOk

`func (o *UpdateIpsecConnectionRequest) GetGreInterfaceAddressOk() (*string, bool)`

GetGreInterfaceAddressOk returns a tuple with the GreInterfaceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreInterfaceAddress

`func (o *UpdateIpsecConnectionRequest) SetGreInterfaceAddress(v string)`

SetGreInterfaceAddress sets GreInterfaceAddress field to given value.

### HasGreInterfaceAddress

`func (o *UpdateIpsecConnectionRequest) HasGreInterfaceAddress() bool`

HasGreInterfaceAddress returns a boolean if a field has been set.

### GetVpnType

`func (o *UpdateIpsecConnectionRequest) GetVpnType() string`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *UpdateIpsecConnectionRequest) GetVpnTypeOk() (*string, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *UpdateIpsecConnectionRequest) SetVpnType(v string)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *UpdateIpsecConnectionRequest) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetRouteBasedIntAddress

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedIntAddress() string`

GetRouteBasedIntAddress returns the RouteBasedIntAddress field if non-nil, zero value otherwise.

### GetRouteBasedIntAddressOk

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedIntAddressOk() (*string, bool)`

GetRouteBasedIntAddressOk returns a tuple with the RouteBasedIntAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedIntAddress

`func (o *UpdateIpsecConnectionRequest) SetRouteBasedIntAddress(v string)`

SetRouteBasedIntAddress sets RouteBasedIntAddress field to given value.

### HasRouteBasedIntAddress

`func (o *UpdateIpsecConnectionRequest) HasRouteBasedIntAddress() bool`

HasRouteBasedIntAddress returns a boolean if a field has been set.

### GetRouteBasedLocal

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedLocal() string`

GetRouteBasedLocal returns the RouteBasedLocal field if non-nil, zero value otherwise.

### GetRouteBasedLocalOk

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedLocalOk() (*string, bool)`

GetRouteBasedLocalOk returns a tuple with the RouteBasedLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedLocal

`func (o *UpdateIpsecConnectionRequest) SetRouteBasedLocal(v string)`

SetRouteBasedLocal sets RouteBasedLocal field to given value.

### HasRouteBasedLocal

`func (o *UpdateIpsecConnectionRequest) HasRouteBasedLocal() bool`

HasRouteBasedLocal returns a boolean if a field has been set.

### GetRouteBasedRemote

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedRemote() string`

GetRouteBasedRemote returns the RouteBasedRemote field if non-nil, zero value otherwise.

### GetRouteBasedRemoteOk

`func (o *UpdateIpsecConnectionRequest) GetRouteBasedRemoteOk() (*string, bool)`

GetRouteBasedRemoteOk returns a tuple with the RouteBasedRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedRemote

`func (o *UpdateIpsecConnectionRequest) SetRouteBasedRemote(v string)`

SetRouteBasedRemote sets RouteBasedRemote field to given value.

### HasRouteBasedRemote

`func (o *UpdateIpsecConnectionRequest) HasRouteBasedRemote() bool`

HasRouteBasedRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


