package main

import "fmt"

func main() {
	fmt.Println("enter any number")
	var num int
	fmt.Scanln(&num)

	fmt.Println("cube of num is :", num*num*num)

	fmt.Println("square of num is :", num*num)

}
