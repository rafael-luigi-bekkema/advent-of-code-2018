package main

import (
	"fmt"
	"io"
	"os"
)

type Grid struct {
	Width, Height    int
	Data             []byte
	OffsetX, OffsetY int
	Changed          bool
}

func (grid *Grid) set(x, y int, value byte) {
	idx := grid.translate(x, y)
	if grid.Data[idx] != value {
		grid.Data[idx] = value
		grid.Changed = true
	}
}

func (grid *Grid) get(x, y int) byte {
	idx := grid.translate(x, y)
	if idx >= len(grid.Data) {
		return offgrid
	}
	return grid.Data[idx]
}

func (grid *Grid) translate(x, y int) int {
	return (y-grid.OffsetY+1)*grid.Width + (x - grid.OffsetX + 1)
}

func (grid *Grid) save() {
	f, err := os.Create("day17_out.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid.print(f)
}

func (grid *Grid) print(out io.Writer) {
	for i, c := range grid.Data {
		if i%grid.Width == 0 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "%c", c)
	}

	fmt.Fprintln(out)
}

func (grid *Grid) Score() int {
	var score int
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if y == 0 || y == grid.Height-1 {
				continue
			}
			if val := grid.Data[y*grid.Width+x]; val == water || val == wet {
				score += 1
			}
		}
	}
	return score
}

var source = Point{x: 500, y: 0}

const sand byte = ' '
const clay byte = '#'
const other byte = '+'
const water byte = '~'
const wet byte = '|'
const offgrid byte = 'o'

func buildGrid(input []string) Grid {
	type Point struct {
		x, y int
	}

	var points []Point

	var minx, maxx, miny, maxy = source.x, source.x, source.y, source.y

	for _, line := range input {
		var c1, c2 rune
		var n1, n2, n3 int
		fmt.Sscanf(line, "%c=%d, %c=%d..%d", &c1, &n1, &c2, &n2, &n3)

		for i := n2; i <= n3; i++ {
			var x, y int
			if c1 == 'x' {
				x, y = n1, i
			} else {
				x, y = i, n1
			}
			points = append(points, Point{x: x, y: y})
			if x < minx {
				minx = x
			}
			if y < miny {
				miny = y
			}
			if x >= maxx {
				maxx = x
			}
			if y >= maxy {
				maxy = y
			}
		}
	}

	height := maxy - miny + 3
	width := maxx - minx + 3

	grid := Grid{
		Width:   width,
		Height:  height,
		Data:    make([]byte, width*height),
		OffsetX: minx,
		OffsetY: miny,
	}

	for i := range grid.Data {
		grid.Data[i] = sand
	}
	for _, point := range points {
		grid.set(point.x, point.y, clay)
	}
	grid.set(source.x, source.y, other)
	return grid
}

func day17a(input []string) int {
	grid := buildGrid(input)
	//grid.print()

	t := 0

	for {
		grid.Changed = false
		drop := source
		var wentLeft, wentRight bool
	middle:
		for {
			var moved bool
		inner:
			for {
				switch grid.get(drop.x, drop.y+1) {
				case sand, wet:
				case offgrid:
					break middle
				default:
					break inner
				}

				drop.y++
				wentLeft, wentRight = false, false
				moved = true
				grid.set(drop.x, drop.y, wet)
			}
			if !wentRight {
				if val := grid.get(drop.x-1, drop.y); val == sand || val == wet {
					wentLeft = true
					drop.x--
					moved = true
					grid.set(drop.x, drop.y, wet)
					continue
				}
			}
			if !wentLeft {
				if val := grid.get(drop.x+1, drop.y); val == sand || val == wet {
					wentRight = true
					drop.x++
					moved = true
					grid.set(drop.x, drop.y, wet)
				}
			}
			if !moved {
				grid.set(drop.x, drop.y, water)
				break
			}
		}

		// grid.print()
		fmt.Printf("t = %d\n", t)

		t++
		// time.Sleep(time.Millisecond * 100)
		if t > 100_000 || !grid.Changed {
			break
		}
	}

	grid.save()
	fmt.Println("Done")

	return grid.Score()
}
