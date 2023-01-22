package main

import "fmt"

func main() {
	fmt.Println("enter principal")
	var principal, si float32
	fmt.Scanln(&principal)

	fmt.Println("enter time")
	var time float32
	fmt.Scanln(&time)

	fmt.Println("enter rate")
	var rate float32
	fmt.Scanln(&rate)

	si = (principal * time * rate) / 100

	fmt.Println("your simple interest is :", si)
}
