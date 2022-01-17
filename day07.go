package main

import (
	"fmt"
)

func day7parse(input []string) map[byte][]byte {
	deps := map[byte][]byte{}
	for _, line := range input {
		var s [2]byte
		Must(fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &s[0], &s[1]))
		deps[s[1]] = append(deps[s[1]], s[0])
		if _, ok := deps[s[0]]; !ok {
			deps[s[0]] = nil
		}
	}
	return deps
}

func day7next(deps map[byte][]byte, done map[byte]bool) byte {
	var dones []byte
deps:
	for item, ideps := range deps {
		for _, dep := range ideps {
			if !done[dep] {
				continue deps
			}
		}
		dones = append(dones, item)
	}
	if len(dones) == 0 {
		return 0
	}
	item := Min(dones...)
	delete(deps, item)
	return item
}

func day7a(input []string) string {
	deps := day7parse(input)
	done := map[byte]bool{}
	result := make([]byte, 0, len(deps))
	for len(deps) > 0 {
		item := day7next(deps, done)
		done[item] = true
		result = append(result, item)
	}
	return string(result)
}

func day7b(input []string, sec, elfs int) int {
	s := -1
	deps := day7parse(input)

	type Work struct {
		what   byte
		doneAt int
	}
	done := map[byte]bool{}
	workers := map[int]*Work{}
	result := make([]byte, 0, len(deps))
	for len(deps) > 0 || len(workers) > 0 {
		s++
		for elf, work := range workers {
			if work.doneAt == s {
				done[work.what] = true
				result = append(result, work.what)
				delete(workers, elf)
			}
		}
		for elf := 0; elf < elfs; elf++ {
			work := workers[elf]
			if work == nil {
				item := day7next(deps, done)
				if item == 0 {
					continue
				}
				d := s + int(item-'A') + sec + 1
				workers[elf] = &Work{what: item, doneAt: d}
			}
		}
	}
	return s
}
