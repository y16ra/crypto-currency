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

// AssetsData struct for AssetsData
type AssetsData struct {
	Amount *string `json:"amount,omitempty"`
	Available *string `json:"available,omitempty"`
	ConversionRate *string `json:"conversionRate,omitempty"`
	Symbol *Symbols `json:"symbol,omitempty"`
}

// NewAssetsData instantiates a new AssetsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsData() *AssetsData {
	this := AssetsData{}
	return &this
}

// NewAssetsDataWithDefaults instantiates a new AssetsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsDataWithDefaults() *AssetsData {
	this := AssetsData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AssetsData) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsData) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AssetsData) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AssetsData) SetAmount(v string) {
	o.Amount = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *AssetsData) GetAvailable() string {
	if o == nil || o.Available == nil {
		var ret string
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsData) GetAvailableOk() (*string, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *AssetsData) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given string and assigns it to the Available field.
func (o *AssetsData) SetAvailable(v string) {
	o.Available = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *AssetsData) GetConversionRate() string {
	if o == nil || o.ConversionRate == nil {
		var ret string
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsData) GetConversionRateOk() (*string, bool) {
	if o == nil || o.ConversionRate == nil {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *AssetsData) HasConversionRate() bool {
	if o != nil && o.ConversionRate != nil {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given string and assigns it to the ConversionRate field.
func (o *AssetsData) SetConversionRate(v string) {
	o.ConversionRate = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AssetsData) GetSymbol() Symbols {
	if o == nil || o.Symbol == nil {
		var ret Symbols
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsData) GetSymbolOk() (*Symbols, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AssetsData) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given Symbols and assigns it to the Symbol field.
func (o *AssetsData) SetSymbol(v Symbols) {
	o.Symbol = &v
}

func (o AssetsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.ConversionRate != nil {
		toSerialize["conversionRate"] = o.ConversionRate
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableAssetsData struct {
	value *AssetsData
	isSet bool
}

func (v NullableAssetsData) Get() *AssetsData {
	return v.value
}

func (v *NullableAssetsData) Set(val *AssetsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsData(val *AssetsData) *NullableAssetsData {
	return &NullableAssetsData{value: val, isSet: true}
}

func (v NullableAssetsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

