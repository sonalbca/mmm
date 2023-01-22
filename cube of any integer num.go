package main

import "fmt"

func main() {
	fmt.Println("enter any number")
	var num int
	fmt.Scanln(&num)
	fmt.Println("cube of any integer number is :", num*num*num)

}
