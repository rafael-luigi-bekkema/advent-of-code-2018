package main

import "testing"

func TestDay14(t *testing.T) {
	TestEqual(t, [10]int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}, day14a(9))
	TestEqual(t, [10]int{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}, day14a(5))
	TestEqual(t, [10]int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}, day14a(18))
	TestEqual(t, [10]int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}, day14a(2018))

	TestEqual(t, [10]int{9, 2, 1, 1, 1, 3, 4, 3, 1, 5}, day14a(77201))

	TestEqual(t, 9, day14b(5, 1, 5, 8, 9))
	TestEqual(t, 5, day14b(0, 1, 2, 4, 5))
	TestEqual(t, 18, day14b(9, 2, 5, 1, 0))
	TestEqual(t, 2018, day14b(5, 9, 4, 1, 4))
	TestEqual(t, 20357548, day14b(0, 7, 7, 2, 0, 1))
}
