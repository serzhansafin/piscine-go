package piscine

/*
import (
	"fmt"
)
*/

func IterativeFactorial(nb int) int {

	if nb < 18 && nb >= 0 {
		result := 1

		for i := 1; i <= nb; i++ {

			result *= i

		}
		return result
	}

	return 0

}

/*
func main() {
	arg := -4738
	fmt.Println(IterativeFactorial(arg))
}
*/
