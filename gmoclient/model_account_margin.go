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

// AccountMargin struct for AccountMargin
type AccountMargin struct {
	// 0 Nomal end
	Status *int32 `json:"status,omitempty"`
	Data *AccountMarginData `json:"data,omitempty"`
	Responsetime *time.Time `json:"responsetime,omitempty"`
}

// NewAccountMargin instantiates a new AccountMargin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMargin() *AccountMargin {
	this := AccountMargin{}
	return &this
}

// NewAccountMarginWithDefaults instantiates a new AccountMargin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMarginWithDefaults() *AccountMargin {
	this := AccountMargin{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountMargin) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMargin) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountMargin) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AccountMargin) SetStatus(v int32) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AccountMargin) GetData() AccountMarginData {
	if o == nil || o.Data == nil {
		var ret AccountMarginData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMargin) GetDataOk() (*AccountMarginData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AccountMargin) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AccountMarginData and assigns it to the Data field.
func (o *AccountMargin) SetData(v AccountMarginData) {
	o.Data = &v
}

// GetResponsetime returns the Responsetime field value if set, zero value otherwise.
func (o *AccountMargin) GetResponsetime() time.Time {
	if o == nil || o.Responsetime == nil {
		var ret time.Time
		return ret
	}
	return *o.Responsetime
}

// GetResponsetimeOk returns a tuple with the Responsetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMargin) GetResponsetimeOk() (*time.Time, bool) {
	if o == nil || o.Responsetime == nil {
		return nil, false
	}
	return o.Responsetime, true
}

// HasResponsetime returns a boolean if a field has been set.
func (o *AccountMargin) HasResponsetime() bool {
	if o != nil && o.Responsetime != nil {
		return true
	}

	return false
}

// SetResponsetime gets a reference to the given time.Time and assigns it to the Responsetime field.
func (o *AccountMargin) SetResponsetime(v time.Time) {
	o.Responsetime = &v
}

func (o AccountMargin) MarshalJSON() ([]byte, error) {
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

type NullableAccountMargin struct {
	value *AccountMargin
	isSet bool
}

func (v NullableAccountMargin) Get() *AccountMargin {
	return v.value
}

func (v *NullableAccountMargin) Set(val *AccountMargin) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMargin) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMargin(val *AccountMargin) *NullableAccountMargin {
	return &NullableAccountMargin{value: val, isSet: true}
}

func (v NullableAccountMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


