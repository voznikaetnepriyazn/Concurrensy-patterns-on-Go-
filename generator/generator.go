package generator

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func GenerateLines(ctx context.Context, filename string) (<-chan string, <-chan error) {
	out := make(chan string, 32)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		file, err := os.Open(filename)
		if err != nil {
			errCh <- fmt.Errorf("open file: %w", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			select {
			case out <- scanner.Text():
				fmt.Printf("successfully sent")
			case <-ctx.Done():
				return
			}
		}

		if err := scanner.Err(); err != nil {
			select {
			case errCh <- fmt.Errorf("scan file: %w", err):
			case <-ctx.Done():
			}
		}
	}()

	return out, errCh
}

//раcпределение нагрузки на разные треды
