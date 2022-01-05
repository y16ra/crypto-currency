# TickersData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ask** | Pointer to **int32** |  | [optional] 
**Bid** | Pointer to **int32** |  | [optional] 
**High** | Pointer to **int32** |  | [optional] 
**Last** | Pointer to **int32** |  | [optional] 
**Low** | Pointer to **int32** |  | [optional] 
**Symbol** | Pointer to [**Symbols**](Symbols.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Volume** | Pointer to **float64** |  | [optional] 

## Methods

### NewTickersData

`func NewTickersData() *TickersData`

NewTickersData instantiates a new TickersData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickersDataWithDefaults

`func NewTickersDataWithDefaults() *TickersData`

NewTickersDataWithDefaults instantiates a new TickersData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *TickersData) GetAsk() int32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *TickersData) GetAskOk() (*int32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *TickersData) SetAsk(v int32)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *TickersData) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetBid

`func (o *TickersData) GetBid() int32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *TickersData) GetBidOk() (*int32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *TickersData) SetBid(v int32)`

SetBid sets Bid field to given value.

### HasBid

`func (o *TickersData) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetHigh

`func (o *TickersData) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *TickersData) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *TickersData) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *TickersData) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLast

`func (o *TickersData) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *TickersData) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *TickersData) SetLast(v int32)`

SetLast sets Last field to given value.

### HasLast

`func (o *TickersData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLow

`func (o *TickersData) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *TickersData) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *TickersData) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *TickersData) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetSymbol

`func (o *TickersData) GetSymbol() Symbols`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TickersData) GetSymbolOk() (*Symbols, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TickersData) SetSymbol(v Symbols)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TickersData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimestamp

`func (o *TickersData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TickersData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TickersData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TickersData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVolume

`func (o *TickersData) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TickersData) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TickersData) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *TickersData) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


