package piscine

/*
import "github.com/01-edu/z01"
*/
/*
import (
	"fmt"
)
*/
func AlphaCount(str string) int {

	counter := 0
	list := []byte(str)

	for _, letter := range list {

		if letter >= 65 && letter <= 90 || letter <= 122 && letter >= 97 {

			counter++
		}
	}

	return counter
}

/*
func main() {
    str := "Hello 78 World!    4455 /"
    nb := AlphaCount(str)
    fmt.Println(nb)
}
*/
