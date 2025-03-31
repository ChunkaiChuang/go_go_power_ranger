package main

import "fmt"

func main() {
	// 整數陣列
	/*
		var numbers [3]int
		numbers[0] = 15
		numbers[1] = 20
		numbers[2] = 25
		fmt.Println(numbers)
		fmt.Println(numbers[1] * 10)
		// 字串陣列
		var names [2]string = [2]string{"Ken", "Vivin"}
		fmt.Println(names)
		// 取得陣列長度
		fmt.Println(len(names))
		fmt.Println(len(numbers))
	*/
	// 巡迴陣列中的元素
	var grades [3]int
	var index int
	// 逐一輸入陣列的資料
	fmt.Println("請輸入成績：")
	for index = 0; index < len(grades); index++ {
		fmt.Scanln(&grades[index])
	}
	// 逐一取得陣列的資料
	var sum int
	for index = 0; index < len(grades); index++ {
		sum = sum + grades[index]
	}
	fmt.Println(sum / len(grades)) // 平均值

}
