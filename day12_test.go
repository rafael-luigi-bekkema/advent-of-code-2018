package main

import "testing"

func TestDay12(t *testing.T) {
	TestEqual(t, 325, day12a(20, Lines(12, "_example")))
	TestEqual(t, 3337, day12a(20, Lines(12)))
	TestEqual(t, 17549, day12a(200, Lines(12)))

	TestEqual(t, 4300000000349, day12a(50_000_000_000, Lines(12)))
}
