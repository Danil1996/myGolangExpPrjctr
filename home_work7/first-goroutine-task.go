package home_work7

import (
	"fmt"
	"math/rand"
	"time"
)

// Яка створює 3 горутини. Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину. Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал. Третя горутина виводить середнє значення на екран.

func FirstTaskRun() {
	rand.Seed(time.Now().UnixNano())

	randomNumbers := make(chan int)
	averageResult := make(chan float64)

	go func() {
		for i := 0; i < 10; i++ {
			randomNumber := rand.Intn(100)
			randomNumbers <- randomNumber
		}
		close(randomNumbers)
	}()

	go func() {
		sum := 0
		count := 0
		for num := range randomNumbers {
			sum += num
			count++
		}
		average := float64(sum) / float64(count)
		averageResult <- average
		close(averageResult)
	}()

	go func() {
		average := <-averageResult
		fmt.Printf("Average value: %.2f\n", average)
	}()

	time.Sleep(1 * time.Second)
}
