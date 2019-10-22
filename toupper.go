package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for index, word := range arr {
		if word >= 'a' && word <= 'z' {
			arr[index] = word - 32
		}
	}
	return string(arr)
}
