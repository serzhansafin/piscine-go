package piscine

func IsLower(str string) bool {
	arr := []rune(str)
	ln := 0
	for i := range arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if arr[i] < 'a' || arr[i] > 'z' {
			return false
		}
	}
	return true
}
