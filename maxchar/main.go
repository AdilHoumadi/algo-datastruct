package main

import "fmt"

func main() {
	c := MaxChar("ab1c1d1e1f1g1")
	fmt.Println(c)
}

func MaxChar(str string) string {
	res := make(map[string]int)
	for _, c := range str {
		val := res[string(c)]
		if val != 0 {
			res[string(c)] = val + 1
		} else {
			res[string(c)] = 1
		}
	}
	max := -1
	letter := ""
	for pos, val := range res {
		if max < val {
			max = val
			letter = pos
		}
	}
	return letter
}
