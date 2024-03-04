// Sizga to'g'ri to'rtburchakning tomonlari a va b beriladi, siz uning yuzasi va perametrini topib quyidagi shartga tekshirishingiz kerak bo'ladi.

// Agar yuzasi perimetridan katta bo'lsa yuzasini aks holda peremetrini chiqaring,

package main

import(
	"fmt"
)

func main(){
	var a,b int
	fmt.Println("enter a,b")
	fmt.Scan(&a,&b)

	P:= 2 * (a + b)
	S:=a*b
	if P>S {
		fmt.Println("P",P)
	}else{
		fmt.Println("S",S)
	}
}