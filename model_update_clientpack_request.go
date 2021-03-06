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
	"fmt"
)

// UpdateClientpackRequest - struct for UpdateClientpackRequest
type UpdateClientpackRequest struct {
	Interface{} *interface{}
}

// interface{}AsUpdateClientpackRequest is a convenience function that returns interface{} wrapped in UpdateClientpackRequest
func Interface{}AsUpdateClientpackRequest(v *interface{}) UpdateClientpackRequest {
	return UpdateClientpackRequest{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateClientpackRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UpdateClientpackRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateClientpackRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateClientpackRequest) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateClientpackRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableUpdateClientpackRequest struct {
	value *UpdateClientpackRequest
	isSet bool
}

func (v NullableUpdateClientpackRequest) Get() *UpdateClientpackRequest {
	return v.value
}

func (v *NullableUpdateClientpackRequest) Set(val *UpdateClientpackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClientpackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClientpackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClientpackRequest(val *UpdateClientpackRequest) *NullableUpdateClientpackRequest {
	return &NullableUpdateClientpackRequest{value: val, isSet: true}
}

func (v NullableUpdateClientpackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClientpackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


