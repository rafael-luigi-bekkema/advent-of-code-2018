package main

import "testing"

func TestDay1(t *testing.T) {
	TestEqual(t, 3, day1a([]string{"+1", "-2", "+3", "+1"}))
	TestEqual(t, 2, day1b([]string{"+1", "-2", "+3", "+1"}))
	TestEqual(t, 0, day1b([]string{"+1", "-1"}))
	TestEqual(t, 10, day1b([]string{"+3", "+3", "+4", "-2", "-4"}))
	TestEqual(t, 5, day1b([]string{"-6", "+3", "+8", "+5", "-6"}))
	TestEqual(t, 14, day1b([]string{"+7", "+7", "-2", "-7", "-4"}))

	file := Lines(1)
	TestEqual(t, 574, day1a(file))
	TestEqual(t, 452, day1b(file))
}
