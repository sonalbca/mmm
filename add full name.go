//concate of your full name

package main

import "fmt"

func main() {
	fmt.Println("enter first name")
	var firstname string
	fmt.Scanln(&firstname)
	fmt.Println("enter second name")
	var secondname string
	fmt.Scanln(&secondname)
	fmt.Println("your full name is :" + firstname + " " + secondname)

}
