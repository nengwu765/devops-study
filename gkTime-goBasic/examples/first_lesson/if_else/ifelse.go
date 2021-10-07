package main

import "fmt"

func main() {
	young(9)
	young(55)

	ifUsingNewVariable(10, 200)
	ifUsingNewVariable(100, 130)
}

func young(age int) {
	if age < 18 {
		println("I am a child")
	} else {
		fmt.Println("I a young man")
	}
}

func ifUsingNewVariable(start int, end int) {
	// 注意在 if 里表达式变量的作用域，只有if-else结构体内有效
	if distance := end - start; distance > 100 {
		fmt.Printf("距离太远了： %d \n", distance)
	} else {
		fmt.Printf("距离不远，来一趟： %d \n", distance)
	}
}
