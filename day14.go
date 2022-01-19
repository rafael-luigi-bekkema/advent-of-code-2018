package main

func day14a(after int) [10]int {
	recipes := []int{3, 7}
	elfs := [2]int{0, 1}
	for len(recipes) < after+10 {
		combo := Sum(recipes[elfs[0]], recipes[elfs[1]])
		var adds []int
		for {
			adds = append(adds, combo%10)
			combo /= 10
			if combo == 0 {
				break
			}
		}
		for i := len(adds) - 1; i >= 0; i-- {
			recipes = append(recipes, adds[i])
		}
		for elf := 0; elf < len(elfs); elf++ {
			mv := 1 + recipes[elfs[elf]]
			elfs[elf] = (elfs[elf] + mv) % len(recipes)
		}
	}
	var res [10]int
	copy(res[:], recipes[after:])
	return res
}

func day14b(after ...uint8) int {
	recipes := make([]uint8, 0, 0)
	recipes = append(recipes, 3, 7)
	elfs := [2]int{0, 1}
	same := func(s1, s2 []uint8) bool {
		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
	adds := make([]uint8, 0, 2)
	for {
		combo := recipes[elfs[0]] + recipes[elfs[1]]
		adds = adds[:0]
		for {
			adds = append(adds, combo%10)
			combo /= 10
			if combo == 0 {
				break
			}
		}
		for i := len(adds) - 1; i >= 0; i-- {
			recipes = append(recipes, adds[i])

			if len(recipes) >= len(after) {
				cmp := recipes[len(recipes)-len(after):]
				if same(cmp, after) {
					return len(recipes) - len(after)
				}
			}
		}
		for elf := 0; elf < len(elfs); elf++ {
			mv := 1 + int(recipes[elfs[elf]])
			elfs[elf] = (elfs[elf] + mv) % len(recipes)
		}
	}
}
