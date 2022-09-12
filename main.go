package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/FedeSpeltini/go-routines/data"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}
	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("Duration: %dms\n", duration)
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	data.FinishedBook(id, m)
	delay := rand.Intn(800)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	wg.Done()
}
