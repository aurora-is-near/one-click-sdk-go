# QuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dry** | **bool** | Flag indicating whether this is a dry run request. If &#x60;true&#x60;, the response will **NOT** contain the following fields: - &#x60;depositAddress&#x60; - &#x60;timeWhenInactive&#x60; - &#x60;deadline&#x60; | 
**DepositMode** | Pointer to **string** | What deposit address mode you will get in the response, most chain supports only &#x60;SIMPLE&#x60; and some(for example &#x60;stellar&#x60;) only &#x60;MEMO&#x60;: - &#x60;SIMPLE&#x60; - usual deposit with only deposit address. - &#x60;MEMO&#x60; - some chains will **REQUIRE** the &#x60;memo&#x60; together with &#x60;depositAddress&#x60; for swap to work. | [optional] [default to "SIMPLE"]
**SwapType** | **string** | How to interpret &#x60;amount&#x60; (and refunds) when performing the swap:  - &#x60;EXACT_INPUT&#x60; — requests the output amount for an exact input.   - If deposit is less than &#x60;amountIn&#x60;, the deposit is refunded by deadline.   - If deposit is above than &#x60;amountIn&#x60;, the swap is processed and the excess is refunded to &#x60;refundTo&#x60; address after swap is complete.  - &#x60;EXACT_OUTPUT&#x60; — requests the input amount for an exact output.   - The quote response would have two fields &#x60;minAmountIn&#x60; and &#x60;maxAmountIn&#x60;.   - If the input is above than &#x60;maxAmountIn&#x60; the swap is processed and the excess is refunded to &#x60;refundTo&#x60; address after swap is complete.   - If the input is less than  &#x60;minAmountIn&#x60;, the deposit is refunded by deadline.  - &#x60;FLEX_INPUT&#x60; — a flexible input amount that allows for partial deposits and variable amounts.   - &#x60;slippage&#x60; applies both to &#x60;amountOut&#x60; and &#x60;amountIn&#x60; and defines an acceptable range (&#x60;minAmountIn&#x60; and &#x60;minAmountOut&#x60;).   - Any amount higher than &#x60;minAmountIn&#x60; is accepted and converted to the output asset as long as &#x60;minAmountOut&#x60; is met.   - The &#x60;amountIn&#x60; can be less, as long as the &#39;slippage + 1%&#39; constraint is met. If the total received by the deadline is below the lower bound, the deposit is refunded.   - If deposits exceed the upper bound, the swap is still processed | 
**SlippageTolerance** | **float32** | Slippage tolerance for the swap. This value is in basis points (1/100th of a percent), e.g. 100 for 1% slippage. | 
**OriginAsset** | **string** | ID of the origin asset. | 
**DepositType** | **string** | Type of deposit address: - &#x60;ORIGIN_CHAIN&#x60; - deposit address on the origin chain. - &#x60;INTENTS&#x60; - the account ID within NEAR Intents to which you should transfer assets. | 
**DestinationAsset** | **string** | ID of the destination asset. | 
**Amount** | **string** | Amount to swap as the base amount. It is interpreted as the input or output amount based on the &#x60;swapType&#x60; flag and is specified in the smallest unit of the currency (e.g., wei for ETH). | 
**RefundTo** | **string** | Address used for refunds. | 
**RefundType** | **string** | Type of refund address: - &#x60;ORIGIN_CHAIN&#x60; - assets are refunded to the &#x60;refundTo&#x60; address on the origin chain. - &#x60;INTENTS&#x60; - assets are refunded to the &#x60;refundTo&#x60; Intents account. | 
**Recipient** | **string** | Recipient address. The format must match &#x60;recipientType&#x60;. | 
**ConnectedWallets** | Pointer to **[]string** | Addresses of connected wallets. | [optional] 
**SessionId** | Pointer to **string** | Unique client session identifier for 1Click. | [optional] 
**VirtualChainRecipient** | Pointer to **string** | EVM address of a transfer recipient in a virtual chain | [optional] 
**VirtualChainRefundRecipient** | Pointer to **string** | EVM address of a refund recipient in a virtual chain | [optional] 
**CustomRecipientMsg** | Pointer to **string** | **HIGHLY EXPERIMENTAL** Message to pass to &#x60;ft_transfer_call&#x60; when withdrawing assets to NEAR.  Otherwise, &#x60;ft_transfer&#x60; will be used.  **WARNING**: Funds will be lost if used with non NEP-141 tokens, in case of insufficient &#x60;storage_deposit&#x60; or if the recipient does not implement &#x60;ft_on_transfer&#x60; method. | [optional] 
**RecipientType** | **string** | Type of recipient address: - &#x60;DESTINATION_CHAIN&#x60; - assets are transferred to the chain of &#x60;destinationAsset&#x60;. - &#x60;INTENTS&#x60; - assets are transferred to an account inside Intents | 
**Deadline** | **time.Time** | Timestamp in ISO format that identifies when the user refund begins if the swap isn&#39;t completed by then. It must exceed the time required for the deposit transaction to be mined. For example, Bitcoin may require around one hour depending on the fees paid. | 
**Referral** | Pointer to **string** | Referral identifier (lowercase only). It will be reflected in the on-chain data and displayed on public analytics platforms. | [optional] 
**QuoteWaitingTimeMs** | Pointer to **float32** | Time in milliseconds the user is willing to wait for a quote from the relay. **If you want to receive the fastest quote - use &#x60;0&#x60; as a value**  | [optional] [default to 3000]
**AppFees** | Pointer to [**[]AppFee**](AppFee.md) | List of recipients and their fees | [optional] 
**Confidentiality** | Pointer to **string** | Quote confidentiality mode echoed by the API | [optional] 
**Insured** | Pointer to **bool** | Whether the quote is insured (echoed by the API on quoteRequest) | [optional] 
**CorrelationId** | Pointer to **string** | Echoed by the API on the nested quoteRequest object (in addition to QuoteResponse.correlationId) | [optional] 

