//go:build !gocove

// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import "testing"

func TestMain(m *testing.M) {
	runMockTests(m)
}
