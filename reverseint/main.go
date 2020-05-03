package main

import (
	"fmt"
	"strconv"
)

func main() {}

func ReverseInt(nb int) int {
	neg := nb < 0
	if neg {
		nb *= -1
	}
	str := strconv.Itoa(nb)
	res := ""
	for _, c := range str {
		res = fmt.Sprintf("%s%s", string(c), res)
	}
	if i, err := strconv.Atoi(res); err == nil {
		if neg {
			return i * -1
		}
		return i
	}
	return -1
}
