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
	MinY             int
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
			fmt.Fprintf(out, "\n%04d: ", i/grid.Width-1)
		}
		fmt.Fprintf(out, "%c", c)
	}

	fmt.Fprintln(out)
}

func (grid *Grid) removeWet() {
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			val := grid.Data[y*grid.Width+x]
			if val != wet {
				continue
			}
			grid.Data[y*grid.Width+x] = sand
		lbl:
			for sx := x-1; sx >= 0; sx-- {
				switch grid.Data[y*grid.Width+sx] {
				default:
					break lbl
				case water:
					grid.Data[y*grid.Width+sx] = sand
				}
			}

		lbl2:
			for sx := x+1; sx < len(grid.Data); sx++ {
				switch grid.Data[y*grid.Width+sx] {
				default:
					break lbl2
				case water:
					grid.Data[y*grid.Width+sx] = sand
				}
			}
		}
	}
}

func (grid *Grid) Score() int {
	var score int
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if y <= grid.MinY || y == grid.Height-1 {
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

	var minx, maxx, miny, maxy = 0, 0, 0, 0

	for idx, line := range input {
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
			if idx == 0 || x < minx {
				minx = x
			}
			if idx == 0 || y < miny {
				miny = y
			}
			if idx == 0 || x >= maxx {
				maxx = x
			}
			if idx == 0 || y >= maxy {
				maxy = y
			}
		}
	}

	realminy := miny

	if source.x < minx {
		minx = source.x
	}
	if source.y < miny {
		miny = source.y
	}

	height := maxy - miny + 3
	width := maxx - minx + 3

	grid := Grid{
		Width:   width,
		Height:  height,
		Data:    make([]byte, width*height),
		OffsetX: minx,
		OffsetY: miny,
		MinY:    realminy,
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

func step(grid *Grid, point Point) {
	grid.set(point.x, point.y, wet)
	// time.Sleep(time.Millisecond * 200)
	// grid.print(os.Stdout)

	switch grid.get(point.x, point.y+1) {
	case sand:
		step(grid, Point{point.x, point.y + 1})
	case offgrid:
		return
	}

lbl:
	for x := point.x; x >= grid.OffsetX; x-- {
		switch grid.get(x, point.y+1) {
		case clay, sand:
			break lbl
		case water:
		default:
			return
		}
	}

lbl2:
	for x := point.x; x < grid.OffsetX+grid.Width; x++ {
		switch grid.get(x, point.y+1) {
		case clay, sand:
			break lbl2
		case water:
		default:
			return
		}
	}

	switch grid.get(point.x-1, point.y) {
	case sand:
		step(grid, Point{point.x - 1, point.y})
		if grid.get(point.x-1, point.y) == water {
			grid.set(point.x, point.y, water)
		}
	case clay:
		grid.set(point.x, point.y, water)
	}

	switch grid.get(point.x+1, point.y) {
	case sand:
		step(grid, Point{point.x + 1, point.y})
		if grid.get(point.x+1, point.y) == water {
			grid.set(point.x, point.y, water)
		}
	case clay:
		grid.set(point.x, point.y, water)
	}
}

func day17a(input []string, save bool) int {
	grid := buildGrid(input)

	step(&grid, Point{source.x, source.y + 1})

	if save {
		grid.save()
	} else {
		grid.print(os.Stdout)
	}

	return grid.Score()
}

func day17b(input []string, save bool) int {
	grid := buildGrid(input)

	step(&grid, Point{source.x, source.y + 1})

	grid.removeWet()

	if save {
		grid.save()
	} else {
		grid.print(os.Stdout)
	}

	return grid.Score()
}
