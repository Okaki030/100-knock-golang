// 05. n-gram
// 与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I am an NLPer"

	// 単語bi-gram
	slice := strings.Split(str, " ")
	i := 0
	for i+2 <= len(slice) {
		fmt.Println(slice[i : i+2])
		i += 2
	}

	// 文字bi-gram
	for i, s := range str {
		fmt.Printf("%c", s)
		if i%2 == 1 {
			fmt.Println()
		}
	}
}
