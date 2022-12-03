package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "test" {
		fmt.Println(day17b([]string{
			"x=495, y=2..7",
			"y=7, x=495..501",
			"x=501, y=3..7",
			"x=498, y=2..4",
			"x=506, y=1..2",
			"x=498, y=10..13",
			"x=504, y=10..13",
			"y=13, x=498..504",
		}, false))
		return
	}
	fmt.Println(day17b(Lines(17), true))
}
