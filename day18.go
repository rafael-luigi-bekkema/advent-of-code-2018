package main

import (
	"fmt"
	"strings"
)

const (
	openGround = '.'
	trees      = '|'
	lumberYard = '#'
)

func day18a(input string, minutes int) int {
	width := strings.Index(input, "\n")
	initial := []byte(strings.ReplaceAll(input, "\n", ""))

	iters := map[string]int{}
	dupeFound := false

	for n := 0; n < minutes; n++ {
		if !dupeFound {
			if n2, ok := iters[string(initial)]; ok {
				every := n - n2
				fmt.Println("Found dupe at:", n, n2, every)
				n += (minutes-n)/every*every - 1
				dupeFound = true
				continue
			}
		}

		iters[string(initial)] = n
		current := make([]byte, len(initial))
		copy(current, initial)

		for i, acre := range initial {
			x := i % width
			y := i / width

			counts := make(map[byte]int, 3)

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					x := x + dx
					y := y + dy
					if x < 0 || x >= width || y < 0 || y >= width {
						continue
					}

					counts[initial[y*width+x]]++
				}
			}

			switch {
			case acre == openGround && counts[trees] >= 3:
				current[i] = trees
			case acre == trees && counts[lumberYard] >= 3:
				current[i] = lumberYard
			case acre == lumberYard && (counts[lumberYard] == 0 || counts[trees] == 0):
				current[i] = openGround
			}
		}

		initial = current
	}

	counts := make(map[byte]int, 3)
	for _, acre := range initial {
		counts[acre]++
	}

	return counts[trees] * counts[lumberYard]
}
