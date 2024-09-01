package main

import (
	"fmt"
)

func printOddNumber() {
	i := 0

	for i < 10 {
		if i%2 != 0 {
			fmt.Println(i)
		}

		i++
	}
}

func printEvenNumber() {
	i := 0

	for i < 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}

		i++
	}
}

func main() {
	go printEvenNumber()
	go printOddNumber()
}
