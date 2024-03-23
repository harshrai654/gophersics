package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func CharCount(file *os.File) {
	input := bufio.NewReader(file)
	letterFreq := make(map[rune]int)
	digitFreq := make(map[rune]int)
	var sizeCount [utf8.UTFMax + 1]int
	invalid := 0

	for {
		r, s, err := input.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && s == 1 {
			invalid++
		}

		switch {
		case unicode.IsDigit(r):
			digitFreq[r]++
		case unicode.IsLetter(r):
			letterFreq[r]++
		}

		sizeCount[s]++
	}

	fmt.Printf("Letter\tCount\n")
	for k, v := range letterFreq {
		fmt.Printf("%q\t%d\n", k, v)
	}

	fmt.Printf("\nDigit\tCount\n")
	for k, v := range digitFreq {
		fmt.Printf("%q\t%d\n", k, v)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range sizeCount {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func WordCount(file *os.File) {
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	wordFreq := make(map[string]int)
	lengthFreq := make(map[int]int)

	for scan.Scan() {
		str := scan.Text()
		wordFreq[str]++
		lengthFreq[len(str)]++
	}

	fmt.Printf("Word\tCount\n")
	for k, v := range wordFreq {
		fmt.Printf("%s\t%d\n", k, v)
	}

	fmt.Printf("Word Length\tCount\n")
	for k, v := range lengthFreq {
		fmt.Printf("%d\t%d\n", k, v)
	}
}
