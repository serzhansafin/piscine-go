package piscine

func IsUpper(str string) bool {
	arr := []rune(str)
	ln := 0
	for i := range arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if arr[i] < 'A' || arr[i] > 'Z' {
			return false
		}
	}
	return true
}
