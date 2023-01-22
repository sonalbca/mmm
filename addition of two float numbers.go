//addition of two float numbers
package main

import "fmt"

func main() {
	var num1, num2 float32
	fmt.Println("enter first float numbers")
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	fmt.Println("addition of two integer numbers is :", num1+num2)
}
