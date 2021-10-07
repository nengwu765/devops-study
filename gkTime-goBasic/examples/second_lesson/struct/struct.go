package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你声明这样的时候， Go 已经帮你分配号内存
	// 不用担心空指针的问题，因为它压根不是一个指针
	var duck4 ToyDuck
	duck4.Swim()

	// duck5 是一个指针
	var duck5 *ToyDuck
	// duck5.Swim()   // panic 空指针异常
	fmt.Printf("%T\n", duck5)

	duck6 := ToyDuck{
		Color:  "yellow",
		Prince: 2,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用（一是不便于阅读，而是接口体字段变更需要同步修改）
	duck7 := ToyDuck{"red", 9}
	duck7.Swim()

	duck8 := ToyDuck{}
	duck8.Color = "black"
	duck8.Swim()

	duck9 := &ToyDuck{}
	duck9.Color = "123"
	duck9.Swim()

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color  string
	Prince uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，有过一群鸭，我是 %s， %d 一只\n", t.Color, t.Prince)
}

