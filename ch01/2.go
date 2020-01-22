// 02. 「パトカー」＋「タクシー」＝「パタトクカシーー」
// 「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．

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
