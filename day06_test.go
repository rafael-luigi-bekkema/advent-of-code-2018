package main

import "testing"

func TestDay6(t *testing.T) {
	TestEqual(t, 17, day6a([]string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}))
	TestEqual(t, 16, day6b([]string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}, 32))
	file := Lines(6)
	TestEqual(t, 3449, day6a(file))
	TestEqual(t, 44868, day6b(file, 10_000))
}
