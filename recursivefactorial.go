package piscine

/*
import (
	"fmt"
)
*/

func RecursiveFactorial(nb int) int {

	if nb > 1 && nb < 18 {
		return nb * RecursiveFactorial(nb-1)
	}
	if nb == 1 {
		return 1
	}
	return 0

}

/*
func main() {
	arg := 3
	fmt.Println(RecursiveFactorial(arg))
}
*/
