// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type APIV1GuestWalletsTopups string                                                                                                                                                  // Always "/api/v1/guest-wallets/topups"
type Compose string                                                                                                                                                                  // Always "compose"
type GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending string // Always "Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."
type HTTPSXquikComAPIV1GuestWalletsStatus string                                                                                                                                     // Always "https://xquik.com/api/v1/guest-wallets/status"
type PaidReads string                                                                                                                                                                // Always "paid_reads"
type Post string                                                                                                                                                                     // Always "POST"
type ProductionWeightNotPublishedByX string                                                                                                                                          // Always "Production weight not published by X"
type Refine string                                                                                                                                                                   // Always "refine"
type Running string                                                                                                                                                                  // Always "running"
type Score string                                                                                                                                                                    // Always "score"
type StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable string                                                                                // Always "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
type Usd string                                                                                                                                                                      // Always "usd"
type XWriteAction string                                                                                                                                                             // Always "x_write_action"

func (c APIV1GuestWalletsTopups) Default() APIV1GuestWalletsTopups {
	return "/api/v1/guest-wallets/topups"
}
func (c Compose) Default() Compose { return "compose" }
func (c GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending) Default() GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending {
	return "Give checkout_url to the user. They must complete payment on Stripe. Never submit payment for them. After payment, poll status_url every poll_after_seconds until latest_purchase.status is no longer pending."
}
func (c HTTPSXquikComAPIV1GuestWalletsStatus) Default() HTTPSXquikComAPIV1GuestWalletsStatus {
	return "https://xquik.com/api/v1/guest-wallets/status"
}
func (c PaidReads) Default() PaidReads { return "paid_reads" }
func (c Post) Default() Post           { return "POST" }
func (c ProductionWeightNotPublishedByX) Default() ProductionWeightNotPublishedByX {
	return "Production weight not published by X"
}
func (c Refine) Default() Refine   { return "refine" }
func (c Running) Default() Running { return "running" }
func (c Score) Default() Score     { return "score" }
func (c StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable) Default() StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable {
	return "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
}
func (c Usd) Default() Usd                   { return "usd" }
func (c XWriteAction) Default() XWriteAction { return "x_write_action" }

func (c APIV1GuestWalletsTopups) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Compose) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c GiveCheckoutURLToTheUserTheyMustCompletePaymentOnStripeNeverSubmitPaymentForThemAfterPaymentPollStatusURLEveryPollAfterSecondsUntilLatestPurchaseStatusIsNoLongerPending) MarshalJSON() ([]byte, error) {
	return marshalString(c)
}
func (c HTTPSXquikComAPIV1GuestWalletsStatus) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c PaidReads) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Post) MarshalJSON() ([]byte, error)                                 { return marshalString(c) }
func (c ProductionWeightNotPublishedByX) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Refine) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Running) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Score) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable) MarshalJSON() ([]byte, error) {
	return marshalString(c)
}
func (c Usd) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c XWriteAction) MarshalJSON() ([]byte, error) { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
