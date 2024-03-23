package xkcd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var comics []Comic

func (comic Comic) String() string {
	return fmt.Sprintf(`Title: %s
Image: %s
Transcript: %s\n`, comic.Title, comic.Img, comic.Transcript)

}

func init() {
	file, err := os.Open("data.json")

	if err != nil {
		log.Fatal("Unable to load data!")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&comics); err != nil {
		log.Fatal("Error loading data to memory from disk!!")
	}

	totalComics := len(comics)

	fmt.Printf("Loaded total %d comics\n", totalComics)
}

type Similarity struct {
	comic *Comic
	score int
}

func StartApp() {
	fmt.Println("Welcome to xkcd")
	fmt.Println("*********************************")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your search query: ")

	for scanner.Scan() {

		input := scanner.Text()

		result := getComics(input)

		fmt.Printf("\nBest comic match with score: %d:\n", result.score)
		fmt.Printf("%s\n", result.comic)
		fmt.Print("\nEnter your search query: ")
	}

}

func getComics(input string) Similarity {
	ans := Similarity{comic: nil, score: 0}
	for _, comic := range comics {
		temp := similarityScore(&comic, input)
		fmt.Printf("Similarity score: %d\n", temp.score)
		if temp.score > ans.score {
			ans = temp
		}
	}
	return ans
}

func similarityScore(comic *Comic, query string) Similarity {
	match := strings.Count(comic.Transcript, query)
	return Similarity{comic: comic, score: match}
}
