// 1.Sizga Selsiy bo'yicha  manfiy bo'lmagan sonberilgan, bu Selsiy bo'yicha haroratni bildiradi.
// Siz Tselsiyni Kelvin va Farengeytga aylantirib, javobni ekranga chiqaring.
// Kelvin = Celsius + 273.15
// Fahrenheit = Celsius * 1.80 + 32.00
package main

import (
	"fmt"
)

func main(){
	var Celsius,Kelvin,Fahrenheit float64
	fmt.Println("Celius manfiy bo'lmagan kiriting")
	fmt.Scan(&Celsius)

	Kelvin=Celsius+273.15
	Fahrenheit=Celsius * 1.80 + 32.00

	
	fmt.Println("Kelvin",Kelvin)
	fmt.Println("Fahrenhei",Fahrenheit)
}