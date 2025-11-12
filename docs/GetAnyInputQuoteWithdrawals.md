# GetAnyInputQuoteWithdrawals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | **string** | ID of the destination asset. | 
**Recipient** | **string** | Recipient address | 
**AffiliateRecipient** | **string** | Affiliate recipient address | 
**Withdrawals** | Pointer to [**AnyInputQuoteWithdrawal**](AnyInputQuoteWithdrawal.md) | Details of withdrawals | [optional] 

## Methods

### NewGetAnyInputQuoteWithdrawals

`func NewGetAnyInputQuoteWithdrawals(asset string, recipient string, affiliateRecipient string, ) *GetAnyInputQuoteWithdrawals`

NewGetAnyInputQuoteWithdrawals instantiates a new GetAnyInputQuoteWithdrawals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnyInputQuoteWithdrawalsWithDefaults

`func NewGetAnyInputQuoteWithdrawalsWithDefaults() *GetAnyInputQuoteWithdrawals`

NewGetAnyInputQuoteWithdrawalsWithDefaults instantiates a new GetAnyInputQuoteWithdrawals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetAnyInputQuoteWithdrawals) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetAnyInputQuoteWithdrawals) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetAnyInputQuoteWithdrawals) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetRecipient

`func (o *GetAnyInputQuoteWithdrawals) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetAnyInputQuoteWithdrawals) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetAnyInputQuoteWithdrawals) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAffiliateRecipient

`func (o *GetAnyInputQuoteWithdrawals) GetAffiliateRecipient() string`

GetAffiliateRecipient returns the AffiliateRecipient field if non-nil, zero value otherwise.

### GetAffiliateRecipientOk

`func (o *GetAnyInputQuoteWithdrawals) GetAffiliateRecipientOk() (*string, bool)`

GetAffiliateRecipientOk returns a tuple with the AffiliateRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRecipient

`func (o *GetAnyInputQuoteWithdrawals) SetAffiliateRecipient(v string)`

SetAffiliateRecipient sets AffiliateRecipient field to given value.


### GetWithdrawals

`func (o *GetAnyInputQuoteWithdrawals) GetWithdrawals() AnyInputQuoteWithdrawal`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *GetAnyInputQuoteWithdrawals) GetWithdrawalsOk() (*AnyInputQuoteWithdrawal, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *GetAnyInputQuoteWithdrawals) SetWithdrawals(v AnyInputQuoteWithdrawal)`

SetWithdrawals sets Withdrawals field to given value.

### HasWithdrawals

`func (o *GetAnyInputQuoteWithdrawals) HasWithdrawals() bool`

HasWithdrawals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


