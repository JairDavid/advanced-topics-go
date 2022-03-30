package functions

import "strings"

func Repeat(n int) func(text string) string {
	return func(text string) string {
		return strings.Repeat(text, n)
	}
}

func Counter() func() int {
	c := 0
	return func() int {
		c += 1
		return c
	}
}
