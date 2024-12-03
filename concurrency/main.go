package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchPlaceholder(text string, done chan bool) {
	url := fmt.Sprintf("https://placehold.co/400x400?text=%s", text)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		done <- false
		return
	}
	defer resp.Body.Close()

	// Read a small portion of the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		done <- false
		return
	}

	// Output part of the response or response status
	fmt.Printf("Fetched URL %s - Status: %s\n", url, resp.Status)
	fmt.Printf("Body sample (length %d): %s...\n\n", len(body), string(body[:50]))
	done <- true
}

func otherWork(done chan bool) {
	// 模擬其他工作，這也需要等待
	for i := 0; i < 5; i++ {
		fmt.Printf("Doing other work... %d\n", i)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func main() {
	texts := []string{"1", "2", "3"}
	doneFetch := make(chan bool)  // 用於接收 fetch 操作的完成信號
	doneOther := make(chan bool)  // 用於接收其他任務的完成信號

	// 启动 fetch 操作
	for _, text := range texts {
		go fetchPlaceholder(text, doneFetch)
	}

	// 启动其他操作
	go otherWork(doneOther)

	// 等待所有 fetch 操作完成
	for range texts {
		<-doneFetch // 接收每次 fetch 操作的完成信號
	}

	// 等待其他工作完成
	<-doneOther

	fmt.Println("All operations completed!")
}