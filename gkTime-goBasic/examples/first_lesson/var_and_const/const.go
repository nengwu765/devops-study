package main

import "fmt"

const internal = "包内可访问"
const (
	External = "包外可访问"
	aa = "新分组哈哈"
)

func main()  {
	const a = "你好"
	fmt.Println(a)
}


