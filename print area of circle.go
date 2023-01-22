package main

import "fmt"

func main() {
	var Areaofcircle, radius float32
	fmt.Println("enter radius")
	fmt.Scanln(&radius)
	Areaofcircle = 22 / 7 * (radius * radius)
	fmt.Println("area of circle is :", Areaofcircle)

}
