package main

import "github.com/01-edu/z01"

func main() {

	var s string = "abcdefghijklmnopqrtwxyz"
	for i := 0; i < len(s); i++  {

		z01.PrintRune(rune(s[i]))
	}
}
