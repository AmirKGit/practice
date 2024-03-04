package main

import "fmt"

func teskarison(num int) int {
	if num < 10 {
		return num
	}
	return (num % 10) * res(num) + teskarison(num / 10)
}

func res(num int) int {
	if num < 10 {
		return 1
	}
	return 10 * res(num / 10)
}

func main() {
	var a int
	fmt.Println("enter num")
	fmt.Scan(&a)
	fmt.Println(teskarison(a))
}
