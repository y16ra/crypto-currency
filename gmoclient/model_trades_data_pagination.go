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
)

// TradesDataPagination struct for TradesDataPagination
type TradesDataPagination struct {
	CurrentPage *int32 `json:"currentPage,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewTradesDataPagination instantiates a new TradesDataPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradesDataPagination() *TradesDataPagination {
	this := TradesDataPagination{}
	return &this
}

// NewTradesDataPaginationWithDefaults instantiates a new TradesDataPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradesDataPaginationWithDefaults() *TradesDataPagination {
	this := TradesDataPagination{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *TradesDataPagination) GetCurrentPage() int32 {
	if o == nil || o.CurrentPage == nil {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataPagination) GetCurrentPageOk() (*int32, bool) {
	if o == nil || o.CurrentPage == nil {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *TradesDataPagination) HasCurrentPage() bool {
	if o != nil && o.CurrentPage != nil {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *TradesDataPagination) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TradesDataPagination) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataPagination) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TradesDataPagination) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TradesDataPagination) SetCount(v int32) {
	o.Count = &v
}

func (o TradesDataPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPage != nil {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableTradesDataPagination struct {
	value *TradesDataPagination
	isSet bool
}

func (v NullableTradesDataPagination) Get() *TradesDataPagination {
	return v.value
}

func (v *NullableTradesDataPagination) Set(val *TradesDataPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableTradesDataPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableTradesDataPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradesDataPagination(val *TradesDataPagination) *NullableTradesDataPagination {
	return &NullableTradesDataPagination{value: val, isSet: true}
}

func (v NullableTradesDataPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradesDataPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

