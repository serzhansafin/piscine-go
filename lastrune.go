package piscine

func LastRune(s string) rune {
	b := 0
	for index := range s {
		b = index
	}
	arr := []rune(s)
	return arr[b]
}
