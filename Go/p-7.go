package main

import (
	"fmt"
	"sort"
)

func main() {

	//問題
	//X 段重ねの鏡餅 (X≥1) とは、X 枚の円形の餅を縦に積み重ねたものであって、どの餅もその真下の餅より直径が小さい（一番下の餅を除く）もののことです。
	//例えば、直径 10、8、6 センチメートルの餅をこの順に下から積み重ねると 3 段重ねの鏡餅になり、餅を一枚だけ置くと 1 段重ねの鏡餅になります。
	//ダックスフンドのルンルンは N 枚の円形の餅を持っていて、そのうち i 枚目の餅の直径は d i​センチメートルです。
	//これらの餅のうち一部または全部を使って鏡餅を作るとき、最大で何段重ねの鏡餅を作ることができるでしょうか。

	//制約
	//1 ≤ N ≤ 100
	//1 ≤ d
	//i ≤ 100
	//入力値はすべて整数である。

	//入力
	//N
	//d1
	//:
	//dN​

	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	res := removeDuplicates(numbers)

	fmt.Println(len(res))
}

func removeDuplicates(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	result := []int{slice[0]}
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}

	return result
}
