package main

import "fmt"

func main() {
	// break
	/*
		var x int = 0
		for x < 5 {
			if x == 3 {
				break
			}
			fmt.Println(x)
			x++
		}
	*/
	// continue
	/*
		var x int
		for x = 0; x < 5; x++ {
			if x%2 == 0 {  // if x is even, skip the rest of the loop
				continue
			}
			fmt.Println(x)
		}
	*/
	//實際範例，時序輸入數字，計算總和，直到輸入0結束
	var result int = 0
	for true {
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		result += n
	}
	fmt.Println("總計：", result)
}
