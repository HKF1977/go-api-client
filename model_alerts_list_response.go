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

// AlertsListResponse struct for AlertsListResponse
type AlertsListResponse struct {
	Response []Alert `json:"response,omitempty"`
}

// NewAlertsListResponse instantiates a new AlertsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsListResponse() *AlertsListResponse {
	this := AlertsListResponse{}
	return &this
}

// NewAlertsListResponseWithDefaults instantiates a new AlertsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsListResponseWithDefaults() *AlertsListResponse {
	this := AlertsListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AlertsListResponse) GetResponse() []Alert {
	if o == nil || o.Response == nil {
		var ret []Alert
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListResponse) GetResponseOk() ([]Alert, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AlertsListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []Alert and assigns it to the Response field.
func (o *AlertsListResponse) SetResponse(v []Alert) {
	o.Response = v
}

func (o AlertsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAlertsListResponse struct {
	value *AlertsListResponse
	isSet bool
}

func (v NullableAlertsListResponse) Get() *AlertsListResponse {
	return v.value
}

func (v *NullableAlertsListResponse) Set(val *AlertsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsListResponse(val *AlertsListResponse) *NullableAlertsListResponse {
	return &NullableAlertsListResponse{value: val, isSet: true}
}

func (v NullableAlertsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


