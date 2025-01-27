package main

import "fmt"

// 指標: 儲存資料的記憶體位址
// 基本思考順序: 1. 宣告變數 2. 取得變數記憶體位址 3. 宣告指標變數 4. 指標變數指向記憶體位址
// func main() {
// 1. 建立存放資料變數
// var x int = 3
// fmt.Println("原來的資料:", x)

// 2. 取得變數記憶體位址: &變數名稱
// var x int = 3
// fmt.Println("資料的記憶體位址", &x)

// 3. 存放到指標變數
// var x int = 3
// 指標資料型態:*資料型態
// var xPtr *int = &x
// fmt.Println(xPtr)

// 4. 反解指標變數
// var x int = 3
// // 反解指標變數: *指標變數名稱
// var xPtr *int = &x
// fmt.Println("反解指標回來的資料", *xPtr)

// var word string = "Hello"
// fmt.Println(word)
// var wordPtr *string = &word
// fmt.Println(wordPtr)
// fmt.Println(*wordPtr)
}
// 結構基礎 Struct
// 結構: 存放其他資料欄位的容器
// 定義 > 實體化

// 定義結構
// type 結構名稱 struct {
// 		欄位名稱 資料型態
//      欄位名稱 資料型態...
// // }
// type Point struct {
// 	x int
// 	y int
// }

// 實體化結構
// 結構名稱{欄位資料, 欄位資料2...}
// 結構名稱{欄位名稱: 資料, 欄位名稱2: 資料2...}
// func main() {
// 	var p1 Point=Point{3, 4}
// 	var p2 Point=Point{y:2, x:1}
// }

// 欄位中的資料
//結構實體.欄位名稱
// type Point struct {
// 	x int
// 	y int
// }
// func main(){
// 	var p1 Point
// 	fmt.Println(p1.x, p1.y)
// }

// 資料型態: 結構是一種資料型態, 例如範例中的Point結構