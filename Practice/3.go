
package main

import(
	"fmt"
)

func main(){
	var A, B, C int

	fmt.Print("A ni kiriting: ")
	fmt.Scan(&A)
	fmt.Print("B ni kiriting: ")
	fmt.Scan(&B)
	fmt.Print("C ni kiriting: ")
	fmt.Scan(&C)

	if A==B && B==C {
		fmt.Println("Teng tomonli")
	}else if A==B || B==C || A==C{
		fmt.Println("teng enli uchburchak")
	}else{
		fmt.Println("tengsiz")
	}


}