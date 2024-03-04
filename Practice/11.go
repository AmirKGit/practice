package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)

	if y == 1918 {
		fmt.Println("26.09.", y)
	} else if y >= 1919 {
		if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
			fmt.Println("12.09.", y)
		} else {
			fmt.Println("13.09.", y)
		}
	} else if y <= 1917 {
		if y%4 == 0 {
			fmt.Println("12.09.", y)
		} else {
			fmt.Println("13.09.", y)
		}
	}
}
