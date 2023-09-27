package main

import (
	"fmt"
)

func main() {

	//問題
	// N 枚のカードがあります. i 枚目のカードには, a i​  という数が書かれています.
	//Alice と Bob は, これらのカードを使ってゲームを行います.
	//ゲームでは, Alice と Bob が交互に 1 枚ずつカードを取っていきます. Alice が先にカードを取ります.
	//2 人がすべてのカードを取ったときゲームは終了し,
	//取ったカードの数の合計がその人の得点になります.
	//2 人とも自分の得点を最大化するように最適な戦略を取った時, Alice は Bob より何点多く取るか求めてください.

	//制約
	//N は 1 以上 100 以下の整数
	//ai (1≤i≤N) は 1 以上 100 以下の整数

	//入力
	//N
	//a1  a2  a3 ... aN​

	//出力
	//両者が最適な戦略を取った時, Alice は Bob より何点多く取るかを出力してください.

	var N, Alice, Bob int
	fmt.Scan(&N)

	deck := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&deck[i])
	}

	for j := 0; len(deck) > 0; j++ {
		card := 0
		card, deck = getMaxCard(deck)
		if j%2 == 0 {
			Alice += card
		} else {
			Bob += card
		}
	}
	fmt.Println(Alice - Bob)
}

func getMaxCard(deck []int) (int, []int) {
	max := deck[0]
	count := 0
	for i := 0; i < len(deck); i++ {
		if deck[i] > max {
			max = deck[i]
			count = i
		}
	}
	deck = append(deck[:count], deck[count+1:]...)
	return max, deck
}
