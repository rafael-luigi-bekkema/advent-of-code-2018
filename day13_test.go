package main

import "testing"

func TestDay13(t *testing.T) {
	example := `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `[1:]
	TestEqual(t, NewResult("7,3", ""), day13a(example))
	example = `
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`[1:]
	TestEqual(t, NewResult("2,0", "6,4"), day13a(example))
	TestEqual(t, NewResult("118,66", "70,129"), day13a(Input(13)))
}
