/*
GMO Coin APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: dev@tricoro.tech
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gmoclient

import (
	"encoding/json"
	"time"
)

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	// 0 Nomal end
	Status *int32 `json:"status,omitempty"`
	// orderid
	Data *string `json:"data,omitempty"`
	Responsetime *time.Time `json:"responsetime,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *InlineResponse200) SetStatus(v int32) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse200) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse200) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *InlineResponse200) SetData(v string) {
	o.Data = &v
}

// GetResponsetime returns the Responsetime field value if set, zero value otherwise.
func (o *InlineResponse200) GetResponsetime() time.Time {
	if o == nil || o.Responsetime == nil {
		var ret time.Time
		return ret
	}
	return *o.Responsetime
}

// GetResponsetimeOk returns a tuple with the Responsetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetResponsetimeOk() (*time.Time, bool) {
	if o == nil || o.Responsetime == nil {
		return nil, false
	}
	return o.Responsetime, true
}

// HasResponsetime returns a boolean if a field has been set.
func (o *InlineResponse200) HasResponsetime() bool {
	if o != nil && o.Responsetime != nil {
		return true
	}

	return false
}

// SetResponsetime gets a reference to the given time.Time and assigns it to the Responsetime field.
func (o *InlineResponse200) SetResponsetime(v time.Time) {
	o.Responsetime = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Responsetime != nil {
		toSerialize["responsetime"] = o.Responsetime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

