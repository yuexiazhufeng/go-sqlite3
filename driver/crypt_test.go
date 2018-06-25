// Copyright (C) 2018 The Go-SQLite3 Authors.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

import (
	"fmt"
	"strings"
	"testing"
)

// TestCryptEncoders to increase coverage
func TestEncoders(t *testing.T) {
	tests := []struct {
		enc      string
		salt     string
		expected string
	}{
		{"sha1", "", "d033e22ae348aeb5660fc2140aec35850c4da997"},
		{"sha256", "", "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"},
		{"sha384", "", "9ca694a90285c034432c9550421b7b9dbd5c0f4b6673f05f6dbce58052ba20e4248041956ee8c9a2ec9f10290cdc0782"},
		{"sha512", "", "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec"},
		{"ssha1", "salt", "9bc7aa55f08fdad935c3f8362d3f48bcf70eb280"},
		{"ssha256", "salt", "f9a81477552594c79f2abc3fc099daa896a6e3a3590a55ffa392b6000412e80b"},
		{"ssha384", "salt", "9ed776b477fcfc1b5e584989e8d770f5e17d98a7643546a63c2b07d4ab00f1348f6b8e73103d3a23554f727136e8c215"},
		{"ssha512", "salt", "3c4a79782143337be4492be072abcfe979dd703c00541a8c39a0f3df4bab2029c050cf46fddc47090b5b04ac537b3e78189e3de16e601e859f95c51ac9f6dafb"},
	}

	for _, e := range tests {
		var enc CryptEncoder
		switch e.enc {
		case "sha1":
			enc = NewSHA1Encoder()
		case "ssha1":
			enc = NewSSHA1Encoder(e.salt)
		case "sha256":
			enc = NewSHA256Encoder()
		case "ssha256":
			enc = NewSSHA256Encoder(e.salt)
		case "sha384":
			enc = NewSHA384Encoder()
		case "ssha384":
			enc = NewSSHA384Encoder(e.salt)
		case "sha512":
			enc = NewSHA512Encoder()
		case "ssha512":
			enc = NewSSHA512Encoder(e.salt)
		}

		if _, ok := enc.(CryptSaltedEncoder); ok {
			if strings.Compare(enc.(CryptSaltedEncoder).Salt(), e.salt) != 0 {
				t.Fatal("Salt Mismatch")
			}
		}

		h := enc.Encode([]byte("admin"), nil)
		if strings.Compare(fmt.Sprintf("%x", h), e.expected) != 0 {
			t.Fatalf("Invalid %s hash: expected: %s; got: %x", strings.ToUpper(e.enc), e.expected, h)
		}

		if e.enc != enc.String() {
			t.Fatalf("Invalid encoder; expected: %s, got: %s", e.enc, enc.String())
		}
	}
}
