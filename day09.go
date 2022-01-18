package main

import "container/ring"

func day9(players, lastMarble int) int {
	cur := ring.New(1)
	cur.Value = 0
	scores := make([]int, players)
	player := -1
	for marble := 1; marble <= lastMarble; marble++ {
		player = (player + 1) % players
		if marble%23 == 0 {
			scores[player] += marble
			cur = cur.Move(-8)
			rem := cur.Unlink(1)
			scores[player] += rem.Value.(int)
			cur = cur.Next()
			continue
		}
		cur = cur.Next()
		nr := ring.New(1)
		nr.Value = marble
		cur.Link(nr)
		cur = nr
	}
	return Max(scores...)
}
