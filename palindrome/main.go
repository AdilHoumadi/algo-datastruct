package main

import "fmt"

func main() {}

func IsPalindrome1(str string) bool {
	res := ""
	for _, c := range str {
		res = fmt.Sprintf("%s%s", string(c), res)
	}
	return res == str
}

func IsPalindrome(str string) bool {
	last := len(str) - 1
	for pos, char := range str {
		if pos == last-pos {
			break
		}
		if string(char) != string(str[last-pos]) {
			return false
		}
	}
	return true
}
