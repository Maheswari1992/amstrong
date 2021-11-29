package main

import "fmt"

func main() {
	var number, remainder, temp int

	result := 0

	fmt.Println("Enter any positive integer : ")
	fmt.Scanln(&number)

	temp = number
	for {
		remainder = number % 10
		remainder = remainder * remainder * remainder
		result = result + remainder
		number = number / 10

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == result {
		fmt.Println(temp, "is a Palindrome")
	} else {
		fmt.Println(temp, "is not a Palindrome")
	}

}
