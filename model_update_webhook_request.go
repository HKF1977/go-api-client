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

// UpdateWebhookRequest struct for UpdateWebhookRequest
type UpdateWebhookRequest struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Events []string `json:"events,omitempty"`
	// Serialized HTTP Post request body
	Body *string `json:"body,omitempty"`
	ValidateCert *bool `json:"validate_cert,omitempty"`
	CustomProperties []CreateWebookRequestCustomPropertiesInner `json:"custom_properties,omitempty"`
	Headers []CreateWebookRequestHeadersInner `json:"headers,omitempty"`
	Parameters []CreateWebookRequestHeadersInner `json:"parameters,omitempty"`
}

// NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookRequest() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *UpdateWebhookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UpdateWebhookRequest) SetBody(v string) {
	o.Body = &v
}

// GetValidateCert returns the ValidateCert field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetValidateCert() bool {
	if o == nil || o.ValidateCert == nil {
		var ret bool
		return ret
	}
	return *o.ValidateCert
}

// GetValidateCertOk returns a tuple with the ValidateCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetValidateCertOk() (*bool, bool) {
	if o == nil || o.ValidateCert == nil {
		return nil, false
	}
	return o.ValidateCert, true
}

// HasValidateCert returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasValidateCert() bool {
	if o != nil && o.ValidateCert != nil {
		return true
	}

	return false
}

// SetValidateCert gets a reference to the given bool and assigns it to the ValidateCert field.
func (o *UpdateWebhookRequest) SetValidateCert(v bool) {
	o.ValidateCert = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetCustomProperties() []CreateWebookRequestCustomPropertiesInner {
	if o == nil || o.CustomProperties == nil {
		var ret []CreateWebookRequestCustomPropertiesInner
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetCustomPropertiesOk() ([]CreateWebookRequestCustomPropertiesInner, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []CreateWebookRequestCustomPropertiesInner and assigns it to the CustomProperties field.
func (o *UpdateWebhookRequest) SetCustomProperties(v []CreateWebookRequestCustomPropertiesInner) {
	o.CustomProperties = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetHeaders() []CreateWebookRequestHeadersInner {
	if o == nil || o.Headers == nil {
		var ret []CreateWebookRequestHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetHeadersOk() ([]CreateWebookRequestHeadersInner, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []CreateWebookRequestHeadersInner and assigns it to the Headers field.
func (o *UpdateWebhookRequest) SetHeaders(v []CreateWebookRequestHeadersInner) {
	o.Headers = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetParameters() []CreateWebookRequestHeadersInner {
	if o == nil || o.Parameters == nil {
		var ret []CreateWebookRequestHeadersInner
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetParametersOk() ([]CreateWebookRequestHeadersInner, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []CreateWebookRequestHeadersInner and assigns it to the Parameters field.
func (o *UpdateWebhookRequest) SetParameters(v []CreateWebookRequestHeadersInner) {
	o.Parameters = v
}

func (o UpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.ValidateCert != nil {
		toSerialize["validate_cert"] = o.ValidateCert
	}
	if o.CustomProperties != nil {
		toSerialize["custom_properties"] = o.CustomProperties
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWebhookRequest struct {
	value *UpdateWebhookRequest
	isSet bool
}

func (v NullableUpdateWebhookRequest) Get() *UpdateWebhookRequest {
	return v.value
}

func (v *NullableUpdateWebhookRequest) Set(val *UpdateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookRequest(val *UpdateWebhookRequest) *NullableUpdateWebhookRequest {
	return &NullableUpdateWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


