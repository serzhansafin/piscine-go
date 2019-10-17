package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	for _, char := range str {

		z01.PrintRune(char)

	}
	z01.PrintRune(10)
}
