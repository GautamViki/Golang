package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}

func main() {
	dataChan := make(chan int)

	go func() { // go-routine that's running on the background thread
		wg := sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() { // go-routine that's running on the background thread
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
			fmt.Println("lllllllllllllllll")
		}
		fmt.Println("jjjjjjjjjjjjjjjj")
		wg.Wait()
		close(dataChan)
	}()
	fmt.Println("ppppppppppppppp")
	for x := range dataChan {
		fmt.Printf("%d\n", x)
	}
	uuidWithHyphen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	ch := make(chan string)
    go pushToChannel(ch)
    for val := range ch {
        fmt.Println(val)
    }
}

func pushToChannel(ch chan string) {
    ch <- "a"
    ch <- "b"
    ch <- "c"
    close(ch)
}