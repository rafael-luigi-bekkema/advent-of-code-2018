package main

import "fmt"

func day7a(input []string) string {
	// "Step F must be finished before step E can begin.",
	items := make([][2]byte, len(input))
	for i, line := range input {
		s := &items[i]
		Must(fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &s[0], &s[1]))

	}
	fmt.Println(items)
	return ""
}
