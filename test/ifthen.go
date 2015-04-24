// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test if statements in various forms.

package main

import (
	"fmt"
)

func assertequal(is, shouldbe int, msg string) {
	if is != shouldbe {
		print("assertion fail", msg, "\n")
		panic(1)
	}
}

func asserttest(test bool, msg string) {
	if !test {
		print("assertion fail", msg, "\n")
		panic(1)
	}
}

func max(a, b int) int {
	return a > b then a else b
}

func min(a, b int) int {
	return a < b then a else b
}

func main() {
	i5 := 5
	i7 := 7

	fmt.Printf("max(%d, %d) = %d\n", i5, i7, max(i5, i7))
	fmt.Printf("max(%d, %d) = %d\n", i7, i5, max(i7, i5))

	fmt.Printf("min(%d, %d) = %d\n", i5, i7, min(i5, i7))
	fmt.Printf("min(%d, %d) = %d\n", i7, i5, min(i7, i5))

	a := i5 == i7 then i5 else i7
	assertequal(a, i7, "i5 == i7")

	a = i5 == i5 then i5 else i7
	assertequal(a, i5, "i5 != i5")

	a = i7 == i7 then i7 else i5
	assertequal(a, i7, "i7 != i7")

	t := i7 > i5 then true else false
	asserttest(t, "i7 > i5")

	p := i7 > i5 then int64(1) else int64(0)
	fmt.Printf("p = %d\n", p)
}