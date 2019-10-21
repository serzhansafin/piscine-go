package piscine

/*
import (
	"fmt"
)
*/

func RecursivePower(nb int, power int) int {

	if power == 0 {
		return 1
	}

	if power >= 0 && power <= 18 {
		return nb * RecursivePower(nb, power-1)
	}
	return 0
}

/*

func main() {
	arg1 := -3
	arg2 := 0
	fmt.Println(RecursivePower(arg1, arg2))
}
*/
