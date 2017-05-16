// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	test := []struct {
		typ  DataCoding
		text []byte
		want []byte
	}{
		{Latin1Type, []byte("áéíóú moço"), []byte("\xe1\xe9\xed\xf3\xfa mo\xe7o")},
		{ISO88595Type, []byte("Тест кодировки"), []byte("\xc2\xd5\xe1\xe2 \xda\xde\xd4\xd8\xe0\xde\xd2\xda\xd8")},
		{UCS2Type, []byte("áéíóú moço"), []byte("\x00\xe1\x00\xe9\x00\xed\x00\xf3\x00\xfa\x00 \x00m\x00o\x00\xe7\x00o")},
		{SilentType, []byte("áéíóú moço"), []byte("")},
	}
	for _, tc := range test {
		have := Encode(tc.typ, tc.text)
		if !bytes.Equal(tc.want, have) {
			t.Fatalf("unexpected text for %#x:\nwant: %q\nhave: %q",
				tc.typ, tc.want, have)
		}
	}
}

func TestDecode(t *testing.T) {
	test := []struct {
		typ  DataCoding
		want []byte
		text []byte
	}{
		{Latin1Type, []byte("áéíóú moço"), []byte("\xe1\xe9\xed\xf3\xfa mo\xe7o")},
		{ISO88595Type, []byte("Тест кодировки"), []byte("\xc2\xd5\xe1\xe2 \xda\xde\xd4\xd8\xe0\xde\xd2\xda\xd8")},
		{UCS2Type, []byte("áéíóú moço"), []byte("\x00\xe1\x00\xe9\x00\xed\x00\xf3\x00\xfa\x00 \x00m\x00o\x00\xe7\x00o")},
		{SilentType, []byte("áéíóú moço"), []byte("áéíóú moço")},
	}
	for _, tc := range test {
		have := Decode(tc.typ, tc.text)
		if !bytes.Equal(tc.want, have) {
			t.Fatalf("unexpected text for %#x:\nwant: %q\nhave: %q",
				tc.typ, tc.want, have)
		}
	}
}
