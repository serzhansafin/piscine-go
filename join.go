package piscine

func Join(strs []string, sep string) string {

	var r string

	length := 0

	for i := range strs {
		length = i
	}

	for j, w := range strs {

		if length != j {
			r += w + sep
		} else {
			r += w
		}
	}
	return r
}
