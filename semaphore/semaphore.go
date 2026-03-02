package semaphore

import (
	"fmt"
	"time"
)

const NumWorkers = 3 //ограничение количества воркеров

func main() {
	tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8"}

	sem := make(chan struct{}, NumWorkers)  //семафор указывает на то, что очередь заполнена, 3 горутины в работе максимально, одна закончила работу - вторая сразу запускается
	done := make(chan struct{}, len(tasks)) //сигнал о завершении горутиной работы

	//при заполнении канала висим
	//в единицу времени работают 3 горутины
	for _, v := range tasks {
		go func(task string) {
			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Printf("Worker %s started\n", task)
			time.Sleep(3 * time.Second)
			fmt.Printf("Worker %s done\n", task)

			done <- struct{}{}
		}(v)
	}

	for range tasks {
		<-done
	}

}
