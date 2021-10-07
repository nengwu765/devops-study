package main

import "fmt"

// 首字母小写，所以是一个包私有的接口
type animal interface {
	// 这里可以有任意多个方法，不过我们一般建议是小接口
	// 即接口里面不会有很多方法
	// 方法声明不需要 func 关键字

	//Sound()
	Eat()
}

type animal2 struct {
	Name string
	Eat func()
}

type Duck interface {
	Swim()
}

// 判断是否实现了接口，如果没有实现接口编译会报错
// var _ animal = (*cat)(nil) // 把nil转化成*cat类型 赋值后即丢弃
var _ animal = &cat{} // 没有实现所有方法，所有编译报错

type cat struct {
	Name string
}

type cat2 struct {
	animal2
}

func (c *cat2) Eat()  {
	fmt.Println("fish")
}

func (c *cat) Eat()  {
	fmt.Println("fish")
}
