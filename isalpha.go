package piscine

func IsAlpha(str string) bool {
	arr := []rune(str)
	check := true
	len := 0

	for i := range arr {
		len = i
	}

	for i := 0; i <= len; i++ {
		if (arr[i] < 'A' || arr[i] > 'Z') &&
			(arr[i] < 'a' || arr[i] > 'z') &&
			(arr[i] < '0' || arr[i] > '9') {
			check = false
			break
		}
	}
	return check
}
