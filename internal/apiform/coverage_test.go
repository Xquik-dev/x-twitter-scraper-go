// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package apiform

import (
	"bytes"
	"io"
	"mime/multipart"
	"reflect"
	"strings"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go/packages/param"
)

type metadataReader struct {
	io.Reader
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

type CoverageEmbedded struct {
	Embedded string `form:"embedded"`
}

type coverageForm struct {
	CoverageEmbedded
	hidden       string `form:"hidden"`
	NoTag        string
	Skip         string            `form:"-"`
	Empty        string            `form:""`
	Omit         string            `form:"omit,omitzero"`
	Rich         param.Opt[string] `form:"rich"`
	JSONFallback string            `json:"json"`
}

type invalidCoverageExtras struct {
	Extras map[int]string `form:"-,extras"`
}

type coverageTags struct {
	All        string `form:"all,required,extras,metadata,omitzero" api:"extrafields,required,metadata"`
	JSON       string `json:"json"`
	NoTag      string
	NonString  int    `form:"number" default:"ignored"`
	OnlyFormat string `form:"formatted" format:"date"`
}

func (metadataReader) Filename() string {
	return `tweet"timeline.json`
}

func (metadataReader) ContentType() string {
	return "application/json"
}

func newCoverageWriter(t *testing.T) (*bytes.Buffer, *multipart.Writer) {
	t.Helper()
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	if err := writer.SetBoundary("coverage"); err != nil {
		t.Fatal(err)
	}
	return &buffer, writer
}

func TestMarshalEntryPointsAndEdgeCases(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := Marshal(map[string]string{"query": "tweet search"}, writer); err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("MarshalRoot", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := MarshalRoot(map[string]string{"query": "timeline"}, writer); err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("nil", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := Marshal(nil, writer); err != nil {
			t.Fatal(err)
		}
		var pointer *Primitives
		if err := Marshal(pointer, writer); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("comma arrays", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := MarshalWithSettings([]int{}, writer, "comma"); err != nil {
			t.Fatal(err)
		}
		if err := MarshalWithSettings([]int{1, 2}, writer, "comma"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("unsupported array format", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := MarshalWithSettings([]int{1}, writer, "invalid"); err == nil {
			t.Fatal("unsupported array format was accepted")
		}
	})

	t.Run("unsupported primitive", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := Marshal(complex(1, 2), writer); err == nil {
			t.Fatal("unsupported primitive was accepted")
		}
	})

	t.Run("non-string map key", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := Marshal(map[int]string{1: "tweet"}, writer); err == nil {
			t.Fatal("non-string map key was accepted")
		}
	})

	t.Run("empty union", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		if err := Marshal(StructUnionWrapper{}, writer); err == nil {
			t.Fatal("empty union was accepted")
		}
	})

	t.Run("reader metadata", func(t *testing.T) {
		buffer, writer := newCoverageWriter(t)
		value := ReaderStruct{
			File: metadataReader{Reader: strings.NewReader("{}")},
		}
		if err := Marshal(value, writer); err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
		if got := buffer.String(); !strings.Contains(got, "application/json") ||
			!strings.Contains(got, `tweet\"timeline.json`) {
			t.Fatalf("multipart metadata = %s", got)
		}
	})
}

func TestWriteExtras(t *testing.T) {
	buffer, writer := newCoverageWriter(t)
	if err := WriteExtras(writer, map[string]any{"query": "tweet"}); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buffer.String(), "tweet") {
		t.Fatalf("multipart extras = %s", buffer.String())
	}

	_, writer = newCoverageWriter(t)
	if err := WriteExtras(writer, map[string]any{"page": 2}); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	writer = multipart.NewWriter(failingWriter{})
	if err := WriteExtras(writer, map[string]any{"query": "tweet"}); err == nil {
		t.Fatal("failing multipart writer accepted an extra field")
	}
}

func TestEncoderBranchContracts(t *testing.T) {
	t.Run("field selection and null option", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		value := coverageForm{
			CoverageEmbedded: CoverageEmbedded{Embedded: "tweet"},
			Rich:             param.Null[string](),
			JSONFallback:     "timeline",
		}
		if err := MarshalWithSettings(value, writer, "repeat"); err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("invalid extras key", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		err := Marshal(invalidCoverageExtras{Extras: map[int]string{1: "tweet"}}, writer)
		if err == nil || !strings.Contains(err.Error(), "non string key") {
			t.Fatalf("invalid extras error = %v", err)
		}
	})

	t.Run("array element error", func(t *testing.T) {
		_, writer := newCoverageWriter(t)
		err := MarshalWithSettings([]complex128{complex(1, 2)}, writer, "repeat")
		if err == nil || !strings.Contains(err.Error(), "unknown type") {
			t.Fatalf("array error = %v", err)
		}
	})

	t.Run("nil interface", func(t *testing.T) {
		var value any
		_, writer := newCoverageWriter(t)
		encode := (encoder{arrayFmt: "repeat"}).newInterfaceEncoder()
		if err := encode("value", reflect.ValueOf(&value).Elem(), writer); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("key formats", func(t *testing.T) {
		tests := map[string][]string{
			"comma":            {"", "item"},
			"repeat":           {"", "item"},
			"brackets":         {"[]", "item[]"},
			"indices:dots":     {"0", "item.1"},
			"indices:brackets": {"0", "item[1]"},
		}
		for format, want := range tests {
			key := (encoder{arrayFmt: format}).arrayKeyEncoder()
			if key == nil {
				t.Fatalf("%s returned no key encoder", format)
			}
			got := []string{key("", 0), key("item", 1)}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%s keys = %v, want %v", format, got, want)
			}
		}
		if key := (encoder{arrayFmt: "invalid"}).arrayKeyEncoder(); key != nil {
			t.Fatal("invalid format returned a key encoder")
		}
	})
}

func TestFormTagBranches(t *testing.T) {
	typ := reflect.TypeOf(coverageTags{})
	all, ok := parseFormStructTag(typ.Field(0))
	if !ok || all.name != "all" || !all.required || !all.extras ||
		!all.metadata || !all.omitzero {
		t.Fatalf("all tag = %+v, %t", all, ok)
	}
	jsonTag, ok := parseFormStructTag(typ.Field(1))
	if !ok || jsonTag.name != "json" {
		t.Fatalf("JSON tag = %+v, %t", jsonTag, ok)
	}
	if _, ok := parseFormStructTag(typ.Field(2)); ok {
		t.Fatal("untagged field was accepted")
	}
	nonString, ok := parseFormStructTag(typ.Field(3))
	if !ok || nonString.defaultValue != nil {
		t.Fatalf("non-string default = %+v, %t", nonString, ok)
	}
	if format, ok := parseFormatStructTag(typ.Field(4)); !ok || format != "date" {
		t.Fatalf("format = %q, %t", format, ok)
	}
}
