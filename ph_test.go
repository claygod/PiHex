// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package PiHex

// PiHex
// Test

import (
	"testing"
)

func TestLimitZero(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 0)) != 0 {
		t.Error("Error two zeros in the argument /pi.Get(0, 0)")
	}
}

func TestLimitCount1(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 5)) != 5 {
		t.Error("It returned the wrong number of digits /pi.Get(0, 5)")
	}
}

func TestLimitCount2(t *testing.T) {
	pi := New()
	if len(pi.Get(5, 5)) != 5 {
		t.Error("It returned the wrong number of digits /pi.Get(5, 5)")
	}
}

func TestLimitNegative1(t *testing.T) {
	pi := New()
	if len(pi.Get(-1, 5)) != 0 {
		t.Error("Adopted by Negative arguments /pi.Get(-1, 5)")
	}
}

func TestLimitNegative2(t *testing.T) {
	pi := New()
	if len(pi.Get(0, -5)) != 0 {
		t.Error("Adopted by Negative arguments /pi.Get(0, -5)")
	}
}

func TestLimitNegative3(t *testing.T) {
	pi := New()
	if len(pi.Get(-1, -5)) != 0 {
		t.Error("Adopted by Negative arguments /pi.Get(-1, -5)")
	}
}

func TestLimitMax1(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 10000001)) != 0 {
		t.Error("The value argument exceeds the limit /pi.Get(0, 10000001)")
	}
}

func TestLimitMax2(t *testing.T) {
	pi := New()
	if len(pi.Get(9999999, 2)) != 0 {
		t.Error("The value argument exceeds the limit /pi.Get(9999999, 2)")
	}
}

func TestLimitResult1(t *testing.T) {
	pi := New()
	tru := []byte{2, 4, 3, 15, 6, 10, 8, 8, 8, 5, 10, 3, 0, 8, 13, 3, 1, 3, 1, 9}
	result := pi.Get(0, len(tru))
	for i := 0; i < len(tru); i++ {
		if tru[i] != result[i] {
			t.Error("Wrong Result ", result[i], " at position ", i, " /pi.Get(0, len(tru)")
		}
	}
}

func TestLimitResult2(t *testing.T) {
	pi := New()
	tru := []byte{6, 12, 6, 5, 14, 5, 2, 12, 11, 4}
	result := pi.Get(1000000, len(tru))
	for i := 0; i < len(tru); i++ {
		if tru[i] != result[i] {
			t.Error("Wrong Result ", result[i], " at position ", i, " /pi.Get(1000000, len(tru)")
		}
	}
}
