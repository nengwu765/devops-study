package main

import "fmt"

func main() {
	name := "Tom"
	age := 17
	// 这个 Api 是返回字符串的，所以大多数我们都是用这个
	str := fmt.Sprintf("hello, I am %s, I am %d years old \n", name, age)
	println(str)

	// 这个是直接输出，一般简单程序 DEBUG 会用它输出到一些信息到控制台
	fmt.Printf("hello, I am %s, I am %d years old \n", name, age)


	replaceHolder()
}

type user struct {
	Name string
	Age  int
}

func replaceHolder() {
	u := &user{
		Name: "Tom",
		Age:  17,
	}

	// %v 只输出所有值
	// %+v 先输出字段名字，再输出该字段的值
	// %#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）
	fmt.Printf("v => %v \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)
}
