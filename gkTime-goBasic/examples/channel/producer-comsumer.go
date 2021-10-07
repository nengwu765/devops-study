package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 生产完成，关闭通道
	close(ch)
}

func consumer(id int, ch <-chan int, done chan<- bool) {
	// 起一个死循环协程
	for {
		fmt.Println(len(ch))
		//fmt.Println(cap(ch))
		value, ok := <-ch  // 三种情况，1-通道阻塞卡住了;2-通道有值，ok为true;3-通道关闭了，ok为false

		if ok {
			time.Sleep(1 * time.Second)
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed\n", id)
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 3)
	coNum := 2
	done := make(chan bool, coNum)

	for i := 0; i < coNum; i++ {
		go consumer(i, ch, done)
	}
	//time.Sleep(5 * time.Second)
	go producer(ch)

	// 阻塞两个消费者，避免主协程退出
	for i := 0; i < coNum; i++ {
		<-done
	}
}
