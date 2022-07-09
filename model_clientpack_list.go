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

// ClientpackList struct for ClientpackList
type ClientpackList struct {
	Name *string `json:"name,omitempty"`
	OverlayIpaddress *string `json:"overlay_ipaddress,omitempty"`
	LinuxOnefile *string `json:"linux_onefile,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ConfSha1 *string `json:"conf_sha1,omitempty"`
	WindowsOnefile *string `json:"windows_onefile,omitempty"`
	OvpnSha1 *string `json:"ovpn_sha1,omitempty"`
	TarballFile *string `json:"tarball_file,omitempty"`
	TarballSha1 *string `json:"tarball_sha1,omitempty"`
	SequentialId *int32 `json:"sequential_id,omitempty"`
	CheckedOut *bool `json:"checked_out,omitempty"`
	ZipSha1 *string `json:"zip_sha1,omitempty"`
	ZipFile *string `json:"zip_file,omitempty"`
	LastConnect *string `json:"last_connect,omitempty"`
	LastDisconnect *string `json:"last_disconnect,omitempty"`
	// VPN compression status. on, off or error message.
	Compression *string `json:"compression,omitempty"`
	ManagerId *string `json:"manager_id,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	Status *string `json:"status,omitempty"`
	Connected *bool `json:"connected,omitempty"`
	// Key, value object of tags
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewClientpackList instantiates a new ClientpackList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackList() *ClientpackList {
	this := ClientpackList{}
	return &this
}

// NewClientpackListWithDefaults instantiates a new ClientpackList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackListWithDefaults() *ClientpackList {
	this := ClientpackList{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientpackList) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientpackList) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientpackList) SetName(v string) {
	o.Name = &v
}

// GetOverlayIpaddress returns the OverlayIpaddress field value if set, zero value otherwise.
func (o *ClientpackList) GetOverlayIpaddress() string {
	if o == nil || o.OverlayIpaddress == nil {
		var ret string
		return ret
	}
	return *o.OverlayIpaddress
}

// GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetOverlayIpaddressOk() (*string, bool) {
	if o == nil || o.OverlayIpaddress == nil {
		return nil, false
	}
	return o.OverlayIpaddress, true
}

// HasOverlayIpaddress returns a boolean if a field has been set.
func (o *ClientpackList) HasOverlayIpaddress() bool {
	if o != nil && o.OverlayIpaddress != nil {
		return true
	}

	return false
}

// SetOverlayIpaddress gets a reference to the given string and assigns it to the OverlayIpaddress field.
func (o *ClientpackList) SetOverlayIpaddress(v string) {
	o.OverlayIpaddress = &v
}

// GetLinuxOnefile returns the LinuxOnefile field value if set, zero value otherwise.
func (o *ClientpackList) GetLinuxOnefile() string {
	if o == nil || o.LinuxOnefile == nil {
		var ret string
		return ret
	}
	return *o.LinuxOnefile
}

// GetLinuxOnefileOk returns a tuple with the LinuxOnefile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetLinuxOnefileOk() (*string, bool) {
	if o == nil || o.LinuxOnefile == nil {
		return nil, false
	}
	return o.LinuxOnefile, true
}

// HasLinuxOnefile returns a boolean if a field has been set.
func (o *ClientpackList) HasLinuxOnefile() bool {
	if o != nil && o.LinuxOnefile != nil {
		return true
	}

	return false
}

// SetLinuxOnefile gets a reference to the given string and assigns it to the LinuxOnefile field.
func (o *ClientpackList) SetLinuxOnefile(v string) {
	o.LinuxOnefile = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClientpackList) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClientpackList) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClientpackList) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConfSha1 returns the ConfSha1 field value if set, zero value otherwise.
func (o *ClientpackList) GetConfSha1() string {
	if o == nil || o.ConfSha1 == nil {
		var ret string
		return ret
	}
	return *o.ConfSha1
}

// GetConfSha1Ok returns a tuple with the ConfSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetConfSha1Ok() (*string, bool) {
	if o == nil || o.ConfSha1 == nil {
		return nil, false
	}
	return o.ConfSha1, true
}

