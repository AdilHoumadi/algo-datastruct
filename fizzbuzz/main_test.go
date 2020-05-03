package main

import (
	"bytes"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	buffer := &bytes.Buffer{}
	FizBuzz(buffer, 15)
	got := buffer.String()
	want := `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
