package main

import "fmt"

func main() {
	fmt.Println("enter any number")
	var num, diff, cube, square int
	fmt.Scanln(&num)
	cube = num * num * num
	fmt.Println("cube of num is :", cube)
	square = num * num
	fmt.Println("square of num is :", square)
	diff = cube - square
	fmt.Println("diff of cube and square is :", diff)

}
