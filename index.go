package piscine

func Index(s string, toFind string) int {
	ln := 0
	ln2 := 0
	for _, c := range s {
		if c == c {
			ln++
		}
	}

	for _, c := range toFind {
		if c == c {
			ln2++
		}
	}
	for i := 0; i < ln; i++ {
		if ln2 != 0 && s[i] == toFind[0] {
			ok := true
			num := 0
			for j := 0; j < ln2; j++ {
				if i+num >= ln || toFind[j] != s[i+num] {
					ok = false
					break
				}
				num++
			}
			if ok == true {
				return i
			}
		}
	}
	if toFind == "" {
		return 0
	}
	return -1
}
