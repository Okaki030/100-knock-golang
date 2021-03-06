// 03. 円周率
// "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Index  int
	Str    string
	Length int
}

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	slice := strings.Split(str, " ")

	// インデックと単語ごとの文字数を取得
	word := make([]Word, 0)
	for i, s := range slice {
		word = append(word, Word{i, s, len(s)})
	}

	// 単語の文字数順にsort
	sort.Slice(word, func(i, j int) bool { return word[i].Length < word[j].Length })
	for _, w := range word {
		fmt.Println(w.Str)
	}
}
