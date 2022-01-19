package main

import "fmt"

type Op struct {
	opcode, a, b, c int
}

func day16funcs(reg []int) map[string]func(Op) {
	return map[string]func(Op){
		"addr": func(op Op) {
			reg[op.c] = reg[op.a] + reg[op.b]
		},
		"addi": func(op Op) {
			reg[op.c] = reg[op.a] + op.b
		},
		"mulr": func(op Op) {
			reg[op.c] = reg[op.a] * reg[op.b]
		},
		"muli": func(op Op) {
			reg[op.c] = reg[op.a] * op.b
		},
		"banr": func(op Op) {
			reg[op.c] = reg[op.a] & reg[op.b]
		},
		"bani": func(op Op) {
			reg[op.c] = reg[op.a] & op.b
		},
		"borr": func(op Op) {
			reg[op.c] = reg[op.a] | reg[op.b]
		},
		"bori": func(op Op) {
			reg[op.c] = reg[op.a] | op.b
		},
		"setr": func(op Op) {
			reg[op.c] = reg[op.a]
		},
		"seti": func(op Op) {
			reg[op.c] = op.a
		},
		"gtir": func(op Op) {
			if op.a > reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"gtri": func(op Op) {
			if reg[op.a] > op.b {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"gtrr": func(op Op) {
			if reg[op.a] > reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqir": func(op Op) {
			if op.a == reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqri": func(op Op) {
			if reg[op.a] == op.b {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
		"eqrr": func(op Op) {
			if reg[op.a] == reg[op.b] {
				reg[op.c] = 1
			} else {
				reg[op.c] = 0
			}
		},
	}
}

func day16a(input []string) int {
	var after, before [4]int
	var op Op
	total := 0
	var reg [4]int
	funcs := day16funcs(reg[:])
	for i := 0; i < len(input); i += 4 {
		line := input[i]
		if len(line) == 0 {
			break
		}
		Must(fmt.Sscanf(line[8:], "[%d, %d, %d, %d]",
			&before[0], &before[1], &before[2], &before[3]))
		Must(fmt.Sscanf(input[i+1], "%d %d %d %d",
			&op.opcode, &op.a, &op.b, &op.c))
		Must(fmt.Sscanf(input[i+2][8:], "[%d, %d, %d, %d]",
			&after[0], &after[1], &after[2], &after[3]))
		count := 0
		for _, f := range funcs {
			reg = before
			f(op)
			if reg == after {
				count++
				if count == 3 {
					total++
					break
				}
			}
		}
	}
	return total
}

func day16b(input []string) int {
	var reg [4]int
	funcs := day16funcs(reg[:])
	progI := -1

	type ins struct {
		after, before [4]int
		op            Op
	}
	var inss []ins
	opfuncs := map[string]int{}
	opfuncs2 := map[int]string{}
	for i := 0; i < len(input); i += 4 {
		line := input[i]
		var in ins
		if len(line) == 0 {
			progI = i + 2
			break
		}
		Must(fmt.Sscanf(line[8:], "[%d, %d, %d, %d]",
			&in.before[0], &in.before[1], &in.before[2], &in.before[3]))
		Must(fmt.Sscanf(input[i+1], "%d %d %d %d",
			&in.op.opcode, &in.op.a, &in.op.b, &in.op.c))
		Must(fmt.Sscanf(input[i+2][8:], "[%d, %d, %d, %d]",
			&in.after[0], &in.after[1], &in.after[2], &in.after[3]))
		inss = append(inss, in)
		opfuncs2[in.op.opcode] = ""
	}

outer:
	for i := 0; i < len(inss); i++ {
		in := inss[i]
		if opfuncs2[in.op.opcode] != "" {
			continue
		}
		count := 0
		var fname string
		for name, f := range funcs {
			if _, ok := opfuncs[name]; ok {
				continue
			}
			copy(reg[:], in.before[:])
			f(in.op)
			if reg == in.after {
				fname = name
				if count == 1 {
					continue outer
				}
				count++
			}
		}
		if count == 1 {
			opfuncs[fname] = in.op.opcode
			opfuncs2[in.op.opcode] = fname
		}
	}

	for opcode, name := range opfuncs2 {
		if name == "" {
			panic(fmt.Sprintln("unknown opcode: ", opcode))
		}
	}

	copy(reg[:], []int{0, 0, 0, 0})
	for i := progI; i < len(input); i++ {
		line := input[i]
		var op Op
		Must(fmt.Sscanf(line, "%d %d %d %d",
			&op.opcode, &op.a, &op.b, &op.c))
		funcs[opfuncs2[op.opcode]](op)
	}
	return reg[0]
}
