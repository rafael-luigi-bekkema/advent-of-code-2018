package main

import (
	"fmt"
	"sort"
)

type Result[T, U any] struct {
	a T
	b U
}

func (r Result[T, U]) String() string {
	return fmt.Sprintf("a=%v/b=%v", r.a, r.b)
}

func NewResult[T, U any](a T, b U) Result[T, U] {
	return Result[T, U]{a, b}
}

func day4a(input []string) Result[int, int] {
	sort.Strings(input)
	type Guard struct {
		id    int
		total int
		min   map[int]int
	}
	guards := map[int]*Guard{}
	var maxTotal, lastFall, maxMinuteCount, maxMinute int
	var guard, maxGuard, maxMinuteGuard *Guard
	for _, line := range input {
		var year, month, day, hour, minute int
		Must(fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute))
		rest := line[19:]
		switch rest[:5] {
		case "Guard":
			var id int
			Must(fmt.Sscanf(rest, "Guard #%d begins shift", &id))
			guard = guards[id]
			if guard == nil {
				guard = &Guard{id: id, min: map[int]int{}}
				guards[id] = guard
			}
		case "falls":
			lastFall = minute
		case "wakes":
			for m := minute - 1; m >= lastFall; m-- {
				guard.total++
				guard.min[m]++
				if guard.min[m] > maxMinuteCount {
					maxMinuteCount = guard.min[m]
					maxMinute = m
					maxMinuteGuard = guard
				}
			}
			if guard.total > maxTotal {
				maxTotal = guard.total
				maxGuard = guard
			}
		default:
			panic("unexpected value: " + rest[:5])
		}
	}
	var maxCount, maxMin int
	for m, count := range maxGuard.min {
		if count > maxCount {
			maxCount = count
			maxMin = m
		}
	}
	return NewResult(maxGuard.id*maxMin, maxMinuteGuard.id*maxMinute)
}
