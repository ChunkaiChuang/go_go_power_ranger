package main

import "fmt"

func add(xPtr *int) {
	*xPtr = *xPtr + 10
	// fmt.Println("Add Function", *xPtr)
}

func main() {
	var a int = 10
	add(&a) // Pass by Reference
	fmt.Println("Main Function", a)
	// 和使用者要求輸入，運用到指標參數 Pass by Pointer
	var msg string
	var msgPtr *string = &msg
	fmt.Scanln(msgPtr) // 傳遞字串變數的指標(記憶體位址)
	fmt.Println(msg)
}

/*
func add(x int) {
	x = x + 10
	fmt.Println("Add Function", x)
}

func main() {
	var a int = 10
	add(a) // Pass by Value
	fmt.Println("Main Function", a)

}
*/
