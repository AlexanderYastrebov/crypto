// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(386 || arm || mips || mipsle)

package cryptobyte

import (
	"fmt"
	"testing"
)

func TestASN1ObjectIdentifier64bit(t *testing.T) {
	testData := []readASN1ObjectIdentifierTest{
		{[]byte{6, 7, 0x55, 0x02, 0x88, 0x80, 0x80, 0x80, 0x00}, true, []int{2, 5, 2, 2147483648}},                                        // 2**31
		{[]byte{6, 11, 0x2a, 0x24, 0xcb, 0x89, 0x90, 0x82, 0x1e, 0x03, 0x01, 0x01, 0x01}, true, []int{1, 2, 36, 20151795998, 3, 1, 1, 1}}, // https://github.com/golang/go/issues/58821
		{[]byte{6, 11, 0x55, 0x02, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, true, []int{2, 5, 2, 9223372036854775807}},      // 2**63-1
		{[]byte{0, 12, 0x55, 0x02, 0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x00}, false, []int{}},                           // 2**63
	}

	for i, test := range testData {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			testASN1ObjectIdentifier(t, test)
		})
	}
}
