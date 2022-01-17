package main

import "fmt"

func day6a(input []string) int {
	coords := make([][2]int, len(input))
	for i, line := range input {
		Must(fmt.Sscanf(line, "%d, %d", &coords[i][0], &coords[i][1]))
	}
	dist := func(c1, c2 [2]int) int {
		return Abs(c1[0]-c2[0]) + Abs(c1[1]-c2[1])
	}
	var minx, maxx, miny, maxy int
	cmap := map[[2]int]byte{}
	for i, c1 := range coords {
		if i == 0 || c1[0] < minx {
			minx = c1[0]
		}
		if i == 0 || c1[1] < miny {
			miny = c1[1]
		}
		if i == 0 || c1[0] > maxx {
			maxx = c1[0]
		}
		if i == 0 || c1[1] > maxy {
			maxy = c1[1]
		}
		cmap[c1] = 'a' + byte(i)
	}
	closest := map[[2]int]int{}
	inf := map[[2]int]bool{}
	printMap := false
	if printMap {
		fmt.Println()
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			c2 := [2]int{x, y}
			if c, ok := cmap[c2]; ok {
				_ = c
				// fmt.Printf("%c", c-('a'-'A'))
				if printMap {
					fmt.Print("X")
				}
				closest[c2]++
				continue
			}
			var mindist int
			var minc [2]int
			var dupe bool
			for _, c := range coords {
				d := dist(c2, c)
				if mindist > 0 && mindist == d {
					dupe = true
				}
				if mindist == 0 || d < mindist {
					mindist = d
					minc = c
					dupe = false
				}
			}
			if mindist > 0 && !dupe {
				if printMap {
					fmt.Print("x")
				}
				// fmt.Printf("%c", cmap[minc])
				if x == minx || x == maxx || y == miny || y == maxy {
					inf[minc] = true
				}
				closest[minc]++
			} else {
				if printMap {
					fmt.Print(".")
				}
			}
		}
		if printMap {
			fmt.Println()
		}
	}
	var maxcount int
	var maxc [2]int
	for c, count := range closest {
		if inf[c] {
			continue
		}
		if count > maxcount {
			maxcount = count
			maxc = c
		}
	}
	_ = maxc
	return maxcount
}

func day6b(input []string, maxdist int) int {
	coords := make([][2]int, len(input))
	for i, line := range input {
		Must(fmt.Sscanf(line, "%d, %d", &coords[i][0], &coords[i][1]))
	}
	dist := func(c1, c2 [2]int) int {
		return Abs(c1[0]-c2[0]) + Abs(c1[1]-c2[1])
	}
	var minx, maxx, miny, maxy int
	for i, c1 := range coords {
		if i == 0 || c1[0] < minx {
			minx = c1[0]
		}
		if i == 0 || c1[1] < miny {
			miny = c1[1]
		}
		if i == 0 || c1[0] > maxx {
			maxx = c1[0]
		}
		if i == 0 || c1[1] > maxy {
			maxy = c1[1]
		}
	}
	var count int
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			var sum int
			c1 := [2]int{x, y}
			for _, c2 := range coords {
				d := dist(c1, c2)
				sum += d
			}
			if sum < maxdist {
				count++
			}
		}
		fmt.Println()
	}
	return count
}
