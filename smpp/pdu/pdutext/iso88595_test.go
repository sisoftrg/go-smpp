// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

import (
	"bytes"
	"testing"
)

func TestISO88595Encoder(t *testing.T) {
	want := []byte("\xc2\xd5\xe1\xe2 \xda\xde\xd4\xd8\xe0\xde\xd2\xda\xd8")
	text := []byte("Тест кодировки")
	s := ISO88595(text)
	if s.Type() != 0x06 {
		t.Fatalf("Unexpected data type; want 0x06, have %d", s.Type())
	}
	have := s.Encode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}

func TestISO88595Decoder(t *testing.T) {
	want := []byte("Тест кодировки")
	text := []byte("\xc2\xd5\xe1\xe2 \xda\xde\xd4\xd8\xe0\xde\xd2\xda\xd8")
	s := ISO88595(text)
	if s.Type() != 0x06 {
		t.Fatalf("Unexpected data type; want 0x06, have %d", s.Type())
	}
	have := s.Decode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}
