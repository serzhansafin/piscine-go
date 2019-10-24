package piscine

/*
import (
	"fmt"
)
*/

func AppendRange(min, max int) []int {
	var ans []int

	for i := min - 1; i < max; i++ {

		if i != max-1 {
			ans = append(ans, i+1)
		}

	}
	return ans
}

/*
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
*/
