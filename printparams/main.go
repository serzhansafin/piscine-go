package main

import "github.com/01-edu/z01"
import (
	"os"
)

func main() {
	params := os.Args
	length := 0

	for index := range params {
		length = index
	}

	for i := 1; i <= length; i++ {
		for _, w := range params[i] {
			z01.PrintRune(w)

		}
		z01.PrintRune('\n')

	}

}
