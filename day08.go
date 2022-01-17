package main

import (
	"strings"
)

func day8a(input string) Result[int, int] {
	type Node struct {
		meta     []int
		children []*Node
	}
	ints := Map(Atoi, strings.Split(input, " "))
	var value func(n *Node) int 
	value = func(n *Node) (sum int) {
		if len(n.children) == 0 {
			return Sum(n.meta...)
		}
		for _, idx := range n.meta {
			if idx < 1 || idx > len(n.children) {
				continue
			}
			sum += value(n.children[idx-1])
		}
		return sum
	}

	var totalMeta int
	var eat func(n, depth int) (res []*Node)
	eat = func(n, depth int) (res []*Node) {
		for i := 0; i < n; i++ {
			var n Node
			nrChild := ints[0]
			nrMeta := ints[1]
			ints = ints[2:]
			n.children = eat(nrChild, depth+1)
			n.meta = ints[:nrMeta]
			totalMeta += Sum(n.meta...)
			ints = ints[nrMeta:]
			res = append(res, &n)
		}
		return
	}
	root := eat(1, 0)
	return NewResult(totalMeta, value(root[0]))
}
