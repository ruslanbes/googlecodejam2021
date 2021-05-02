// Copyright (c) 2021 ruslanbes. All rights reserved.
//
// Google Code Jam Tester

// "github.com/rhysd/go-fakeio"
package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/rhysd/go-fakeio"
)

func runMain(t *testing.T) string {
	data, err := ioutil.ReadFile("in.txt")
	if err != nil {
		t.Errorf("File reading error %v", err)
		panic(err)
	}
	fake := fakeio.Stdout().StdinBytes(data)
	fake.CloseStdin()

	main()

	out, err := fake.String()
	if err != nil {
		t.Errorf("Reading buffer error: %v", err)
		panic(err)
	}

	fake.Restore()
	return out
}

func saveOut(t *testing.T, out string) {
	fileOut, err := os.Create("out.go.txt")
	if err != nil {
		t.Errorf("Creating output file error: %v", err)
		panic(err)
	}
	defer fileOut.Close()

	fileOut.WriteString(out)
}

func validateOut(t *testing.T, out string) {
	got := strings.Split(out, "\n")

	file, err := os.Open("out.txt")
	if err != nil {
		t.Errorf("failed to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		if scanner.Text() != got[i] {
			t.Errorf("Got: %v, want: %v", got[i], scanner.Text())
		}
		i++
	}
}

func Test_main(t *testing.T) {
	out := runMain(t)
	saveOut(t, out)
	validateOut(t, out)
}
