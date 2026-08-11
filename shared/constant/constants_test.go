// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package constant

import (
	"encoding/json"
	"testing"
)

func TestConstantsMarshalDefaults(t *testing.T) {
	tests := []struct {
		name  string
		value json.Marshaler
		want  string
	}{
		{"guest wallet topups", APIV1GuestWalletsTopups(""), `"/api/v1/guest-wallets/topups"`},
		{"compose", Compose(""), `"compose"`},
		{"wallet status", HTTPSXquikComAPIV1GuestWalletsStatus(""), `"https://xquik.com/api/v1/guest-wallets/status"`},
		{"paid reads", PaidReads(""), `"paid_reads"`},
		{"post", Post(""), `"POST"`},
		{"production weight", ProductionWeightNotPublishedByX(""), `"Production weight not published by X"`},
		{"refine", Refine(""), `"refine"`},
		{"score", Score(""), `"score"`},
		{
			"store credentials",
			StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable(""),
			`"Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."`,
		},
		{"currency", Usd(""), `"usd"`},
		{"write action", XWriteAction(""), `"x_write_action"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
