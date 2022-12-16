package main

import "fmt"

type Op19 struct {
	name    string
	a, b, c int
}

func (op Op19) String() string {
	return fmt.Sprintf("%d = %s %d %d", op.c, op.name, op.b, op.c)
}

func day19funcs(reg []int) map[string]func(Op19) {
	return map[string]func(Op19){
		"addr": func(op Op19) {
			reg[op.c] = reg[op.a] + reg[op.b]
		},
		"addi": func(op Op19) {
			reg[op.c] = reg[op.a] + op.b
		},
		"mulr": func(op Op19) {
			reg[op.c] = reg[op.a] * reg[op.b]
		},
		"muli": func(op Op19) {
			reg[op.c] = reg[op.a] * op.b
		},
		"banr": func(op Op19) {
			reg[op.c] = reg[op.a] & reg[op.b]
		},
		"bani": func(op Op19) {
			reg[op.c] = reg[op.a] & op.b
		},
		"borr": func(op Op19) {
			reg[op.c] = reg[op.a] | reg[op.b]
		},
		"bori": func(op Op19) {
			reg[op.c] = reg[op.a] | op.b
		},
		"setr": func(op Op19) {
			reg[op.c] = reg[op.a]
		},
		"seti": func(op Op19) {
			reg[op.c] = op.a
		},
		"gtir": func(op Op19) {
			if op.a > reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"gtri": func(op Op19) {
			if reg[op.a] > op.b {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"gtrr": func(op Op19) {
			if reg[op.a] > reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqir": func(op Op19) {
			if op.a == reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqri": func(op Op19) {
			if reg[op.a] == op.b {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqrr": func(op Op19) {
			if reg[op.a] == reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
	}
}

// func day19a(input []string, reg0 int) int {
// 	var reg [6]int
// 	reg[0] = reg0
// 	funcs := day19funcs(reg[:])
// 	var ip int
// 	var ops []Op19
// 	for _, line := range input {
// 		if line[0] == '#' {
// 			// #ip 0
// 			Must(fmt.Sscanf(line, "#ip %d", &ip))
// 			continue
// 		}
// 		var op Op19
// 		Must(fmt.Sscanf(line, "%s %d %d %d", &op.name, &op.a, &op.b, &op.c))
// 		ops = append(ops, op)
// 	}

// 	for pntr := reg[ip]; ; pntr++ {
// 		reg[ip] = pntr
// 		if pntr >= len(ops) {
// 			break
// 		}
// 		op := ops[pntr]

// 		fmt.Printf("ip=%02d %v %v", pntr, reg, op)
// 		funcs[op.name](op)
// 		fmt.Printf(" %v\n", reg)

// 		// fmt.Println(reg)
// 		pntr = reg[ip]
// 	}

// 	return reg[0]
// }

func day19(input []string, a bool) int {
	target := 10551315
	if a {
		target = 915
	}

	zero := 0
	count := 0
	count2 := 1

	for count2 <= target {
		count = 1

		for count <= target {
			if count2*count == target {
				zero += count2
			}
			count++
		}

		count2++

		// This right here skips most of the work
		for target%count2 != 0 && count2 <= target {
			count2++
		}
	}

	return zero
}

// func tmp() {
// 	ip := 0
// 	reg0 := 1

// 	ip += 16 // goto ins 16+1
// 	ip *= ip // goto ins 256
// }

/*
#ip 2
00 -> ip += 16 (goto 17)

01 -> reg3 = 1
02 -> reg5 = 1
03 -> reg4 = reg3 * reg5
04 -> reg4 = reg4 == reg1
05 -> ip += reg4
06 -> ip += 1

07 -> reg0 += reg3 // if reg3 * reg5 == 10551315

08 -> reg5 +=  1
09 -> reg4 = reg5 > reg1
10 -> ip += reg4
11 -> ip = 2 // goto 3 if reg5 <= 10551315

12 -> reg3 += 1
13 -> reg4 = reg3 > reg1
14 -> ip += reg4
15 -> ip = 1

16 -> ip *= ip //====> EXIT WHEN reg3 > 10551315

17 -> reg1 += 2
18 -> reg1 *= reg1
19 -> reg1 *= ip (19)
20 -> reg1 *= 11
21 -> reg4 += 3
22 -> reg4 *= ip (22)
23 -> reg4 += 13
24 -> reg1 += reg4   // block purpose: reg1 = 915

25 -> ip += reg0 (if b, skip next line)
26 -> ip = 0

27 -> reg4 = ip (27)
28 -> reg4 *= ip (28)
29 -> reg4 += ip (29)
30 -> reg4 *= ip (30)
31 -> reg4 *= 14
32 -> reg4 *= ip (32)
33 -> reg1 += reg4   // block purpose: reg1 = 10551315

34 -> reg0 = 0
35 -> ip = 0 // goto 1
*/
