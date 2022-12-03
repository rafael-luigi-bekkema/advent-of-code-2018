package main

import "fmt"

func main() {
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
	fmt.Println(day18a(initial, 1_000_000_000))

	// fmt.Println(day18a(Input(18), 1_000_000_000))
}
