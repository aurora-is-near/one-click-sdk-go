# SubmitDepositTxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | **string** | Transaction hash of your deposit | 
**DepositAddress** | **string** | Deposit address for the quote | 
**NearSenderAccount** | Pointer to **string** | Sender account (used only for NEAR blockchain) | [optional] 
**Memo** | Pointer to **string** | Memo (use if deposit was submitted with one) | [optional] 

## Methods

### NewSubmitDepositTxRequest

`func NewSubmitDepositTxRequest(txHash string, depositAddress string, ) *SubmitDepositTxRequest`

NewSubmitDepositTxRequest instantiates a new SubmitDepositTxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDepositTxRequestWithDefaults

`func NewSubmitDepositTxRequestWithDefaults() *SubmitDepositTxRequest`

NewSubmitDepositTxRequestWithDefaults instantiates a new SubmitDepositTxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *SubmitDepositTxRequest) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *SubmitDepositTxRequest) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *SubmitDepositTxRequest) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetDepositAddress

`func (o *SubmitDepositTxRequest) GetDepositAddress() string`

GetDepositAddress returns the DepositAddress field if non-nil, zero value otherwise.

### GetDepositAddressOk

`func (o *SubmitDepositTxRequest) GetDepositAddressOk() (*string, bool)`

GetDepositAddressOk returns a tuple with the DepositAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddress

`func (o *SubmitDepositTxRequest) SetDepositAddress(v string)`

SetDepositAddress sets DepositAddress field to given value.


### GetNearSenderAccount

`func (o *SubmitDepositTxRequest) GetNearSenderAccount() string`

GetNearSenderAccount returns the NearSenderAccount field if non-nil, zero value otherwise.

### GetNearSenderAccountOk

`func (o *SubmitDepositTxRequest) GetNearSenderAccountOk() (*string, bool)`

GetNearSenderAccountOk returns a tuple with the NearSenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearSenderAccount

`func (o *SubmitDepositTxRequest) SetNearSenderAccount(v string)`

SetNearSenderAccount sets NearSenderAccount field to given value.

### HasNearSenderAccount

`func (o *SubmitDepositTxRequest) HasNearSenderAccount() bool`

HasNearSenderAccount returns a boolean if a field has been set.

### GetMemo

`func (o *SubmitDepositTxRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *SubmitDepositTxRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *SubmitDepositTxRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *SubmitDepositTxRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


