# AssetsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to [**Symbols**](Symbols.md) |  | [optional] 

## Methods

### NewAssetsData

`func NewAssetsData() *AssetsData`

NewAssetsData instantiates a new AssetsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsDataWithDefaults

`func NewAssetsDataWithDefaults() *AssetsData`

NewAssetsDataWithDefaults instantiates a new AssetsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AssetsData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AssetsData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AssetsData) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AssetsData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAvailable

`func (o *AssetsData) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AssetsData) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AssetsData) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AssetsData) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetConversionRate

`func (o *AssetsData) GetConversionRate() string`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *AssetsData) GetConversionRateOk() (*string, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *AssetsData) SetConversionRate(v string)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *AssetsData) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetSymbol

`func (o *AssetsData) GetSymbol() Symbols`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AssetsData) GetSymbolOk() (*Symbols, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AssetsData) SetSymbol(v Symbols)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AssetsData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


