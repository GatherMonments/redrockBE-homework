package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	c := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				c.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终计数:", c.Value())
}

// 预期输出:
// 启动 100 个协程,每个协程累加 10 次...
// 最终计数: 1000
// 提示:
// - 一定要用 defer mu.Unlock()
// - 对比:不加锁的情况下结果会不正确
