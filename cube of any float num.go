package main

import "fmt"

func main() {
	fmt.Println("enter any float number")
	var num float32
	fmt.Scanln(&num)
	fmt.Println("cube of float numbers is :", num*num*num)

}
