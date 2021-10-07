package main

import (
	"fmt"
)

func main() {
	// 创建一个预估容量是2的map
	// 通过make为slice、map或chan初始化并返回引用（T）。和new的区别：初始化一个指向类型零值的指针（*T）
	// slice： slice, len, cap
	// map: map, cap
	// channel: channel, cap
	m := make(map[string]string, 2)

	// 木有指定预估容量
	m1 := make(map[string]string)

	channel := make(chan int)
	channel1 := make(chan int, 1)
	channel2 := make(chan int, 10)

	go func() {
		//for i := 0; i < 10; i++ {
			channel2 <- 10
			//channel <- i
		//}
		//close(channel2)
	}()
	//time.Sleep(3 * time.Second)

	fmt.Println(len(channel))
	fmt.Println(len(channel1))
	fmt.Println(len(channel2))

	//select {
	////case <-channel:
	////	fmt.Println(<-channel)
	//case <-channel2:
	//	fmt.Println(<-channel2)
	//}

	i := <- channel2
	fmt.Println(i)


	// 直接初始化
	m2 := map[string]string{
		"Tom": "Jerry",
	}

	m["hello"] = "world"
	m1["hello"] = "world"
	m2["hello"] = "world"

	val := m["hello"]
	fmt.Println(val)

	// 再次取值，使用两个返回值，后面的ok会告诉你map有没有这个key值
	val, ok := m2["xxx"]
	if !ok {
		fmt.Println("key xxx not found")
	}

	for key, val := range m {
		fmt.Printf("%s => %s\n", key, val)
	}

}
