package piscine

/*
import (
	"fmt"
)
*/

func Fibonacci(index int) int {
	if index == 1 || index == 2 {
		return 1
	}
	if index > 2 || index < 100 {
		return Fibonacci(index-2) + Fibonacci(index-1)
	}
	return 0
}

/*
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
*/
