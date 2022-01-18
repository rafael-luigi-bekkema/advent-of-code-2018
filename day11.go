package main

import "fmt"

func day11pow(x, y, serialNo int) int {
	rackID := x + 10
	pow := (rackID*y + serialNo) * rackID
	pow = (pow/100)%10 - 5
	return pow
}

func day11a(serialNo int) string {
	const size = 300
	var maxsum int
	var maxpoint Point
	first := true
	for y := 1; y <= size-2; y++ {
		for x := 1; x <= size-2; x++ {
			sum := 0
			for dy := 0; dy < 3; dy++ {
				for dx := 0; dx < 3; dx++ {
					sum += day11pow(x+dx, y+dy, serialNo)
				}
			}
			if first || sum > maxsum {
				maxsum = sum
				maxpoint = Point{x, y}
				first = false
			}
		}
	}
	return fmt.Sprintf("%d,%d", maxpoint.x, maxpoint.y)
}

func day11b(serialNo int) string {
	const size = 300
	var points [size][size]int
	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			points[y-1][x-1] = day11pow(x, y, serialNo)
		}
	}
	var prev [size][size]int
	var maxpoint Point
	var maxsum, maxsize int
	first := true
	for s := 0; s < size; s++ {
		for y := 1; y <= size-s; y++ {
			for x := 1; x <= size-s; x++ {
				sum := prev[y-1][x-1]
				for dy := 0; dy <= s; dy++ {
					sum += points[y+dy-1][x+s-1]
				}
				for dx := 0; dx < s; dx++ {
					sum += points[y+s-1][x+dx-1]
				}
				prev[y-1][x-1] = sum
				if first || sum > maxsum {
					maxsum = sum
					maxpoint = Point{x, y}
					maxsize = s + 1
					first = false
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", maxpoint.x, maxpoint.y, maxsize)
}
