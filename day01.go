package main

func day1a(input []string) (total int) {
	for _, line := range input {
		total += Atoi(line)
	}
	return
}

func day1b(input []string) (total int) {
	nums := make([]int, len(input))
	for i, line := range input {
		nums[i] = Atoi(line)
	}
	hist := map[int]bool{0: true}
	for {
		for _, num := range nums {
			total += num
			if _, ok := hist[total]; ok {
				return total
			}
			hist[total] = true
		}
	}
}
