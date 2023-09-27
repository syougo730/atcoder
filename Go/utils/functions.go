package main

import (
	"sort"
)

// 昇順に整数スライスをソートする関数
func sortAscending(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// 降順に整数スライスをソートする関数
func sortDescending(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

// 整数型のスライスから重複した値を削除する関数
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
