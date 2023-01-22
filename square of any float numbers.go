package main

import "fmt"

func main() {
	fmt.Println("enter any float numbers")
	var num float32
	fmt.Scanln(&num)

	fmt.Println("square of any float numbers is :", num*num)

}
