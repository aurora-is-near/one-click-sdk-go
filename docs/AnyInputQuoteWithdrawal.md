# AnyInputQuoteWithdrawal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Withdrawal status | 
**AmountOutFormatted** | **string** | Amount out in readable format | 
**AmountOutUsd** | **string** | Amount out in USD | 
**AmountOut** | **string** | Amount out in smallest unit | 
**WithdrawFeeFormatted** | **string** | Withdrawal fee in readable format | 
**WithdrawFee** | **string** | Withdrawal fee in smallest unit | 
**WithdrawFeeUsd** | **string** | Withdrawal fee in USD | 
**Timestamp** | **string** | Timestamp of withdrawal | 
**Hash** | **string** | Transaction hash | 

## Methods

### NewAnyInputQuoteWithdrawal

`func NewAnyInputQuoteWithdrawal(status string, amountOutFormatted string, amountOutUsd string, amountOut string, withdrawFeeFormatted string, withdrawFee string, withdrawFeeUsd string, timestamp string, hash string, ) *AnyInputQuoteWithdrawal`

NewAnyInputQuoteWithdrawal instantiates a new AnyInputQuoteWithdrawal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnyInputQuoteWithdrawalWithDefaults

`func NewAnyInputQuoteWithdrawalWithDefaults() *AnyInputQuoteWithdrawal`

NewAnyInputQuoteWithdrawalWithDefaults instantiates a new AnyInputQuoteWithdrawal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AnyInputQuoteWithdrawal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnyInputQuoteWithdrawal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnyInputQuoteWithdrawal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountOutFormatted

`func (o *AnyInputQuoteWithdrawal) GetAmountOutFormatted() string`

GetAmountOutFormatted returns the AmountOutFormatted field if non-nil, zero value otherwise.

### GetAmountOutFormattedOk

`func (o *AnyInputQuoteWithdrawal) GetAmountOutFormattedOk() (*string, bool)`

GetAmountOutFormattedOk returns a tuple with the AmountOutFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutFormatted

`func (o *AnyInputQuoteWithdrawal) SetAmountOutFormatted(v string)`

SetAmountOutFormatted sets AmountOutFormatted field to given value.


### GetAmountOutUsd

`func (o *AnyInputQuoteWithdrawal) GetAmountOutUsd() string`

GetAmountOutUsd returns the AmountOutUsd field if non-nil, zero value otherwise.

### GetAmountOutUsdOk

`func (o *AnyInputQuoteWithdrawal) GetAmountOutUsdOk() (*string, bool)`

GetAmountOutUsdOk returns a tuple with the AmountOutUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutUsd

`func (o *AnyInputQuoteWithdrawal) SetAmountOutUsd(v string)`

SetAmountOutUsd sets AmountOutUsd field to given value.


### GetAmountOut

`func (o *AnyInputQuoteWithdrawal) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *AnyInputQuoteWithdrawal) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *AnyInputQuoteWithdrawal) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.


### GetWithdrawFeeFormatted

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFeeFormatted() string`

GetWithdrawFeeFormatted returns the WithdrawFeeFormatted field if non-nil, zero value otherwise.

### GetWithdrawFeeFormattedOk

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFeeFormattedOk() (*string, bool)`

GetWithdrawFeeFormattedOk returns a tuple with the WithdrawFeeFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFeeFormatted

`func (o *AnyInputQuoteWithdrawal) SetWithdrawFeeFormatted(v string)`

SetWithdrawFeeFormatted sets WithdrawFeeFormatted field to given value.


### GetWithdrawFee

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFee() string`

GetWithdrawFee returns the WithdrawFee field if non-nil, zero value otherwise.

### GetWithdrawFeeOk

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFeeOk() (*string, bool)`

GetWithdrawFeeOk returns a tuple with the WithdrawFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFee

`func (o *AnyInputQuoteWithdrawal) SetWithdrawFee(v string)`

SetWithdrawFee sets WithdrawFee field to given value.


### GetWithdrawFeeUsd

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFeeUsd() string`

GetWithdrawFeeUsd returns the WithdrawFeeUsd field if non-nil, zero value otherwise.

### GetWithdrawFeeUsdOk

`func (o *AnyInputQuoteWithdrawal) GetWithdrawFeeUsdOk() (*string, bool)`

GetWithdrawFeeUsdOk returns a tuple with the WithdrawFeeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFeeUsd

`func (o *AnyInputQuoteWithdrawal) SetWithdrawFeeUsd(v string)`

SetWithdrawFeeUsd sets WithdrawFeeUsd field to given value.


### GetTimestamp

`func (o *AnyInputQuoteWithdrawal) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AnyInputQuoteWithdrawal) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AnyInputQuoteWithdrawal) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetHash

`func (o *AnyInputQuoteWithdrawal) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AnyInputQuoteWithdrawal) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AnyInputQuoteWithdrawal) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


