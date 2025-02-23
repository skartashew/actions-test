package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
			wg.Done()
		}(i) // Передаем значение i в горутину
	}
	}

	wg.Wait() // Ожидаем завершения всех горутин
}
