package main

func day2a(input []string) int {
	var alltwos, allthrees int
	for _, line := range input {
		counts := map[byte]int{}
		for _, c := range []byte(line) {
			counts[c]++
		}

		var twos, threes bool
		for _, count := range counts {
			if count == 2 {
				twos = true
			}
			if count == 3 {
				threes = true
			}
		}
		if twos {
			alltwos++
		}
		if threes {
			allthrees++
		}
	}
	return alltwos * allthrees
}

func day2b(input []string) string {
	for i, w1 := range input {
		for _, w2 := range input[i+1:] {
			diff := 0
			diffpos := -1
			for j, ch := range []byte(w1) {
				if w2[j] != ch {
					diff++
					diffpos = j
				}
			}
			if diff == 1 {
				return w1[:diffpos]+w2[diffpos+1:]
			}
		}
	}
	return ""
}
