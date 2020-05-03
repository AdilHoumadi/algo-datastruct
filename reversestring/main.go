package main

import "fmt"

func main() {}

func Reverse(str string) string {
	i := len(str)
	data := ""
	for {
		i = i - 1
		data = fmt.Sprintf("%s%s", data, string(str[i]))
		if i == 0 {
			break
		}
	}
	return data
}

func Reverse1(str string) string {
	data := ""
	for _, char := range str {
		data = fmt.Sprintf("%s%s", string(char), data)
	}
	return data
}
