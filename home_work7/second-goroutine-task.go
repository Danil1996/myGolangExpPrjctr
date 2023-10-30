package home_work7

import (
	"fmt"
	"math/rand"
	"time"
)

//  Яка створює 2 горутини. Перша горутина генерує випадкові числа в заданому діапазоні й надсилає їх через канал у другу горутину. Друга горутина отримує випадкові числа і знаходить найбільше й найменше число, після чого надсилає його назад у першу горутину через канал. Перша горутина виводить найбільше й найменше числа на екран.

func SecondTaskRun() {
	rand.Seed(time.Now().UnixNano())

	randomNumbers := make(chan int)
	extremes := make(chan struct {
		min, max int
	})

	go func() {
		for i := 0; i < 10; i++ {
			randomNumber := rand.Intn(100)
			randomNumbers <- randomNumber
		}
		close(randomNumbers)
	}()

	go func() {
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

	extremeValues := <-extremes
	fmt.Printf("Min number: %d\n", extremeValues.min)
	fmt.Printf("Max number: %d\n", extremeValues.max)
}
