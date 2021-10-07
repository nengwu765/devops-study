package main

import "fmt"

func main()  {
	son := Son{
		Parent{},
	}

	son.SayHello()
	son.Parent.SayHello()

	son2 := Son{}
	// 都是调用父亲的方法，此处返回p.Name()都是 parent
	son2.SayHello()
	son2.Parent.SayHello()
}

type Parent struct {

}

// 一旦调用到父亲的方法时，p.Name()会被父亲接管，此时和组合子结构体就木有什么关系了
func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

// 这个Son的方法没有的化，则 p.Name()会被父亲接管，则p.Name()都是 parent
//func (s Son) SayHello() {
//	fmt.Println("I am " + s.Name())
//}


//func (s Son) SayHello() {
//	fmt.Println("I am " + s.Name())
//}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

func (s Son) Name() string {
	return "Son"
}



