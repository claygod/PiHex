package PiHex

// PiHex
// Configuration settings
// Copyright © 2016-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// Editable parameters
const (
	// The number of hexadecimal digits generated in one step.
	// For 64-bit architecture can not exceed nine.
	// If you have doubts about the accuracy, you can reduce this number
	// (it's a bit to increase the work program).
	STEP = 9 // or 8
)

// Non-editable parameters
const (
	// The limit for calculation
	CLIMIT = 10000000
	// The number of digits exponent
	NTP = 25
	// Special option
	EPS = 1e-17
)
