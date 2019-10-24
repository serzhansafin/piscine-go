package piscine

/*
import (
	"fmt"
)
*/
/*
func SplitWhiteSpaces(str string) []string {

	var index int
	//my_str := ""


		for j := range str {
			index = j
		}

	my_str := make([]string, index+1)

		for n, w := range str {

			if w != ' ' {
				my_str[n] = string(w)

			}
		}

	return my_str

}

func main() {
	str := "Hello how are you?"
	fmt.Println(SplitWhiteSpaces(str))
}
*/

func SplitWhiteSpaces(str string) []string {
	i := 0
	s := ""
	len := 0

	for i, value := range str {
		if (value == ' ' || value == '\n' || value == '\t') && (i != 0 && (str[i-1] != ' ' && str[i-1] != '\n' && str[i-1] != '\t')) {
			len++
		}
	}

	array := make([]string, len+1)
	for _, value := range str {
		if value == ' ' || value == '\n' || value == '\t' {
			if s != "" {
				array[i] = s
				s = ""
				i++
			}
		} else {
			s = s + string(value)
		}
	}
	if s != "" {
		array[i] = s
	}
	return array
}

func main() {
	str := "Hello how are you?"
	fmt.Println(SplitWhiteSpaces(str))
}
