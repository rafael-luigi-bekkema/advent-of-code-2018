package main

import "testing"

func TestDay7(t *testing.T) {
	example := []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}
	TestEqual(t, "CABDFE", day7a(example))
	TestEqual(t, 15, day7b(example, 0, 2))
	file := Lines(7)
	TestEqual(t, "HEGMPOAWBFCDITVXYZRKUQNSLJ", day7a(file))
	TestEqual(t, 1226, day7b(file, 60, 5))
}
