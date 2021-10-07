package main

import "fmt"

func main() {
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法，以下这一句会编译报错
	//fake.Swim()
	fake.FakeSwim()

	// 转换成Fish
	td := Fish(fake)
	// 真的变成了鱼
	td.Swim()

	sFake := StrangeFake{}
	// 调用自己的方法
	sFake.Swim()

	td = Fish(sFake)
	// 真的变成了鱼
	td.Swim()
}

// 新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Println("我是山寨鱼，嘎嘎噶")
}

type StrangeFake FakeFish

func (f StrangeFake) Swim() {
	fmt.Println("我是华强北超级山寨鱼，嘎嘎噶")
}

type Fish struct {

}

func (f Fish) Swim()  {
	fmt.Println("我是一条鱼，假装自己是一只鸭子")
}