// HasConfSha1 returns a boolean if a field has been set.
func (o *ClientpackList) HasConfSha1() bool {
	if o != nil && o.ConfSha1 != nil {
		return true
	}

	return false
}

// SetConfSha1 gets a reference to the given string and assigns it to the ConfSha1 field.
func (o *ClientpackList) SetConfSha1(v string) {
	o.ConfSha1 = &v
}

// GetWindowsOnefile returns the WindowsOnefile field value if set, zero value otherwise.
func (o *ClientpackList) GetWindowsOnefile() string {
	if o == nil || o.WindowsOnefile == nil {
		var ret string
		return ret
	}
	return *o.WindowsOnefile
}

// GetWindowsOnefileOk returns a tuple with the WindowsOnefile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetWindowsOnefileOk() (*string, bool) {
	if o == nil || o.WindowsOnefile == nil {
		return nil, false
	}
	return o.WindowsOnefile, true
}

// HasWindowsOnefile returns a boolean if a field has been set.
func (o *ClientpackList) HasWindowsOnefile() bool {
	if o != nil && o.WindowsOnefile != nil {
		return true
	}

	return false
}

// SetWindowsOnefile gets a reference to the given string and assigns it to the WindowsOnefile field.
func (o *ClientpackList) SetWindowsOnefile(v string) {
	o.WindowsOnefile = &v
}

// GetOvpnSha1 returns the OvpnSha1 field value if set, zero value otherwise.
func (o *ClientpackList) GetOvpnSha1() string {
	if o == nil || o.OvpnSha1 == nil {
		var ret string
		return ret
	}
	return *o.OvpnSha1
}

// GetOvpnSha1Ok returns a tuple with the OvpnSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetOvpnSha1Ok() (*string, bool) {
	if o == nil || o.OvpnSha1 == nil {
		return nil, false
	}
	return o.OvpnSha1, true
}

// HasOvpnSha1 returns a boolean if a field has been set.
func (o *ClientpackList) HasOvpnSha1() bool {
	if o != nil && o.OvpnSha1 != nil {
		return true
	}

	return false
}

// SetOvpnSha1 gets a reference to the given string and assigns it to the OvpnSha1 field.
func (o *ClientpackList) SetOvpnSha1(v string) {
	o.OvpnSha1 = &v
}

// GetTarballFile returns the TarballFile field value if set, zero value otherwise.
func (o *ClientpackList) GetTarballFile() string {
	if o == nil || o.TarballFile == nil {
		var ret string
		return ret
	}
	return *o.TarballFile
}

// GetTarballFileOk returns a tuple with the TarballFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetTarballFileOk() (*string, bool) {
	if o == nil || o.TarballFile == nil {
		return nil, false
	}
	return o.TarballFile, true
}

// HasTarballFile returns a boolean if a field has been set.
func (o *ClientpackList) HasTarballFile() bool {
	if o != nil && o.TarballFile != nil {
		return true
	}

	return false
}

// SetTarballFile gets a reference to the given string and assigns it to the TarballFile field.
func (o *ClientpackList) SetTarballFile(v string) {
	o.TarballFile = &v
}

// GetTarballSha1 returns the TarballSha1 field value if set, zero value otherwise.
func (o *ClientpackList) GetTarballSha1() string {
	if o == nil || o.TarballSha1 == nil {
		var ret string
		return ret
	}
	return *o.TarballSha1
}

// GetTarballSha1Ok returns a tuple with the TarballSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetTarballSha1Ok() (*string, bool) {
	if o == nil || o.TarballSha1 == nil {
		return nil, false
	}
	return o.TarballSha1, true
}

// HasTarballSha1 returns a boolean if a field has been set.
func (o *ClientpackList) HasTarballSha1() bool {
	if o != nil && o.TarballSha1 != nil {
		return true
	}

	return false
}

// SetTarballSha1 gets a reference to the given string and assigns it to the TarballSha1 field.
func (o *ClientpackList) SetTarballSha1(v string) {
	o.TarballSha1 = &v
}

// GetSequentialId returns the SequentialId field value if set, zero value otherwise.
func (o *ClientpackList) GetSequentialId() int32 {
	if o == nil || o.SequentialId == nil {
		var ret int32
		return ret
	}
	return *o.SequentialId
}

