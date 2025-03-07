package main //可執行程式必須使用 main 封包
import "fmt" //載入內建 fmt 封包，用來做基本輸出輸入
func main() { //建立 main() 函式，程式的進入點
	// 執行程式時，從這邊開始
	// 輸出訊息到終端：fmt.Println(資料, .....)
	/*
		fmt.Println(3)    // 整數int
		fmt.Println(3.14) // 浮點數float64
		fmt.Println("中文") // 字串string
		fmt.Println(true) // 布林值 bool
		fmt.Println('a')  // 字符 rune
	*/
	// 變數使用
	var x int              // 宣告變數
	x = 4                  // 把4存到變數x
	fmt.Println(x)         // 使用變數x
	x = 10                 // 覆蓋資料
	fmt.Println(x)         // 使用變數x
	var f float64 = 3.1415 // 宣告變數f 順便放進資料
	fmt.Println(f)         // 使用變數f
	var s string = "Hello"
	fmt.Println(s)
	var test bool = true
	fmt.Println(test)
	var c rune = 'a'
	fmt.Println(c)
}
