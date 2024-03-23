package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const urlFormat = "https://xkcd.com/%d/info.0.json"

type Comic struct {
	Month      string `json:"month"`
	Num        uint32 `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type ComicResult struct {
	comic *Comic
	err   error
}

func InitDB() {
	var wg sync.WaitGroup
	resultChannel := make(chan ComicResult, 100)
	const comicCount = 2000
	startTime := time.Now()

	file, err := os.Create("data.json")

	if err != nil {
		log.Fatal("Unable to create file!")
	}

	file.Write([]byte("["))

	go func() {
		for i := uint32(1); i < comicCount; i++ {
			wg.Add(1)
			go func(i uint32) {
				defer wg.Done()
				comic, err := fetchComic(i)
				result := ComicResult{comic: comic, err: err}

				if comic == nil || err != nil {
					return
				}
				resultChannel <- result
			}(i)
		}
		wg.Wait()
		close(resultChannel)
		file.Write([]byte("]"))
		file.Close()
	}()

	for result := range resultChannel {
		if result.comic == nil {
			continue
		}

		if result.err != nil {
			log.Fatal(result.err)
			continue
		}

		jsonBytes, err := json.Marshal(result.comic)

		if err != nil {
			log.Printf("Invalid JSON: %v\n", result.comic)
		}

		file.Write(jsonBytes)
		file.Write([]byte(","))
		file.Sync()

		fmt.Printf("Comic-%d Persisted to disk.\n", result.comic.Num)

	}

	endTime := time.Now()
	fmt.Printf("\nFetched %d comics in %f seconds\n", comicCount, endTime.Sub(startTime).Seconds())

}

func fetchComic(i uint32) (*Comic, error) {
	resp, err := http.Get(fmt.Sprintf(urlFormat, i))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid comic number %d\n", i)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "Invalid comic number %d\n", i)
		return nil, err
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
