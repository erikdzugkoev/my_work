package main

import (
	"fmt"
	"unicode"
)

func Capitalize(s string) []string {
	a, b := []rune(s), []rune(s)

	for i := range a {
		if i%2 == 0 {
			a[i] = unicode.ToUpper(a[i])
		} else {
			b[i] = unicode.ToUpper(b[i])
			fmt.Println([]string{string(a), string(b)})
		}
	}
	return []string{string(a), string(b)}
}
func main() {
	Capitalize("abcdef")

}
