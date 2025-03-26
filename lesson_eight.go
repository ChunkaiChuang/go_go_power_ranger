package main

import "fmt"

func test(msg string) {
	fmt.Println(msg)
}
func add(n1 int, n2 int) {
	var result int = n1 + n2
	fmt.Println(result)
}

// 計算1+2+3+...+max的總和
func sum(max int) {
	var result int = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n
	}
	fmt.Println(result)
}
func main() {
	// 基礎函式語法演練
	/*
		test("hey")
		test("go go power ranger")
		add(4, 3)
		add(-2, 5)
	*/
	//計算1+2+3+...+max的總和
	sum(100) //....+100
	sum(50)  //....+50
	sum(10)  //....+10

}
