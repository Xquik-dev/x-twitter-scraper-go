//go:build gocove

// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"testing"

	"github.com/srvgit/gocove/cov"
)

func TestMain(m *testing.M) {
	defer cov.Flush()
	runMockTests(m)
}
