package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	backNumber := 7 - int(input[0]-'0')

	if err := ioutil.WriteFile("output.txt", []byte(fmt.Sprint(backNumber)), 0644); err != nil {
		log.Fatal(err)
	}
}
