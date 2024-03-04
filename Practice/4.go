package main

import(
	"fmt"
)

// 4.n butun soni berilgan :
//  Agar n  3 va 5 ga boʻlinadigan boʻlsa "FizzBuzz".
// Agar n 3 ga bo'linadigan bo'lsa "Fizz".
// Agar n 5 ga bo'linadigan bo'lsa "Buzz".
// Agar yuqoridagi shartlarning hech biri to'g'ri bo'lmasa n sonini chiqaring.


func main(){
	var n int;
	fmt.Println("n sonini kirit ")
	fmt.Scan(&n)

	if n%3==0 && n%5==0{
		fmt.Println("FizzBuzz")
	}else if n%5==0 {
		fmt.Println("Buzz")
	}else if n%3==0 {
		fmt.Println("Fizz")
	}else{
		fmt.Println(n)
	}
}