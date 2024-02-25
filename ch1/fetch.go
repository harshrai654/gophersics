package ch1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			log.Panicf("Unable to fetch from %s. | error: %s\n", url, err)
		}

		fmt.Printf("Status: %d\n", resp.StatusCode)
		size, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Panicf("Error copying response: %s\n", err)
		}

		fmt.Printf("Response size: %d Bytes\n", size)
	}
}
