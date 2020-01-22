// 01. 「パタトクカシーー」
// 「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．

package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {

	s := strings.Split(str, "")

	var reverseStr string
	for i := len(s) - 1; i >= 0; {
		reverseStr += s[i]
		i--
	}

	return strings.Join(s, "")
}

func main() {
	str := "stressed"

	fmt.Println(Reverse(str))
}
