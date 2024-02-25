package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Solve() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Solve12() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}