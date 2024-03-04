// 7. A ,B,C sonlari berilgan. A ni qiymati C ga , C ni qiymati B ga va B ni qiymatini A ga almashtiring.

package main

import(
	"fmt"
)

func main(){
	var A, B, C,F int

	fmt.Print("A ni kiriting: ")
	fmt.Scan(&A)
	fmt.Print("B ni kiriting: ")
	fmt.Scan(&B)
	fmt.Print("C ni kiriting: ")
	fmt.Scan(&C)

	F=A
	A=C
	C=B
	B=F

	fmt.Println(A,C,B)


}