package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {

	// 解析命令行参数
	flag.Parse()

	//对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
	//指针变量的值是指针地址。
	//对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。

	// 输出命令行参数
	fmt.Println(mode)
	fmt.Println(*mode)
}