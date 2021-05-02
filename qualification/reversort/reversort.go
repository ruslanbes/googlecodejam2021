// Copyright (c) 2021 ruslanbes. All rights reserved.
//
// Google Code Jam 2021 Qualification Round - Problem A. Reversort
// https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d0a5c
//
// Time:  O(N^2)
// Space: O(1)
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitToInts(str string) []int {
	strs := strings.Split(str, " ")
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func minIndex(ints []int, i int) int {
	mi := i
	for k := i; k < len(ints); k++ {
		if ints[k] < ints[mi] {
			mi = k
		}
	}
	return mi
}

func reverse(ints []int, i, j int) {
	for i < j {
		ints[i], ints[j] = ints[j], ints[i]
		i++
		j--
	}
}

func readLine(sc *bufio.Scanner) string {
	if sc.Scan() {
		return sc.Text()
	} else {
		panic("EOF")
	}
}

func printCase(t int, res interface{}) {
	fmt.Printf("Case #%d: %v\n", t, res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	T, _ := strconv.Atoi(readLine(scanner))
	for t := 0; t < T; t++ {
		N, _ := strconv.Atoi(readLine(scanner))
		L := splitToInts(readLine(scanner))
		res := 0
		for i := 0; i < N-1; i++ {
			mi := minIndex(L, i)
			res += mi - i + 1
			reverse(L, i, mi)
		}
		printCase(t+1, res)
	}

}
