package main

import (
	"concurrensy-patterns/fan-out_fan-in/fan_in"
	"concurrensy-patterns/fan-out_fan-in/fan_out"
	"fmt"
)

func main() {
	input := make(chan int)

	go func() {
		defer close(input)
		for i := 1; i < 10; i++ {
			input <- i
		}
	}()

	workers := fan_out.FanOut(input, 3)

	results := fan_in.FanIn(workers...)

	for result := range results {
		fmt.Printf("got result: %d\n", result)
	}
}
