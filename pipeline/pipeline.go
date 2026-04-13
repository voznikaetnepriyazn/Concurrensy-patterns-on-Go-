package pipeline

import (
	"fmt"
)

type Result struct {
	Value int
	Err   error
}

func generateWithError(nums ...int) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for _, n := range nums {
			if n < 0 {
				out <- Result{Err: fmt.Errorf("negative number: %d", n)}
				return
			}
			out <- Result{Value: n}
		}
	}()
	return out
}

func squareWithError(in <-chan Result) chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for r := range in {
			if r.Err != nil {
				out <- r
				continue
			}
			out <- Result{Value: r.Value * r.Value}
		}
	}()
	return out
}

func main() {
	for result := range squareWithError(generateWithError(2, -3, 4)) {
		if result.Err != nil {
			fmt.Printf("Error: %v\n", result.Err)
			continue
		}
		fmt.Printf("Result: %d\n", result.Value)
	}
}
