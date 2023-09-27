package main

import (
	"fmt"
)

func main() {

	//問題
	//1 以上 N 以下の整数のうち、10 進法での各桁の和が A 以上 B 以下であるものの総和を求めてください。

	//入力
	//N A B

	//出力
	//1 以上 N 以下の整数のうち、10 進法での各桁の和が A 以上 B 以下であるものの総和を出力せよ。

	var N, A, B int
	fmt.Scan(&N, &A, &B)

	res := 0

	for i := 1; i <= N; i++ {
		j := i
		sum := 0
		for j > 0 { // 各桁の和sumを求める
			sum += j % 10
			j /= 10
		}
		if sum >= A && sum <= B {
			res += i
		}
	}

	fmt.Println(res)
}
