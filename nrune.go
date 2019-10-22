package piscine

func NRune(s string, n int) rune {
	b := 0
	for index := range s {
		b = index
	}
	if (b+1) >= n && n > 0 {
		arr := []rune(s)
		return (arr[n-1])
	} else {
		return 0
	}
}
