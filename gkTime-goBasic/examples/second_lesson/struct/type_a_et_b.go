package main

import "fmt"

func main() {
	var n = fakeNews{
		Name: "hello",
	}
	n.Report()
}

type news struct {
	Name string
}

func (d news) Report() {
	fmt.Println("I am news: " + d.Name)
}

// 只是一个别名
type fakeNews = news
