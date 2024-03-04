package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var input string
	fmt.Println("Enter a number (up to 1000):")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	uzbekOnes := [...]string{"", "bir", "ikki", "uch", "to'rt", "besh", "olti", "yetti", "sakkiz", "to'qqiz"}
	uzbekTens := [...]string{"", "o'n", "yigirma", "o'ttiz", "qirq", "ellik", "oltmish", "yetmish", "sakson", "to'qson"}
	uzbekHundreds := [...]string{"", "yuz", "ming"}

	var result string
	if n == 0 {
		result = "nol"
	} else {
		hundreds := n / 100
		if hundreds > 0 {
			result += uzbekHundreds[hundreds] + " "
			n %= 100
		}

		tens := n / 10
		if tens > 1 {
			result += uzbekTens[tens] + " "
			n %= 10
		}

		if n > 0 {
			result += uzbekOnes[n]
		}
	}

	fmt.Println("Number in Uzbek Latin script:", result)
}
