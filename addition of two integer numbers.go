// addition of two integer numbers
package main

import "fmt"

func main() {
	fmt.Println("enter first number")
	var num1, num2 int
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	fmt.Println("addition of two integer numbers is :", num1+num2)

}
