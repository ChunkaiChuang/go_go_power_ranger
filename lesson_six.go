package main

import "fmt"

func main() {
	// 基本迴圈使用
	/*
		var x int = 0
		for x < 3 {
			fmt.Println(x)
			x++
		}
	*/
	/*
		var x int
		for x = 0; x < 3; x++ {
			fmt.Println(x)
		}
	*/
	// 計算1+2+3+...+50 的結果
	var result int = 0
	var x int = 1
	for x <= 50 {
		// fmt.Println("result:", result, " x:", x)
		result = result + x
		x++
	}
	fmt.Println(result)
}
