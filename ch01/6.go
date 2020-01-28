// 06. 集合
// "paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ

package main

import (
	"fmt"
	"strings"
)

// 文字bi-gram
func wordBiGram(strArray []string) []string {

	biGram := make([]string, 0)

	for _, v := range strArray {
		if 2 <= len(v) {
			for i := 0; i < len(v)-1; i++ {
				biGram = append(biGram, fmt.Sprint(v[i:i+2], " "))
			}
		}
	}

	return biGram
}

// スライスの中に含まれていない場合false
func containsStrNot(unionWords []string, searchWord string) bool {

	for _, word := range unionWords {
		if word == searchWord {
			return false
		}
	}

	return true
}

// 重複削除
func removeDuplicate1(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

// 和集合
func union(x, y []string) []string {

	unionWords := removeDuplicate1(x)

	for _, searchWord := range y {
		if containsStrNot(unionWords, searchWord) {
			unionWords = append(unionWords, searchWord)
		}
	}

	return unionWords
}

// // 積集合
// func intersection(x, y []string) []string {

// }

// // 差集合
// func diffrenceSet(x, y []string) []string {

// }

func main() {

	strArray1 := strings.Fields("paraparaparadise")
	strArray2 := strings.Fields("paragraph")

	x := wordBiGram(strArray1)
	y := wordBiGram(strArray2)

	fmt.Println(union(x, y))
}
