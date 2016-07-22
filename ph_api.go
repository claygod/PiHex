// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package PiHex

// PiHex
// API

import (
	//"fmt"
	"runtime"
)

// PiHex Library generates a hexadecimal number sequence in the number Pi in
// the range from 0 to 10,000,000. To calculate using "Bailey-Borwein-Plouffe"
// algorithm, instructions that was published by David September 17, 2006.

// New - create a new PiHex-struct
func New() *pi {
	pi := &pi{}
	pi.genExp()
	pi.ch = make(chan float64)
	return pi
}

// Get - 'num' hexadecimal digits in the number Pi since 'start' position.
// If the inputs exceed the permissible values, it returns an empty slice.
//  Arguments:
// start - start number
// num - how to calculate the numbers
//  Return:
// slice bytes bytes 0 to 15
func (pi *pi) GetHex(start int, num int) []byte {
	var out []byte
	if start <= CLIMIT &&
		start+num < CLIMIT &&
		start >= 0 &&
		num > 0 {
		numcpu := runtime.NumCPU()
		runtime.GOMAXPROCS(numcpu)
		for i := 0; i < num; i = i + STEP {
			c := num - i
			if c > STEP {
				out = append(out, pi.genHex(start+i, STEP)...)
			} else {
				out = append(out, pi.genHex(start+i, c)...)
			}
		}
	}
	return out
}
