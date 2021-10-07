package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 创建一个包含三个元素，容量为4的切片
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 7)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	// 推荐通过make创建切片的方式： s := make([]int, 0, n)
	s3 := make([]int, 4) // 创建一个包含四个元素， 容量为4的切片
	// 等价于 s3 := make([]int, 4, 4)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	s3 = append(s3, 6)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	fmt.Printf("s3[2]: %d \n", s3[2])

	fmt.Println("====================")

	//subSlice()
	shareSlice()
}

func subSlice() {
	// 子切片，左闭右开原则，注意下标均从0开始
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))
}

func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199, 222, 444)
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
}
