package main

import (
	"fmt"
	"testing"
)

func TestDay9(t *testing.T) {
	TestEqual(t, 32, day9(9, 25))
	TestEqual(t, 8317, day9(10, 1618))
	TestEqual(t, 146373, day9(13, 7999))
	TestEqual(t, 2764, day9(17, 1104))
	TestEqual(t, 54718, day9(21, 6111))
	TestEqual(t, 37305, day9(30, 5807))

	file := Input(9)
	var players, lastMarble int
	Must(fmt.Sscanf(file, "%d players; last marble is worth %d points", &players, &lastMarble))
	TestEqual(t, 418237, day9(players, lastMarble))
	TestEqual(t, 3505711612, day9(players, lastMarble*100))
}
