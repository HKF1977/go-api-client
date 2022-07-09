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

// FirewallRuleTupleInner - struct for FirewallRuleTupleInner
type FirewallRuleTupleInner struct {
	Int32 *int32
	String *string
}

// int32AsFirewallRuleTupleInner is a convenience function that returns int32 wrapped in FirewallRuleTupleInner
func Int32AsFirewallRuleTupleInner(v *int32) FirewallRuleTupleInner {
	return FirewallRuleTupleInner{
		Int32: v,
	}
}

// stringAsFirewallRuleTupleInner is a convenience function that returns string wrapped in FirewallRuleTupleInner
func StringAsFirewallRuleTupleInner(v *string) FirewallRuleTupleInner {
	return FirewallRuleTupleInner{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FirewallRuleTupleInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(FirewallRuleTupleInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(FirewallRuleTupleInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FirewallRuleTupleInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FirewallRuleTupleInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableFirewallRuleTupleInner struct {
	value *FirewallRuleTupleInner
	isSet bool
}

func (v NullableFirewallRuleTupleInner) Get() *FirewallRuleTupleInner {
	return v.value
}

func (v *NullableFirewallRuleTupleInner) Set(val *FirewallRuleTupleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRuleTupleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRuleTupleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRuleTupleInner(val *FirewallRuleTupleInner) *NullableFirewallRuleTupleInner {
	return &NullableFirewallRuleTupleInner{value: val, isSet: true}
}

func (v NullableFirewallRuleTupleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRuleTupleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


