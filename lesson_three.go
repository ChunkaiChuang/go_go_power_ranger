package main

import "fmt"

func main() {
	// &變數名稱：取得變數的指摽（Pointer）
	/*
		var x int
		fmt.Print("請輸入一個數字：")
		fmt.Scanln(&x)
		fmt.Println(x)
	*/
	// 整合練習：書錄兩個數字，並輸出數字相乘的結果
	/*
		var x int
		var y int
		fmt.Print("請輸入第一個數字：")
		fmt.Scanln(&x)
		fmt.Print("請輸入第二個數字：")
		fmt.Scanln(&y)
		var result int = x + y
		fmt.Println("兩個數字相加的結果是：", result)
	*/
	var x int
	var y int
	fmt.Print("輸入兩個數字，用空格格開：")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Println("兩個數字相加的結果是：", result)
}
