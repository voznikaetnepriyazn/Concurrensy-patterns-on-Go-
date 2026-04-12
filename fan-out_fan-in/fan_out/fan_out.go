package fan_out

import "time"

func FanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		outputs[i] = worker(input)
	}

	return outputs
}

func worker(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for num := range input {
			result := complexComputation(num)
			output <- result
		}
	}()

	return output
}

func complexComputation(n int) int {
	time.Sleep(100 * time.Millisecond)
	return n * n
}
