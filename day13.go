package main

import (
	"fmt"
	"sort"
	"strings"
)

func day13a(input string) Result[string, string] {
	// left := Point{-1,0}
	// right := Point{1,0}
	// up := Point{0,-1}
	// down := Point{0,1}
	type Direction int
	const (
		up Direction = iota
		right
		down
		left
	)

	type Train struct {
		Point
		nr   int
		turn int
		dir  Direction
	}
	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))
	var trains []*Train
	for i, line := range lines {
		grid[i] = []byte(line)
		for x, c := range grid[i] {
			var dir Direction
			switch c {
			case '^':
				dir = up
				grid[i][x] = '|'
			case '>':
				dir = right
				grid[i][x] = '-'
			case 'v':
				dir = down
				grid[i][x] = '|'
			case '<':
				dir = left
				grid[i][x] = '-'
			default:
				continue
			}
			trains = append(trains, &Train{Point: Point{x, i}, nr: len(trains) + 1, turn: 2, dir: dir})
		}
	}
	render := func() {
		fmt.Println()
		for y := range grid {
		x:
			for x, c := range grid[y] {
				for _, train := range trains {
					if train.x == x && train.y == y {
						switch train.dir {
						case up:
							fmt.Print("^")
						case right:
							fmt.Print(">")
						case down:
							fmt.Print("v")
						case left:
							fmt.Print("<")
						}
						continue x
					}
				}
				fmt.Printf("%c", c)
			}
			fmt.Println()
		}
	}
	trainCount := len(trains)
	turn := -1
	firstCrash, lastRemain := "", ""
outer:
	for {
		turn++
		trains = Remove(trains, nil)
		if trainCount == 1 {
			lastRemain = fmt.Sprintf("%d,%d", trains[0].x, trains[0].y)
			break outer
		}
		sort.Slice(trains, func(i, j int) bool {
			if trains[i].y == trains[j].y {
				return trains[i].x < trains[j].x
			}
			return trains[i].y < trains[j].y
		})
		for tid := 0; tid < len(trains); tid++ {
			train := trains[tid]
			if train == nil {
				continue
			}
			c := grid[train.y][train.x]
			switch c {
			case '+':
				train.turn = (train.turn + 1) % 3
				switch train.turn {
				case 0: // left
					train.dir = (train.dir + 3) % 4
				case 1: // straight
				case 2: // right
					train.dir = (train.dir + 1) % 4
				}
				// turn
			case '\\':
				switch train.dir {
				case up: // up
					train.dir = left
				case right: // right
					train.dir = down
				case down: // down
					train.dir = right
				case left: // left
					train.dir = up
				default:
					panic("oops")
				}
			case '/':
				switch train.dir {
				case up: // up
					train.dir = right
				case right: // right
					train.dir = up
				case down: // down
					train.dir = left
				case left: // left
					train.dir = down
				default:
					panic("oops")
				}
			case '-', '|': // straight ahead
			default:
				panic("expected train symbol")
			}

			switch train.dir {
			case up: // up
				train.y--
			case right: // right
				train.x++
			case down: // down
				train.y++
			case left: // left
				train.x--
			}
			for tid2, train2 := range trains {
				if train2 == nil || tid == tid2 {
					continue
				}
				if train.x == train2.x && train.y == train2.y {
					if firstCrash == "" {
						firstCrash = fmt.Sprintf("%d,%d", train.x, train.y)
					}
					// fmt.Printf("Crash! Step: %d Trains left: %d\n", turn, len(trains))
					trains[tid] = nil
					trains[tid2] = nil
					trainCount -= 2
					if trainCount == 0 {
						break outer
					}
					break
				}
			}
		}
		_ = render
	}
	return NewResult(firstCrash, lastRemain)
}