## Methods

### NewQuoteRequest

`func NewQuoteRequest(dry bool, swapType string, slippageTolerance float32, originAsset string, depositType string, destinationAsset string, amount string, refundTo string, refundType string, recipient string, recipientType string, deadline time.Time, ) *QuoteRequest`

NewQuoteRequest instantiates a new QuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRequestWithDefaults

`func NewQuoteRequestWithDefaults() *QuoteRequest`

NewQuoteRequestWithDefaults instantiates a new QuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDry

`func (o *QuoteRequest) GetDry() bool`

GetDry returns the Dry field if non-nil, zero value otherwise.

### GetDryOk

`func (o *QuoteRequest) GetDryOk() (*bool, bool)`

GetDryOk returns a tuple with the Dry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDry

`func (o *QuoteRequest) SetDry(v bool)`

SetDry sets Dry field to given value.


### GetDepositMode

`func (o *QuoteRequest) GetDepositMode() string`

GetDepositMode returns the DepositMode field if non-nil, zero value otherwise.

### GetDepositModeOk

`func (o *QuoteRequest) GetDepositModeOk() (*string, bool)`

GetDepositModeOk returns a tuple with the DepositMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositMode

`func (o *QuoteRequest) SetDepositMode(v string)`

SetDepositMode sets DepositMode field to given value.

### HasDepositMode

`func (o *QuoteRequest) HasDepositMode() bool`

HasDepositMode returns a boolean if a field has been set.

### GetSwapType

`func (o *QuoteRequest) GetSwapType() string`

GetSwapType returns the SwapType field if non-nil, zero value otherwise.

### GetSwapTypeOk

`func (o *QuoteRequest) GetSwapTypeOk() (*string, bool)`

GetSwapTypeOk returns a tuple with the SwapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapType

`func (o *QuoteRequest) SetSwapType(v string)`

SetSwapType sets SwapType field to given value.


### GetSlippageTolerance

`func (o *QuoteRequest) GetSlippageTolerance() float32`

