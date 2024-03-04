// 9.Berilgan x, y, z sonlari ichidan eng kichigini aniqlash algoritmini tuzing.
package main

import(
	"fmt"
)

func main(){
	var x, y, z int

	fmt.Println("x ni kiriting: ")
	fmt.Scan(&x)
	fmt.Println("y ni kiriting: ")
	fmt.Scan(&y)
	fmt.Println("z ni kiriting: ")
	fmt.Scan(&z)

	if x<y && x<z {
		fmt.Println(x)
	}else if y<x && y<z {
		fmt.Println(y)
	}else {
		fmt.Println(z)
	}
}