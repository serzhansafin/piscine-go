package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for index, letter := range arr {
		if letter >= 'A' && letter <= 'Z' {
			arr[index] = letter + 32
		}
	}
	return string(arr)
}