// GetSequentialIdOk returns a tuple with the SequentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetSequentialIdOk() (*int32, bool) {
	if o == nil || o.SequentialId == nil {
		return nil, false
	}
	return o.SequentialId, true
}

// HasSequentialId returns a boolean if a field has been set.
func (o *ClientpackList) HasSequentialId() bool {
	if o != nil && o.SequentialId != nil {
		return true
	}

	return false
}

// SetSequentialId gets a reference to the given int32 and assigns it to the SequentialId field.
func (o *ClientpackList) SetSequentialId(v int32) {
	o.SequentialId = &v
}

// GetCheckedOut returns the CheckedOut field value if set, zero value otherwise.
func (o *ClientpackList) GetCheckedOut() bool {
	if o == nil || o.CheckedOut == nil {
		var ret bool
		return ret
	}
	return *o.CheckedOut
}

// GetCheckedOutOk returns a tuple with the CheckedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetCheckedOutOk() (*bool, bool) {
	if o == nil || o.CheckedOut == nil {
		return nil, false
	}
	return o.CheckedOut, true
}

// HasCheckedOut returns a boolean if a field has been set.
func (o *ClientpackList) HasCheckedOut() bool {
	if o != nil && o.CheckedOut != nil {
		return true
	}

	return false
}

// SetCheckedOut gets a reference to the given bool and assigns it to the CheckedOut field.
func (o *ClientpackList) SetCheckedOut(v bool) {
	o.CheckedOut = &v
}

// GetZipSha1 returns the ZipSha1 field value if set, zero value otherwise.
func (o *ClientpackList) GetZipSha1() string {
	if o == nil || o.ZipSha1 == nil {
		var ret string
		return ret
	}
	return *o.ZipSha1
}

// GetZipSha1Ok returns a tuple with the ZipSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetZipSha1Ok() (*string, bool) {
	if o == nil || o.ZipSha1 == nil {
		return nil, false
	}
	return o.ZipSha1, true
}

// HasZipSha1 returns a boolean if a field has been set.
func (o *ClientpackList) HasZipSha1() bool {
	if o != nil && o.ZipSha1 != nil {
		return true
	}

	return false
}

// SetZipSha1 gets a reference to the given string and assigns it to the ZipSha1 field.
func (o *ClientpackList) SetZipSha1(v string) {
	o.ZipSha1 = &v
}

// GetZipFile returns the ZipFile field value if set, zero value otherwise.
func (o *ClientpackList) GetZipFile() string {
	if o == nil || o.ZipFile == nil {
		var ret string
		return ret
	}
	return *o.ZipFile
}

// GetZipFileOk returns a tuple with the ZipFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetZipFileOk() (*string, bool) {
	if o == nil || o.ZipFile == nil {
		return nil, false
	}
	return o.ZipFile, true
}

// HasZipFile returns a boolean if a field has been set.
func (o *ClientpackList) HasZipFile() bool {
	if o != nil && o.ZipFile != nil {
		return true
	}

	return false
}

// SetZipFile gets a reference to the given string and assigns it to the ZipFile field.
func (o *ClientpackList) SetZipFile(v string) {
	o.ZipFile = &v
}

// GetLastConnect returns the LastConnect field value if set, zero value otherwise.
func (o *ClientpackList) GetLastConnect() string {
	if o == nil || o.LastConnect == nil {
		var ret string
		return ret
	}
	return *o.LastConnect
}

// GetLastConnectOk returns a tuple with the LastConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetLastConnectOk() (*string, bool) {
	if o == nil || o.LastConnect == nil {
		return nil, false
	}
	return o.LastConnect, true
}

// HasLastConnect returns a boolean if a field has been set.
func (o *ClientpackList) HasLastConnect() bool {
	if o != nil && o.LastConnect != nil {
		return true
	}

	return false
}

// SetLastConnect gets a reference to the given string and assigns it to the LastConnect field.
func (o *ClientpackList) SetLastConnect(v string) {
	o.LastConnect = &v
}

// GetLastDisconnect returns the LastDisconnect field value if set, zero value otherwise.
func (o *ClientpackList) GetLastDisconnect() string {
	if o == nil || o.LastDisconnect == nil {
		var ret string
		return ret
	}
	return *o.LastDisconnect
}

