// 8.Kvadratning tomoni a berilgan. Uning perimetri P = 4*a va yuzasi S= a*a aniglansin

package main

import(
	"fmt"
)

func main(){
	var a,P,S int
	fmt.Println("enter a")
	fmt.Scan(&a)

	P=4*a
	fmt.Println("PERIMETR",P)
	S=a*a
	fmt.Println("YUZASI",S)
}