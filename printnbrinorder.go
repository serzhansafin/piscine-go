package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	arr := [10]int{}

	if n >= 0 && n <= 9 {

		z01.PrintRune(rune(n + 48))

	} else {
		for {
			if n == 0 {
				break
			}
			arr[n%10]++
			n = n / 10
		}
		for index, num := range arr {
			if num != 0 {
				for i := 0; i < num; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}

	//z01.PrintRune(10)

}

/*
func main() {
    PrintNbrInOrder(321)
    PrintNbrInOrder(0)
    PrintNbrInOrder(321)
}
*/
