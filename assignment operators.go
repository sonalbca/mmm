package main

import "fmt"

func main() {
	var p int = 45
	var q int = 50
	// '=' (simple assignment)
	p = q
	fmt.Println(p)

	// '+=' (add assignment)
	p += q
	fmt.Println(p)

	// '-=' (subtract assignment)
	p -= q
	fmt.Println(p)

	// '*=' (multiply assignment)
	p *= q
	fmt.Println(p)

	// '/=' (Division assignment)
	p /= q
	fmt.Println(p)

	// '%=' (modulus assignment)
	p %= q
	fmt.Println(p)

}
