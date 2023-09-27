package main

import (
	"fmt"
	"time"
)

const s string = "constant"

func main() {
	//fmt.Println("hello world")
	// fmt.Println("1+1 =", 1+1)
	// fmt.Println("7.0/3.0 =", 7.0/3.0)
	// var a = "initial"
	// fmt.Println(a)
	// var b, c int = 1, 2
	// fmt.Println(b, c)
	// var e int // 0
	// fmt.Println(e)
	// f := "apple" // var f = "apple"
	// fmt.Println(f)

	// const n = 500000000
	// const d = 3e20 / n
	// fmt.Println(s)
	// fmt.Println(d)
	// fmt.Println(int64(d))
	// fmt.Println(math.Sin(n)) // サインを求める

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the Weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a bool")
		case int:
			fmt.Println("Im an int")
		default:
			fmt.Println("Dont know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
