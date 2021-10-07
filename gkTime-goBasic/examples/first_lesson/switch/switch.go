package main

import "fmt"

func main()  {
	chooseFruit("apple")
	chooseFruit("pear")
	chooseFruit("orange")
}

func chooseFruit(fruit string)  {
	switch fruit {
	case "apple":
		fmt.Println("this is an apple")
	case "banana", "pear":
		fmt.Println("this is a banana or pear")
	default:
		fmt.Printf("I dont known %s\n", fruit)
	}
}
