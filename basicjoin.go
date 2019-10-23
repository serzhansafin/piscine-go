package main

import (
	"fmt"
)

func BasicJoin(strs []string) string {

	var r string
	for _, w := range strs {
		r += w
	}
	return r
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(toConcat))
}
