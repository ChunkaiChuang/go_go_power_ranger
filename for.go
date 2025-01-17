package main

import "fmt"

func main() {
	// var x int = 0
	// for x < 3 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// var x int
	// for x = 0; x < 5; x += 2 {
	// 	fmt.Println(x)
	// }
	// var result int = 0
	// var x int = 1
	// for x <= 3 {
	// 	result = result + x
	// 	x++
	// }
	// fmt.Println(result)
	// var x int = 0
	// for x < 3 {
	// 	if x == 1 {
	// 		break // 強迫中斷迴圈
	// 	}
	// 	fmt.Println(x)
	// 	x++
	// }
	// var x int
	// for x = 0; x < 3; x++ {
	// 	if x == 1 {
	// 		continue //強迫回到開頭
	// 	}
	// 	fmt.Println(x)
	// }
	// break

	// var x int = 0
	// for x < 5 {
	// 	if x == 3 {
	// 		break
	// 	}
	// 	fmt.Println(x)
	// 	x++
	// }
	// continue
	var x int
	for x = 0; x < 5; x++ {
		if x == 2 {
			continue
		}
		fmt.Println(x)
	}
}
