// Copyright 2013 Beego Authors
// Copyright 2014 The Macaron Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package session

import (
	"bytes"
	"crypto/rand"
	"io"

	"github.com/unknwon/com"
	"github.com/vmihailenco/msgpack/v5"
)

// MessagePack doesn't require type registration like gob

// Encode encodes obj with MessagePack
func Encode(obj map[interface{}]interface{}) ([]byte, error) {
	return msgpack.Marshal(obj)
}

// Decode decodes bytes to obj
func Decode(encoded []byte) (out map[interface{}]interface{}, err error) {
	out, err = msgpack.NewDecoder(bytes.NewReader(encoded)).DecodeUntypedMap()
	return out, err
}

// NOTE: A local copy in case of underlying package change
var alphanum = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// generateRandomKey creates a random key with the given strength.
func generateRandomKey(strength int) []byte {
	k := make([]byte, strength)
	if n, err := io.ReadFull(rand.Reader, k); n != strength || err != nil {
		return com.RandomCreateBytes(strength, alphanum...)
	}
	return k
}
