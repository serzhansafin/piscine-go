package piscine

func IsPrintable(str string) bool {
	arr := []rune(str)
	ln := 0
	for i := range arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if arr[i] < 32 {
			return false
		}
	}
	return true
}
