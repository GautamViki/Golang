package handler

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type routine struct{}

func NewGoroutinesHandler() *routine {
	return &routine{}
}

// creating deadlock conditions
// 1.remove defer wg.Done()
// 2.accept WaitGroup as a referece not value EX. wg sync.WaitGroup
func (r *routine) Helper_Deadlock(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(str, i)
	}
}

// To prevent race condition using mutex
// use shared resource in between mutex.Lock() and mutex.Unlock()
func (r *routine) Helper_Race_Condition_1(counter *int32, str string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*counter = *counter + 1
	mutex.Unlock()
	defer wg.Done()
}

// two prevent race condition using atomic function and pass the counter value with incremented value
func (r *routine) Helper_Race_Condition_2(counter *int32, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(counter, 1)
	// *counter++
}
