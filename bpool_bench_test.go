package bpool

import (
	"io"
	"os"
	"sync"
	"testing"
)

func BenchmarkBytePool(b *testing.B) {
	bp := NewBytePoolCap(500, 1024, 1024)
	for i := 0; i < b.N; i++ {
		opBytePool(bp)
	}
}

func BenchmarkSyncPool(b *testing.B) {
	p := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024, 1024)
		},
	}
	for i := 0; i < b.N; i++ {
		opSyncPool(p)
	}
}

func BenchmarkRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opRaw()
	}
}

func opBytePool(bp *BytePoolCap) {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func(bp *BytePoolCap) {
			buffer := bp.Get()
			defer bp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(bp)
	}
	wg.Wait()
}

func opSyncPool(sp *sync.Pool) {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func(sp *sync.Pool) {
			buffer := sp.Get().([]byte)
			defer sp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(sp)
	}
	wg.Wait()
}

func opRaw() {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			buffer := make([]byte, 1024, 1024)
			mockReadFile(buffer)
			wg.Done()
		}()
	}
	wg.Wait()
}

func mockReadFile(b []byte) {
	f, _ := os.Open("./bpool.go")
	for {
		n, err := io.ReadFull(f, b)
		if n == 0 || err == io.EOF {
			break
		}
	}
}
