package oneclick

import (
	"encoding/json"
	"testing"
)

func TestQuoteResponseUnmarshal__should_accept_echoed_correlationId(t *testing.T) {
	body := []byte(`{
		"correlationId": "top-id",
		"timestamp": "2026-08-19T13:49:58Z",
		"signature": "sig",
		"quoteRequest": {
			"dry": false,
			"swapType": "EXACT_INPUT",
			"slippageTolerance": 100,
			"originAsset": "nep141:wrap.near",
			"depositType": "ORIGIN_CHAIN",
			"destinationAsset": "nep245:v2_1.omni.hot.tg:143_11111111111111111111",
			"amount": "1",
			"refundTo": "account.near",
			"refundType": "ORIGIN_CHAIN",
			"recipient": "0xb5E899e65380f2Ee5845Cc893d5A77d838Edf697",
			"recipientType": "DESTINATION_CHAIN",
			"deadline": "2026-08-19T14:04:58Z",
			"correlationId": "request-id"
		},
		"quote": {
			"correlationId": "quote-id",
			"depositAddress": "foo.near",
			"amountIn": "1",
			"amountInFormatted": "1",
			"amountInUsd": "1",
			"minAmountIn": "1",
			"amountOut": "1",
			"amountOutFormatted": "1",
			"amountOutUsd": "1",
			"minAmountOut": "1",
			"timeEstimate": 30
		}
	}`)

	var resp QuoteResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if resp.GetCorrelationId() != "top-id" {
		t.Fatalf("top correlationId: got %q", resp.GetCorrelationId())
	}
	qr := resp.GetQuoteRequest()
	if got := qr.GetCorrelationId(); got != "request-id" {
		t.Fatalf("quoteRequest.correlationId: got %q", got)
	}
	q := resp.GetQuote()
	if !q.HasDepositAddress() || q.GetDepositAddress() != "foo.near" {
		t.Fatalf("depositAddress: got %q", q.GetDepositAddress())
	}
}
