package main

import (
	"fmt"
)

func main()  {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = add(s, 4, 5)
	fmt.Printf("%#v \n", s)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = add(s, 1, 9)
	fmt.Printf("%#v \n", s)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = add(s, 6, 13)
	fmt.Printf("%#v \n", s)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = myDelete(s, 2)
	fmt.Printf("%#v \n", s)

	// 结果应该是9, 2, 4, 7, 13
	s = myDelete(s, 0)
	fmt.Printf("%#v \n", s)

	// 结果应该是9, 2, 4, 7
	s = myDelete(s, 4)
	fmt.Printf("%#v \n", s)
}

func add(s []int, index int, value int) []int {
	// 需要判断index的合法性，避免数组越界
	s1 := s[:index]
	s2 := s[index:]
	// 注意不可复用s1切片进行append操作，因为子切片和父切片很可能共享了底层数组（没有扩容，则它们的机构也就没有发生变化）
	//s1 = append(s1, value)  // 操作后，很可能改变了 s[index] 位置的值，则s2的第一个下标值[0]也就随之改变了
	//s1 = append(s1, s2...)
	//for i, i2 := range s1 {
	//	fmt.Printf("%d ==> %d\n", i, i2)
	//}
	//for i, i2 := range s2 {
	//	fmt.Printf("%d --> %d\n", i, i2)
	//}
	//os.Exit(33)
	var newS []int
	newS = append(newS, s1...)
	newS = append(newS, value)
	newS = append(newS, s2...)
	return newS
}

func myDelete(s []int, index int) []int {
	s1 := s[:index]
	s2 := s[index+1:]
	s1 = append(s1, s2...)
	return s1
}