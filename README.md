# PiHex
PiHex Library generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000. To calculate using "Bailey-Borwein-Plouffe" algorithm, instructions that was published by David H. Bailey September 17, 2006.

[![API documentation](https://godoc.org/github.com/claygod/PiHex?status.svg)](https://godoc.org/github.com/claygod/PiHex)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/PiHex)](https://goreportcard.com/report/github.com/claygod/PiHex)

# Usage

An example of using the PiHex Library:
```go
package main

import (
	"fmt"
	"github.com/claygod/PiHex"
)

func main() {
	pi := PiHex.New()
	fmt.Print("The first 20 digits of Pi (hexadecimal): ", pi.Get(0, 20))
}
```

# Settings

In the configuration file, you can change the constant STEP. This constant determines the amount generated in one step numbers. The reduction leads to a constant increase in the operating time of the program.

Attention! This constant can not be more than 9! Limitation due to the 64-bit library architecture.

The configuration file [config.go](https://github.com/claygod/PiHex/blob/master/ph_config.go)

# Perfomance

To optimize the run-time program, highly loaded sections of the library are performed in parallel (4 goroutines).

# API

Methods:
-  *New* - create a new PiHex
-  *Get* - receiving a sequence of hexadecimal digits starting at the specified position and in the right quantity.

Example:
`
pi := PiHex.New()
x :=  pi.Get(1000, 5)
`

# Algorithm

The Bailey–Borwein–Plouffe formula (BBP formula) is a spigot algorithm for computing the nth binary digit of Pi using base 16 math. The formula can directly calculate the value of any given digit of π without calculating the preceding digits. The BBP is a summation-style formula that was discovered in 1995 by Simon Plouffe and was named after the authors of the paper in which the formula was published, David H. Bailey, Peter Borwein, and Simon Plouffe.

# Implementation

Plays Library is based on the publication "The BBP Algorithm for Pi" of David H. Baileyon September 17, 2006: http://www.davidhbailey.com/dhbpapers/bbp-alg.pdf

