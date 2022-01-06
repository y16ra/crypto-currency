# Trades

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 0 Nomal end | [optional] 
**Data** | Pointer to [**TradesData**](TradesData.md) |  | [optional] 
**Responsetime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTrades

`func NewTrades() *Trades`

NewTrades instantiates a new Trades object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradesWithDefaults

`func NewTradesWithDefaults() *Trades`

NewTradesWithDefaults instantiates a new Trades object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Trades) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Trades) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Trades) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Trades) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Trades) GetData() TradesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Trades) GetDataOk() (*TradesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Trades) SetData(v TradesData)`

SetData sets Data field to given value.

### HasData

`func (o *Trades) HasData() bool`

HasData returns a boolean if a field has been set.

### GetResponsetime

`func (o *Trades) GetResponsetime() time.Time`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *Trades) GetResponsetimeOk() (*time.Time, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *Trades) SetResponsetime(v time.Time)`

SetResponsetime sets Responsetime field to given value.

### HasResponsetime

`func (o *Trades) HasResponsetime() bool`

HasResponsetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


