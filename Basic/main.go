package main

import "fmt"

func main() {
	const number = 50
	var name = "Bob"
	var another_number int

	fmt.Println("Please give other number")
	fmt.Scan(&another_number)
	fmt.Printf("namelol %v %v\n", another_number, name)
	fmt.Printf("we have: %v\n", number-another_number)
}
