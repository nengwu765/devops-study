package main

import "fmt"

func main() {
	a1 := [3]int{2, 3, 4}
	fmt.Printf("a1: %v, len: %d, cap: %d \n", a1, len(a1), cap(a1))

	var a2 [3]int
	fmt.Printf("a2: %v, len: %d, cap: %d \n", a2, len(a2), cap(a2))

	//a1 = append(a1, a2)  数组不支持 append 操作

	fmt.Printf("a2[2]： %d \n", a2[2])
}
