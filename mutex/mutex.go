package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {

	w := sync.WaitGroup{}
	w.Add(3)
	go func() {
		m.Lock()
		w.Done()
		fmt.Println("go thread1")
		m.Unlock()
	}()

	go func() {
		for {
			if m.TryLock() {
				fmt.Println("go thread2")
				w.Done()
				break
			}
		}
		defer m.Unlock()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		m.Lock()
		defer m.Unlock()
		fmt.Println("go thread3")
		w.Done()
	}()

	w.Wait()
}
