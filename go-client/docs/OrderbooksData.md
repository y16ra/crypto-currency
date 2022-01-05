# OrderbooksData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to [**[]OrderbooksDataAsks**](OrderbooksDataAsks.md) |  | [optional] 
**Bids** | Pointer to [**[]OrderbooksDataAsks**](OrderbooksDataAsks.md) |  | [optional] 
**Symbol** | Pointer to [**Symbols**](Symbols.md) |  | [optional] 

## Methods

### NewOrderbooksData

`func NewOrderbooksData() *OrderbooksData`

NewOrderbooksData instantiates a new OrderbooksData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderbooksDataWithDefaults

`func NewOrderbooksDataWithDefaults() *OrderbooksData`

NewOrderbooksDataWithDefaults instantiates a new OrderbooksData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *OrderbooksData) GetAsks() []OrderbooksDataAsks`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *OrderbooksData) GetAsksOk() (*[]OrderbooksDataAsks, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *OrderbooksData) SetAsks(v []OrderbooksDataAsks)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *OrderbooksData) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *OrderbooksData) GetBids() []OrderbooksDataAsks`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *OrderbooksData) GetBidsOk() (*[]OrderbooksDataAsks, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *OrderbooksData) SetBids(v []OrderbooksDataAsks)`

SetBids sets Bids field to given value.

### HasBids

`func (o *OrderbooksData) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderbooksData) GetSymbol() Symbols`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderbooksData) GetSymbolOk() (*Symbols, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderbooksData) SetSymbol(v Symbols)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderbooksData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


