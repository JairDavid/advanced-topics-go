package functions

import "strings"

func Repeat(n int) func(text string) string {
	return func(text string) string {
		return strings.Repeat(text, n)
	}
}
