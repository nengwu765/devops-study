package main

import (
	"errors"
	"fmt"
)

func main()  {
	r, e := fibonacci(10)
	if e != nil {
		panic(e)
	}
	fmt.Println(r)
}

func fibonacci(n int) (int, error)  {
	// 判断n的值为正整数
	if n < 0 {
		return 0, errors.New("n need >= 0")
	}

	// 执行递归
	switch n {
	case 0:
		return 0, nil
	case 1:
		return 1, nil
	default:
		n1, e1 := fibonacci(n-1)
		if e1 != nil {
			return 0, e1
		}
		n2, e2 := fibonacci(n-2)
		if e2 != nil {
			return 0, e1
		}
		return n1 + n2, nil
	}
}
