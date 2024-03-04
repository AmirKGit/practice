// .3 xonali son berilgan. Berilgan sonning raqamlari toqmi, yo’qmi aniqlash kerak!
package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("enter 3 digit number")
	fmt.Scan(&a)

	if a < 100 || a > 999 {
		fmt.Println("wrong")
		return
	}

	hundred := a / 100
	ten := (a / 10) % 10
	one := a % 10
	if hundred%2 == 1 || ten%2 == 1 || one%2 == 1 {
		fmt.Println("bor")
	} else {
		fmt.Println("yo'q")
	}
	fmt.Println(hundred, ten, one)
}
