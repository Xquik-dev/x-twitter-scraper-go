// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"reflect"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go/internal/reflectvalue"
)

type TimeMarshaler interface {
	MarshalJSONWithTimeLayout(string) []byte
}

func TimeLayout(fmt string) string {
	switch fmt {
	case "", "date-time":
		return time.RFC3339
	case "date":
		return time.DateOnly
	default:
		return fmt
	}
}

var timeType = reflect.TypeFor[time.Time]()

func timeEncoder(e *encodeState, v reflect.Value, opts encOpts) {
	t, err := reflectvalue.As[time.Time](v)
	if err != nil {
		e.error(err)
	}
	stringEncoder(e, reflect.ValueOf(t.Format(TimeLayout(opts.timefmt))), opts)
}

func continueWithTimeFmt(timefmt string, k encoderFunc) encoderFunc {
	return func(e *encodeState, v reflect.Value, opts encOpts) {
		opts.timefmt = timefmt
		k(e, v, opts)
	}
}

func marshalWithTimeLayout(m Marshaler, layout string) ([]byte, error) {
	if tm, ok := m.(TimeMarshaler); ok {
		if b := tm.MarshalJSONWithTimeLayout(layout); b != nil {
			return b, nil
		}
	}
	return m.MarshalJSON()
}
