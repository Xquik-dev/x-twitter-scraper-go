// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package constant

import (
	"encoding/json"
	"testing"
)

func TestConstantsMarshalDefaults(t *testing.T) {
	tests := map[string]struct {
		value json.Marshaler
		want  string
	}{
		"guest wallet topups": {
			value: APIV1GuestWalletsTopups(""),
			want:  `"/api/v1/guest-wallets/topups"`,
		},
		"compose": {
			value: Compose(""),
			want:  `"compose"`,
		},
		"checkout instructions": {
			value: GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending(""),
			want:  `"Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."`,
		},
		"wallet status": {
			value: HTTPSXquikComAPIV1GuestWalletsStatus(""),
			want:  `"https://xquik.com/api/v1/guest-wallets/status"`,
		},
		"paid reads": {
			value: PaidReads(""),
			want:  `"paid_reads"`,
		},
		"post": {
			value: Post(""),
			want:  `"POST"`,
		},
		"production weight": {
			value: ProductionWeightNotPublishedByX(""),
			want:  `"Production weight not published by X"`,
		},
		"refine": {
			value: Refine(""),
			want:  `"refine"`,
		},
		"running": {
			value: Running(""),
			want:  `"running"`,
		},
		"score": {
			value: Score(""),
			want:  `"score"`,
		},
		"store credentials": {
			value: StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable(""),
			want:  `"Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."`,
		},
		"currency": {
			value: Usd(""),
			want:  `"usd"`,
		},
		"write action": {
			value: XWriteAction(""),
			want:  `"x_write_action"`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := test.value.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != test.want {
				t.Fatalf("MarshalJSON() = %s, want %s", got, test.want)
			}
		})
	}

	if got := ValueOf[Compose](); got != "compose" {
		t.Fatalf("ValueOf[Compose]() = %q", got)
	}
	custom, err := Compose("custom").MarshalJSON()
	if err != nil || string(custom) != `"custom"` {
		t.Fatalf("custom constant = %s, %v", custom, err)
	}
}
