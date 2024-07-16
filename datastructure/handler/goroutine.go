package handler

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread managed by the Go runtime.
To start a new goroutine, you use the 'go' keyword followed by a function call.
The function will run concurrently with the calling function.
*/

func GoRoutineSay(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
