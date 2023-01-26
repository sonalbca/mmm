package main

import "fmt"

func main() {
	fmt.Println("enter width")
	var width, area float32
	fmt.Scanln(&width)

	fmt.Println("enter length")
	var length float32
	fmt.Scanln(&length)

	area = width * length
	fmt.Println("area of rectangle is :", area)

}
