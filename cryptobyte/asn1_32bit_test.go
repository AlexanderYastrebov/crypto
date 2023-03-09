// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || arm || mips || mipsle

package cryptobyte

import (
	"fmt"
	"testing"
)

func TestASN1ObjectIdentifier32bit(t *testing.T) {
	testData := []readASN1ObjectIdentifierTest{
		{[]byte{6, 7, 0x55, 0x02, 0x88, 0x80, 0x80, 0x80, 0x00}, false, []int{}}, // 2**31
	}

	for i, test := range testData {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			testASN1ObjectIdentifier(t, test)
		})
	}
}
