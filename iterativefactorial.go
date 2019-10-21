package piscine

/*
import (
	"fmt"
)
*/

func IterativeFactorial(nb int) int {

	result := 1

	for i := 1; i <= nb; i++ {

		result *= i

	}
	return result
}

/*
func main() {
	arg := -6
	fmt.Println(IterativeFactorial(arg))
}
*/
