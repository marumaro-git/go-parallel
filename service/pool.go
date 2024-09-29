package service

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// NOTE: sync.Poolは一時的なオブジェクトを保持するための機能
// sync.PoolはGCが実行されるときに一時的なオブジェクトを解放する
// 適切な使い方をすると、メモリの再利用ができるため、パフォーマンスが向上する

func connectService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmService() *sync.Pool {
	p := &sync.Pool{
		New: connectService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func Sample() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		server, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			panic(err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("failed to accept: %v", err)
				continue
			}
			connectService()
			fmt.Println("Connected")
			conn.Close()
		}
	}()
	return &wg
}

func PoolSample() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmService()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("failed to accept: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Println("Connected")
			// Putはdeferで行った方が確実
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}