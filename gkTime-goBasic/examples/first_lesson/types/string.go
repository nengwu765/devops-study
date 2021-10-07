package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("He said: \"Hello Go\"")

	// 长的，复杂的，比如：json
	fmt.Println(`hello
a b c
"xyz"`)

	fmt.Println(len("你好"))  // 6
	fmt.Println(utf8.RuneCountInString("你好")) // 2
	fmt.Println(len("你好ab")) // 8
	fmt.Println(utf8.RuneCountInString("你好ab")) // 4

	// 反正遇到计算字符个数的，比如名字长度、博客多长这种字符个数，记得使用：utf8.RuneCountInString
}
