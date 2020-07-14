package main

import "sync"

const M = 10
func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(M)
	for i := 0; i < M; i++ {
		go func() {
			defer wg.Done()
			mu.Lock() //多线程锁住
			m[i] = i
			mu.Unlock()
		}()
	}
    wg.Wait()
    println(len(m))
}

