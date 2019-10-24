package piscine

/*
import (
	"fmt"
)
*/

func MakeRange(min, max int) []int {
	res := max - min
	var ans []int

	if max <= min {

		return nil

	}

	ans = make([]int, res)

	for i := 0; i < res; i++ {

		ans[i] = min
		min++

	}

	return ans
}

/*
func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
*/
