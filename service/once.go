package service

import "sync"

// NOTE: sync.Onceは一度だけ実行するための機能
func OnceSample() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
			// Do関数は1度だけ実行されるため2回目以降は実行されない
			once.Do(increment)
		}()
	}

	increments.Wait()
	println("Count is", count)
}
