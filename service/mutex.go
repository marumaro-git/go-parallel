package service

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// NOTE: Mutexは排他制御を行うための機能

func MutexSample1() {

	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incremented: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		// NOTE: panicが発生した場合、Unlockが呼ばれないため、deferを使ってUnlockを呼び出す
		defer lock.Unlock()
		count--
		fmt.Printf("Decremented: %d\n", count)
	}

	// インクリメント
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// デクリメント
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()

	fmt.Println("Done")
}

// NOTE: RWMutexは読み込みと書き込みを分けて排他制御を行うための機能
func RWMutexSample1() {

	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tREMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, &m), test(count, &m, &m))
	}
}
