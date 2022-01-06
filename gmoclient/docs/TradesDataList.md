# TradesDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **int32** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTradesDataList

`func NewTradesDataList() *TradesDataList`

NewTradesDataList instantiates a new TradesDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradesDataListWithDefaults

`func NewTradesDataListWithDefaults() *TradesDataList`

NewTradesDataListWithDefaults instantiates a new TradesDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *TradesDataList) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TradesDataList) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TradesDataList) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TradesDataList) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSide

`func (o *TradesDataList) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *TradesDataList) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *TradesDataList) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *TradesDataList) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSize

`func (o *TradesDataList) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TradesDataList) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TradesDataList) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *TradesDataList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTimestamp

`func (o *TradesDataList) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TradesDataList) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TradesDataList) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TradesDataList) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


