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

// PutUploadLdapAuthCert200Response struct for PutUploadLdapAuthCert200Response
type PutUploadLdapAuthCert200Response struct {
	Response *string `json:"response,omitempty"`
}

// NewPutUploadLdapAuthCert200Response instantiates a new PutUploadLdapAuthCert200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutUploadLdapAuthCert200Response() *PutUploadLdapAuthCert200Response {
	this := PutUploadLdapAuthCert200Response{}
	return &this
}

// NewPutUploadLdapAuthCert200ResponseWithDefaults instantiates a new PutUploadLdapAuthCert200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutUploadLdapAuthCert200ResponseWithDefaults() *PutUploadLdapAuthCert200Response {
	this := PutUploadLdapAuthCert200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PutUploadLdapAuthCert200Response) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutUploadLdapAuthCert200Response) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PutUploadLdapAuthCert200Response) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *PutUploadLdapAuthCert200Response) SetResponse(v string) {
	o.Response = &v
}

func (o PutUploadLdapAuthCert200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullablePutUploadLdapAuthCert200Response struct {
	value *PutUploadLdapAuthCert200Response
	isSet bool
}

func (v NullablePutUploadLdapAuthCert200Response) Get() *PutUploadLdapAuthCert200Response {
	return v.value
}

func (v *NullablePutUploadLdapAuthCert200Response) Set(val *PutUploadLdapAuthCert200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutUploadLdapAuthCert200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutUploadLdapAuthCert200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutUploadLdapAuthCert200Response(val *PutUploadLdapAuthCert200Response) *NullablePutUploadLdapAuthCert200Response {
	return &NullablePutUploadLdapAuthCert200Response{value: val, isSet: true}
}

func (v NullablePutUploadLdapAuthCert200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutUploadLdapAuthCert200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


