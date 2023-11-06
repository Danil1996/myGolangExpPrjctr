package home_work7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SecondTaskRun() {
	rand.Seed(time.Now().UnixNano())

	randomNumbers := make(chan int)
	extremes := make(chan struct {
		min, max int
	})

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			randomNumber := rand.Intn(100)
			randomNumbers <- randomNumber
		}
		close(randomNumbers)

		someValue := <-extremes
		fmt.Printf("Min number: %d\n", someValue.min)
		fmt.Printf("Max number: %d\n", someValue.max)
	}()

	go func() {
		defer wg.Done()
		min := 100
		max := 0

		for num := range randomNumbers {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}

		extremes <- struct {
			min, max int
		}{min, max}
		close(extremes)
	}()

	wg.Wait()
}
