package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Some of these are bogus, to test errors
var urls = []string{"https://www.smh.com.au",
	"https://www.abc.net.au",
	"https://www.google.com.au",
	"https://www.golang.org",
	"https://www.theguardian.com/au",
	"https://junk.karl.lol",
	"https://www.netflix.com",
	"https://www.karl.lol",
	"https://www.amazon.com.au",
	"https://rambler.ru",
	"https://yahoo.com",
	"https://china.com.cn",
	"https://typepad.com",
	"https://netscape.com",
	"https://timesonline.co.uk",
	"https://over-blog.com",
	"https://google.co.uk",
	"https://blogger.com",
	"https://altervista.org",
	"https://howstuffworks.com",
	"https://admin.ch",
	"https://netlog.com",
	"https://ted.com",
	"https://amazonaws.com",
	"https://mail.ru",
	"https://google.com.hk",
	"https://walmart.com",
	"https://techcrunch.com",
	"https://cocolog-nifty.com"}

type RequestMessage struct {
	url string
}

func main() {
	startTime := time.Now()

	// Create a WaitGroup so we can wait for all goroutines to finish
	wg := sync.WaitGroup{}

	// Create a channel which will send requests to workers
	c := make(chan RequestMessage)

	// Create n worker threads
	for i := 0; i < 10; i++ {
		go startWorker(i, c, &wg)
	}

	// Post each work request to the channel
	for _, url := range urls {
		wg.Add(1)
		c <- RequestMessage{url}
	}

	fmt.Println("Waiting to finish up")
	wg.Wait()
	fmt.Println("All done!")

	finishTime := time.Now()

	fmt.Printf("Time taken: %v\n", finishTime.Sub(startTime))
}

func startWorker(id int, c chan RequestMessage, wg *sync.WaitGroup) {
	for {
		request := <-c
		fetch(id, request.url)
		wg.Done()
	}
}

func fetch(id int, url string) {
	// Call the website
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%d: %s\n", id, err)
		return
	}
	// Convert to a byte slice
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d: website %s has content length %d bytes\n", id, url, len(content))
	response.Body.Close()
}
