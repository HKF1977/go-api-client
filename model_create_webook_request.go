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

// CreateWebookRequest struct for CreateWebookRequest
type CreateWebookRequest struct {
	Name string `json:"name"`
	Url *string `json:"url,omitempty"`
	Events []string `json:"events,omitempty"`
	// Serialized HTTP Post request body
	Body *string `json:"body,omitempty"`
	ValidateCert *bool `json:"validate_cert,omitempty"`
	CustomProperties []CreateWebookRequestCustomPropertiesInner `json:"custom_properties,omitempty"`
	Headers []CreateWebookRequestHeadersInner `json:"headers,omitempty"`
	Parameters []CreateWebookRequestHeadersInner `json:"parameters,omitempty"`
}

// NewCreateWebookRequest instantiates a new CreateWebookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebookRequest(name string) *CreateWebookRequest {
	this := CreateWebookRequest{}
	this.Name = name
	return &this
}

// NewCreateWebookRequestWithDefaults instantiates a new CreateWebookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebookRequestWithDefaults() *CreateWebookRequest {
	this := CreateWebookRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWebookRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWebookRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateWebookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *CreateWebookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CreateWebookRequest) SetBody(v string) {
	o.Body = &v
}

// GetValidateCert returns the ValidateCert field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetValidateCert() bool {
	if o == nil || o.ValidateCert == nil {
		var ret bool
		return ret
	}
	return *o.ValidateCert
}

// GetValidateCertOk returns a tuple with the ValidateCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetValidateCertOk() (*bool, bool) {
	if o == nil || o.ValidateCert == nil {
		return nil, false
	}
	return o.ValidateCert, true
}

// HasValidateCert returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasValidateCert() bool {
	if o != nil && o.ValidateCert != nil {
		return true
	}

	return false
}

// SetValidateCert gets a reference to the given bool and assigns it to the ValidateCert field.
func (o *CreateWebookRequest) SetValidateCert(v bool) {
	o.ValidateCert = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetCustomProperties() []CreateWebookRequestCustomPropertiesInner {
	if o == nil || o.CustomProperties == nil {
		var ret []CreateWebookRequestCustomPropertiesInner
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetCustomPropertiesOk() ([]CreateWebookRequestCustomPropertiesInner, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []CreateWebookRequestCustomPropertiesInner and assigns it to the CustomProperties field.
func (o *CreateWebookRequest) SetCustomProperties(v []CreateWebookRequestCustomPropertiesInner) {
	o.CustomProperties = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetHeaders() []CreateWebookRequestHeadersInner {
	if o == nil || o.Headers == nil {
		var ret []CreateWebookRequestHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetHeadersOk() ([]CreateWebookRequestHeadersInner, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []CreateWebookRequestHeadersInner and assigns it to the Headers field.
func (o *CreateWebookRequest) SetHeaders(v []CreateWebookRequestHeadersInner) {
	o.Headers = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateWebookRequest) GetParameters() []CreateWebookRequestHeadersInner {
	if o == nil || o.Parameters == nil {
		var ret []CreateWebookRequestHeadersInner
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebookRequest) GetParametersOk() ([]CreateWebookRequestHeadersInner, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateWebookRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []CreateWebookRequestHeadersInner and assigns it to the Parameters field.
func (o *CreateWebookRequest) SetParameters(v []CreateWebookRequestHeadersInner) {
	o.Parameters = v
}

func (o CreateWebookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableCreateWebookRequest struct {
	value *CreateWebookRequest
	isSet bool
}

func (v NullableCreateWebookRequest) Get() *CreateWebookRequest {
	return v.value
}

func (v *NullableCreateWebookRequest) Set(val *CreateWebookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebookRequest(val *CreateWebookRequest) *NullableCreateWebookRequest {
	return &NullableCreateWebookRequest{value: val, isSet: true}
}

func (v NullableCreateWebookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


