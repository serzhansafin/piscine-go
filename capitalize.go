package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	if arr[0] >= 'a' && arr[0] <= 'z' {
		arr[0] = rune(arr[0] - 32)
	}

	len_a := 0

	for index := range arr {
		len_a = index
	}

	for i := 1; i <= len_a; i++ {
		if (arr[i-1] < 'A' || arr[i-1] > 'Z') &&
			(arr[i] >= 'a' && arr[i] <= 'z') &&
			(arr[i-1] < 'a' || arr[i-1] > 'z') &&
			(arr[i-1] < '0' || arr[i-1] > '9') {
			arr[i] = rune(arr[i] - 32)
		}
		if (arr[i-1] >= 'A' && arr[i-1] <= 'Z') &&
			(arr[i] >= 'A' && arr[i] <= 'Z') {
			arr[i] = rune(arr[i] + 32)
		}
		if (arr[i-1] >= 'a' && arr[i-1] <= 'z') &&
			(arr[i] >= 'A' && arr[i] <= 'Z') {
			arr[i] = rune(arr[i] + 32)
		}
		if (arr[i-1] >= '0' && arr[i-1] <= '9') &&
			(arr[i] >= 'A' && arr[i] <= 'Z') {
			arr[i] = rune(arr[i] + 32)
		}

	}

	return string(arr)
}
