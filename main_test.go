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
		go func() {
			time.Sleep(100 * time.Millisecond) // Добавляем задержку
			fmt.Println(i)                     // Используем переменную i, захваченную по ссылке
			wg.Done()
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
}
