package dijkstra

import (
	"math"
)

type Graph[T comparable] struct {
	nodes []T
	edges map[T]map[T]float64
}

func (g *Graph[T]) AddNode(n T) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph[T]) AddEdge(n, e T, v float64) {
	if g.edges == nil {
		g.edges = make(map[T]map[T]float64)
	}
	if g.edges[n] == nil {
		g.edges[n] = make(map[T]float64)
	}
	g.edges[n][e] = v
}

var inf = math.Inf(0)

func in[T comparable](items []T, subject T) bool {
	for _, item := range items {
		if item == subject {
			return true
		}
	}
	return false
}

func Run[T comparable](graph Graph[T], source T) (dist map[T]float64, prev map[T]T) {
	Q := make([]T, 0, len(graph.nodes))
	dist = map[T]float64{}
	prev = map[T]T{}
	for _, node := range graph.nodes {
		dist[node] = inf
		Q = append(Q, node)
	}
	dist[source] = 0

	for len(Q) > 0 {
		var minI int
		for i, u := range Q {
			if i == 0 || dist[u] < dist[Q[minI]] {
				minI = i
			}
		}
		u := Q[minI]
		Q[minI], Q[len(Q)-1] = Q[len(Q)-1], Q[minI]
		Q = Q[:len(Q)-1]

		for v, l := range graph.edges[u] {
			if !in(Q, v) {
				continue
			}
			alt := dist[u] + l
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return dist, prev
}
