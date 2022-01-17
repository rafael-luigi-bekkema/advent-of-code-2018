package main

import "fmt"

type Claim struct {
	id                       int
	top, left, width, height int
	right, bottom            int
}

func (c Claim) String() string {
	return fmt.Sprintf("#%d @ l:%d,t:%d: w:%dxh:%d", c.id, c.left, c.top, c.width, c.height)
}

type day3result struct {
	overlap, soloID int
}

func day3a(lines []string) day3result {
	claims := make([]Claim, len(lines))
	for i, line := range lines {
		c := &claims[i]
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.width, &c.height)
		c.right = c.left + c.width
		c.bottom = c.top + c.height
	}
	var total int
	grid := map[[2]int]int{}
	for _, c1 := range claims {
		for x := c1.left; x < c1.left+c1.width; x++ {
			for y := c1.top; y < c1.top+c1.height; y++ {
				grid[[2]int{x, y}]++
				if grid[[2]int{x, y}] == 2 {
					total++
				}
			}
		}
	}

	var solo int
claim:
	for _, c1 := range claims {
		for x := c1.left; x < c1.left+c1.width; x++ {
			for y := c1.top; y < c1.top+c1.height; y++ {
				if grid[[2]int{x, y}] != 1 {
					continue claim
				}
			}
		}
		solo = c1.id
	}
	return day3result{overlap: total, soloID: solo}
}
