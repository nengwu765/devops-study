package main

import "fmt"

func main()  {
	//c1 := new(Concrete1)
	//var c1 = new(Concrete1)
	// 申明了一个指针，没有对应实例，会出现空指针异常
	var c1 *Concrete1
	fmt.Println(c1.Name)
}

// Swimming 会游泳的
type Swimming interface {
	Swim()
}

// Duck 会游泳，所以组合了Swimming
type Duck interface {
	Swimming
}

type Base struct {
	Name string
}

type Concrete1 struct {
	Base
}

type Concrete2 struct {
	*Base
}

func (c Concrete1) SayHello()  {
	// c.Name 直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)

	// 调用了被组合的
	c.Base.SayHello()
}

// 方法接收器，可以是结构体或者指针
func (b *Base) SayHello()  {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}

