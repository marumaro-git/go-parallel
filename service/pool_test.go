package service

import (
	"io"
	"net"
	"testing"
)

func init() {
	damonStart := PoolSample()
	damonStart.Wait()
	
	damonStart2 := Sample()
	damonStart2.Wait()
}

func PoolBenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("failed to dial: %v", err)
		}
		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("failed to read: %v", err)
		}

		conn.Close()
	}
}

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			b.Fatalf("failed to dial: %v", err)
		}
		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("failed to read: %v", err)
		}

		conn.Close()
	}
}
