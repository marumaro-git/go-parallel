package service

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: WaitGroupはgoroutineの完了を待つための機能

func WaitGroupSample1() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("World")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("Done")
}

func WaitGroupSample2() {

	// NOTE: Addの呼び出しはgoroutineの外で行う（goroutine内で行うと、Addの呼び出しを待たずに次の行が実行されるため）
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %d!\n", id)
	}
	var wg sync.WaitGroup
	const numGreeters = 5
	wg.Add(numGreeters)

	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
