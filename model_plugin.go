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

// Plugin struct for Plugin
type Plugin struct {
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	TagName *string `json:"tag_name,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMsg *string `json:"status_msg,omitempty"`
	ImportId *string `json:"import_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Containers []PluginContainersInner `json:"containers,omitempty"`
	RunningContainers *int32 `json:"running_containers,omitempty"`
	Data *PluginData `json:"data,omitempty"`
}

// NewPlugin instantiates a new Plugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlugin() *Plugin {
	this := Plugin{}
	return &this
}

// NewPluginWithDefaults instantiates a new Plugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginWithDefaults() *Plugin {
	this := Plugin{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Plugin) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Plugin) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Plugin) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plugin) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plugin) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *Plugin) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *Plugin) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *Plugin) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *Plugin) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plugin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plugin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plugin) SetName(v string) {
	o.Name = &v
}

// GetTagName returns the TagName field value if set, zero value otherwise.
func (o *Plugin) GetTagName() string {
	if o == nil || o.TagName == nil {
		var ret string
		return ret
	}
	return *o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetTagNameOk() (*string, bool) {
	if o == nil || o.TagName == nil {
		return nil, false
	}
	return o.TagName, true
}

// HasTagName returns a boolean if a field has been set.
func (o *Plugin) HasTagName() bool {
	if o != nil && o.TagName != nil {
		return true
	}

	return false
}

// SetTagName gets a reference to the given string and assigns it to the TagName field.
func (o *Plugin) SetTagName(v string) {
	o.TagName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Plugin) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Plugin) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Plugin) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMsg returns the StatusMsg field value if set, zero value otherwise.
func (o *Plugin) GetStatusMsg() string {
	if o == nil || o.StatusMsg == nil {
		var ret string
		return ret
	}
	return *o.StatusMsg
}

// GetStatusMsgOk returns a tuple with the StatusMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetStatusMsgOk() (*string, bool) {
	if o == nil || o.StatusMsg == nil {
		return nil, false
	}
	return o.StatusMsg, true
}

// HasStatusMsg returns a boolean if a field has been set.
func (o *Plugin) HasStatusMsg() bool {
	if o != nil && o.StatusMsg != nil {
		return true
	}

	return false
}

// SetStatusMsg gets a reference to the given string and assigns it to the StatusMsg field.
func (o *Plugin) SetStatusMsg(v string) {
	o.StatusMsg = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Plugin) GetImportId() string {
	if o == nil || o.ImportId == nil {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetImportIdOk() (*string, bool) {
	if o == nil || o.ImportId == nil {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Plugin) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *Plugin) SetImportId(v string) {
	o.ImportId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Plugin) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Plugin) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Plugin) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Plugin) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Plugin) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Plugin) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Plugin) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Plugin) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Plugin) SetUrl(v string) {
	o.Url = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *Plugin) GetContainers() []PluginContainersInner {
	if o == nil || o.Containers == nil {
		var ret []PluginContainersInner
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetContainersOk() ([]PluginContainersInner, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *Plugin) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []PluginContainersInner and assigns it to the Containers field.
func (o *Plugin) SetContainers(v []PluginContainersInner) {
	o.Containers = v
}

// GetRunningContainers returns the RunningContainers field value if set, zero value otherwise.
func (o *Plugin) GetRunningContainers() int32 {
	if o == nil || o.RunningContainers == nil {
		var ret int32
		return ret
	}
	return *o.RunningContainers
}

// GetRunningContainersOk returns a tuple with the RunningContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetRunningContainersOk() (*int32, bool) {
	if o == nil || o.RunningContainers == nil {
		return nil, false
	}
	return o.RunningContainers, true
}

// HasRunningContainers returns a boolean if a field has been set.
func (o *Plugin) HasRunningContainers() bool {
	if o != nil && o.RunningContainers != nil {
		return true
	}

	return false
}

// SetRunningContainers gets a reference to the given int32 and assigns it to the RunningContainers field.
func (o *Plugin) SetRunningContainers(v int32) {
	o.RunningContainers = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Plugin) GetData() PluginData {
	if o == nil || o.Data == nil {
		var ret PluginData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetDataOk() (*PluginData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Plugin) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PluginData and assigns it to the Data field.
func (o *Plugin) SetData(v PluginData) {
	o.Data = &v
}

func (o Plugin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TagName != nil {
		toSerialize["tag_name"] = o.TagName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMsg != nil {
		toSerialize["status_msg"] = o.StatusMsg
	}
	if o.ImportId != nil {
		toSerialize["import_id"] = o.ImportId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.RunningContainers != nil {
		toSerialize["running_containers"] = o.RunningContainers
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePlugin struct {
	value *Plugin
	isSet bool
}

func (v NullablePlugin) Get() *Plugin {
	return v.value
}

func (v *NullablePlugin) Set(val *Plugin) {
	v.value = val
	v.isSet = true
}

func (v NullablePlugin) IsSet() bool {
	return v.isSet
}

func (v *NullablePlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlugin(val *Plugin) *NullablePlugin {
	return &NullablePlugin{value: val, isSet: true}
}

func (v NullablePlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


