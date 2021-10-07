package main

import "fmt"

func vals(a int, b int) (int, int) {
	return  a+1, b+1
}

func main() {
	a,b := vals(4, 5)
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals(6, 7)
	fmt.Println(c)
}
