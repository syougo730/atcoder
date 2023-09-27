package main

import (
	"fmt"
)

func main() {

	//問題
	// 黒板に N 個の正の整数 A 1 ,...,A N​ が書かれています．
	//すぬけ君は，黒板に書かれている整数がすべて偶数であるとき，次の操作を行うことができます．
	//黒板に書かれている整数すべてを，2 で割ったものに置き換える．すぬけ君は最大で何回操作を行うことができるかを求めてください．

	//入力
	//N​
	//A1 A2​ ... AN​

	// 制約
	// 1 ≤ N ≤ 200
	// 1 ≤ Ai ≤ 10^9

	var n int
	var count = 0
	fmt.Scan(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	for {
		if searchOdd(n, numbers) {
			break
		}
		numbers = halve(n, numbers)
		count++
	}

	fmt.Println(count)
}

func searchOdd(n int, numbers []int) bool {
	for i := 0; i < n; i++ {
		if numbers[i]%2 == 1 {
			return true
		}
	}
	return false
}

func halve(n int, numbers []int) []int {
	for i := 0; i < n; i++ {
		numbers[i] = numbers[i] / 2
	}
	return numbers
}
