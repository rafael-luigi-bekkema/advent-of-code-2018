package main

import "testing"

func TestDay17(t *testing.T) {
	TestEqual(t, 57, day17a([]string{
		"x=495, y=2..7",
		"y=7, x=495..501",
		"x=501, y=3..7",
		"x=498, y=2..4",
		"x=506, y=1..2",
		"x=498, y=10..13",
		"x=504, y=10..13",
		"y=13, x=498..504",
	}))
}
