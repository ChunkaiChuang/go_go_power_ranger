package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person = Person{"嘿嘿", 22}
	fmt.Println(p1.name, p1.age)
	var p2 Person = Person{age: 33, name: "哈哈"}
	p2.name = "呵呵"
	fmt.Println(p2.name, p2.age)
}
