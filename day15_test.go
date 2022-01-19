package main

import "testing"

func TestDay15(t *testing.T) {
	example := `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`[1:]
	TestEqual(t, 27730, day15a(example))
	TestEqual(t, 4988, day15b(example))
	example = `
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`[1:]
	TestEqual(t, 36334, day15a(example))
	example = `
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`[1:]
	TestEqual(t, 39514, day15a(example))
	TestEqual(t, 31284, day15b(example))
	example = `
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`[1:]
	TestEqual(t, 27755, day15a(example))
	TestEqual(t, 3478, day15b(example))
	example = `
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`[1:]
	TestEqual(t, 28944, day15a(example))
	TestEqual(t, 6474, day15b(example))
	example = `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`[1:]
	TestEqual(t, 18740, day15a(example))
	TestEqual(t, 1140, day15b(example))

	TestEqual(t, 189000, day15a(Input(15)))
	TestEqual(t, 38512, day15b(Input(15)))
}
