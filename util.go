package main

import (
	"bufio"
	"constraints"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestEqual[T comparable](t *testing.T, expect, result T, message ...string) {
	t.Helper()
	var msg string
	if len(message) > 0 {
		msg = message[0]
	}
	t.Run(msg, func(t *testing.T) {
		t.Helper()
		if result != expect {
			t.Fatalf("expected %v, got %v", expect, result)
		}
	})
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Atoi(input string) int {
	i := Must(strconv.Atoi(input))
	return i
}

func Abs[T constraints.Signed | constraints.Float](v T) T {
	if v < 0 {
		return v * -1
	}
	return v
}

const inputDir = "."

func inputPath(day int, suffix ...string) string {
	var s string
	if len(suffix) > 0 {
		s = strings.Join(suffix, "")
	}
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d%s.txt", inputDir, day, s)))
}

func Input(day int, suffix ...string) string {
	data := Must(os.ReadFile(inputPath(day, suffix...)))
	return strings.TrimRight(string(data), "\n")
}

func InputReader(day int, suffix ...string) io.ReadSeekCloser {
	return Must(os.Open(inputPath(day, suffix...)))
}

func InputScanner(day int, suffix ...string) (s *bufio.Scanner, close func() error) {
	f := InputReader(day, suffix...)
	s = bufio.NewScanner(f)
	return s, f.Close
}

func Lines(day int, suffix ...string) []string {
	b, close := InputScanner(day, suffix...)
	defer close()
	var lines []string
	for b.Scan() {
		lines = append(lines, b.Text())
	}
	return lines
}

func IntLines(day int, suffix ...string) (lines []int) {
	b, close := InputScanner(day, suffix...)
	defer close()
	for b.Scan() {
		lines = append(lines, Atoi(b.Text()))
	}
	return lines
}

func Values[T any, K comparable](m map[K]T) []T {
	result := make([]T, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func Map[T, U any](f func(v T) U, items []T) []U {
	out := make([]U, len(items))
	for i := range items {
		out[i] = f(items[i])
	}
	return out
}

func Reduce[T, U any](f func(U, T) U, items []T, init U) U {
	var acc U
	for _, item := range items {
		acc = f(acc, item)
	}
	return acc
}

func MinMax[T constraints.Ordered](values []T) (min, max T) {
	for i, v := range values {
		if i == 0 || v < min {
			min = v
		}
		if i == 0 || v > max {
			max = v
		}
	}
	return min, max
}

func Max[T constraints.Ordered](values ...T) (max T) {
	for i, v := range values {
		if i == 0 || v > max {
			max = v
		}
	}
	return
}

func Min[T constraints.Ordered](values ...T) (min T) {
	for i, v := range values {
		if i == 0 || v < min {
			min = v
		}
	}
	return
}

func In[T comparable](item T, items ...T) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func CopySlice[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

func Remove[T comparable](slice []T, remove ...T) []T {
	result := make([]T, 0, len(slice))
	for _, val := range slice {
		if !In(val, remove...) {
			result = append(result, val)
		}
	}
	return result
}

func NewGrid[T any](size int, init T) [][]T {
	res := make([][]T, size)
	for i := 0; i < size; i++ {
		res[i] = make([]T, size)
		for j := 0; j < size; j++ {
			res[i][j] = init
		}
	}
	return res
}
