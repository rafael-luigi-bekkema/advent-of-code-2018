package main

import "testing"

func TestDay18a(t *testing.T) {
	initial := `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`
	expect := 1147
	result := day18a(initial, 10)

	TestEqual(t, expect, result)
}

func TestDay18a2(t *testing.T) {
	expect := 558960
	result := day18a(Input(18), 10)

	TestEqual(t, expect, result)
}

func TestDay18b(t *testing.T) {
	initial := `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`
	expect := 0
	result := day18a(initial, 1_000_000_000)

	TestEqual(t, expect, result)
}

func TestDay18b2(t *testing.T) {
	expect := 207900
	result := day18a(Input(18), 1_000_000_000)

	TestEqual(t, expect, result)
}
