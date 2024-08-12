package main

import (
	"fmt"
	"goroutines/handler"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(200000)
	var mutex sync.Mutex
	goroutine := handler.NewGoroutinesHandler()
	// go goroutine.Helper_Deadlock("hello", &wg)
	// go goroutine.Helper_Deadlock("welcome", &wg)
	// go goroutine.Helper_Deadlock("Bengaluru", &wg)
	var counter int32
	for i := 0; i < 100000; i++ {
		go goroutine.Helper_Race_Condition_1(&counter, "Helper_Race_Condition_1 :", &wg, &mutex)
	}

	for i := 0; i < 100000; i++ {
		go goroutine.Helper_Race_Condition_2(&counter, "Helper_Race_Condition_2", &wg)
	}

	wg.Wait()
	fmt.Println("Helper_Race_Condition_1 Counter : ", counter)
}
