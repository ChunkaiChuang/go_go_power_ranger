package main

// 計算 1+2+3+....+max的函式包裝
// func sum(max int) {
// 	var result int = 0
// 	var n int
// 	for n = 1; n <= max; n++ {
// 		result += n // result = result + n
// 	}
// 	fmt.Println(result)
// }
// func main() {
// 	sum(100) //...+100
// 	sum(50)  //...+50
// 	sum(10)  //...+10
// }

// 結束函式；return
// 執行return後，函式結束
// func show(msg string) {
// 	if msg == "Hello" {
// 		return
// 	}
// 	fmt.Println(msg)
// }

// // 呼叫show函式
// func main() {
// 	show("Hello")
// 	show("你好")
// }

//	func 函式名稱(參數列表)回傳值資料型態{
//		函式內部程式碼}
//
// return 回傳值，須符合定義中的資料型態}
// func mutiply(n1 int, n2 int) int {
// 	var result int = n1 * n2
// 	return result
// }
// func main() {
// 	var x int = mutiply(3, 4)
// 	fmt.Println(x)
// }

// 多個回傳值
//
//	func 函式名稱(參數列表)(資料型態,資料型態,... ){
//		函式內部程式碼
//		return 回傳值1,回傳值2,...
//	}
// func test() (int, string) {
// 	return 10, "Test"
// }
// func main() {
// 	var x int
// 	var s string
// 	x, s = test()
// 	fmt.Println(x, s)
// }