// GetLastDisconnectOk returns a tuple with the LastDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetLastDisconnectOk() (*string, bool) {
	if o == nil || o.LastDisconnect == nil {
		return nil, false
	}
	return o.LastDisconnect, true
}

// HasLastDisconnect returns a boolean if a field has been set.
func (o *ClientpackList) HasLastDisconnect() bool {
	if o != nil && o.LastDisconnect != nil {
		return true
	}

	return false
}

// SetLastDisconnect gets a reference to the given string and assigns it to the LastDisconnect field.
func (o *ClientpackList) SetLastDisconnect(v string) {
	o.LastDisconnect = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ClientpackList) GetCompression() string {
	if o == nil || o.Compression == nil {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetCompressionOk() (*string, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ClientpackList) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *ClientpackList) SetCompression(v string) {
	o.Compression = &v
}

// GetManagerId returns the ManagerId field value if set, zero value otherwise.
func (o *ClientpackList) GetManagerId() string {
	if o == nil || o.ManagerId == nil {
		var ret string
		return ret
	}
	return *o.ManagerId
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetManagerIdOk() (*string, bool) {
	if o == nil || o.ManagerId == nil {
		return nil, false
	}
	return o.ManagerId, true
}

// HasManagerId returns a boolean if a field has been set.
func (o *ClientpackList) HasManagerId() bool {
	if o != nil && o.ManagerId != nil {
		return true
	}

	return false
}

// SetManagerId gets a reference to the given string and assigns it to the ManagerId field.
func (o *ClientpackList) SetManagerId(v string) {
	o.ManagerId = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *ClientpackList) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *ClientpackList) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *ClientpackList) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClientpackList) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClientpackList) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClientpackList) SetStatus(v string) {
	o.Status = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *ClientpackList) GetConnected() bool {
	if o == nil || o.Connected == nil {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetConnectedOk() (*bool, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *ClientpackList) HasConnected() bool {
	if o != nil && o.Connected != nil {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *ClientpackList) SetConnected(v bool) {
	o.Connected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClientpackList) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackList) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClientpackList) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ClientpackList) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o ClientpackList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverlayIpaddress != nil {
		toSerialize["overlay_ipaddress"] = o.OverlayIpaddress
	}
	if o.LinuxOnefile != nil {
		toSerialize["linux_onefile"] = o.LinuxOnefile
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ConfSha1 != nil {
		toSerialize["conf_sha1"] = o.ConfSha1
	}
	if o.WindowsOnefile != nil {
		toSerialize["windows_onefile"] = o.WindowsOnefile
	}
	if o.OvpnSha1 != nil {
		toSerialize["ovpn_sha1"] = o.OvpnSha1
	}
	if o.TarballFile != nil {
		toSerialize["tarball_file"] = o.TarballFile
	}
	if o.TarballSha1 != nil {
		toSerialize["tarball_sha1"] = o.TarballSha1
	}
	if o.SequentialId != nil {
		toSerialize["sequential_id"] = o.SequentialId
	}
	if o.CheckedOut != nil {
		toSerialize["checked_out"] = o.CheckedOut
	}
	if o.ZipSha1 != nil {
		toSerialize["zip_sha1"] = o.ZipSha1
	}
	if o.ZipFile != nil {
		toSerialize["zip_file"] = o.ZipFile
	}
	if o.LastConnect != nil {
		toSerialize["last_connect"] = o.LastConnect
	}
	if o.LastDisconnect != nil {
		toSerialize["last_disconnect"] = o.LastDisconnect
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.ManagerId != nil {
		toSerialize["manager_id"] = o.ManagerId
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Connected != nil {
		toSerialize["connected"] = o.Connected
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackList struct {
	value *ClientpackList
	isSet bool
}

func (v NullableClientpackList) Get() *ClientpackList {
	return v.value
}

func (v *NullableClientpackList) Set(val *ClientpackList) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackList) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackList(val *ClientpackList) *NullableClientpackList {
	return &NullableClientpackList{value: val, isSet: true}
}

func (v NullableClientpackList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


