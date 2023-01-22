package main

import "fmt"

func main() {
	fmt.Println("enter any num")
	var num, square, cube, sum int
	fmt.Scanln(&num)
	square = num * num
	fmt.Println("square of num is ", square)
	cube = num * num * num
	fmt.Println("cube of num is ", cube)
	sum = square + cube
	fmt.Println("sum of square and cube of num is :", sum)

}
