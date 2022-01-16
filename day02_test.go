package main

import "testing"

func TestDay2(t *testing.T) {
	TestEqual(t, 12, day2a([]string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee	",
		"ababab",
	}))
	TestEqual(t, "fgij", day2b([]string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}))
	file := Lines(2)
	TestEqual(t, 5952, day2a(file))
	TestEqual(t, "krdmtuqjgwfoevnaboxglzjph", day2b(file))
}
