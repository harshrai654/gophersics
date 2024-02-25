package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func FetchAll() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan string) {
	startTIme := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Unable to fetch from %s. | error: %s\n", url, err)
		return
	}

	status := resp.StatusCode

	size, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("Error copying response: %s\n", err)
		return
	}

	duration := time.Since(startTIme).Milliseconds()
	ch <- fmt.Sprintf("URL: %s | Status: %d | Time: %d ms |  Response size: %d Bytes  \n", url, status, duration, size)

}
