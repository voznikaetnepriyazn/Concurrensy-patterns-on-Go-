package donechan

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) //пустaя структура не занимает места
	//в него прилетают значения

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("gorounite 1 stopping")
				return
			default:
				fmt.Println("goroutine 1 working...")
				time.Sleep(500 + time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("goroutine 2 stopping")
				return
			default:
				fmt.Println("goroutine 2 working...")
				time.Sleep(500 + time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	/*instead of string 40 -
	done <- struct{}{} - only one goroutine will's been stopped*/
	close(done)
	time.Sleep(100 * time.Millisecond)
}
