package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

func day10a(input []string) Result[string, int] {
	type Star struct {
		pos Point
		vel Point
	}
	stars := make([]Star, len(input))
	for i, line := range input {
		s := &stars[i]
		Must(fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>",
			&s.pos.x, &s.pos.y, &s.vel.x, &s.vel.y))
	}
	minmax := func() (minx, maxx, miny, maxy int) {
		minx, maxx = stars[0].pos.x, stars[0].pos.x
		miny, maxy = stars[0].pos.y, stars[0].pos.y
		for _, star := range stars[1:] {
			if star.pos.x < minx {
				minx = star.pos.x
			}
			if star.pos.x > maxx {
				maxx = star.pos.x
			}
			if star.pos.y < miny {
				miny = star.pos.y
			}
			if star.pos.y > maxy {
				maxy = star.pos.y
			}
		}
		return minx, maxx, miny, maxy
	}
	var minsize int
	render := func() string {
		minx, maxx, miny, maxy := minmax()
		size := (maxy - miny) * (maxx - minx)
		if minsize == 0 || size < minsize {
			minsize = size
		}
		if minsize == 0 || size <= minsize {
			return ""
		}
		// This happens when size has increased for the first time
		// Rewind 1 step and render output
		for i := range stars {
			stars[i].pos.x -= stars[i].vel.x
			stars[i].pos.y -= stars[i].vel.y
		}
		var s strings.Builder
		smap := make(map[Point]bool, len(stars))
		for _, star := range stars {
			smap[star.pos] = true
		}
		minx, maxx, miny, maxy = minmax()
		for y := miny; y <= maxy; y++ {
			for x := minx; x <= maxx; x++ {
				if smap[Point{x, y}] {
					fmt.Fprint(&s, "#")
					continue
				}
				fmt.Fprint(&s, ".")
			}
			if y != maxy {
				fmt.Fprintln(&s)
			}
		}
		return s.String()
	}
	var out string
	var sec int
	for s := 1; s <= 1000_000; s++ {
		for i := range stars {
			stars[i].pos.x += stars[i].vel.x
			stars[i].pos.y += stars[i].vel.y
		}
		out = render()
		if out != "" {
			sec = s - 1
			break
		}
	}
	return NewResult(out, sec)
}
