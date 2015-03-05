package main

import "fmt"

func main() {
	sum := 0
	termOne := 1
	termTwo := 2
	for termOne < 4000000 {
		if termOne % 2 == 0 { sum += termOne }
		nextTerm := termOne + termTwo
		termOne = termTwo
		termTwo = nextTerm
	}
	fmt.Println("Problem 2: The sum of the even-valued terms of the Fibonacci sequence whos values do not exceed four million is", sum)
}
