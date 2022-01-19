package main

import "testing"

func TestDay16(t *testing.T) {
	TestEqual(t, 1, day16a([]string{
		"Before: [3, 2, 1, 1]",
		"9 2 1 2",
		"After:  [3, 2, 2, 1]",
	}))
	file := Lines(16)
	TestEqual(t, 542, day16a(file))
	TestEqual(t, 575, day16b(file))
}
