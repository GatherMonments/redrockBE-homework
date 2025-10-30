package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	results <- fmt.Sprintf("%s 下载完成", filename)
}

func main() {
	var wg sync.WaitGroup
	myList := []string{"file1.zip", "file2.pdf", "file3.mp4"}
	channel := make(chan string, 3)
	for _, file := range myList {
		wg.Add(1)
		go download(file, &wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for f := range channel {
		fmt.Printf("%s\n", f)
	}

	fmt.Println("所有文件下载完成！")
}

// 预期输出:
// 开始下载 3 个文件...
// file1.zip 下载完成
// file2.pdf 下载完成
// file3.mp4 下载完成
// 所有文件下载完成!
// 提示:
// - 记得在启动 goroutine 前 wg.Add(1)
// - 用另一个 goroutine 等待并关闭 channel:

//   go func() {
//       wg.Wait()
//       close(results)
//   }()
