package main

import "fmt"

func main() {
	fmt.Printf("%v\n", byte(0&1))
	fmt.Printf("%v\n", byte(1&1))
	fmt.Printf("%v\n", byte(2&1))
	fmt.Printf("%v\n", byte(3&1))

	fmt.Printf("%v\n", byte(1>>8))

}
