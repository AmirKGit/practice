package main

import "fmt"

func main() {
	var a, temp int
	fmt.Println("enter num")
	fmt.Scan(&a)

	s := a
	var res int

	for a > 0 {
		temp = a % 10
		res = (res * 10) + temp
		a = a / 10
	}
	if s == res {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}
