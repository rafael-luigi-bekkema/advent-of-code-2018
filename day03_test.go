package main

import "testing"

func TestDay3(t *testing.T) {
	TestEqual(t, day3result{4, 3}, day3a([]string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}))
	TestEqual(t, day3result{103806, 625}, day3a(Lines(3)))
}