GetSlippageTolerance returns the SlippageTolerance field if non-nil, zero value otherwise.

### GetSlippageToleranceOk

`func (o *QuoteRequest) GetSlippageToleranceOk() (*float32, bool)`

GetSlippageToleranceOk returns a tuple with the SlippageTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageTolerance

`func (o *QuoteRequest) SetSlippageTolerance(v float32)`

SetSlippageTolerance sets SlippageTolerance field to given value.


### GetOriginAsset

`func (o *QuoteRequest) GetOriginAsset() string`

GetOriginAsset returns the OriginAsset field if non-nil, zero value otherwise.

### GetOriginAssetOk

`func (o *QuoteRequest) GetOriginAssetOk() (*string, bool)`

GetOriginAssetOk returns a tuple with the OriginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAsset

`func (o *QuoteRequest) SetOriginAsset(v string)`

SetOriginAsset sets OriginAsset field to given value.


### GetDepositType

`func (o *QuoteRequest) GetDepositType() string`

GetDepositType returns the DepositType field if non-nil, zero value otherwise.

### GetDepositTypeOk

`func (o *QuoteRequest) GetDepositTypeOk() (*string, bool)`

GetDepositTypeOk returns a tuple with the DepositType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositType

`func (o *QuoteRequest) SetDepositType(v string)`

SetDepositType sets DepositType field to given value.


### GetDestinationAsset

`func (o *QuoteRequest) GetDestinationAsset() string`

GetDestinationAsset returns the DestinationAsset field if non-nil, zero value otherwise.

### GetDestinationAssetOk

`func (o *QuoteRequest) GetDestinationAssetOk() (*string, bool)`

GetDestinationAssetOk returns a tuple with the DestinationAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAsset

`func (o *QuoteRequest) SetDestinationAsset(v string)`

SetDestinationAsset sets DestinationAsset field to given value.


### GetAmount

`func (o *QuoteRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuoteRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuoteRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRefundTo

`func (o *QuoteRequest) GetRefundTo() string`

GetRefundTo returns the RefundTo field if non-nil, zero value otherwise.

### GetRefundToOk

`func (o *QuoteRequest) GetRefundToOk() (*string, bool)`

GetRefundToOk returns a tuple with the RefundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTo

`func (o *QuoteRequest) SetRefundTo(v string)`

SetRefundTo sets RefundTo field to given value.


### GetRefundType

`func (o *QuoteRequest) GetRefundType() string`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *QuoteRequest) GetRefundTypeOk() (*string, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *QuoteRequest) SetRefundType(v string)`

SetRefundType sets RefundType field to given value.


### GetRecipient

