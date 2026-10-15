// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package ssestream

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const maxLineSize = 16 << 20

// Decoder reads server-sent events one at a time.
type Decoder interface {
	Event() Event
	Next() bool
	Close() error
	Err() error
}

// Event is 1 server-sent event with its type and its joined data lines.
type Event struct {
	Type string
	Data []byte
}

// NewDecoder reads text/event-stream events from the body of res. It returns nil
// when res or its body is nil.
func NewDecoder(res *http.Response) Decoder {
	if res == nil || res.Body == nil {
		return nil
	}
	scanner := bufio.NewScanner(res.Body)
	scanner.Buffer(nil, maxLineSize)
	return &eventStreamDecoder{body: res.Body, scanner: scanner}
}

type eventStreamDecoder struct {
	body    io.ReadCloser
	scanner *bufio.Scanner
	event   Event
	err     error
	started bool
}

func (d *eventStreamDecoder) Next() bool {
	eventType := ""
	var data bytes.Buffer
	for d.scanner.Scan() {
		line := d.scanner.Bytes()
		if !d.started {
			line = bytes.TrimPrefix(line, []byte("\uFEFF"))
			d.started = true
		}
		if len(line) == 0 {
			if data.Len() == 0 {
				eventType = ""
				continue
			}
			d.event = Event{Type: eventType, Data: bytes.TrimSuffix(data.Bytes(), []byte("\n"))}
			return true
		}
		name, value, _ := bytes.Cut(line, []byte(":"))
		value = bytes.TrimPrefix(value, []byte(" "))
		switch string(name) {
		case "event":
			eventType = string(value)
		case "data":
			data.Write(value)
			data.WriteByte('\n')
		}
	}
	d.err = d.scanner.Err()
	return false
}

func (d *eventStreamDecoder) Event() Event { return d.event }

func (d *eventStreamDecoder) Close() error { return d.body.Close() }

func (d *eventStreamDecoder) Err() error { return d.err }

// Stream decodes the data of each server-sent event as 1 value of T. Events
// without data, such as keep-alive comments, are skipped.
//
//	for stream.Next() {
//		event := stream.Current()
//	}
//	if err := stream.Err(); err != nil {
//		return err
//	}
type Stream[T any] struct {
	decoder Decoder
	current T
	err     error
}

// NewStream returns a stream over decoder. A non-nil err ends the stream before
// its first value, and Err returns it.
func NewStream[T any](decoder Decoder, err error) *Stream[T] {
	return &Stream[T]{decoder: decoder, err: err}
}

// Next reads the next value. It returns false when the stream ends or fails.
func (s *Stream[T]) Next() bool {
	if s.err != nil || s.decoder == nil {
		return false
	}
	if !s.decoder.Next() {
		s.err = s.decoder.Err()
		return false
	}
	var next T
	if s.err = json.Unmarshal(s.decoder.Event().Data, &next); s.err != nil {
		return false
	}
	s.current = next
	return true
}

// Current returns the value that the last call to Next read.
func (s *Stream[T]) Current() T { return s.current }

// Err returns the error that ended the stream, or nil after a clean end.
func (s *Stream[T]) Err() error { return s.err }

// Close closes the response body behind the stream.
func (s *Stream[T]) Close() error {
	if s.decoder == nil {
		return nil
	}
	return s.decoder.Close()
}
