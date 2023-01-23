package main

import "fmt"

func main() {
	fmt.Println("enter radius")
	var radius, cm float32
	fmt.Scanln(&radius)
	cm = 2 * 22 / 7 * radius
	fmt.Println("circumference of circle is :", cm)
}
