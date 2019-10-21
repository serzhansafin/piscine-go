package main


import (
	"fmt"
)


func IterativePower(nb int, power int) int {

	if nb < 18 && power >= 0 {
		result := 1
		for i := 1; i <= power; i++ {
			result *= nb
		}
		return result
	}

	return 0

}


func main() {
	arg1 := -3
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}

