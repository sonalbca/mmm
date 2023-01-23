package main

import "fmt"

func main() {
	fmt.Println("enter base")
	var base, height, AreaOfTriangle float32
	fmt.Scanln(&base)

	fmt.Println("enter height")
	fmt.Scanln(&height)

	AreaOfTriangle = (base * height) / 2
	fmt.Println("area of triangle is :", AreaOfTriangle)
}
