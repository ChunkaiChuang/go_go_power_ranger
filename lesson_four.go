package main

import "fmt"

func main() {
	// 算術運算：+, -, *, /
	var x int
	x = 3*3 + 10
	fmt.Println(x)
	// 指定運算：+=, -=, *=, /=
	x = 5
	x -= 2
	fmt.Println(x)
	// 單元運算：++(加一), --(減一)
	x++
	fmt.Println(x)
	// 比較運算：>, <, >=, <=, ==, !=
	var result bool = 4 == 4
	fmt.Println(result)
	// 邏輯運算：&&(and), ||(or), !(反運算)
	// &&(and)：兩者皆為true，結果為true
	// ||(or)：兩者其中一個為true，結果為true
	// !(反運算)：true變false，false變true
	result = false && true
	fmt.Println(result)
}
