package main

import "fmt"

// 計算 1+2+3+....+max的函式包裝
func sum(max int) {
	var result int = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n // result = result + n
	}
	fmt.Println(result)
}
func main() {
	sum(100) //...+100
	sum(50)  //...+50
	sum(10)  //...+10
}
