# AccountMargin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 0 Nomal end | [optional] 
**Data** | Pointer to [**AccountMarginData**](AccountMarginData.md) |  | [optional] 
**Responsetime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountMargin

`func NewAccountMargin() *AccountMargin`

NewAccountMargin instantiates a new AccountMargin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMarginWithDefaults

`func NewAccountMarginWithDefaults() *AccountMargin`

NewAccountMarginWithDefaults instantiates a new AccountMargin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountMargin) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountMargin) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountMargin) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountMargin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *AccountMargin) GetData() AccountMarginData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountMargin) GetDataOk() (*AccountMarginData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountMargin) SetData(v AccountMarginData)`

SetData sets Data field to given value.

### HasData

`func (o *AccountMargin) HasData() bool`

HasData returns a boolean if a field has been set.

### GetResponsetime

`func (o *AccountMargin) GetResponsetime() time.Time`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *AccountMargin) GetResponsetimeOk() (*time.Time, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *AccountMargin) SetResponsetime(v time.Time)`

SetResponsetime sets Responsetime field to given value.

### HasResponsetime

`func (o *AccountMargin) HasResponsetime() bool`

HasResponsetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


