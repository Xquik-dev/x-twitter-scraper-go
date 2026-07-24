// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apiform

type Marshaler interface {
	MarshalMultipart() ([]byte, string, error)
}
