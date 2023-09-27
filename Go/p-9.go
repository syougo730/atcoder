package main

import (
	"fmt"
	"strings"
)

func main() {

	//問題
	// 英小文字からなる文字列 S が与えられます。
	//Tが空文字列である状態から始め、以下の操作を好きな回数繰り返すことで S=T とすることができるか判定してください。
	//・T の末尾に dream dreamer erase eraser のいずれかを追加する。

	//入力
	//S

	// 制約
	// 1 ≦∣ S ∣≦ 10^5

	var S string
	var list [4]string = [4]string{"dream", "dreamer", "erase", "eraser"}
	fmt.Scan(&S)

	for {
		flg := false
		for i := 0; i < len(list); i++ {
			if strings.HasSuffix(S, list[i]) {
				flg = true
				S = strings.TrimSuffix(S, list[i])
			}
		}
		if !flg {
			fmt.Println("NO")
			return
		}
		if len(strings.TrimSpace(S)) == 0 {
			fmt.Println("YES")
			return
		}
	}
}
