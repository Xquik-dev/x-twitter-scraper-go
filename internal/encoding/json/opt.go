// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package json

type Option func(*encOpts)

func WithSkipCompaction(b bool) Option {
	return func(eos *encOpts) {
		eos.skipCompaction = true
	}
}

func (eos encOpts) apply(opts ...Option) encOpts {
	for _, opt := range opts {
		opt(&eos)
	}
	return eos
}
