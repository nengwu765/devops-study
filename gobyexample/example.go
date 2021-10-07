// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以
// 使用带一个 `default` 子句的 `select` 来实现_非阻塞_ 的
// 发送、接收，甚至是非阻塞的多路 `select`。

package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)

	// 我们可以在 `default` 前使用多个 `case` 子句来实现
	// 一个多路的非阻塞的选择器。这里我们试图在 `messages`
	// 和 `signals` 上同时使用非阻塞的接受操作。
	go func() {
		for {
			msg := <-messages
			fmt.Println("received message", msg)
			return
		}
	}()
	time.Sleep(time.Second * 1)

	// 一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	//messages <- "dddd"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

}
