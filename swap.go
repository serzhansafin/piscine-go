package piscine

func Swap(a *int, b *int) {

	var temp int
	temp = *b
	*b = *a
	*a = temp

}
