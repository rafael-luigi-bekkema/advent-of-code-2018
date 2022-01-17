package main

import "testing"

func TestDay5(t *testing.T) {
	TestEqual(t, 10, day5a("dabAcCaCBAcCcaDA"))
	TestEqual(t, 4, day5b("dabAcCaCBAcCcaDA"))
	file := Input(5)
	TestEqual(t, 10368, day5a(file))
	TestEqual(t, 4122, day5b(file))
}
