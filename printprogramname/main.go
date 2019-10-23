package main

import "github.com/01-edu/z01"
import (
	"os"
)

/*
import (
	"fmt"
	"os"
)
*/
func main() {

	name := os.Args
	for _, n := range name[0] {
		z01.PrintRune(n)
	}
	z01.PrintRune(10)
}
