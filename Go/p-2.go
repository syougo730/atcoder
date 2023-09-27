package main

import (
	"fmt"
)

func main() {

	//問題
	// すぬけ君は 1,2,3 の番号がついた 3 つのマスからなるマス目を持っています。 各マスには 0 か 1 が書かれており、
	// マス i には s i が書かれています。
	// すぬけ君は 1 が書かれたマスにビー玉を置きます。 ビー玉が置かれるマスがいくつあるか求めてください。

	//入力
	//s1s2s3

	// 制約
	// s1 ​,s2 ​,s3は1 あるいは 0

	var num string
	fmt.Scan(&num)
	res := 0
	for i := 0; i < len(num); i++ {
		if num[i] == '1' {
			res++
		}
	}
	fmt.Println(res)
}
