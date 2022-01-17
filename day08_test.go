package main

import "testing"

func TestDay8(t *testing.T) {
	TestEqual(t, NewResult(138, 66), day8a("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"))
	file := Input(8)
	TestEqual(t, NewResult(49180, 20611), day8a(file))
}
