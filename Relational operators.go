package main

import (
	"fmt"
)

func main() {
	p := 34
	q := 20
	//'=='(Equal to)
	result1 := p == q
	fmt.Println(result1)

	//'!='(Not equal To)
	result2 := p != q
	fmt.Println(result2)

	//'<'(Less Than)
	result3 := p < q
	fmt.Println(result3)

	//'>'(Greater Than)
	result4 := p > q
	fmt.Println(result4)

	//'>='(Greater Than Equal to)
	result5 := p >= q
	fmt.Println(result5)

	//'<='(Less Than Equal to)
	result6 := p <= q
	fmt.Println(result6)

}
