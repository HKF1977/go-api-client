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

// PacketMonitor struct for PacketMonitor
type PacketMonitor struct {
	// Name of packet monitor. Must be conntrack for type=conntrack as only one conntrack monitor can run at a time
	Name *string `json:"name,omitempty"`
	// conntrack, netflow or pcap
	Type *string `json:"type,omitempty"`
	Interface *string `json:"interface,omitempty"`
	// filter strings are particular to the type of packet monitor. For instance, \"-p 8000\" for pcap
	Filter *string `json:"filter,omitempty"`
	// Indicates length of time to run capture for. Can be forever or some string parsable by the Linux date command
	Duration *string `json:"duration,omitempty"`
	// must be file if pcap or conntrack. Otherwise a host should be specified with the prefix \"host\". E.g. \"host:10.0.3.2:4000\"
	Destination *string `json:"destination,omitempty"`
}

// NewPacketMonitor instantiates a new PacketMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPacketMonitor() *PacketMonitor {
	this := PacketMonitor{}
	return &this
}

// NewPacketMonitorWithDefaults instantiates a new PacketMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPacketMonitorWithDefaults() *PacketMonitor {
	this := PacketMonitor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PacketMonitor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PacketMonitor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PacketMonitor) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PacketMonitor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PacketMonitor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PacketMonitor) SetType(v string) {
	o.Type = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *PacketMonitor) GetInterface() string {
	if o == nil || o.Interface == nil {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetInterfaceOk() (*string, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *PacketMonitor) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *PacketMonitor) SetInterface(v string) {
	o.Interface = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PacketMonitor) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PacketMonitor) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PacketMonitor) SetFilter(v string) {
	o.Filter = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PacketMonitor) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PacketMonitor) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *PacketMonitor) SetDuration(v string) {
	o.Duration = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *PacketMonitor) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketMonitor) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *PacketMonitor) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *PacketMonitor) SetDestination(v string) {
	o.Destination = &v
}

func (o PacketMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullablePacketMonitor struct {
	value *PacketMonitor
	isSet bool
}

func (v NullablePacketMonitor) Get() *PacketMonitor {
	return v.value
}

func (v *NullablePacketMonitor) Set(val *PacketMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullablePacketMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullablePacketMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePacketMonitor(val *PacketMonitor) *NullablePacketMonitor {
	return &NullablePacketMonitor{value: val, isSet: true}
}

func (v NullablePacketMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePacketMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


