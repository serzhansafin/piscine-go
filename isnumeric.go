package piscine

func IsNumeric(str string) bool {
	arr := []rune(str)
	ln := 0
	for i := range arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if arr[i] < '0' || arr[i] > '9' {
			return false
		}
	}
	return true
}
