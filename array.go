package main

import "fmt"

// 陣列: 按照順序,存放多個相同型態資料的容器
// 陣列資料型態: [長度]資料型態
// 貞烈的操作: 1.建立陣列 2.給定資料 3.巡迴陣列
func main() {
	// // 整數陣列 --> 預設0
	// var numbers [3]int
	// numbers[0] = 15
	// numbers[1] = 20
	// numbers[2] = 25
	// fmt.Println(numbers)
	// fmt.Println(numbers[1])

	// // 字串陣列 --> 預設空字串
	// var names [2]string = [2]string{"Eric", "Amy"}
	// fmt.Println(names)
	// // 取得陣列長度
	// fmt.Println(len(numbers))
	// fmt.Println(len(names))

	// 巡迴陣列中的資料
	var grades [4]int = [4]int{60, 90, 75, 10}
	var sum int
	var index int
	for index = 0; index < len(grades); index++ {
		// fmt.Println(grades[index])
		sum += grades[index]
	}
	fmt.Println(sum / len(grades)) // 平均值
}
