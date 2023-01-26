package main

import "fmt"

func main() {
	p := 34
	q := 20
	// & (bitwise AND)
	result1 := p & q
	fmt.Println(result1)

	// | (bitwise OR)
	result2 := p | q
	fmt.Println(result2)

	// ^ (bitwise XOR)
	result3 := p ^ q
	fmt.Println(result3)

	// << (left shift)
	result4 := p << 1
	fmt.Println(result4)

	// >> (right shift)
	result5 := p >> 1
	fmt.Println(result5)

	// &^ (AND NOT)
	result6 := p &^ q
	fmt.Println(result6)
}
