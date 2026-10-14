// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package apijson

type status uint8

const (
	missing status = iota
	null
	invalid
	valid
)

type Field struct {
	raw    string
	status status
}

func (j Field) IsNull() bool    { return j.status <= null }
func (j Field) IsMissing() bool { return j.status == missing }
func (j Field) IsInvalid() bool { return j.status == invalid }
func (j Field) Raw() string     { return j.raw }
