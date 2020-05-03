package main

import (
	"fmt"
	"io"
)

func main() {}

func FizBuzz(out io.Writer, n int) {
	index := 1
	for {
		if index > n {
			break
		}
		if index%3 == 0 && index%5 == 0 {
			fmt.Fprint(out, "FizzBuzz")
		} else if index%3 == 0 {
			fmt.Fprint(out, "Fizz")
		} else if index%5 == 0 {
			fmt.Fprint(out, "Buzz")
		} else {
			fmt.Fprint(out, index)
		}
		fmt.Fprintln(out, "")
		index++
	}
}
