package main

import "fmt"

var Global = "全局变量"

var local = "包变量"

var (
	First = "abc"
	// 不注明类型，默认为 int
	second int32 = 16
)

func main()  {
	a := 13
	fmt.Println(a)
	a = 14
	fmt.Printf("%T \n", a)

	var b int32 = 24
	fmt.Printf("%T \n", b)

	var d int
	fmt.Println(d)

}
