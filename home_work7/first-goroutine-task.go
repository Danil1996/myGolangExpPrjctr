package home_work7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FirstTaskRun() {
	rand.Seed(time.Now().UnixNano())

	randomNumbers := make(chan int)
	averageResult := make(chan float64)

	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			randomNumber := rand.Intn(100)
			randomNumbers <- randomNumber
		}

		close(randomNumbers)
	}()

	go func() {
		defer wg.Done()

		sum := 0
		count := 0

		for num := range randomNumbers {
			sum += num
			count++
		}

		average := float64(sum) / float64(count)

		averageResult <- average
	}()

	go func() {
		defer wg.Done()

		average := <-averageResult

		fmt.Printf("Average value: %.2f\n", average)
	}()

	wg.Wait()
}
