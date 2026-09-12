// SPDX-License-Identifier: BSD-3-Clause

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import stdjson "encoding/json"

// RawMessage lets the upstream tests use Go's original implementation.
type RawMessage = stdjson.RawMessage
