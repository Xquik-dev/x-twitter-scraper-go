// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package constant

import (
	shimjson "github.com/Xquik-dev/x-twitter-scraper-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type.
//
//	constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type APIV1GuestWalletsTopups string
type Active string
type Canceled string
type Compose string
type Deleting string
type Failed string
type HTTPSXquikComAPIV1GuestWalletsStatus string
type PaidReads string
type Pending string
type Point string
type Post string
type Refine string
type RequiresEmailCode string
type Score string
type SourceDefaultProductionValueCanVary string
type StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable string
type Success string
type Succeeded string
type Usd string
type XAccountConnectionAttempt string
type XAccountConnectionChallenge string
type XWriteAction string

func (c APIV1GuestWalletsTopups) Default() APIV1GuestWalletsTopups {
	return "/api/v1/guest-wallets/topups"
}
func (c Active) Default() Active     { return "active" }
func (c Canceled) Default() Canceled { return "canceled" }
func (c Compose) Default() Compose   { return "compose" }
func (c Deleting) Default() Deleting { return "deleting" }
func (c Failed) Default() Failed     { return "failed" }
func (c HTTPSXquikComAPIV1GuestWalletsStatus) Default() HTTPSXquikComAPIV1GuestWalletsStatus {
	return "https://xquik.com/api/v1/guest-wallets/status"
}
func (c PaidReads) Default() PaidReads                 { return "paid_reads" }
func (c Pending) Default() Pending                     { return "pending" }
func (c Point) Default() Point                         { return "Point" }
func (c Post) Default() Post                           { return "POST" }
func (c Refine) Default() Refine                       { return "refine" }
func (c RequiresEmailCode) Default() RequiresEmailCode { return "requires_email_code" }
func (c Score) Default() Score                         { return "score" }
func (c SourceDefaultProductionValueCanVary) Default() SourceDefaultProductionValueCanVary {
	return "Source default; production value can vary"
}
func (c StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable) Default() StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable {
	return "Store api_key and the Idempotency-Key securely before sharing checkout_url. No email recovery is available."
}
func (c Success) Default() Success     { return "success" }
func (c Succeeded) Default() Succeeded { return "succeeded" }
func (c Usd) Default() Usd             { return "usd" }
func (c XAccountConnectionAttempt) Default() XAccountConnectionAttempt {
	return "x_account_connection_attempt"
}
func (c XAccountConnectionChallenge) Default() XAccountConnectionChallenge {
	return "x_account_connection_challenge"
}
func (c XWriteAction) Default() XWriteAction { return "x_write_action" }

func (c APIV1GuestWalletsTopups) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Active) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Canceled) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Compose) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Deleting) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Failed) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c HTTPSXquikComAPIV1GuestWalletsStatus) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c PaidReads) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Pending) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Point) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c Post) MarshalJSON() ([]byte, error)                                 { return marshalString(c) }
func (c Refine) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c RequiresEmailCode) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Score) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c SourceDefaultProductionValueCanVary) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c StoreAPIKeyAndTheIdempotencyKeySecurelyBeforeSharingCheckoutURLNoEmailRecoveryIsAvailable) MarshalJSON() ([]byte, error) {
	return marshalString(c)
}
func (c Success) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Succeeded) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Usd) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c XAccountConnectionAttempt) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c XAccountConnectionChallenge) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c XWriteAction) MarshalJSON() ([]byte, error)                { return marshalString(c) }

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
