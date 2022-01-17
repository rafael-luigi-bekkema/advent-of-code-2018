package main

const capdiff = 'a' - 'A'

func day5reduce(polymer []byte) []byte {
	match := func(a, b byte) bool {
		return a == b-capdiff || a-capdiff == b
	}
	for i := 0; i < len(polymer)-1; i++ {
		if match(polymer[i], polymer[i+1]) {
			copy(polymer[i:], polymer[i+2:])
			polymer = polymer[:len(polymer)-2]
			i -= 2
			if i < -1 {
				i = -1
			}
		}
	}
	return polymer
}

func day5a(input string) int {
	polymer := []byte(input)
	polymer = day5reduce(polymer)
	return len(polymer)
}

func day5b(input string) int {
	polymer := []byte(input)
	remain := map[byte]bool{}
	for _, c := range polymer {
		if 'a' <= c && c <= 'z' {
			remain[c] = true
		} else {
			remain[c+capdiff] = true
		}
	}
	var minlen int
	for c := range remain {
		polymer := Remove(polymer, c, c-capdiff)
		polymer = day5reduce(polymer)
		if minlen == 0 || len(polymer) < minlen {
			minlen = len(polymer)
		}
	}
	return minlen
}
