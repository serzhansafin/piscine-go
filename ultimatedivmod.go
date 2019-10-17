package piscine

func UltimateDivMod(a *int, b *int) {
	var temp, temp1 int

	temp = *a / *b
	temp1 = *a % *b
	*a = temp
	*b = temp1

}
