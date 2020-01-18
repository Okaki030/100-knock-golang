package main

import (
	"fmt"
)

func main() {
	str := "パタトクカシーー"

	for i, r := range str {
		if i%2 == 1 {
			fmt.Printf("%c", r)
		}
	}
}
