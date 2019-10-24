package piscine

/*
import (
	"fmt"
)
*/
func ConcatParams(args []string) string {

	var index int
	var res string

	for i := range args {

		index = i
	}

	for i := 0; i <= index; i++ {
		if i != index {
			res = res + args[i] + "\n"
		} else {
			res = res + args[i]

		}

	}
	return res
}

/*
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
*/
