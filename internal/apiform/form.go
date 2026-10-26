// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package apiform encodes request params as multipart form data.
package apiform

import (
	"bytes"
	"mime/multipart"
)

type Marshaler interface {
	MarshalMultipart() ([]byte, string, error)
}

func MarshalMultipart[T any](value T) ([]byte, string, error) {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	if err := MarshalWithSettings(value, writer, "repeat"); err != nil {
		return nil, "", err
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return buffer.Bytes(), writer.FormDataContentType(), nil
}
