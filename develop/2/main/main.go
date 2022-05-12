package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(word string) string {
	prev, res := "", ""
	notEsc := true

	for _, v := range word {
		if string(v) == `\` && notEsc { // && notEsc for \\
			notEsc = false
			continue
		}

		if !(unicode.IsDigit(v) && notEsc) {
			//letter
			prev = string(v)
			res += string(v)
		} else {
			//digit
			n, _ := strconv.Atoi(string(v))
			for i := 0; i < n-1; i++ {
				res += prev
			}
		}
		notEsc = true
	}
	return res
}

func main() {
	word := `qwe\4\5`

	fmt.Println(Unpack(word))
}
