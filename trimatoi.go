package piscine

/*
import "github.com/01-edu/z01"
*/
/*
import (
	"fmt")
*/
func TrimAtoi(s string) int {

	arr := []rune(s)

	c := 0
	num_our := 0
	sign := 1
	minus := false

	for _, word := range arr {

		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {

				c++
			}
			num_our = num_our*10 + c
			c = 0

		}

		if word == '-' && num_our == 0 {
			minus = true
		}
	}
	if minus == true {
		sign = -1
	}

	return num_our * sign

}

/*
func main() {
	s := "12345"
	s2 := "str123ing45"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "sd+x1fa2W3s4"
	s6 := "sd-x1fa2W3s4"
	s7 := "sdx1-fa2W3s4"

	n := TrimAtoi(s)
	n2 := TrimAtoi(s2)
	n3 := TrimAtoi(s3)
	n4 := TrimAtoi(s4)
	n5 := TrimAtoi(s5)
	n6 := TrimAtoi(s6)
	n7 := TrimAtoi(s7)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}
*/
