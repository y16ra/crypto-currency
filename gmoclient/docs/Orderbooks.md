# Orderbooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 0 Nomal end | [optional] 
**Data** | Pointer to [**OrderbooksData**](OrderbooksData.md) |  | [optional] 
**Responsetime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrderbooks

`func NewOrderbooks() *Orderbooks`

NewOrderbooks instantiates a new Orderbooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderbooksWithDefaults

`func NewOrderbooksWithDefaults() *Orderbooks`

NewOrderbooksWithDefaults instantiates a new Orderbooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Orderbooks) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Orderbooks) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Orderbooks) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Orderbooks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Orderbooks) GetData() OrderbooksData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Orderbooks) GetDataOk() (*OrderbooksData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Orderbooks) SetData(v OrderbooksData)`

SetData sets Data field to given value.

### HasData

`func (o *Orderbooks) HasData() bool`

HasData returns a boolean if a field has been set.

### GetResponsetime

`func (o *Orderbooks) GetResponsetime() time.Time`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *Orderbooks) GetResponsetimeOk() (*time.Time, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *Orderbooks) SetResponsetime(v time.Time)`

SetResponsetime sets Responsetime field to given value.

### HasResponsetime

`func (o *Orderbooks) HasResponsetime() bool`

HasResponsetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


