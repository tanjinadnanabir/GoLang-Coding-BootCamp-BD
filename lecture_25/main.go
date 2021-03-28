package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// booring("boring job...")

	var wg sync.WaitGroup

	wg.Add(1)
	wg.Add(1)
	//wg.Add(2)

	go cooking("rice", &wg)
	go cooking("curry", &wg)

	// time.Sleep(time.Second * 20)
	wg.Wait()
}

// func booring(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(msg, i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func cooking(msg string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
