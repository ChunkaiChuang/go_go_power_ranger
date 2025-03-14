package main

import "fmt"

func main() {
	// 基本語法練習
	/*
		var x int
		fmt.Print("請輸入 100: ")
		fmt.Scanln(&x)
		if x == 100 {
			fmt.Println("幹得漂亮！！")
		} else {
			fmt.Println("你眼瞎嗎！！")
		}
	*/
	// 簡易情境：ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Print("請問想領多少錢: ")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("太少了啦！！")
	} else if money <= 100000 {
		fmt.Println("喔虧！！")
	} else {
		fmt.Println("太多了啦！！")
	}
}
