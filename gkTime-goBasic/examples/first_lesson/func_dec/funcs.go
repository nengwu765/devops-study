package main

import "fmt"

func main() {
	a := fun0("Tom")
	fmt.Println(a)

	b, c := fun1("a", 18)
	fmt.Println(b)
	fmt.Println(c)

	_, d := fun2("a", "b")
	fmt.Println(d)

	fun4("Hello", 19, "CUICUI", "DaMing")
	s := []string{"AAA", "BBB"}
	fun4("hello", 20, s...)
}

func fun0(name string) string {
	return "Hello, " + name
}

func fun1(a string, b int) (int, string) {
	return 0, "你好"
}

// 可忽略age， name，直接返回  --  不推荐这种写法，可读性太差
func fun2(s string, b string) (age int, name string) {
	age = 19
	name = "Tom"
	return
	//return 19, "Tom"
	//return age, name  // 推荐
}

func fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "你好"
	return
}

func fun4(a string, b int, names ...string) {
	for _, name := range names {
		fmt.Printf("不定参数： %s \n", name)
	}
}
