package PiHex

// PiHex
// Work
// Copyright © 2016-2020 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"fmt"
	"math"
)

// Pi structure
type Pi struct {
	// Table exponents
	exponent *[NTP]float64
	// Channel to return the results of the goroutines
	ch chan float64
}

// genHex - generate 'num' hexadecimal digits in the number Pi since 'start' position
// The most time consuming calculations are performed in parallel. Arguments:
// start - start number
// num - how to calculate the numbers
func (pi *Pi) genHex(start int, num int) []byte {
	var d1 float64
	var d2 float64
	// Start of the goroutines
	go pi.series(1, start, 4.)
	go pi.series(4, start, 2.)
	go pi.series(5, start, 1.)
	go pi.series(6, start, 1.)
	// Getting the results of the goroutines
	for i := 0; i < 4; i++ {
		d2 = <-pi.ch

		if d2 > 8. {
			d1 += d2
		} else {
			d1 -= d2
		}
	}

	pid := d1 - 10.
	pid = pid - float64(int(pid)) + 1.

	return pi.ihex(pid, num)
}

// genExp - generate table exponents
func (pi *Pi) genExp() {
	var exp [NTP]float64
	exp[0] = 0.
	exp[1] = 1.

	for i := 2; i < NTP; i++ {
		exp[i] = 2. * exp[i-1]
	}

	pi.exponent = &exp
}

// series - This routine evaluates the series  sum 16^(id-k)/(8*k+m)
// using the modular exponentiation technique. Arguments:
// m - 1,4,5,6 only
// id - start number
// id - coefficient
func (pi *Pi) series(m int, id int, kf float64) {
	var ak, p, s, t float64
	//  Sum the series up to id
	for k := 0; k < id; k++ {
		ak = 8*float64(k) + float64(m)
		p = float64(id) - float64(k)
		t = pi.expm(p, ak)
		s += t / ak
		s = s - float64(int(s))
	}
	//  Compute a few terms where k >= id
	for k := id; k <= id+100; k++ {
		ak = 8*float64(k) + float64(m)
		t = math.Pow(16., float64(id-k)) / ak

		if t < EPS {
			break
		}

		s = s + t
		s = s - float64(int(s))
	}

	s = s * kf

	if kf == 4. {
		s += 10.
	}

	pi.ch <- s
}

func (pi *Pi) expm(p float64, ak float64) float64 {
	var p1, pt, r float64

	if ak == 1. {
		return 0.
	}
	//  Find the greatest power of two less than or equal to p
	i := 0
	for ; i < NTP; i++ {
		if pi.exponent[i] > p {
			break
		}
	}
	pt = pi.exponent[i-1]
	p1 = p
	r = 1.
	//  Perform binary exponentiation algorithm modulo ak
	for j := 1; j <= i; j++ {
		if p1 >= pt {
			r = 16. * r
			r = r - float64(float64(int(r/ak))*ak)
			p1 = p1 - pt
		}

		pt = 0.5 * pt

		if pt >= 1. {
			r = r * r
			r = r - float64(float64(int(r/ak))*ak)
		}
	}

	return r
}

func (pi *Pi) ihex(x float64, num int) []byte {
	var out []byte
	var y float64
	y = math.Abs(x)

	for i := 0; i < num; i++ {
		y = 16. * (y - math.Floor(y))
		out = append(out, byte(y))
	}

	return out
}
