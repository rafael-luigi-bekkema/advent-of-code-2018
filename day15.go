package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2018/dijkstra"
)

type UnitType uint8
type Unit struct {
	num    int
	typ    UnitType
	x, y   int
	hp, ap int
}

func (u *Unit) String() string {
	return fmt.Sprintf("unit %02d (%d,%d) hp:%d", u.num, u.x, u.y, u.hp)
}

func day15a(input string) int {
	res, _ := day15(input, 3, false)
	return res
}

func day15b(input string) int {
	ap := 3
	for {
		ap++
		res, ok := day15(input, ap, true)
		if ok {
			return res
		}
	}
}

func day15(input string, elfAP int, stopOnElfDeath bool) (int, bool) {
	// logfile := Must(os.CreateTemp("", "day15"))
	// defer logfile.Close()
	log.SetFlags(0)
	// log.SetOutput(logfile)
	log.SetOutput(io.Discard)

	const (
		elf    UnitType = 'E'
		goblin          = 'G'
	)
	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))
	var units []*Unit
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for j, c := range []byte(line) {
			switch c {
			case 'G':
				units = append(units, &Unit{num: len(units), typ: goblin, x: j, y: i, hp: 200, ap: 3})
			case 'E':
				units = append(units, &Unit{num: len(units), typ: elf, x: j, y: i, hp: 200, ap: elfAP})
			}
			grid[i][j] = c
		}
	}

	render := func() {
		fmt.Println()
		for y := range grid {
			for x := range grid[y] {
				fmt.Printf("%c", grid[y][x])
			}
			fmt.Println()
		}
	}
	buildGraph := func(source Point, enemy UnitType) (dijkstra.Graph[Point], []Point) {
		var graph dijkstra.Graph[Point]
		var targets []Point
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == byte(enemy) {
					for _, off := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
						dx, dy := x+off[0], y+off[1]
						if dx < 0 || dx >= len(grid[0]) || dy < 0 || dy >= len(grid) || grid[dy][dx] != '.' {
							continue
						}
						targets = append(targets, Point{dx, dy})
					}
					continue
				}
				if grid[y][x] != '.' && !(x == source.x && y == source.y) {
					continue
				}
				node := Point{x, y}
				graph.AddNode(node)
				for _, off := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
					dx, dy := x+off[0], y+off[1]
					if dx < 0 || dx >= len(grid[0]) || dy < 0 || dy >= len(grid) || grid[dy][dx] != '.' {
						continue
					}
					graph.AddEdge(node, Point{dx, dy}, 1)
				}
			}
		}
		return graph, targets
	}

	findTarget := func(unit *Unit, enemy UnitType) (*Unit, bool) {
		var adjacents []*Unit
		var hasEnemies bool
		for _, e := range units {
			if e.typ != enemy {
				continue
			}
			hasEnemies = true
			for _, off := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				dx, dy := unit.x+off[0], unit.y+off[1]
				if dx < 0 || dx >= len(grid[0]) || dy < 0 || dy >= len(grid) || grid[dy][dx] == '#' {
					continue
				}
				if e.x == dx && e.y == dy {
					adjacents = append(adjacents, e)
					break
				}
			}
		}
		if len(adjacents) == 0 {
			return nil, hasEnemies
		}
		sort.Slice(adjacents, func(i, j int) bool {
			if adjacents[i].hp == adjacents[j].hp {
				if adjacents[i].y == adjacents[j].y {
					return adjacents[i].x < adjacents[j].x
				}
				return adjacents[i].y < adjacents[j].y
			}
			return adjacents[i].hp < adjacents[j].hp
		})
		// if adjacents[0].num == 24 || adjacents[0].num == 28 {
		// 	fmt.Println(adjacents)
		// }
		// fmt.Println(adjacents)
		return adjacents[0], hasEnemies
	}

	round := 0
outer:
	for {
		sort.Slice(units, func(i, j int) bool {
			if units[i].y == units[j].y {
				return units[i].x < units[j].x
			}
			return units[i].y < units[j].y
		})
		for i := 0; i < len(units); i++ {
			unit := units[i]
			enemy := elf
			if unit.typ == elf {
				enemy = goblin
			}

			target, hasEnemies := findTarget(unit, enemy)
			if !hasEnemies {
				// fmt.Println("No more enemies! Round:", round)
				break outer
			}
			if target == nil {
				source := Point{unit.x, unit.y}
				graph, targets := buildGraph(source, enemy)
				dist, prev := dijkstra.Run(graph, source)
				type tardist struct {
					dist   float64
					target Point
				}
				var td []tardist
				for _, target := range targets {
					if dist[target] != math.Inf(1) {
						td = append(td, tardist{dist[target], target})
					}
				}
				if len(td) == 0 {
					continue
				}
				sort.Slice(td, func(i, j int) bool {
					if td[i].dist == td[j].dist {
						if td[i].target.y == td[j].target.y {
							return td[i].target.x < td[j].target.x
						}
						return td[i].target.y < td[j].target.y
					}
					return td[i].dist < td[j].dist
				})

				p := td[0].target
				for prev[p] != source {
					p = prev[p]
				}
				if prev[p] != source {
					panic("did not find move")
				}

				log.Printf("Move %d to %d,%d", unit.num, p.x, p.y)
				grid[unit.y][unit.x] = '.'
				unit.x, unit.y = p.x, p.y
				grid[unit.y][unit.x] = byte(unit.typ)
				target, _ = findTarget(unit, enemy)
				if target == nil {
					continue
				}
			}

			target.hp -= unit.ap
			log.Printf("Attack %d -> %d (hp: %d)", unit.num, target.num, target.hp)
			if target.hp <= 0 {
				// fmt.Printf("Unit %02d (%c) died on round %d (left %d)\n", target.num, target.typ, round, len(units))
				if stopOnElfDeath && target.typ == elf {
					return 0, false
				}
				grid[target.y][target.x] = '.'
				units = Remove(units, target)
				for i < len(units) && units[i] != unit {
					i--
				}
			}
		}
		round++
	}
	_ = render

	hp := 0
	for _, unit := range units {
		hp += unit.hp
	}

	return round * hp, true
}
