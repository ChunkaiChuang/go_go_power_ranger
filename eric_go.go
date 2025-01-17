package main

func main() {
	// 建置：go build 程式檔案名稱
	// 執行程式：輸入執行檔名

	// fmt.Println(3)    // 整數int
	// fmt.Println(3.14) // 浮點數float64
	// fmt.Println("中文") // 字串string
	// fmt.Println(true) // 布林值 bool
	// fmt.Println('a')  // 字符 rune

	// 變數使用
	// var x int // 宣告變數
	// x = 4
	// fmt.Println(x)
	// x = 10 // 覆蓋資料
	// fmt.Println(x)
	// fmt.Scanln(&x)

	// fmt.Println(3, "Hello", true)
	// // &變數名稱：取得變數的指標(Ｐointer)
	// var x int
	// fmt.Print("請輸入一個整數：")
	// fmt.Scanln(&x)
	// fmt.Println(x + 1)

	// var x int
	// var y int
	// fmt.Print("請輸入第一個數字： ")
	// fmt.Scanln(&x)
	// fmt.Print("請輸入第二個數字： ")
	// fmt.Scanln(&y)
	// var result = x + y
	// fmt.Println("結果是：", result)

	// var x int
	// var y int
	// fmt.Print("請輸入兩個數字，用空格隔開： ")
	// fmt.Scanln(&x, &y)
	// var result = x + y
	// fmt.Println("結果是...：", result)

	// if 布林值{
	// 	若布林值為true，執行程式區塊
	// }else{
	// 	若布林值為false，執行程式區塊
	// }

	// var x int
	// fmt.Print("請輸入100: ")
	// fmt.Scanln(&x)
	// if x == 100 {
	// 	fmt.Print("漂亮！！做得好！！")
	// } else {
	// 	fmt.Println("錯了啦！臭87")
	// }

	// var money int
	// fmt.Println("請問想多少錢")
	// fmt.Scanln(&money)
	// if money < 100 {
	// 	fmt.Println("Too Few")
	// } else if money <= 100000 {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("Too Much")
	// }
	// fmt.Println("執行完畢")

	// for 迴圈
	// for 布林值{
	// 	若布林值為true，執行程式區塊
	// 	執行完程式區塊後，再次檢查布林值
	// }
	// var x int = 0
	// for x < 3 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for 初始化命令;布林值;重複執行後的命令{
	// 	若布林值為true，執行程式區塊
	// 	執行完程式區塊後，再次檢查布林值
	// }
	// var x int
	// for x = 0; x < 3; x++ {
	// 	fmt.Println(x)
	// }

}
