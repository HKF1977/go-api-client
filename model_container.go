/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 5.1.5
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Container struct for Container
type Container struct {
	Id *string `json:"Id,omitempty"`
	Created *string `json:"Created,omitempty"`
	Path *string `json:"Path,omitempty"`
	RestartCount *int32 `json:"RestartCount,omitempty"`
	Args []string `json:"Args,omitempty"`
	AppArmorProfile *string `json:"AppArmorProfile,omitempty"`
	Image *string `json:"Image,omitempty"`
	Config NullableContainerConfig `json:"Config,omitempty"`
	State *ContainerState `json:"State,omitempty"`
	NetworkSettings *ContainerNetworkSettings `json:"NetworkSettings,omitempty"`
	ResolvConfPath *string `json:"ResolvConfPath,omitempty"`
	HostnamePath *string `json:"HostnamePath,omitempty"`
	HostsPath *string `json:"HostsPath,omitempty"`
	LogPath *string `json:"LogPath,omitempty"`
	Name *string `json:"Name,omitempty"`
	Platform *string `json:"Platform,omitempty"`
	Driver *string `json:"Driver,omitempty"`
	ExecDriver *string `json:"ExecDriver,omitempty"`
	MountLabel *string `json:"MountLabel,omitempty"`
	ProcessLabel *string `json:"ProcessLabel,omitempty"`
	ExecIDs []string `json:"ExecIDs,omitempty"`
	Volumes map[string]interface{} `json:"Volumes,omitempty"`
	VolumesRW map[string]interface{} `json:"VolumesRW,omitempty"`
	HostConfig map[string]interface{} `json:"HostConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Container Container

// NewContainer instantiates a new Container object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainer() *Container {
	this := Container{}
	return &this
}

// NewContainerWithDefaults instantiates a new Container object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWithDefaults() *Container {
	this := Container{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Container) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Container) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Container) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Container) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Container) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Container) SetCreated(v string) {
	o.Created = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Container) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Container) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Container) SetPath(v string) {
	o.Path = &v
}

// GetRestartCount returns the RestartCount field value if set, zero value otherwise.
func (o *Container) GetRestartCount() int32 {
	if o == nil || o.RestartCount == nil {
		var ret int32
		return ret
	}
	return *o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetRestartCountOk() (*int32, bool) {
	if o == nil || o.RestartCount == nil {
		return nil, false
	}
	return o.RestartCount, true
}

// HasRestartCount returns a boolean if a field has been set.
func (o *Container) HasRestartCount() bool {
	if o != nil && o.RestartCount != nil {
		return true
	}

	return false
}

// SetRestartCount gets a reference to the given int32 and assigns it to the RestartCount field.
func (o *Container) SetRestartCount(v int32) {
	o.RestartCount = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *Container) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetArgsOk() ([]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *Container) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *Container) SetArgs(v []string) {
	o.Args = v
}

// GetAppArmorProfile returns the AppArmorProfile field value if set, zero value otherwise.
func (o *Container) GetAppArmorProfile() string {
	if o == nil || o.AppArmorProfile == nil {
		var ret string
		return ret
	}
	return *o.AppArmorProfile
}

// GetAppArmorProfileOk returns a tuple with the AppArmorProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetAppArmorProfileOk() (*string, bool) {
	if o == nil || o.AppArmorProfile == nil {
		return nil, false
	}
	return o.AppArmorProfile, true
}

// HasAppArmorProfile returns a boolean if a field has been set.
func (o *Container) HasAppArmorProfile() bool {
	if o != nil && o.AppArmorProfile != nil {
		return true
	}

	return false
}

// SetAppArmorProfile gets a reference to the given string and assigns it to the AppArmorProfile field.
func (o *Container) SetAppArmorProfile(v string) {
	o.AppArmorProfile = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Container) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Container) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Container) SetImage(v string) {
	o.Image = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetConfig() ContainerConfig {
	if o == nil || o.Config.Get() == nil {
		var ret ContainerConfig
		return ret
	}
	return *o.Config.Get()
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetConfigOk() (*ContainerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config.Get(), o.Config.IsSet()
}

// HasConfig returns a boolean if a field has been set.
func (o *Container) HasConfig() bool {
	if o != nil && o.Config.IsSet() {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NullableContainerConfig and assigns it to the Config field.
func (o *Container) SetConfig(v ContainerConfig) {
	o.Config.Set(&v)
}
// SetConfigNil sets the value for Config to be an explicit nil
func (o *Container) SetConfigNil() {
	o.Config.Set(nil)
}

// UnsetConfig ensures that no value is present for Config, not even an explicit nil
func (o *Container) UnsetConfig() {
	o.Config.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Container) GetState() ContainerState {
	if o == nil || o.State == nil {
		var ret ContainerState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetStateOk() (*ContainerState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Container) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ContainerState and assigns it to the State field.
func (o *Container) SetState(v ContainerState) {
	o.State = &v
}

// GetNetworkSettings returns the NetworkSettings field value if set, zero value otherwise.
func (o *Container) GetNetworkSettings() ContainerNetworkSettings {
	if o == nil || o.NetworkSettings == nil {
		var ret ContainerNetworkSettings
		return ret
	}
	return *o.NetworkSettings
}

// GetNetworkSettingsOk returns a tuple with the NetworkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetNetworkSettingsOk() (*ContainerNetworkSettings, bool) {
	if o == nil || o.NetworkSettings == nil {
		return nil, false
	}
	return o.NetworkSettings, true
}

// HasNetworkSettings returns a boolean if a field has been set.
func (o *Container) HasNetworkSettings() bool {
	if o != nil && o.NetworkSettings != nil {
		return true
	}

	return false
}

// SetNetworkSettings gets a reference to the given ContainerNetworkSettings and assigns it to the NetworkSettings field.
func (o *Container) SetNetworkSettings(v ContainerNetworkSettings) {
	o.NetworkSettings = &v
}

// GetResolvConfPath returns the ResolvConfPath field value if set, zero value otherwise.
func (o *Container) GetResolvConfPath() string {
	if o == nil || o.ResolvConfPath == nil {
		var ret string
		return ret
	}
	return *o.ResolvConfPath
}

// GetResolvConfPathOk returns a tuple with the ResolvConfPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetResolvConfPathOk() (*string, bool) {
	if o == nil || o.ResolvConfPath == nil {
		return nil, false
	}
	return o.ResolvConfPath, true
}

// HasResolvConfPath returns a boolean if a field has been set.
func (o *Container) HasResolvConfPath() bool {
	if o != nil && o.ResolvConfPath != nil {
		return true
	}

	return false
}

// SetResolvConfPath gets a reference to the given string and assigns it to the ResolvConfPath field.
func (o *Container) SetResolvConfPath(v string) {
	o.ResolvConfPath = &v
}

// GetHostnamePath returns the HostnamePath field value if set, zero value otherwise.
func (o *Container) GetHostnamePath() string {
	if o == nil || o.HostnamePath == nil {
		var ret string
		return ret
	}
	return *o.HostnamePath
}

// GetHostnamePathOk returns a tuple with the HostnamePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetHostnamePathOk() (*string, bool) {
	if o == nil || o.HostnamePath == nil {
		return nil, false
	}
	return o.HostnamePath, true
}

// HasHostnamePath returns a boolean if a field has been set.
func (o *Container) HasHostnamePath() bool {
	if o != nil && o.HostnamePath != nil {
		return true
	}

	return false
}

// SetHostnamePath gets a reference to the given string and assigns it to the HostnamePath field.
func (o *Container) SetHostnamePath(v string) {
	o.HostnamePath = &v
}

// GetHostsPath returns the HostsPath field value if set, zero value otherwise.
func (o *Container) GetHostsPath() string {
	if o == nil || o.HostsPath == nil {
		var ret string
		return ret
	}
	return *o.HostsPath
}

// GetHostsPathOk returns a tuple with the HostsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetHostsPathOk() (*string, bool) {
	if o == nil || o.HostsPath == nil {
		return nil, false
	}
	return o.HostsPath, true
}

// HasHostsPath returns a boolean if a field has been set.
func (o *Container) HasHostsPath() bool {
	if o != nil && o.HostsPath != nil {
		return true
	}

	return false
}

// SetHostsPath gets a reference to the given string and assigns it to the HostsPath field.
func (o *Container) SetHostsPath(v string) {
	o.HostsPath = &v
}

// GetLogPath returns the LogPath field value if set, zero value otherwise.
func (o *Container) GetLogPath() string {
	if o == nil || o.LogPath == nil {
		var ret string
		return ret
	}
	return *o.LogPath
}

// GetLogPathOk returns a tuple with the LogPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetLogPathOk() (*string, bool) {
	if o == nil || o.LogPath == nil {
		return nil, false
	}
	return o.LogPath, true
}

// HasLogPath returns a boolean if a field has been set.
func (o *Container) HasLogPath() bool {
	if o != nil && o.LogPath != nil {
		return true
	}

	return false
}

// SetLogPath gets a reference to the given string and assigns it to the LogPath field.
func (o *Container) SetLogPath(v string) {
	o.LogPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Container) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Container) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Container) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *Container) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *Container) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *Container) SetPlatform(v string) {
	o.Platform = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *Container) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *Container) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *Container) SetDriver(v string) {
	o.Driver = &v
}

// GetExecDriver returns the ExecDriver field value if set, zero value otherwise.
func (o *Container) GetExecDriver() string {
	if o == nil || o.ExecDriver == nil {
		var ret string
		return ret
	}
	return *o.ExecDriver
}

// GetExecDriverOk returns a tuple with the ExecDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetExecDriverOk() (*string, bool) {
	if o == nil || o.ExecDriver == nil {
		return nil, false
	}
	return o.ExecDriver, true
}

// HasExecDriver returns a boolean if a field has been set.
func (o *Container) HasExecDriver() bool {
	if o != nil && o.ExecDriver != nil {
		return true
	}

	return false
}

// SetExecDriver gets a reference to the given string and assigns it to the ExecDriver field.
func (o *Container) SetExecDriver(v string) {
	o.ExecDriver = &v
}

// GetMountLabel returns the MountLabel field value if set, zero value otherwise.
func (o *Container) GetMountLabel() string {
	if o == nil || o.MountLabel == nil {
		var ret string
		return ret
	}
	return *o.MountLabel
}

// GetMountLabelOk returns a tuple with the MountLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetMountLabelOk() (*string, bool) {
	if o == nil || o.MountLabel == nil {
		return nil, false
	}
	return o.MountLabel, true
}

// HasMountLabel returns a boolean if a field has been set.
func (o *Container) HasMountLabel() bool {
	if o != nil && o.MountLabel != nil {
		return true
	}

	return false
}

// SetMountLabel gets a reference to the given string and assigns it to the MountLabel field.
func (o *Container) SetMountLabel(v string) {
	o.MountLabel = &v
}

// GetProcessLabel returns the ProcessLabel field value if set, zero value otherwise.
func (o *Container) GetProcessLabel() string {
	if o == nil || o.ProcessLabel == nil {
		var ret string
		return ret
	}
	return *o.ProcessLabel
}

// GetProcessLabelOk returns a tuple with the ProcessLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetProcessLabelOk() (*string, bool) {
	if o == nil || o.ProcessLabel == nil {
		return nil, false
	}
	return o.ProcessLabel, true
}

// HasProcessLabel returns a boolean if a field has been set.
func (o *Container) HasProcessLabel() bool {
	if o != nil && o.ProcessLabel != nil {
		return true
	}

	return false
}

// SetProcessLabel gets a reference to the given string and assigns it to the ProcessLabel field.
func (o *Container) SetProcessLabel(v string) {
	o.ProcessLabel = &v
}

// GetExecIDs returns the ExecIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetExecIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExecIDs
}

// GetExecIDsOk returns a tuple with the ExecIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetExecIDsOk() ([]string, bool) {
	if o == nil || o.ExecIDs == nil {
		return nil, false
	}
	return o.ExecIDs, true
}

// HasExecIDs returns a boolean if a field has been set.
func (o *Container) HasExecIDs() bool {
	if o != nil && o.ExecIDs != nil {
		return true
	}

	return false
}

// SetExecIDs gets a reference to the given []string and assigns it to the ExecIDs field.
func (o *Container) SetExecIDs(v []string) {
	o.ExecIDs = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetVolumes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetVolumesOk() (map[string]interface{}, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Container) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given map[string]interface{} and assigns it to the Volumes field.
func (o *Container) SetVolumes(v map[string]interface{}) {
	o.Volumes = v
}

// GetVolumesRW returns the VolumesRW field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetVolumesRW() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumesRW
}

// GetVolumesRWOk returns a tuple with the VolumesRW field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetVolumesRWOk() (map[string]interface{}, bool) {
	if o == nil || o.VolumesRW == nil {
		return nil, false
	}
	return o.VolumesRW, true
}

// HasVolumesRW returns a boolean if a field has been set.
func (o *Container) HasVolumesRW() bool {
	if o != nil && o.VolumesRW != nil {
		return true
	}

	return false
}

// SetVolumesRW gets a reference to the given map[string]interface{} and assigns it to the VolumesRW field.
func (o *Container) SetVolumesRW(v map[string]interface{}) {
	o.VolumesRW = v
}

// GetHostConfig returns the HostConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetHostConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.HostConfig
}

// GetHostConfigOk returns a tuple with the HostConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetHostConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.HostConfig == nil {
		return nil, false
	}
	return o.HostConfig, true
}

// HasHostConfig returns a boolean if a field has been set.
func (o *Container) HasHostConfig() bool {
	if o != nil && o.HostConfig != nil {
		return true
	}

	return false
}

// SetHostConfig gets a reference to the given map[string]interface{} and assigns it to the HostConfig field.
func (o *Container) SetHostConfig(v map[string]interface{}) {
	o.HostConfig = v
}

func (o Container) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Created != nil {
		toSerialize["Created"] = o.Created
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.RestartCount != nil {
		toSerialize["RestartCount"] = o.RestartCount
	}
	if o.Args != nil {
		toSerialize["Args"] = o.Args
	}
	if o.AppArmorProfile != nil {
		toSerialize["AppArmorProfile"] = o.AppArmorProfile
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Config.IsSet() {
		toSerialize["Config"] = o.Config.Get()
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.NetworkSettings != nil {
		toSerialize["NetworkSettings"] = o.NetworkSettings
	}
	if o.ResolvConfPath != nil {
		toSerialize["ResolvConfPath"] = o.ResolvConfPath
	}
	if o.HostnamePath != nil {
		toSerialize["HostnamePath"] = o.HostnamePath
	}
	if o.HostsPath != nil {
		toSerialize["HostsPath"] = o.HostsPath
	}
	if o.LogPath != nil {
		toSerialize["LogPath"] = o.LogPath
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}
	if o.Driver != nil {
		toSerialize["Driver"] = o.Driver
	}
	if o.ExecDriver != nil {
		toSerialize["ExecDriver"] = o.ExecDriver
	}
	if o.MountLabel != nil {
		toSerialize["MountLabel"] = o.MountLabel
	}
	if o.ProcessLabel != nil {
		toSerialize["ProcessLabel"] = o.ProcessLabel
	}
	if o.ExecIDs != nil {
		toSerialize["ExecIDs"] = o.ExecIDs
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}
	if o.VolumesRW != nil {
		toSerialize["VolumesRW"] = o.VolumesRW
	}
	if o.HostConfig != nil {
		toSerialize["HostConfig"] = o.HostConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Container) UnmarshalJSON(bytes []byte) (err error) {
	varContainer := _Container{}

	if err = json.Unmarshal(bytes, &varContainer); err == nil {
		*o = Container(varContainer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Created")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "RestartCount")
		delete(additionalProperties, "Args")
		delete(additionalProperties, "AppArmorProfile")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Config")
		delete(additionalProperties, "State")
		delete(additionalProperties, "NetworkSettings")
		delete(additionalProperties, "ResolvConfPath")
		delete(additionalProperties, "HostnamePath")
		delete(additionalProperties, "HostsPath")
		delete(additionalProperties, "LogPath")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Platform")
		delete(additionalProperties, "Driver")
		delete(additionalProperties, "ExecDriver")
		delete(additionalProperties, "MountLabel")
		delete(additionalProperties, "ProcessLabel")
		delete(additionalProperties, "ExecIDs")
		delete(additionalProperties, "Volumes")
		delete(additionalProperties, "VolumesRW")
		delete(additionalProperties, "HostConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainer struct {
	value *Container
	isSet bool
}

func (v NullableContainer) Get() *Container {
	return v.value
}

func (v *NullableContainer) Set(val *Container) {
	v.value = val
	v.isSet = true
}

func (v NullableContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainer(val *Container) *NullableContainer {
	return &NullableContainer{value: val, isSet: true}
}

func (v NullableContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


