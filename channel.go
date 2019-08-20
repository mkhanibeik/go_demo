package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func channel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Got %d\n", i)
	}
}

func channelChallenge() {
	ch := make(chan string)

	urls := []string{"https://golang.org", "https://api.github.com", "https://httpbin.org/xml"}
	go func() {
		for _, url := range urls {
			go returnTypeWithChannel(url, ch)
		}
	}()

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}

func returnTypeWithChannel(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR %s", err)
		out <- fmt.Sprintf("ERROR %s", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func selectChannel() {
	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("Got %f\n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout!")
	}
}
