package main

import (
	"strings"
)

func day12a(n int, input []string) int {
	right := Map([]byte(input[0][15:]), func(val byte) bool {
		return val == '#'
	})
	patterns := map[[5]bool]bool{}
	for _, line := range input[2:] {
		from, to, ok := strings.Cut(line, " => ")
		if !ok {
			panic("oops")
		}
		var val [5]bool
		for i, c := range []byte(from) {
			val[i] = c == '#'
		}
		patterns[val] = to == "#"
	}
	var left []bool
	get := func(i int) bool {
		if i < 0 {
			i := (i * -1) - 1
			if i >= len(left) {
				return false
			}
			return left[i]
		}
		if i >= len(right) {
			return false
		}
		return right[i]
	}
	render := func(gen int) (string, int) {
		var out []byte
		var firstOn bool
		firstI := -1
		for i, c := range right {
			if c {
				if !firstOn {
					firstI = i
				}
				firstOn = true
				out = append(out, '#')
			} else if firstOn {
				out = append(out, '.')
			}
		}
		return string(out), firstI
	}
	summ := func(left, right []bool, idiff int) int {
		var sum int
		for i, v := range right {
			if v {
				sum += i + idiff
			}
		}
		for i, v := range left {
			if v {
				sum += (i + 1) * -1
			}
		}
		return sum
	}
	hist := map[string]bool{}
	var repeat string
	var repeatIdiff int
	for gen := 1; gen <= n; gen++ {
		nright := make([]bool, len(right)+2)
		for i := 0; i < len(right)+2; i++ {
			var val [5]bool
			for j := 0; j < 5; j++ {
				val[j] = get(i + j - 2)
			}
			nright[i] = patterns[val]
		}
		nleft := make([]bool, len(left)+2)
		for i := 0; i < len(left)+2; i++ {
			var val [5]bool
			for j := 0; j < 5; j++ {
				val[j] = get((i+1)*-1 + j - 2)
			}
			nleft[i] = patterns[val]
		}
		left, right = nleft, nright
		for len(left) > 0 && !left[len(left)-1] {
			left = left[:len(left)-1]
		}
		for len(right) > 0 && !right[len(right)-1] {
			right = right[:len(right)-1]
		}
		g, firstI := render(gen)
		if hist[g] && len(left) == 0 {
			repeat = g
			repeatIdiff = gen - firstI
			break
		}
		hist[g] = true
	}
	if repeat != "" {
		return summ(nil, Map([]byte(repeat), func(c byte) bool {
			return c == '#'
		}), n-repeatIdiff)
	}
	return summ(left, right, 0)
}
