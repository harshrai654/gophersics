package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	input := bufio.NewScanner(os.Stdin)
	freqMap := make(map[string]int)

	for input.Scan() {
		freqMap[input.Text()]++
	}

	for line, count := range freqMap {
		fmt.Printf("%s\t%d\n", line, count)
	}
}