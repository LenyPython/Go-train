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

	var array [10]string
	var array2 = [10]string{}
	array[0] = array2[3]

	fmt.Printf("Array 1:%v\n", array)
	fmt.Printf("Array length:%v\n", len(array))

}
