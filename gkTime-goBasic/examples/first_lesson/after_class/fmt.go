package main

import "fmt"

func main() {
	//nums := []int {1, 3, 5, 6}
	//index := searchInsert(nums, 5)
	//fmt.Println(index)

	s := printNumWith2(66.66666)
	fmt.Println(s)

	x := printBytes([]byte("这是一个Demo"))
	fmt.Println(x)
	fmt.Printf("%%v%v", x)
}

func searchInsert(nums []int, target int) int {
	return getI(nums, target)
}

func getI(nums []int, target int) int {
	return getIndex(len(nums), func(i int) bool {return nums[i] >= target})
}

func getIndex(len int, f func(int) bool) int {
	i, j := 0, len
	for i < j {
		mid := (i+j)/2
		println(mid)
		if f(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

// 输出两位小数
func printNumWith2(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func printBytes(data []byte) string {
	h := fmt.Sprintf("%x", data)
	return h
}
