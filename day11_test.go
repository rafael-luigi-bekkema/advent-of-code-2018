package main

import "testing"

func TestDay11(t *testing.T) {
	TestEqual(t, 4, day11pow(3, 5, 8))
	TestEqual(t, -5, day11pow(122, 79, 57))
	TestEqual(t, 0, day11pow(217, 196, 39))
	TestEqual(t, 4, day11pow(101, 153, 71))

	TestEqual(t, "33,45", day11a(18))
	TestEqual(t, "21,61", day11a(42))

	TestEqual(t, "235,18", day11a(5153))

	TestEqual(t, "90,269,16", day11b(18))
	TestEqual(t, "232,251,12", day11b(42))

	TestEqual(t, "236,227,12", day11b(5153))
}
