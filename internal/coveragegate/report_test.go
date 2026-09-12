// Copyright the Xquik contributors.
// SPDX-License-Identifier: Apache-2.0

package coveragegate

import (
	"maps"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

func reportFiles() fstest.MapFS {
	return fstest.MapFS{
		"coverage.out":   {Data: []byte("mode: set\nexample/a.go:1.1,2.2 2 1\nexample/a.go:1.1,2.2 2 0\nexample/a.go:3.1,4.2 1 0\n")},
		"coverage.xml":   {Data: []byte(`<coverage lines-covered="2" lines-valid="3"><packages><package name="example"><classes><class filename="a.go"><methods><method line-rate="1"/><method line-rate="0"/><method line-rate="0.5"/></methods></class></classes></package></packages></coverage>`)},
		"branches.json":  {Data: []byte(`{"packages":[{"import_path":"example","files":[{"name":"a.go","blocks":[{"idx":1,"line":1,"col":1,"end_line":2,"end_col":2}],"branches":[{"block_idx":1}]}]}]}`)},
		"branches.cover": {Data: []byte("mode: set\nexample/a.go:1.1,2.2 2 1\n")},
	}
}

func TestMeasure(t *testing.T) {
	got, err := Measure(reportFiles())
	want := map[string]Metric{"statements": {3, 2, 0}, "functions": {3, 2, 0}, "lines": {3, 2, 0}, "branches": {1, 1, 0}}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("Measure() = %#v, %v; want %#v", got, err, want)
	}
	for name := range reportFiles() {
		t.Run("missing/"+name, func(t *testing.T) {
			files := reportFiles()
			delete(files, name)
			if _, err := Measure(files); err == nil || !strings.Contains(err.Error(), name) {
				t.Fatalf("missing %s error = %v", name, err)
			}
		})
	}
	for _, test := range []struct{ name, file, data string }{
		{"bad profile", "coverage.out", "broken"},
		{"bad XML", "coverage.xml", "<coverage"},
		{"wrong XML root", "coverage.xml", "<unrelated/>"},
		{"bad branch JSON", "branches.json", "{"},
		{"bad branch profile", "branches.cover", "broken"},
		{"negative method rate", "coverage.xml", "-1"},
		{"excessive method rate", "coverage.xml", "2"},
		{"NaN method rate", "coverage.xml", "NaN"},
		{"infinite method rate", "coverage.xml", "+Inf"},
	} {
		t.Run(test.name, func(t *testing.T) {
			files := reportFiles()
			data := test.data
			if strings.Contains(test.name, "method rate") {
				data = strings.Replace(string(files[test.file].Data), `line-rate="1"`, `line-rate="`+data+`"`, 1)
			}
			files[test.file] = &fstest.MapFile{Data: []byte(data)}
			if _, err := Measure(files); err == nil {
				t.Fatal("invalid report passed")
			}
		})
	}
}

func TestBaseline(t *testing.T) {
	files := fstest.MapFS{"report.json": {Data: []byte(`{"Revision":"parent","Scope":"scope","Total":{"lines":{"Total":3,"Covered":2}}}`)}}
	got, err := Baseline(files, "parent", "scope")
	if err != nil || got.Total["lines"] != (Metric{3, 2, 0}) {
		t.Fatalf("Baseline() = %#v, %v", got, err)
	}
	for _, pair := range [][2]string{{"", "scope"}, {"parent", ""}, {"other", "scope"}, {"parent", "other"}} {
		if _, err := Baseline(files, pair[0], pair[1]); err == nil {
			t.Fatalf("mismatched baseline passed: %v", pair)
		}
	}
	for _, data := range [][]byte{nil, []byte("{"), []byte(`{"Revision":true}`)} {
		invalid := maps.Clone(files)
		if data == nil {
			delete(invalid, "report.json")
		} else {
			invalid["report.json"] = &fstest.MapFile{Data: data}
		}
		if _, err := Baseline(invalid, "parent", "scope"); err == nil {
			t.Fatal("invalid baseline passed")
		}
	}
}

func TestMeasureRejectsMismatchedInventories(t *testing.T) {
	for _, name := range []string{"omitted source", "unexpected source", "duplicate source"} {
		t.Run(name, func(t *testing.T) {
			files := reportFiles()
			switch name {
			case "omitted source":
				files["coverage.out"].Data = append(files["coverage.out"].Data, []byte("example/missing.go:1.1,2.1 1 0\n")...)
			case "unexpected source":
				files["coverage.xml"].Data = []byte(strings.Replace(string(files["coverage.xml"].Data), `filename="a.go"`, `filename="other.go"`, 1))
			case "duplicate source":
				files["coverage.xml"].Data = []byte(strings.Replace(string(files["coverage.xml"].Data), "</classes>", `<class filename="a.go"/></classes>`, 1))
			}
			if _, err := Measure(files); err == nil {
				t.Fatal("mismatched source inventory passed")
			}
		})
	}
}
