package main

import "fmt"

func main() {
	fmt.Println("enter any numbers")
	var num int
	fmt.Scanln(&num)
	fmt.Println("square of any numbers is :", num*num)
}
