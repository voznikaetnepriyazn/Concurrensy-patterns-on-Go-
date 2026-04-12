package fan_in

import (
	"sync"
)

func FanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)

		go func(ch <-chan int) {
			defer wg.Done()

			for val := range ch {
				output <- val
			}
		}(input)
	}

	return output
}
