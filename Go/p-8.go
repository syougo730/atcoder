package main

import (
	"fmt"
)

func main() {

	//問題
	// 日本でよく使われる紙幣は、10000 円札、5000 円札、1000 円札です。以下、「お札」とはこれらのみを指します。
	//青橋くんが言うには、彼が祖父から受け取ったお年玉袋にはお札が N 枚入っていて、合計で Y 円だったそうですが、嘘かもしれません。
	//このような状況がありうるか判定し、ありうる場合はお年玉袋の中身の候補を一つ見つけてください。
	//なお、彼の祖父は十分裕福であり、お年玉袋は十分大きかったものとします。

	//制約
	//1 ≤ N ≤ 2000
	//1000 ≤ Y ≤ 2×10^7
	//N は整数である。
	//Y は 1000 の倍数である。

	//入力
	//N Y

	//出力
	//N 枚のお札の合計金額が Y 円となることがありえない場合は、-1 -1 -1 と出力せよ。
	//N 枚のお札の合計金額が Y 円となることがありうる場合は、
	//そのような N 枚のお札の組み合わせの一例を「10000 円札 x 枚、5000 円札 y 枚、1000 円札 z 枚」として、
	//x、y、z を空白で区切って出力せよ。複数の可能性が考えられるときは、そのうちどれを出力してもよい。

	var N, sum int
	fmt.Scan(&N, &sum)

	i := 0
	for {
		sum_i := i * 1000
		if sum_i == sum && i == N {
			fmt.Println(0, 0, i)
			return
		}
		if sum_i < sum && i < N {
			j := 0
			for {
				sum_j := sum_i + j*5000
				if sum_j == sum && (i+j) == N {
					fmt.Println(0, j, i)
					return
				}
				if sum_j < sum && (i+j) < N {
					k := 1
					for {
						sum_k := sum_j + k*10000
						if sum_k == sum && i+j+k == N {
							fmt.Println(k, j, i)
							return
						}
						if sum_k < sum && (i+j+k) < N {
							k++
						} else {
							break
						}
					}
					j++
				} else {
					break
				}
			}
		} else {
			break
		}
		i++
	}

	fmt.Println("-1 -1 -1")
}
