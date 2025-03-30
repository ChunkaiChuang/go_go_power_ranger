package main

import "fmt"

func main() {
	/*
		// 1.建立存放資料的變數
		var x int = 5
		fmt.Println("原來的資料", x)
		// 2.取得記憶體位址: &變數名稱
		fmt.Println("記憶體位址", &x)
		// 3.存放到指標變數, 注意指標資料型態: *資料型態
		var xPtr *int = &x
		fmt.Println("指標變數", xPtr)
		// 4.反解指標變數: *指標變數名稱
		fmt.Println("反解指標變數", *xPtr)
	*/
	var word string = "Hello"
	fmt.Println("原來的資料", word)
	var wordPtr *string = &word
	fmt.Println("指標變數", wordPtr)
	fmt.Println("反解指標變數", *wordPtr)

}
