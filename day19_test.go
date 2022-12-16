package main

import "testing"

func TestDay19a(t *testing.T) {
	expect := 1488
	result := day19(Lines(19), true)
	TestEqual(t, expect, result)
}

func TestDay19b(t *testing.T) {
	expect := 17427456
	result := day19(Lines(19), false)
	TestEqual(t, expect, result)
}
