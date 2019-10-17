package piscine

func StrLep(str string) int {

	sentence := []rune(str)
	var num int
	for index := range sentence {

		num = index
	}
	return num + 1
}

func StrRev(s string) string {

	reverse := []rune(s)
	for i, j := 0, StrLep(s)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)

}
