package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()

	for s := range squares { //resulting channel
		fmt.Println(s)
	}
}

//раcпределение нагрузки на разные треды
