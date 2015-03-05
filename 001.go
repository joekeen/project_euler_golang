package main

import "fmt"
import "time"

func main() {
	startTime := time.Now()
	sum := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	defer timeTrack(startTime)
	fmt.Println("Problem 1: The sum of all the multiples of 3 or 5 below 1000 is", sum)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("took", elapsed)
}
