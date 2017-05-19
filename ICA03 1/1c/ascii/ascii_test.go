// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package ascii

import (
	"testing"
	"unicode"
)

func TestGreetingASCII(t *testing.T) {
	b := GreetingASCII()
	if isASCII(b) == false {
		t.Errorf("FAIL")
	}
}

func isASCII(t2 string) bool {
	for _, i := range t2 {
		if i > unicode.MaxASCII {
			return false
		}
	}
	return true
}
