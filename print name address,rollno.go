package main

import "fmt"

func main() {
	var name, address, rollno string
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	fmt.Println("enter your address")
	fmt.Scanln(&address)
	fmt.Println("enter rollno")
	fmt.Scanln(&rollno)
	fmt.Println("your name is :", name)
	fmt.Println("your address is :", address)
	fmt.Println("your rollno is :", rollno)
}
