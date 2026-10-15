// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package ssestream

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type value struct {
	N int `json:"n"`
}

func response(body io.Reader) *http.Response {
	return &http.Response{Body: io.NopCloser(body)}
}

const twoEvents = ": connected\n\nid: 42\nevent: tweet.new\ndata: {\"n\":1}\n\n: keep-alive\n\nevent: ignored\n\nid: 43\r\nevent: tweet.reply\r\ndata: {\"n\":\r\ndata:2}\r\n\r\nretry: 1000\n\n"

func TestDecoderJoinsDataAndSkipsComments(t *testing.T) {
	decoder := NewDecoder(response(strings.NewReader(twoEvents)))
	var got []Event
	for decoder.Next() {
		got = append(got, decoder.Event())
	}
	if decoder.Err() != nil || decoder.Close() != nil {
		t.Fatalf("err = %v", decoder.Err())
	}
	if len(got) != 2 ||
		got[0].Type != "tweet.new" || string(got[0].Data) != `{"n":1}` ||
		got[1].Type != "tweet.reply" || string(got[1].Data) != "{\"n\":\n2}" {
		t.Fatalf("events = %q", got)
	}
}

func TestDecoderSkipsALeadingByteOrderMark(t *testing.T) {
	decoder := NewDecoder(response(strings.NewReader("\uFEFFevent: tweet.new\ndata: {\"n\":1}\n\n")))
	if !decoder.Next() || decoder.Event().Type != "tweet.new" || string(decoder.Event().Data) != `{"n":1}` {
		t.Fatalf("event = %q, err = %v", decoder.Event(), decoder.Err())
	}
}

func TestStreamDecodesEachEvent(t *testing.T) {
	stream := NewStream[value](NewDecoder(response(strings.NewReader(twoEvents))), nil)
	var got []int
	for stream.Next() {
		got = append(got, stream.Current().N)
	}
	if stream.Err() != nil || stream.Close() != nil || len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("values = %v, err = %v", got, stream.Err())
	}
}

func TestStreamDropsAnUnfinishedEvent(t *testing.T) {
	stream := NewStream[value](NewDecoder(response(strings.NewReader("data: {\"n\":1}\n\ndata: {\"n\":2}"))), nil)
	var got []int
	for stream.Next() {
		got = append(got, stream.Current().N)
	}
	if stream.Err() != nil || len(got) != 1 || got[0] != 1 {
		t.Fatalf("values = %v, err = %v", got, stream.Err())
	}
}

func TestStreamStopsOnInvalidData(t *testing.T) {
	stream := NewStream[value](NewDecoder(response(strings.NewReader("data: {\"n\":1}\n\ndata: not json\n\ndata: {\"n\":3}\n\n"))), nil)
	if !stream.Next() || stream.Current().N != 1 {
		t.Fatalf("first value = %v, err = %v", stream.Current(), stream.Err())
	}
	if stream.Next() || stream.Err() == nil || stream.Next() {
		t.Fatalf("invalid data did not end the stream, err = %v", stream.Err())
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("connection reset") }

func TestStreamReportsReadErrors(t *testing.T) {
	stream := NewStream[value](NewDecoder(response(failingReader{})), nil)
	if stream.Next() || stream.Err() == nil || stream.Err().Error() != "connection reset" {
		t.Fatalf("err = %v", stream.Err())
	}
}

func TestStreamWithoutResponse(t *testing.T) {
	if NewDecoder(nil) != nil || NewDecoder(&http.Response{}) != nil {
		t.Fatal("a missing response or body must give no decoder")
	}
	failed := NewStream[value](nil, errors.New("401 Unauthorized"))
	if failed.Next() || failed.Err() == nil || failed.Err().Error() != "401 Unauthorized" || failed.Close() != nil {
		t.Fatalf("err = %v", failed.Err())
	}
	empty := NewStream[value](nil, nil)
	if empty.Next() || empty.Err() != nil || empty.Close() != nil {
		t.Fatalf("err = %v", empty.Err())
	}
}
