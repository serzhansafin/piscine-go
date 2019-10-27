package piscine

import "github.com/01-edu/z01"

func printStr(str string) {
	arrayStr := []rune(str)

	lengthOfStr := 0

	for ind := range arrayStr {
		lengthOfStr = ind
	}

	for i := 0; i < lengthOfStr; i++ {

		z01.PrintRune(string(arrayStr[i]))

	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	lengthOfArg := 0

	sentence := os.Args
	for index := range sentence {
		lengthOfArg = index
	}

	if isEven(lengthOfArg) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