`func (o *QuoteRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *QuoteRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *QuoteRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetConnectedWallets

`func (o *QuoteRequest) GetConnectedWallets() []string`

GetConnectedWallets returns the ConnectedWallets field if non-nil, zero value otherwise.

### GetConnectedWalletsOk

`func (o *QuoteRequest) GetConnectedWalletsOk() (*[]string, bool)`

GetConnectedWalletsOk returns a tuple with the ConnectedWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWallets

`func (o *QuoteRequest) SetConnectedWallets(v []string)`

SetConnectedWallets sets ConnectedWallets field to given value.

### HasConnectedWallets

`func (o *QuoteRequest) HasConnectedWallets() bool`

HasConnectedWallets returns a boolean if a field has been set.

### GetSessionId

`func (o *QuoteRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *QuoteRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *QuoteRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *QuoteRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetVirtualChainRecipient

`func (o *QuoteRequest) GetVirtualChainRecipient() string`

GetVirtualChainRecipient returns the VirtualChainRecipient field if non-nil, zero value otherwise.

### GetVirtualChainRecipientOk

`func (o *QuoteRequest) GetVirtualChainRecipientOk() (*string, bool)`

GetVirtualChainRecipientOk returns a tuple with the VirtualChainRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChainRecipient

`func (o *QuoteRequest) SetVirtualChainRecipient(v string)`

SetVirtualChainRecipient sets VirtualChainRecipient field to given value.

### HasVirtualChainRecipient

`func (o *QuoteRequest) HasVirtualChainRecipient() bool`

HasVirtualChainRecipient returns a boolean if a field has been set.

### GetVirtualChainRefundRecipient

`func (o *QuoteRequest) GetVirtualChainRefundRecipient() string`

GetVirtualChainRefundRecipient returns the VirtualChainRefundRecipient field if non-nil, zero value otherwise.

### GetVirtualChainRefundRecipientOk

`func (o *QuoteRequest) GetVirtualChainRefundRecipientOk() (*string, bool)`

GetVirtualChainRefundRecipientOk returns a tuple with the VirtualChainRefundRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChainRefundRecipient

`func (o *QuoteRequest) SetVirtualChainRefundRecipient(v string)`

SetVirtualChainRefundRecipient sets VirtualChainRefundRecipient field to given value.

### HasVirtualChainRefundRecipient

`func (o *QuoteRequest) HasVirtualChainRefundRecipient() bool`

HasVirtualChainRefundRecipient returns a boolean if a field has been set.

### GetCustomRecipientMsg

`func (o *QuoteRequest) GetCustomRecipientMsg() string`

GetCustomRecipientMsg returns the CustomRecipientMsg field if non-nil, zero value otherwise.

### GetCustomRecipientMsgOk

`func (o *QuoteRequest) GetCustomRecipientMsgOk() (*string, bool)`

GetCustomRecipientMsgOk returns a tuple with the CustomRecipientMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRecipientMsg

`func (o *QuoteRequest) SetCustomRecipientMsg(v string)`

SetCustomRecipientMsg sets CustomRecipientMsg field to given value.

### HasCustomRecipientMsg

`func (o *QuoteRequest) HasCustomRecipientMsg() bool`

HasCustomRecipientMsg returns a boolean if a field has been set.

### GetRecipientType

`func (o *QuoteRequest) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *QuoteRequest) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *QuoteRequest) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.


### GetDeadline

`func (o *QuoteRequest) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *QuoteRequest) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *QuoteRequest) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetReferral

`func (o *QuoteRequest) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *QuoteRequest) GetReferralOk() (*string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *QuoteRequest) SetReferral(v string)`

SetReferral sets Referral field to given value.

### HasReferral

`func (o *QuoteRequest) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### GetQuoteWaitingTimeMs

`func (o *QuoteRequest) GetQuoteWaitingTimeMs() float32`

GetQuoteWaitingTimeMs returns the QuoteWaitingTimeMs field if non-nil, zero value otherwise.

### GetQuoteWaitingTimeMsOk

`func (o *QuoteRequest) GetQuoteWaitingTimeMsOk() (*float32, bool)`

GetQuoteWaitingTimeMsOk returns a tuple with the QuoteWaitingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteWaitingTimeMs

`func (o *QuoteRequest) SetQuoteWaitingTimeMs(v float32)`

SetQuoteWaitingTimeMs sets QuoteWaitingTimeMs field to given value.

### HasQuoteWaitingTimeMs

`func (o *QuoteRequest) HasQuoteWaitingTimeMs() bool`

HasQuoteWaitingTimeMs returns a boolean if a field has been set.

### GetAppFees

`func (o *QuoteRequest) GetAppFees() []AppFee`

GetAppFees returns the AppFees field if non-nil, zero value otherwise.

### GetAppFeesOk

`func (o *QuoteRequest) GetAppFeesOk() (*[]AppFee, bool)`

GetAppFeesOk returns a tuple with the AppFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppFees

`func (o *QuoteRequest) SetAppFees(v []AppFee)`

SetAppFees sets AppFees field to given value.

### HasAppFees

`func (o *QuoteRequest) HasAppFees() bool`

HasAppFees returns a boolean if a field has been set.

### GetCorrelationId

`func (o *QuoteRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *QuoteRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *QuoteRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *QuoteRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


