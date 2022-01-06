# TradesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**TradesDataPagination**](TradesDataPagination.md) |  | [optional] 
**List** | Pointer to [**[]TradesDataList**](TradesDataList.md) |  | [optional] 

## Methods

### NewTradesData

`func NewTradesData() *TradesData`

NewTradesData instantiates a new TradesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradesDataWithDefaults

`func NewTradesDataWithDefaults() *TradesData`

NewTradesDataWithDefaults instantiates a new TradesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *TradesData) GetPagination() TradesDataPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TradesData) GetPaginationOk() (*TradesDataPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TradesData) SetPagination(v TradesDataPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TradesData) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetList

`func (o *TradesData) GetList() []TradesDataList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *TradesData) GetListOk() (*[]TradesDataList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *TradesData) SetList(v []TradesDataList)`

SetList sets List field to given value.

### HasList

`func (o *TradesData) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


